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
	path          string
	pkgPath       string
	metas         []Meta
	importTracker ImportTracker
	packageParser *PackageParser
	tpl           *template.Template
}

//TODO 这个后续再增加更多的可定制化
func NewTemplateGenerator(workdir string, pkgPath string, metas []Meta, templateText string,
	funcMap map[string]any) (*TemplateGenerator, error) {

	packageParser := NewPackageParser(workdir)
	err := packageParser.Load(pkgPath)
	if err != nil {
		return nil, err
	}
	path := packageParser.Path(pkgPath)
	importTracker := NewDefaultImportTracker(pkgPath)
	templateFunc := NewTemplateFunc(importTracker, packageParser, pkgPath, metas)
	tpl, err := template.New("TemplateGen").
		Funcs(sprig.GenericFuncMap()).
		Funcs(templateFunc.FuncMap()).
		Funcs(funcMap).
		Parse(templateText)
	return &TemplateGenerator{
		path:          path,
		pkgPath:       pkgPath,
		metas:         metas,
		importTracker: importTracker,
		packageParser: packageParser,
		tpl:           tpl,
	}, nil
}

func (gen *TemplateGenerator) GetMetaNames() []string {
	metas := gen.metas
	metaNames := make([]string, 0, len(metas))
	for _, meta := range metas {
		metaNames = append(metaNames, meta.Name())
	}
	return metaNames
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

func (gen *TemplateGenerator) generate(writer io.Writer) error {
	bodyBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	err := gen.tpl.Execute(bodyBuf, map[string]any{})
	if err != nil {
		return err
	}

	headerBuf := bytes.NewBuffer(make([]byte, 0, 1024))
	gen.writerHeader(headerBuf)

	fileBuf := bytes.NewBuffer(make([]byte, 0, headerBuf.Len()+bodyBuf.Len()))
	_, _ = fileBuf.ReadFrom(headerBuf)
	_, _ = fileBuf.ReadFrom(bodyBuf)

	fileBytes := fileBuf.Bytes()
	fmtFileBytes, err := imports.Process(gen.outputFile(), fileBytes, nil)

	if err != nil {
		_, _ = writer.Write(fileBytes)
		return err
	}
	_, err = writer.Write(fmtFileBytes)
	return err
}

func (gen *TemplateGenerator) outputFile() string {
	return gen.path + "/x_gen.go"
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
