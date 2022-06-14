package meta

import (
	"fmt"
	"go/token"
	"sort"
	"strings"
)

type ImportTracker interface {
	Import(pkgPath string) string
	LocalNameOf(packagePath string) string
	PathOf(localName string) (string, bool)
	ImportLines() []string
}

type DefaultImportTracker struct {
	pkgPathToName map[string]string
	nameToPkgPath map[string]string
	localPkgPath  string
}

func NewDefaultImportTracker(localPkgPath string) *DefaultImportTracker {
	return &DefaultImportTracker{
		pkgPathToName: map[string]string{},
		nameToPkgPath: map[string]string{},
		localPkgPath:  localPkgPath,
	}
}

func (tracker *DefaultImportTracker) Import(pkgPath string) string {
	if tracker.localPkgPath == pkgPath {
		return ""
	}
	name, ok := tracker.pkgPathToName[pkgPath]
	if ok {
		return name
	}
	name = PkgNameOrAlias(tracker, pkgPath)
	tracker.pkgPathToName[pkgPath] = name
	tracker.nameToPkgPath[name] = pkgPath
	return name
}

func (tracker *DefaultImportTracker) LocalNameOf(pkgPath string) string {
	return tracker.pkgPathToName[pkgPath]
}

func (tracker *DefaultImportTracker) PathOf(localName string) (string, bool) {
	name, ok := tracker.nameToPkgPath[localName]
	return name, ok
}

func (tracker *DefaultImportTracker) ImportLines() []string {
	var importPaths []string
	for name, pkgPath := range tracker.nameToPkgPath {
		var importPath string
		if name == pkgPath || strings.HasSuffix(pkgPath, "/"+name) {
			importPath = fmt.Sprintf("\"%s\"", pkgPath)
		} else {
			importPath = fmt.Sprintf("%s \"%s\"", name, pkgPath)
		}
		importPaths = append(importPaths, importPath)
	}
	sort.Sort(sort.StringSlice(importPaths))
	return importPaths
}

func PkgNameOrAlias(tracker ImportTracker, pkgPath string) string {

	dirs := strings.Split(pkgPath, "/")
	for n := len(dirs) - 1; n >= 0; n-- {
		// follow kube convention of not having anything between directory names
		name := strings.Join(dirs[n:], "")
		name = strings.Replace(name, "_", "", -1)
		// These characters commonly appear in import paths for go
		// packages, but aren't legal go names. So we'll sanitize.
		name = strings.Replace(name, ".", "", -1)
		name = strings.Replace(name, "-", "", -1)
		if _, found := tracker.PathOf(name); found {
			// This name collides with some other package
			continue
		}

		// If the import name is a Go keyword, prefix with an underscore.
		if token.Lookup(name).IsKeyword() {
			name = "_" + name
		}
		return name
	}
	panic("can't find import for " + pkgPath)
}
