package meta

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig/v3"
	"golang.org/x/tools/imports"
	"io"
	"os"
	"strings"
	"text/template"
)

func ScanToGenerate() error {
	return nil
}

type PkgGenerator interface {
	Path() string
	PkgPath() string
	PackageParser() *PackageParser
	MetaParser() *Parser
	PkgFunctions() *functions
	ImportTracker() ImportTracker
	Metas() []Meta
	GetMetaNames() []string
	Write(writer io.Writer) error
	Print() error
	Generate() error
}

type TplPkgGenerator struct {
	path            string
	pkgPath         string
	metas           []Meta
	packageParser   *PackageParser
	metaParser      *Parser
	pkgFunctions    *functions
	importTracker   ImportTracker
	funcMapProvider func(generator *TplPkgGenerator) map[string]any
	tpl             *template.Template

	outputFilePrefix string
	outputFileSuffix string
	outputFileName   string
}

func NewTplPkgGenerator(path string, templateText string, options ...TGOption) (tplPkgGen PkgGenerator, err error) {

	tplPkgGenImpl := &TplPkgGenerator{
		metas: []Meta{},
		funcMapProvider: func(generator *TplPkgGenerator) map[string]any {
			return map[string]any{}
		},
		outputFilePrefix: "zz_",
		outputFileSuffix: "_gen",
	}
	tplPkgGenImpl.path = path
	for _, option := range options {
		option(tplPkgGenImpl)
	}

	if tplPkgGenImpl.packageParser == nil {
		tplPkgGenImpl.packageParser = NewPackageParser()
	}

	if tplPkgGenImpl.metaParser == nil {
		tplPkgGenImpl.metaParser = NewParser(tplPkgGenImpl.packageParser, tplPkgGenImpl.pkgPath)
	}
	tplPkgGenImpl.metaParser.AddMeta(tplPkgGenImpl.metas...)

	err = tplPkgGenImpl.packageParser.Load(path)
	if err != nil {
		return nil, err
	}
	tplPkgGenImpl.pkgPath = tplPkgGenImpl.packageParser.PkgPath(path)

	if tplPkgGenImpl.importTracker == nil {
		tplPkgGenImpl.importTracker = NewDefaultImportTracker(tplPkgGenImpl.pkgPath)
	}

	functions := newFunctions(tplPkgGenImpl.packageParser, tplPkgGenImpl.metaParser, tplPkgGenImpl.importTracker, tplPkgGenImpl.pkgPath)
	tplPkgGenImpl.pkgFunctions = functions

	tpl, err := template.New("TemplateGen").
		Funcs(sprig.GenericFuncMap()).
		Funcs(functions.FuncMap()).
		Funcs(tplPkgGenImpl.funcMapProvider(tplPkgGenImpl)).
		Parse(templateText)

	if err != nil {
		return
	}
	tplPkgGenImpl.tpl = tpl
	tplPkgGen = tplPkgGenImpl
	return
}

type TGOption func(gen *TplPkgGenerator)

func WithPackageParser(packageParser *PackageParser) TGOption {
	return func(gen *TplPkgGenerator) {
		gen.packageParser = packageParser
	}
}

func WithMetaParser(metaParser *Parser) TGOption {
	return func(gen *TplPkgGenerator) {
		gen.metaParser = metaParser
	}
}

func WithMetas(metas []Meta) TGOption {
	return func(gen *TplPkgGenerator) {
		gen.metas = metas
	}
}

func WithImportTracker(tracker ImportTracker) TGOption {
	return func(gen *TplPkgGenerator) {
		gen.importTracker = tracker
	}
}

func WithFuncMapProvider(provider func(generator *TplPkgGenerator) map[string]any) TGOption {
	return func(gen *TplPkgGenerator) {
		gen.funcMapProvider = provider
	}
}

func WithOutputFilePrefix(prefix string) TGOption {
	return func(gen *TplPkgGenerator) {
		gen.outputFilePrefix = prefix
	}
}

func (gen *TplPkgGenerator) Path() string {
	return gen.path
}

func (gen *TplPkgGenerator) PkgPath() string {
	return gen.pkgPath
}

func (gen *TplPkgGenerator) PackageParser() *PackageParser {
	return gen.packageParser
}

func (gen *TplPkgGenerator) MetaParser() *Parser {
	return gen.metaParser
}

func (gen *TplPkgGenerator) PkgFunctions() *functions {
	return gen.pkgFunctions
}

func (gen *TplPkgGenerator) ImportTracker() ImportTracker {
	return gen.importTracker
}

func (gen *TplPkgGenerator) Metas() []Meta {
	return gen.metas
}

func (gen *TplPkgGenerator) GetMetaNames() []string {
	metas := gen.metas
	metaNames := make([]string, 0, len(metas))
	for _, meta := range metas {
		metaNames = append(metaNames, meta.Directive())
	}
	return metaNames
}

func (gen *TplPkgGenerator) Bytes() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := gen.generate(buffer)
	return buffer.Bytes(), err
}

func (gen *TplPkgGenerator) Write(writer io.Writer) error {
	err := gen.generate(writer)
	return err
}

func (gen *TplPkgGenerator) Print() error {
	err := gen.generate(os.Stdout)
	return err
}

func (gen *TplPkgGenerator) Generate() error {
	file, err := os.Create(gen.outputFile())
	if err != nil {
		return err
	}
	return gen.generate(file)
}

func (gen *TplPkgGenerator) generate(writer io.Writer) (err error) {

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

func (gen *TplPkgGenerator) outputFile() string {
	pkg := gen.packageParser.Package(gen.pkgPath)
	outputFileName := gen.outputFileName
	if len(outputFileName) == 0 {
		if len(gen.metas) > 0 {
			directiveParts := strings.Split(gen.metas[0].Directive(), ":")
			if len(directiveParts) > 1 {
				outputFileName = strings.Join(directiveParts[:len(directiveParts)-1], "_")
			} else {
				outputFileName = directiveParts[0]
			}
		} else {
			outputFileName = pkg.Name
		}
		outputFileName = gen.outputFilePrefix + outputFileName + gen.outputFileSuffix
	}

	return fmt.Sprintf("%s/%s.go", gen.path, outputFileName)
}

func (gen *TplPkgGenerator) writerHeader(writer *bytes.Buffer) {
	gen.writeBuildTag(writer)
	gen.writePackage(writer)
	gen.writeImports(writer)
}

func (gen *TplPkgGenerator) writeBuildTag(writer *bytes.Buffer) {
	metaNames := strings.Join(gen.GetMetaNames(), ",")
	comment := fmt.Sprintf(GeneratedComment, metaNames)
	_, _ = writer.WriteString(comment)
	_, _ = writer.WriteString("\n\n")

	buildTag := fmt.Sprintf("//+build !%s\n\n", GeneratedBuildTag)
	_, _ = writer.WriteString(buildTag)
	_, _ = writer.WriteString("\n\n")
}

func (gen *TplPkgGenerator) writePackage(writer *bytes.Buffer) {
	pkg := gen.packageParser.Package(gen.pkgPath)
	_, _ = writer.WriteString("package ")
	_, _ = writer.WriteString(pkg.Name)
	_, _ = writer.WriteRune('\n')
}

func (gen *TplPkgGenerator) writeImports(writer *bytes.Buffer) {
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
