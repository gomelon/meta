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

type TemplateGenerator struct {
	path            string
	pkgPath         string
	metas           []Meta
	packageParser   *PackageParser
	metaParser      *MetaParser
	importTracker   ImportTracker
	funcMapProvider func(generator *TemplateGenerator) map[string]any
	tpl             *template.Template
}

type TGOption func(gen *TemplateGenerator)

func WithPackageParser(packageParser *PackageParser) TGOption {
	return func(gen *TemplateGenerator) {
		gen.packageParser = packageParser
	}
}

func WithMetaParser(metaParser *MetaParser) TGOption {
	return func(gen *TemplateGenerator) {
		gen.metaParser = metaParser
	}
}

func WithMetas(metas []Meta) TGOption {
	return func(gen *TemplateGenerator) {
		gen.metas = metas
	}
}

func WithImportTracker(tracker ImportTracker) TGOption {
	return func(gen *TemplateGenerator) {
		gen.importTracker = tracker
	}
}

func WithFuncMapProvider(provider func(generator *TemplateGenerator) map[string]any) TGOption {
	return func(gen *TemplateGenerator) {
		gen.funcMapProvider = provider
	}
}

func NewTemplateGenerator(path string, templateText string, options ...TGOption) (gen *TemplateGenerator, err error) {

	gen = &TemplateGenerator{
		metas: []Meta{},
		funcMapProvider: func(generator *TemplateGenerator) map[string]any {
			return map[string]any{}
		},
	}
	gen.path = path
	for _, option := range options {
		option(gen)
	}

	if gen.packageParser == nil {
		gen.packageParser = NewPackageParser()
	}

	if gen.metaParser == nil {
		gen.metaParser = NewMetaParser(gen.packageParser, gen.pkgPath, gen.metas)
	}

	err = gen.packageParser.Load(path)
	if err != nil {
		return nil, err
	}
	gen.pkgPath = gen.packageParser.PkgPath(path)

	if gen.importTracker == nil {
		gen.importTracker = NewDefaultImportTracker(gen.pkgPath)
	}

	functions := newFunctions(gen.packageParser, gen.metaParser, gen.importTracker, gen.pkgPath)

	tpl, err := template.New("TemplateGen").
		Funcs(sprig.GenericFuncMap()).
		Funcs(functions.FuncMap()).
		Funcs(gen.funcMapProvider(gen)).
		Parse(templateText)

	if err != nil {
		return
	}
	gen.tpl = tpl
	return gen, nil
}

func (gen *TemplateGenerator) Path() string {
	return gen.path
}

func (gen *TemplateGenerator) PkgPath() string {
	return gen.pkgPath
}

func (gen *TemplateGenerator) PackageParser() *PackageParser {
	return gen.packageParser
}

func (gen *TemplateGenerator) MetaParser() *MetaParser {
	return gen.metaParser
}

func (gen *TemplateGenerator) ImportTracker() ImportTracker {
	return gen.importTracker
}

func (gen *TemplateGenerator) Metas() []Meta {
	return gen.metas
}

func (gen *TemplateGenerator) GetMetaNames() []string {
	metas := gen.metas
	metaNames := make([]string, 0, len(metas))
	for _, meta := range metas {
		metaNames = append(metaNames, meta.Directive())
	}
	return metaNames
}

func (gen *TemplateGenerator) Bytes() ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	err := gen.generate(buffer)
	return buffer.Bytes(), err
}

func (gen *TemplateGenerator) Write(writer io.Writer) error {
	err := gen.generate(writer)
	return err
}

func (gen *TemplateGenerator) Print() error {
	err := gen.generate(os.Stdout)
	return err
}

func (gen *TemplateGenerator) Generate() error {
	file, err := os.Create(gen.outputFile())
	if err != nil {
		return err
	}
	return gen.generate(file)
}

func (gen *TemplateGenerator) generate(writer io.Writer) (err error) {

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

func (gen *TemplateGenerator) outputFile() string {
	pkg := gen.packageParser.Package(gen.pkgPath)
	return fmt.Sprintf("%s/%s_gen.go", gen.path, pkg.Name)
}

func (gen *TemplateGenerator) writerHeader(writer *bytes.Buffer) {
	gen.writeBuildTag(writer)
	gen.writePackage(writer)
	gen.writeImports(writer)
}

func (gen *TemplateGenerator) writeBuildTag(writer *bytes.Buffer) {
	metaNames := strings.Join(gen.GetMetaNames(), ",")
	comment := fmt.Sprintf(GeneratedComment, metaNames)
	_, _ = writer.WriteString(comment)
	_, _ = writer.WriteString("\n\n")

	buildTag := fmt.Sprintf("//+build !%s\n\n", GeneratedBuildTag)
	_, _ = writer.WriteString(buildTag)
	_, _ = writer.WriteString("\n\n")
}

func (gen *TemplateGenerator) writePackage(writer *bytes.Buffer) {
	pkg := gen.packageParser.Package(gen.pkgPath)
	_, _ = writer.WriteString("package ")
	_, _ = writer.WriteString(pkg.Name)
	_, _ = writer.WriteRune('\n')
}

func (gen *TemplateGenerator) writeImports(writer *bytes.Buffer) {
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
