package meta

import (
	"fmt"
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
		"objectByPkgPathAndName":     f.ObjectByPkgPathAndName,
		"assignableToCtx":            f.AssignableToCtx,
		"assignableTo":               f.AssignableTo,
		"exported":                   f.Exported,
		"objectPlace":                f.ObjectPlace,
		"objectType":                 f.ObjectType,
		"typeName":                   f.TypeName,
		"underlyingType":             f.UnderlyingType,
		"structs":                    f.Structs,
		"interfaces":                 f.Interfaces,
		"methods":                    f.Methods,
		"interfaceMethods":           f.InterfaceMethods,
		"structMethods":              f.StructMethods,
		"params":                     f.Params,
		"firstParam":                 f.FirstParam,
		"results":                    f.Results,
		"firstResult":                f.FirstResult,
		"lastResult":                 f.LastResult,
		"hasErrorResult":             f.HasErrorResult,
		"filterByMeta":               f.FilterByMeta,
		"filterByMetaExpr":           f.FilterByMetaExpr,
		"hasMeta":                    f.HasMeta,
		"filterByMethodContainsMeta": f.FilterByMethodContainsMeta,
		"hasMethodContainsMeta":      f.HasMethodContainsMeta,
		"filterObjects":              f.FilterObjects,
		"indirect":                   f.Indirect,
		"declare":                    f.Declare,
		"declareType":                f.DeclareType,
		"typeString":                 f.TypeString,
		"initType":                   f.InitType,
		"methodSignature":            f.MethodSignature,
		"import":                     f.Import,
		"objectMetaGroups":           f.ObjectMetaGroups,
		"objectMetaGroup":            f.ObjectMetaGroup,
		"multipleLines":              f.MultipleLines,
		"fail":                       f.Fail,
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

func (f *functions) ObjectByPkgPathAndName(pkgPath, typeName string) types.Object {
	return f.packageParser.ObjectByPkgPathAndName(pkgPath, typeName)
}

func (f *functions) AssignableToCtx(v types.Type) bool {
	return f.packageParser.AssignableToCtx(v)
}

func (f *functions) AssignableTo(v, t types.Type) bool {
	return f.packageParser.AssignableTo(v, t)
}

func (f *functions) Exported(object types.Object) bool {
	return object.Exported()
}

func (f *functions) ObjectPlace(object types.Object) Place {
	return f.packageParser.ObjectPlace(object)
}

func (f *functions) ObjectType(object types.Object) types.Type {
	return object.Type()
}

func (f *functions) TypeName(typ types.Type) string {
	return f.packageParser.TypeName(typ)
}

func (f *functions) UnderlyingType(typ types.Type) types.Type {
	return f.packageParser.UnderlyingType(typ)
}

func (f *functions) Structs() []types.Object {
	return f.filterByPlace(PlaceStruct)
}

func (f *functions) Interfaces() []types.Object {
	return f.filterByPlace(PlaceInterface)
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

func (f *functions) Params(methodOrFunc types.Object) []types.Object {
	return f.packageParser.Params(methodOrFunc)
}

func (f *functions) FirstParam(methodOrFunc types.Object) types.Object {
	return f.packageParser.FirstParam(methodOrFunc)
}

func (f *functions) Results(methodOrFunc types.Object) []types.Object {
	return f.packageParser.Results(methodOrFunc)
}

func (f *functions) FirstResult(methodOrFunc types.Object) types.Object {
	return f.packageParser.FirstResult(methodOrFunc)
}

func (f *functions) LastResult(methodOrFunc types.Object) types.Object {
	return f.packageParser.LastResult(methodOrFunc)
}

func (f *functions) HasErrorResult(methodOrFunc types.Object) bool {
	return f.packageParser.HasErrorResult(methodOrFunc)
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

func (f *functions) FilterObjects(filterFuncName string, objects []types.Object) []types.Object {
	filter := f.FuncMap()[filterFuncName].(func(types.Object) bool)
	result := make([]types.Object, 2)
	for _, object := range objects {
		if !filter(object) {
			continue
		}
		result = append(result, object)
	}
	return result
}

func (f *functions) Indirect(typ types.Type) types.Type {
	return f.packageParser.Indirect(typ)
}

func (f *functions) Declare(object types.Object) string {
	if object == nil {
		return ""
	}
	if len(object.Name()) > 0 {
		return object.Name() + " " + f.DeclareType(object)
	} else {
		return f.DeclareType(object)
	}
}

func (f *functions) DeclareType(object types.Object) string {
	switch f.ObjectPlace(object) {
	case PlaceInterfaceMethod, PlaceStructMethod, PlaceFunc:
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
	case PlaceVar, PlaceFuncVar:
		return f.TypeString(object.Type())
	}
	return ""
}

func (f *functions) TypeString(typ types.Type) string {
	return types.TypeString(typ, f.typeQualifier)
}

func (f *functions) InitType(typ types.Type) string {
	var result string
	switch typ := typ.(type) {
	case *types.Pointer:
		result = "&" + f.InitType(typ.Elem())
	case *types.Struct:
		result = "{}"
	case *types.Basic:
		switch typ.Kind() {
		case types.String:
			result = "\"\""
		case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
			types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
			types.Uintptr,
			types.Float32, types.Float64:
			result = typ.Name() + "(0)"
		case types.Bool:
		case types.Complex64, types.Complex128:
			result = typ.Name() + "(0,0)"
		default:
			panic(fmt.Errorf("meta: unsupported type to init, type=%s", typ.Name()))
		}
	case *types.Slice:
		result = f.TypeString(typ) + "{}"
	case *types.Map:
		//TODO wait complete
	case *types.Named:
		result = f.TypeString(typ) + f.InitType(typ.Underlying())
	case *types.Tuple:
		//TODO wait complete
	case *types.Array:
		//TODO wait complete
	case *types.Chan:
		//TODO wait complete
	default:
		result = typ.String()
	}
	return result
}

func (f *functions) MethodSignature(object types.Object) *types.Signature {
	return object.Type().(*types.Signature)
}

func (f *functions) Import(pkgPath string) string {
	return f.importTracker.Import(pkgPath)
}

func (f *functions) ObjectMetaGroups(object types.Object, metaNames ...string) map[string]Group {
	return f.metaParser.ObjectMetaGroups(object, metaNames...)
}

func (f *functions) ObjectMetaGroup(object types.Object, metaName string) Group {
	return f.metaParser.ObjectMetaGroup(object, metaName)
}

func (f *functions) MultipleLines(linePrefix, lineSuffix, line string) string {
	return strings.ReplaceAll(line, "\n", lineSuffix+"\"+\n"+linePrefix+"\"")
}

func (f *functions) Fail(err string, args ...any) (any, error) {
	return nil, fmt.Errorf(err, args...)
}

func (f *functions) filterByPlace(place Place) []types.Object {
	scope := f.pkg.Types.Scope()
	var result []types.Object
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		objectPlace := f.packageParser.ObjectPlace(object)
		if objectPlace&place > 0 {
			result = append(result, object)
		}
	}
	return result
}
