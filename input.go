package meta

import (
	"golang.org/x/tools/go/packages"
)

type Input struct {
	pkg     *packages.Package
	pkgPath string
	//Consts     []*Const
	//Vars       []*Var
	Interfaces []*Interface
	Structs    []*Struct
	// Members    []*Member
	// Methods    []*Method
	//Funcs []*Func
}

func NewInput(pkg *packages.Package, pkgPath string) *Input {
	return &Input{
		pkg:        pkg,
		pkgPath:    pkgPath,
		Interfaces: []*Interface{},
		Structs:    []*Struct{},
	}
}

func (i *Input) Path() string {
	return i.pkg.PkgPath
}

func (i *Input) PkgPath() string {
	return i.pkgPath
}

func (i *Input) IsEmpty() bool {
	return len(i.Interfaces) == 0 && len(i.Structs) == 0
}
