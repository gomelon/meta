package meta

import (
	"github.com/antonmedv/expr"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

type functions struct {
	importTracker ImportTracker
	metaParser    *MetaParser
	packageParser *PackageParser
	pkg           *packages.Package
	pkgPath       string
	typeQualifier types.Qualifier
}

func newFunctions(packageParser *PackageParser, metaParser *MetaParser,
	importTracker ImportTracker, pkgPath string) *functions {

	return &functions{
		importTracker: importTracker,
		metaParser:    metaParser,
		packageParser: packageParser,
		pkg:           packageParser.Package(pkgPath),
		pkgPath:       pkgPath,
		typeQualifier: func(p *types.Package) string {
			return importTracker.Import(p.Path())
		},
	}
}

func (f *functions) FuncMap() map[string]any {
	return map[string]any{
		"name":                       f.Name,
		"package":                    f.Package,
		"objectType":                 f.ObjectType,
		"structs":                    f.Structs,
		"interfaces":                 f.Interfaces,
		"methods":                    f.Methods,
		"interfaceMethods":           f.InterfaceMethods,
		"structMethods":              f.StructMethods,
		"filterByMeta":               f.FilterByMeta,
		"filterByMetaExpr":           f.FilterByMetaExpr,
		"hasMeta":                    f.HasMeta,
		"filterByMethodContainsMeta": f.FilterByMethodContainsMeta,
		"hasMethodContainsMeta":      f.HasMethodContainsMeta,
		"declare":                    f.Declare,
		"sign":                       f.Sign,
		"methodSignature":            f.MethodSignature,
		"import":                     f.Import,
	}
}

func (f *functions) Name(in any) string {
	var name string
	switch in := in.(type) {
	case *packages.Package:
		name = in.Name
	case types.Object:
		name = in.Name()
	}
	return name
}

func (f *functions) Package() *packages.Package {
	return f.pkg
}

func (f *functions) ObjectType(object types.Object) Type {
	return f.packageParser.ObjectType(object)
}

func (f *functions) Structs() []types.Object {
	return f.filterByType(TypeStruct)
}

func (f *functions) Interfaces() []types.Object {
	return f.filterByType(TypeInterface)
}

func (f *functions) Methods(object types.Object) []types.Object {
	return f.packageParser.Methods(object)
}

func (f *functions) InterfaceMethods(object types.Object) []types.Object {
	return f.packageParser.InterfaceMethods(object)
}

func (f *functions) StructMethods(object types.Object) []types.Object {
	return f.packageParser.StructMethods(object)
}

func (f *functions) FilterByMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if f.HasMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (f *functions) FilterByMetaExpr(metaName string, exprStr string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if !f.HasMeta(metaName, object) {
			continue
		}
		metas := f.metaParser.ObjectMetaGroup(object, metaName)
		output, err := expr.Eval(exprStr, map[string]any{
			"metas":  metas,
			"object": objects,
		})
		if err != nil {
			panic(err)
		}
		if output != true {
			continue
		}
		filteredObjects = append(filteredObjects, object)
	}
	return filteredObjects
}

func (f *functions) HasMeta(metaName string, object types.Object) bool {
	metas := f.metaParser.ObjectMetaGroups(object, metaName)
	return len(metas) > 0
}

func (f *functions) FilterByMethodContainsMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if f.HasMethodContainsMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (f *functions) HasMethodContainsMeta(metaName string, object types.Object) bool {
	methods := f.Methods(object)
	for _, method := range methods {
		if f.HasMeta(metaName, method) {
			return true
		}
	}
	return false
}

func (f *functions) Declare(object types.Object) string {
	if len(object.Name()) > 0 {
		return object.Name() + " " + f.Sign(object)
	} else {
		return f.Sign(object)
	}
}

func (f *functions) Sign(object types.Object) string {
	switch f.ObjectType(object) {
	case TypeInterfaceMethod, TypeStructMethod, TypeFunc:
		signature := f.MethodSignature(object)
		builder := strings.Builder{}
		builder.Grow(32)
		builder.WriteRune('(')
		params := signature.Params()
		for i, l := 0, params.Len(); i < l; i++ {
			param := params.At(i)
			builder.WriteString(f.Declare(param))
			builder.WriteRune(',')
		}
		builder.WriteRune(')')

		builder.WriteRune(' ')

		builder.WriteRune('(')
		results := signature.Results()
		for i, l := 0, results.Len(); i < l; i++ {
			result := results.At(i)
			builder.WriteString(f.Declare(result))
			builder.WriteRune(',')
		}
		builder.WriteRune(')')

		return builder.String()
	case TypeVar, TypeFuncVar:
		return types.TypeString(object.Type(), f.typeQualifier)
	}
	return ""
}

func (f *functions) MethodSignature(object types.Object) *types.Signature {
	return object.Type().(*types.Signature)
}

func (f *functions) Import(pkgPath string) string {
	return f.importTracker.Import(pkgPath)
}

func (f *functions) filterByType(typ Type) []types.Object {
	scope := f.pkg.Types.Scope()
	var result []types.Object
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		objectType := f.packageParser.ObjectType(object)
		if objectType&typ > 0 {
			result = append(result, object)
		}
	}
	return result
}
