package meta

import (
	"github.com/gomelon/melon/pathx"
	"io/ioutil"
)

type scanner struct {
	moduleSrcPath     string
	err               error
	generatorBuilders []*generatorBuilder
}

func Scan(rootPath string) *scanner {
	moduleSrcPath, err := pathx.ModuleSrcPath(rootPath)
	return &scanner{
		moduleSrcPath: moduleSrcPath,
		err:           err,
	}
}

func (s *scanner) Generate() error {
	if s.err != nil {
		return s.err
	}

	for _, gb := range s.generatorBuilders {
		err := pathx.AntScanThenDo(s.moduleSrcPath, true,
			func(absPath, relPath string) error {
				generator, err := NewTplPkgGenerator(absPath, gb.templateText, gb.options...)
				if err != nil {
					return err
				}
				err = generator.Generate()
				if err != nil {
					return err
				}
				return nil
			},
			gb.patterns...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *scanner) LocalTemplate(templatePath string) *generatorBuilder {
	if s.err != nil {
		return &generatorBuilder{scanner: s}
	}
	templateBytes, err := ioutil.ReadFile(s.moduleSrcPath + templatePath)
	if err != nil {
		return &generatorBuilder{scanner: s}
	}
	return &generatorBuilder{
		templateText: string(templateBytes),
		scanner:      s,
	}
}

func (s *scanner) TemplateText(templateText string) *generatorBuilder {
	if s.err != nil {
		return &generatorBuilder{scanner: s}
	}
	return &generatorBuilder{
		templateText: templateText,
		scanner:      s,
	}
}

type generatorBuilder struct {
	templateText string
	patterns     []string
	options      []TPGOption
	scanner      *scanner
}

func (g *generatorBuilder) Patterns(patterns ...string) *generatorBuilder {
	g.patterns = patterns
	return g
}

func (g *generatorBuilder) Options(patterns ...string) *generatorBuilder {
	g.patterns = patterns
	return g
}

func (g *generatorBuilder) Build() *scanner {
	g.scanner.generatorBuilders = append(g.scanner.generatorBuilders, g)
	return g.scanner
}
