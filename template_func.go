package meta

import (
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

type TemplateFunc struct {
	importTracker ImportTracker
	metaParser    *MetaParser
	packageParser *PackageParser
	pkg           *packages.Package
	pkgPath       string
	typeQualifier types.Qualifier
}

func NewTemplateFunc(importTracker ImportTracker, packageParser *PackageParser,
	pkgPath string, metas []Meta) *TemplateFunc {

	metaParser := NewMetaParser(packageParser, pkgPath, metas)
	return &TemplateFunc{
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

func (tf *TemplateFunc) FuncMap() map[string]any {
	return map[string]any{
		"name":                       tf.Name,
		"package":                    tf.Package,
		"objectType":                 tf.ObjectType,
		"structs":                    tf.Structs,
		"interfaces":                 tf.Interfaces,
		"methods":                    tf.Methods,
		"interfaceMethods":           tf.InterfaceMethods,
		"structMethods":              tf.StructMethods,
		"filterByMeta":               tf.FilterByMeta,
		"hasMeta":                    tf.HasMeta,
		"filterByMethodContainsMeta": tf.FilterByMethodContainsMeta,
		"hasMethodContainsMeta":      tf.HasMethodContainsMeta,
		"declare":                    tf.Declare,
		"sign":                       tf.Sign,
		"methodSignature":            tf.MethodSignature,
	}
}

func (tf *TemplateFunc) Name(in any) string {
	var name string
	switch in := in.(type) {
	case *packages.Package:
		name = in.Name
	case types.Object:
		name = in.Name()
	}
	return name
}

func (tf *TemplateFunc) Package() *packages.Package {
	return tf.pkg
}

func (tf *TemplateFunc) ObjectType(object types.Object) Type {
	return tf.packageParser.ObjectType(object)
}

func (tf *TemplateFunc) Structs() []types.Object {
	return tf.filterByType(TypeStruct)
}

func (tf *TemplateFunc) Interfaces() []types.Object {
	return tf.filterByType(TypeInterface)
}

func (tf *TemplateFunc) Methods(object types.Object) []types.Object {
	return tf.packageParser.Methods(object)
}

func (tf *TemplateFunc) InterfaceMethods(object types.Object) []types.Object {
	return tf.packageParser.InterfaceMethods(object)
}

func (tf *TemplateFunc) StructMethods(object types.Object) []types.Object {
	return tf.packageParser.StructMethods(object)
}

func (tf *TemplateFunc) FilterByMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if tf.HasMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (tf *TemplateFunc) HasMeta(metaName string, object types.Object) bool {
	metas := tf.metaParser.ObjectMetaGroups(object, metaName)
	return len(metas) > 0
}

func (tf *TemplateFunc) FilterByMethodContainsMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if tf.HasMethodContainsMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (tf *TemplateFunc) HasMethodContainsMeta(metaName string, object types.Object) bool {
	methods := tf.Methods(object)
	for _, method := range methods {
		if tf.HasMeta(metaName, method) {
			return true
		}
	}
	return false
}

func (tf *TemplateFunc) Declare(object types.Object) string {
	if len(object.Name()) > 0 {
		return object.Name() + " " + tf.Sign(object)
	} else {
		return tf.Sign(object)
	}
}

func (tf *TemplateFunc) Sign(object types.Object) string {
	switch tf.ObjectType(object) {
	case TypeInterfaceMethod, TypeStructMethod, TypeFunc:
		signature := tf.MethodSignature(object)
		builder := strings.Builder{}
		builder.Grow(32)
		builder.WriteRune('(')
		params := signature.Params()
		for i, l := 0, params.Len(); i < l; i++ {
			param := params.At(i)
			builder.WriteString(tf.Declare(param))
			builder.WriteRune(',')
		}
		builder.WriteRune(')')

		builder.WriteRune(' ')

		builder.WriteRune('(')
		results := signature.Results()
		for i, l := 0, results.Len(); i < l; i++ {
			result := results.At(i)
			builder.WriteString(tf.Declare(result))
			builder.WriteRune(',')
		}
		builder.WriteRune(')')

		return builder.String()
	case TypeVar, TypeFuncVar:
		return types.TypeString(object.Type(), tf.typeQualifier)
	}
	return ""
}

func (tf *TemplateFunc) MethodSignature(object types.Object) *types.Signature {
	return object.Type().(*types.Signature)
}

func (tf *TemplateFunc) filterByType(typ Type) []types.Object {
	scope := tf.pkg.Types.Scope()
	var result []types.Object
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		objectType := tf.packageParser.ObjectType(object)
		if objectType&typ > 0 {
			result = append(result, object)
		}
	}
	return result
}
