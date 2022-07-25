package meta

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"golang.org/x/tools/imports"
	"io"
	"os"
	"path/filepath"
	"text/template"
)

type TmplPkgGenerator struct {
	path            string
	pkgPath         string
	pkgParser       *PkgParser
	metaParser      *Parser
	pkgFunctions    *functions
	importTracker   ImportTracker
	funcMapProvider func(generator *TmplPkgGenerator) map[string]any
	tpl             *template.Template

	outputFilePrefix   string
	outputFileSuffix   string
	outputFilename     string
	outputFullFilename string
}

func NewTplPkgGenerator(path string, templateText string, options ...TPGOption) (gen *TmplPkgGenerator, err error) {

	gen = &TmplPkgGenerator{
		path:             path,
		outputFilePrefix: DefaultOutputFilePrefix,
		outputFileSuffix: DefaultOutputFileSuffix,
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

	if gen.funcMapProvider == nil {
		gen.funcMapProvider = func(generator *TmplPkgGenerator) map[string]any {
			return map[string]any{}
		}
	}

	err = gen.pkgParser.Load(path)
	if err != nil {
		return
	}
	gen.pkgPath = gen.pkgParser.PkgPath(path)

	if gen.importTracker == nil {
		gen.importTracker = NewDefaultImportTracker(gen.pkgPath)
	}

	functions := newFunctions(gen.pkgParser, gen.metaParser, gen.importTracker, gen.pkgPath)
	gen.pkgFunctions = functions

	tpl, err := template.New("TemplateGen").
		Funcs(sprig.GenericFuncMap()).
		Funcs(functions.FuncMap()).
		Funcs(gen.funcMapProvider(gen)).
		Parse(templateText)
	gen.tpl = tpl
	return
}

type TPGOption func(gen *TmplPkgGenerator)

func WithPkgParser(pkgParser *PkgParser) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.pkgParser = pkgParser
	}
}

func WithMetaParser(metaParser *Parser) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.metaParser = metaParser
	}
}

func WithImportTracker(tracker ImportTracker) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.importTracker = tracker
	}
}

func WithFuncMapProvider(provider func(generator *TmplPkgGenerator) map[string]any) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.funcMapProvider = provider
	}
}

func WithOutputFilePrefix(prefix string) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.outputFilePrefix = prefix
	}
}

func WithOutputFileSuffix(suffix string) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.outputFileSuffix = suffix
	}
}

func WithOutputFilename(filename string) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.outputFilename = filename
	}
}

func WithOutputFullFilename(fullFilename string) TPGOption {
	return func(gen *TmplPkgGenerator) {
		gen.outputFullFilename = fullFilename
	}
}

func (gen *TmplPkgGenerator) Path() string {
	return gen.path
}

func (gen *TmplPkgGenerator) PkgPath() string {
	return gen.pkgPath
}

func (gen *TmplPkgGenerator) PkgParser() *PkgParser {
	return gen.pkgParser
}

func (gen *TmplPkgGenerator) MetaParser() *Parser {
	return gen.metaParser
}

func (gen *TmplPkgGenerator) PkgFunctions() *functions {
	return gen.pkgFunctions
}

func (gen *TmplPkgGenerator) ImportTracker() ImportTracker {
	return gen.importTracker
}

func (gen *TmplPkgGenerator) Bytes() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := gen.generate(buffer)
	return buffer.Bytes(), err
}

func (gen *TmplPkgGenerator) Write(writer io.Writer) error {
	err := gen.generate(writer)
	return err
}

func (gen *TmplPkgGenerator) Print() error {
	err := gen.generate(os.Stdout)
	return err
}

func (gen *TmplPkgGenerator) Generate() error {
	file, err := os.Create(gen.outputFile())
	if err != nil {
		return err
	}
	return gen.generate(file)
}

func (gen *TmplPkgGenerator) generate(writer io.Writer) (err error) {

	bodyBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	err = gen.tpl.Execute(bodyBuf, map[string]any{})
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

func (gen *TmplPkgGenerator) outputFile() string {
	pkg := gen.pkgParser.Package(gen.pkgPath)
	outputFullFilename := gen.outputFullFilename
	if len(outputFullFilename) == 0 {
		if len(gen.outputFilename) == 0 {
			outputFullFilename = gen.outputFilePrefix + pkg.Name + gen.outputFileSuffix
		} else {
			outputFullFilename = gen.outputFilePrefix + gen.outputFilename + gen.outputFileSuffix
		}
	}

	return fmt.Sprintf("%s%s%s.go", gen.path, string(filepath.Separator), outputFullFilename)
}

func (gen *TmplPkgGenerator) writerHeader(writer *bytes.Buffer) {
	//gen.writeBuildTag(writer)
	gen.writePackage(writer)
	gen.writeImports(writer)
}

func (gen *TmplPkgGenerator) writeBuildTag(writer *bytes.Buffer) {
	_, _ = writer.WriteString(GeneratedBuildTag)
	_, _ = writer.WriteString("\n\n")

	buildTag := fmt.Sprintf("//+build !%s\n\n", GeneratedBuildTag)
	_, _ = writer.WriteString(buildTag)
	_, _ = writer.WriteString("\n\n")
}

func (gen *TmplPkgGenerator) writePackage(writer *bytes.Buffer) {
	pkg := gen.pkgParser.Package(gen.pkgPath)
	_, _ = writer.WriteString("package ")
	_, _ = writer.WriteString(pkg.Name)
	_, _ = writer.WriteRune('\n')
}

func (gen *TmplPkgGenerator) writeImports(writer *bytes.Buffer) {
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
