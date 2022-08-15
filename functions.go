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
	metaParser    *Parser
	pkgParser     *PkgParser
	pkg           *packages.Package
	pkgPath       string
	typeQualifier types.Qualifier
}

func newFunctions(gen *TmplPkgGen) *functions {

	return &functions{
		importTracker: gen.importTracker,
		metaParser:    gen.metaParser,
		pkgParser:     gen.pkgParser,
		pkg:           gen.pkgParser.Package(gen.pkgPath),
		pkgPath:       gen.pkgPath,
		typeQualifier: gen.typeQualifier,
	}
}

func (f *functions) FuncMap() map[string]any {
	return map[string]any{
		"name":                  f.Name,
		"fullName":              f.FullName,
		"package":               f.Package,
		"object":                f.pkgParser.Object,
		"assignableToCtx":       f.pkgParser.AssignableToCtx,
		"assignableTo":          f.pkgParser.AssignableTo,
		"anonymousAssign":       f.pkgParser.AnonymousAssign,
		"anonymousAssignTo":     f.pkgParser.AnonymousAssignTo,
		"exported":              f.Exported,
		"objectPlace":           f.ObjectPlace,
		"objectType":            f.ObjectType,
		"typeName":              f.pkgParser.TypeName,
		"underlyingType":        f.pkgParser.UnderlyingType,
		"structs":               f.Structs,
		"interfaces":            f.Interfaces,
		"functions":             f.Functions,
		"methods":               f.pkgParser.Methods,
		"params":                f.pkgParser.Params,
		"firstParam":            f.pkgParser.FirstParam,
		"results":               f.pkgParser.Results,
		"firstResult":           f.pkgParser.FirstResult,
		"lastResult":            f.pkgParser.LastResult,
		"hasErrorResult":        f.pkgParser.HasErrorResult,
		"filterByMeta":          f.metaParser.FilterByMeta,
		"filterByMetaExpr":      f.FilterByMetaExpr,
		"hasMeta":               f.metaParser.HasMeta,
		"filterByMethodHasMeta": f.metaParser.FilterByMethodHasMeta,
		"hasMethodHasMeta":      f.metaParser.HasMethodHasMeta,
		"metaProp":              f.metaParser.HasMethodHasMeta,
		"filterObjects":         f.FilterObjects,
		"indirect":              f.pkgParser.Indirect,
		"declare":               f.Declare,
		"declareType":           f.DeclareType,
		"typeString":            f.TypeString,
		"initType":              f.InitType,
		"methodSignature":       f.MethodSignature,
		"import":                f.Import,
		"objectMetaGroups":      f.metaParser.ObjectMetaGroups,
		"objectMetaGroup":       f.metaParser.ObjectMetaGroup,
		"objectMeta":            f.metaParser.ObjectMeta,
		"multipleLines":         f.MultipleLines,
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

func (f *functions) FullName(in any) string {
	var name string
	switch in := in.(type) {
	case *packages.Package:
		name = in.Name
	case types.Object:
		name = in.Pkg().Path() + "." + in.Name()
	}
	return name
}

func (f *functions) Package() *packages.Package {
	return f.pkg
}

func (f *functions) Exported(object types.Object) bool {
	return object.Exported()
}

func (f *functions) ObjectPlace(object types.Object) Place {
	return f.pkgParser.ObjectPlace(object)
}

func (f *functions) ObjectType(object types.Object) types.Type {
	return object.Type()
}

func (f *functions) Structs() []types.Object {
	return f.pkgParser.Structs(f.pkgPath)
}

func (f *functions) Interfaces() []types.Object {
	return f.pkgParser.Interfaces(f.pkgPath)
}

func (f *functions) Functions() []types.Object {
	return f.pkgParser.Functions(f.pkgPath)
}

func (f *functions) FilterByMetaExpr(metaName string, exprStr string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if !f.metaParser.HasMeta(metaName, object) {
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
			result = "false"
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

func (f *functions) MultipleLines(line string) string {
	return "\"" + strings.ReplaceAll(line, "\n", "\\n\"+\n\"") + "\""
}
