package meta

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"go/types"
	"golang.org/x/tools/imports"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type TmplPkgGenFactory struct {
	templateText string
	options      []TPGOption
}

func NewTmplPkgGenFactory(templateText string, options ...TPGOption) *TmplPkgGenFactory {
	return &TmplPkgGenFactory{templateText: templateText, options: options}
}

func (t *TmplPkgGenFactory) Create(absPath, relPath string) (PkgGen, error) {
	return NewTmplPkgGen(absPath, t.templateText, t.options...)
}

type TmplPkgGen struct {
	path             string
	pkgPath          string
	pkgParser        *PkgParser
	metaParser       *Parser
	pkgFunctions     *functions
	importTracker    ImportTracker
	typeQualifier    types.Qualifier
	funcMap          template.FuncMap
	funcMapFactories []TPGFuncMapFactory
	tpl              *template.Template

	outputFilePrefix   string
	outputFileSuffix   string
	outputFilename     string
	outputFullFilename string
}

type TPGFuncMapFactory func(generator *TmplPkgGen) template.FuncMap

func NewTmplPkgGen(path string, templateText string, options ...TPGOption) (gen *TmplPkgGen, err error) {

	gen = &TmplPkgGen{
		path:             path,
		outputFilePrefix: DefaultOutputFilePrefix,
		outputFileSuffix: DefaultOutputFileSuffix,
		funcMap:          map[string]any{},
	}

	for _, option := range options {
		option(gen)
	}

	if gen.metaParser == nil {
		if gen.pkgParser == nil {
			gen.metaParser = defaultParser
		} else {
			gen.metaParser = NewParser(gen.pkgParser)
		}
	}

	if gen.pkgParser == nil {
		gen.pkgParser = defaultPkgParser
	}

	err = gen.pkgParser.Load(path)
	if err != nil {
		return
	}
	gen.pkgPath = gen.pkgParser.PkgPath(path)

	if gen.importTracker == nil {
		gen.importTracker = NewDefaultImportTracker(gen.pkgPath)
	}
	gen.typeQualifier = func(p *types.Package) string {
		return gen.importTracker.Import(p.Path())
	}

	functions := newFunctions(gen)
	gen.pkgFunctions = functions

	for _, factory := range gen.funcMapFactories {
		funcMap := factory(gen)
		for name, fn := range funcMap {
			gen.funcMap[name] = fn
		}
	}
	tpl, err := template.New("TemplateGen").
		Funcs(sprig.GenericFuncMap()).
		Funcs(functions.FuncMap()).
		Funcs(gen.funcMap).
		Parse(templateText)
	gen.tpl = tpl
	return
}

func (gen *TmplPkgGen) PkgParser() *PkgParser {
	return gen.pkgParser
}

func (gen *TmplPkgGen) MetaParser() *Parser {
	return gen.metaParser
}

func (gen *TmplPkgGen) PkgFunctions() *functions {
	return gen.pkgFunctions
}

func (gen *TmplPkgGen) ImportTracker() ImportTracker {
	return gen.importTracker
}

func (gen *TmplPkgGen) PkgPath() string {
	return gen.pkgPath
}

type TPGOption func(gen *TmplPkgGen)

func WithPkgParser(pkgParser *PkgParser) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.pkgParser = pkgParser
	}
}

func WithMetaParser(metaParser *Parser) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.metaParser = metaParser
	}
}

func WithImportTracker(tracker ImportTracker) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.importTracker = tracker
	}
}

func WithFuncMap(funcMap template.FuncMap) TPGOption {
	return func(gen *TmplPkgGen) {
		for name, fn := range funcMap {
			gen.funcMap[name] = fn
		}
	}
}

func WithFuncMapFactory(factory TPGFuncMapFactory) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.funcMapFactories = append(gen.funcMapFactories, factory)
	}
}

func WithOutputFilePrefix(prefix string) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.outputFilePrefix = prefix
	}
}

func WithOutputFileSuffix(suffix string) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.outputFileSuffix = suffix
	}
}

func WithOutputFilename(filename string) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.outputFilename = filename
	}
}

func WithOutputFullFilename(fullFilename string) TPGOption {
	return func(gen *TmplPkgGen) {
		gen.outputFullFilename = fullFilename
	}
}

func (gen *TmplPkgGen) Bytes() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := gen.generate(
		func() (io.Writer, error) {
			return buffer, nil
		},
	)
	return buffer.Bytes(), err

}

func (gen *TmplPkgGen) Write(writer io.Writer) error {
	return gen.generate(
		func() (io.Writer, error) {
			return writer, nil
		},
	)
}

func (gen *TmplPkgGen) Print() error {
	return gen.generate(
		func() (io.Writer, error) {
			return os.Stdout, nil
		},
	)
}

func (gen *TmplPkgGen) Generate() error {
	return gen.generate(
		func() (io.Writer, error) {
			return os.Create(gen.outputFile())
		},
	)
}

func (gen *TmplPkgGen) generate(writerFunc func() (io.Writer, error)) (err error) {

	bodyBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	err = gen.tpl.Execute(bodyBuf, map[string]any{})
	if err != nil || len(strings.TrimSpace(bodyBuf.String())) == 0 {
		return
	}

	writer, err := writerFunc()
	if err != nil {
		return
	}

	headerBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	gen.writerHeader(headerBuf)

	fileBuf := bytes.NewBuffer(make([]byte, 0, headerBuf.Len()+bodyBuf.Len()))
	_, _ = fileBuf.ReadFrom(headerBuf)
	_, _ = fileBuf.ReadFrom(bodyBuf)

	fileBytes := fileBuf.Bytes()
	//format code and optimize import
	fmtFileBytes, err := imports.Process(gen.outputFile(), fileBytes, nil)

	if err != nil {
		_, _ = writer.Write(fileBytes)
		return
	}
	_, err = writer.Write(fmtFileBytes)
	return
}

func (gen *TmplPkgGen) outputFile() string {
	pkg := gen.pkgParser.Package(gen.pkgPath)
	outputFullFilename := gen.outputFullFilename
	if len(outputFullFilename) == 0 {
		if len(gen.outputFilename) == 0 {
			outputFullFilename = gen.outputFilePrefix + pkg.Name + gen.outputFileSuffix
		} else {
			outputFullFilename = gen.outputFilePrefix + gen.outputFilename + gen.outputFileSuffix
		}
	}

	return strings.Join([]string{gen.path, string(filepath.Separator), outputFullFilename, ".go"}, "")
}

func (gen *TmplPkgGen) writerHeader(writer *bytes.Buffer) {
	//gen.writeBuildTag(writer)
	gen.writePackage(writer)
	gen.writeImports(writer)
}

func (gen *TmplPkgGen) writeBuildTag(writer *bytes.Buffer) {
	_, _ = writer.WriteString(GeneratedBuildTag)
	_, _ = writer.WriteString("\n\n")

	buildTag := fmt.Sprintf("//+build !%s\n\n", GeneratedBuildTag)
	_, _ = writer.WriteString(buildTag)
	_, _ = writer.WriteString("\n\n")
}

func (gen *TmplPkgGen) writePackage(writer *bytes.Buffer) {
	pkg := gen.pkgParser.Package(gen.pkgPath)
	_, _ = writer.WriteString("package ")
	_, _ = writer.WriteString(pkg.Name)
	_, _ = writer.WriteRune('\n')
}

func (gen *TmplPkgGen) writeImports(writer *bytes.Buffer) {
	lines := gen.importTracker.ImportLines()

	if len(lines) == 0 {
		return
	}

	_, _ = writer.WriteString("import (\n")
	for _, line := range lines {
		_, _ = writer.WriteString(line)
		_, _ = writer.WriteRune('\n')
	}
	_, _ = writer.WriteString(")\n")
}
