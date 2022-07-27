package meta

import (
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

type PathMatcher func(absPath, relPath string, patterns ...string) (bool, error)

var RegexOrPathMatcher PathMatcher = func(absPath, relPath string, patterns ...string) (match bool, err error) {
	if len(patterns) == 0 || (len(patterns) == 1 && len(patterns[0]) == 0) {
		match = true
		return
	}
	for _, pattern := range patterns {
		match, err = regexp.MatchString(pattern, relPath)
		if err != nil {
			return
		}
		if match {
			return
		}
	}
	return
}

type scanner struct {
	rootPath             string
	err                  error
	pkgGenFactoryHolders []*pkgGenFactoryHolder
}

func ScanFor(rootPath string) *scanner {
	return &scanner{
		rootPath: rootPath,
		err:      nil,
	}
}

func ScanCurrentMod() *scanner {
	wd, err := os.Getwd()
	if err != nil {
		return nil
	}
	moduleSrcPath, err := ModuleSrcPath(wd)
	return &scanner{
		rootPath: moduleSrcPath,
		err:      err,
	}
}

func (s *scanner) Generate() error {
	if s.err != nil {
		return s.err
	}

	for _, factoryHolder := range s.pkgGenFactoryHolders {
		factory := factoryHolder.factory
		err := filepath.WalkDir(s.rootPath, func(absPath string, dirEntry fs.DirEntry, err error) error {
			if err != nil {
				return err
			}

			if !dirEntry.IsDir() {
				return nil
			}

			hasGoFile, err := HasGoFile(absPath)
			if err != nil {
				return err
			}
			if !hasGoFile {
				return nil
			}

			relPath, err := filepath.Rel(s.rootPath, absPath)
			if err != nil {
				return err
			}

			match, err := factoryHolder.pathMatcher(absPath, relPath, factoryHolder.patterns...)
			if err != nil {
				return err
			}
			if !match {
				return nil
			}
			generator, err := factory.Create(absPath, relPath)
			if err != nil {
				return err
			}
			return generator.Generate()
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *scanner) LocalTemplate(templatePath string) *tmplPkgGenFactoryBuilder {
	var templateText string
	if s.err == nil {
		templateBytes, err := ioutil.ReadFile(s.rootPath + templatePath)
		templateText = string(templateBytes)
		s.err = err
	}
	return s.TemplateText(templateText)
}

func (s *scanner) TemplateText(templateText string) *tmplPkgGenFactoryBuilder {
	if s.err != nil {
		return &tmplPkgGenFactoryBuilder{scanner: s}
	}
	return &tmplPkgGenFactoryBuilder{
		templateText: templateText,
		pathMatcher:  RegexOrPathMatcher,
		scanner:      s,
	}
}

func (s *scanner) PkgGenFactory(factory PkgGenFactory) *pkgGenFactoryHolderBuilder {
	return &pkgGenFactoryHolderBuilder{
		factory:     factory,
		pathMatcher: RegexOrPathMatcher,
		scanner:     s,
	}
}

type tmplPkgGenFactoryBuilder struct {
	templateText string
	pathMatcher  PathMatcher
	patterns     []string
	options      []TPGOption
	scanner      *scanner
}

func (g *tmplPkgGenFactoryBuilder) RegexOr(patterns ...string) *tmplPkgGenFactoryBuilder {
	g.pathMatcher = RegexOrPathMatcher
	g.patterns = patterns
	return g
}

func (g *tmplPkgGenFactoryBuilder) Options(options ...TPGOption) *tmplPkgGenFactoryBuilder {
	g.options = options
	return g
}

func (g *tmplPkgGenFactoryBuilder) And() *scanner {
	factory := &TmplPkgGenFactory{
		templateText: g.templateText,
		options:      g.options,
	}
	g.scanner.pkgGenFactoryHolders = append(g.scanner.pkgGenFactoryHolders, &pkgGenFactoryHolder{
		factory:     factory,
		patterns:    g.patterns,
		pathMatcher: RegexOrPathMatcher,
	})
	return g.scanner
}

type pkgGenFactoryHolderBuilder struct {
	factory     PkgGenFactory
	patterns    []string
	pathMatcher PathMatcher
	scanner     *scanner
}

func (g *pkgGenFactoryHolderBuilder) RegexOr(patterns ...string) *pkgGenFactoryHolderBuilder {
	g.pathMatcher = RegexOrPathMatcher
	g.patterns = patterns
	return g
}

func (g *pkgGenFactoryHolderBuilder) And() *scanner {
	g.scanner.pkgGenFactoryHolders = append(g.scanner.pkgGenFactoryHolders, &pkgGenFactoryHolder{
		factory:     g.factory,
		patterns:    g.patterns,
		pathMatcher: RegexOrPathMatcher,
	})
	return g.scanner
}

type pkgGenFactoryHolder struct {
	factory     PkgGenFactory
	patterns    []string
	pathMatcher PathMatcher
}
