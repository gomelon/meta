package meta

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

type Place uint32

const (
	PlaceUnknown       = 0
	PlaceConst   Place = 1 << iota
	PlaceVar
	PlaceInterface
	PlaceStruct
	PlaceField
	PlaceInterfaceMethod
	PlaceStructMethod
	PlaceFunc
	//TODO seperate param and result
	PlaceFuncVar //func/method param or result, because can't distinguish on named result
)

type PackageParser struct {
	pkgPathToPkg         map[string]*packages.Package
	importedPkgPathToPkg map[string]*packages.Package
	posToComments        sync.Map // key=types.Object.Pos(),value=Comments []string
	pathToPkgPath        map[string]string
	pkgPathToPath        map[string]string
	objectToPlace        sync.Map // key=types.Object,value=Place
	fileSet              *token.FileSet
}

func NewPackageParser() *PackageParser {
	return &PackageParser{
		pkgPathToPkg:         map[string]*packages.Package{},
		importedPkgPathToPkg: map[string]*packages.Package{},
		pathToPkgPath:        map[string]string{},
		pkgPathToPath:        map[string]string{},
		fileSet:              token.NewFileSet(),
	}
}

//Load path, path may be relative path/absolute path/package path
func (pp *PackageParser) Load(paths ...string) (err error) {

	unloadedPaths := make([]string, 0, len(paths))
	for _, path := range paths {
		if pp.pkgPathToPath[path] != "" {
			continue
		}
		if pp.pathToPkgPath[path] != "" {
			continue
		}
		absPath, _ := filepath.Abs(path)
		if pp.pathToPkgPath[absPath] != "" {
			continue
		}
		unloadedPaths = append(unloadedPaths, path)
	}

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		//TODO 如果其它地方的代码有依赖生成的文件,但又不加入解析,是否有问题?
		BuildFlags: []string{"-tags", GeneratedBuildTag},
		Fset:       pp.fileSet,
		ParseFile: func(fileSet *token.FileSet, filename string, src []byte) (astFile *ast.File, err error) {
			mode := parser.ParseComments
			astFile, err = parser.ParseFile(fileSet, filename, src, mode)

			if err != nil {
				return
			}

			err = pp.doPosToComments(astFile)
			return
		},
		Tests: false,
	}

	packageList, err := packages.Load(cfg, unloadedPaths...)
	if err != nil {
		return
	}
	for _, pkg := range packageList {
		if len(pkg.GoFiles) == 0 {
			return fmt.Errorf("missing go file in %s", pkg.PkgPath)
		}
		pp.pkgPathToPkg[pkg.PkgPath] = pkg
		goFile := pkg.GoFiles[0]
		goFilePath := path.Dir(goFile)
		pp.pathToPkgPath[goFilePath] = pkg.PkgPath
		pp.pkgPathToPath[pkg.PkgPath] = goFilePath

		for pkgPath, importedPkg := range pkg.Imports {
			pp.importedPkgPathToPkg[pkgPath] = importedPkg
		}
	}
	return
}

func (pp *PackageParser) Package(pkgPath string) *packages.Package {
	pkg := pp.pkgPathToPkg[pkgPath]
	if pkg == nil {
		pkg = pp.importedPkgPathToPkg[pkgPath]
	}
	return pkg
}

func (pp *PackageParser) Path(pkgPath string) string {
	return pp.pkgPathToPath[pkgPath]
}

func (pp *PackageParser) PkgPath(path string) string {
	return pp.pathToPkgPath[path]
}

func (pp *PackageParser) ObjectByPkgPathAndName(pkgPath, typeName string) types.Object {
	err := pp.Load(pkgPath)
	if err != nil {
		panic(fmt.Errorf("can't load pakcage %s", pkgPath))
	}
	pkg := pp.Package(pkgPath)
	if pkg == nil {
		return nil
	}
	return pkg.Types.Scope().Lookup(typeName)
}

// AssignableToCtx reports whether a value of type V is assignable to a variable of type T.
// The behavior of AssignableTo is undefined if V or T is an uninstantiated generic type.
func (pp *PackageParser) AssignableToCtx(v types.Type) bool {
	ctxObject := pp.ObjectByPkgPathAndName("context", "Context")
	return pp.AssignableTo(v, ctxObject.Type())
}

func (pp *PackageParser) AssignableTo(v, t types.Type) bool {
	if pp.TypeName(t) == TypeNameNamed {
		t = t.Underlying()
	}
	namedV, ok := v.(*types.Named)
	if !ok {
		return types.AssignableTo(v, t)
	}
	namedVObj := namedV.Obj()
	namedVObjPkgPath := namedVObj.Pkg().Path()
	namedVObjName := namedVObj.Name()
	vObject := pp.ObjectByPkgPathAndName(namedVObjPkgPath, namedVObjName)
	return types.AssignableTo(vObject.Type(), t)
}

func (pp *PackageParser) Methods(object types.Object) []types.Object {
	switch pp.ObjectPlace(object) {
	case PlaceInterface:
		return pp.InterfaceMethods(object)
	case PlaceStruct:
		return pp.StructMethods(object)
	default:
		return []types.Object{}
	}
}

func (pp *PackageParser) InterfaceMethods(object types.Object) []types.Object {
	itf := object.Type().Underlying().(*types.Interface)
	numMethods := itf.NumMethods()
	methods := make([]types.Object, 0, numMethods)
	for i := 0; i < numMethods; i++ {
		methods = append(methods, itf.Method(i))
	}
	return methods
}

func (pp *PackageParser) StructMethods(object types.Object) []types.Object {
	namedObject := object.Type().(*types.Named)
	numMethods := namedObject.NumMethods()
	methods := make([]types.Object, 0, numMethods)
	for i := 0; i < numMethods; i++ {
		methods = append(methods, namedObject.Method(i))
	}
	return methods
}

func (pp *PackageParser) Params(methodOrFunc types.Object) []types.Object {
	signature, ok := methodOrFunc.Type().(*types.Signature)
	if !ok {
		panic(fmt.Errorf("package parser: object isn't a method [object=%s]", methodOrFunc.Name()))
	}
	params := signature.Params()
	length := params.Len()
	result := make([]types.Object, 0, length)

	for i := 0; i < length; i++ {
		result = append(result, params.At(i))
	}
	return result
}

func (pp *PackageParser) FirstParam(methodOrFunc types.Object) types.Object {
	params := pp.Params(methodOrFunc)
	if len(params) == 0 {
		return nil
	}
	return params[0]
}

func (pp *PackageParser) Results(methodOrFunc types.Object) []types.Object {
	signature, ok := methodOrFunc.Type().(*types.Signature)
	if !ok {
		panic(fmt.Errorf("package parser: object isn't a method [object=%s]", methodOrFunc.Name()))
	}
	params := signature.Results()
	length := params.Len()
	result := make([]types.Object, 0, length)

	for i := 0; i < length; i++ {
		result = append(result, params.At(i))
	}
	return result
}

func (pp *PackageParser) FirstResult(methodOrFunc types.Object) types.Object {
	results := pp.Results(methodOrFunc)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (pp *PackageParser) LastResult(methodOrFunc types.Object) types.Object {
	results := pp.Results(methodOrFunc)
	if len(results) == 0 {
		return nil
	}
	return results[len(results)-1]
}

func (pp *PackageParser) HasErrorResult(methodOrFunc types.Object) bool {
	lastResult := pp.LastResult(methodOrFunc)
	if lastResult == nil {
		return false
	}
	return lastResult.Type().String() == "error"
}

func (pp *PackageParser) Indirect(typ types.Type) types.Type {
	pointer := typ.(*types.Pointer)
	return pointer.Elem().(types.Type)
}

func (pp *PackageParser) ObjectPlace(object types.Object) (objectPlace Place) {
	objectPlaceValue, ok := pp.objectToPlace.Load(object)
	if ok {
		objectPlace = objectPlaceValue.(Place)
		return
	}
	switch object := object.(type) {
	case *types.Const:
		objectPlace = PlaceConst
	case *types.Var:
		if object.IsField() {
			objectPlace = PlaceField
		} else if parent := object.Parent(); parent != nil {
			if strings.HasPrefix(parent.String(), "package") {
				objectPlace = PlaceVar
			} else if strings.HasPrefix(parent.String(), "function") {
				objectPlace = PlaceFuncVar
			}
		} else {
			objectPlace = PlaceFuncVar
		}
	case *types.Func:
		signature := object.Type().(*types.Signature)
		receiver := signature.Recv()
		if receiver == nil {
			objectPlace = PlaceFunc
		} else {
			receiverPointer, ok := receiver.Type().(*types.Pointer)
			var receiverType types.Type
			if ok {
				receiverType = receiverPointer.Elem().Underlying()
			} else {
				receiverType = receiver.Type().Underlying()
			}
			switch receiverType.(type) {
			case *types.Struct:
				objectPlace = PlaceStructMethod
			case *types.Interface:
				objectPlace = PlaceInterfaceMethod
			default:
				objectPlace = PlaceUnknown
			}
		}
	case *types.TypeName:
		switch object.Type().Underlying().(type) {
		case *types.Interface:
			objectPlace = PlaceInterface
		case *types.Struct:
			objectPlace = PlaceStruct
		default:
			objectPlace = PlaceUnknown
		}
	default:
		objectPlace = PlaceUnknown
	}
	pp.objectToPlace.Store(object, objectPlace)
	return
}

const (
	TypeNamePointer   = "Pointer"
	TypeNameStruct    = "Struct"
	TypeNameInterface = "Interface"
	TypeNameSignature = "Signature"
	TypeNameBasic     = "Basic"
	TypeNameSlice     = "Slice"
	TypeNameMap       = "Map"
	TypeNameNamed     = "Named"
	TypeNameTuple     = "Tuple" //see types.Tuple
	TypeNameArray     = "Array"
	TypeNameChan      = "Chan"
)

func (pp *PackageParser) TypeName(typ types.Type) string {
	var typName string
	switch typ := typ.(type) {
	case *types.Pointer:
		typName = TypeNamePointer
	case *types.Struct:
		typName = TypeNameStruct
	case *types.Interface:
		typName = TypeNameInterface
	case *types.Signature:
		typName = TypeNameSignature
	case *types.Basic:
		typName = TypeNameBasic
	case *types.Slice:
		typName = TypeNameSlice
	case *types.Map:
		typName = TypeNameMap
	case *types.Named:
		typName = TypeNameNamed
	case *types.Tuple:
		typName = TypeNameTuple
	case *types.Array:
		typName = TypeNameArray
	case *types.Chan:
		typName = TypeNameChan
	default:
		typName = typ.String()
	}
	return typName
}

func (pp *PackageParser) UnderlyingType(typ types.Type) types.Type {
	switch typ := typ.(type) {
	case *types.Basic, *types.Struct, *types.Interface, *types.Chan, *types.Signature:
		return typ
	case *types.Named:
		return pp.UnderlyingType(typ.Underlying())
	case *types.Pointer:
		return pp.UnderlyingType(typ.Elem())
	case *types.Slice:
		return pp.UnderlyingType(typ.Elem())
	case *types.Map:
		return pp.UnderlyingType(typ.Elem())
	default:
		panic(fmt.Errorf("msql fail: unsupported result type from query"))
	}
}

func (pp *PackageParser) Comments(pos token.Pos) []string {
	value, ok := pp.posToComments.Load(pos)
	if !ok {
		return []string{}
	}
	return value.([]string)
}

func (pp *PackageParser) doPosToComments(astFile *ast.File) error {
	commentMap := ast.NewCommentMap(pp.fileSet, astFile, astFile.Comments)
	for astNode, commentGroups := range commentMap {
		var nodeIdentPos token.Pos
		switch node := astNode.(type) {
		case *ast.ValueSpec:
			nodeIdentPos = node.Names[0].Pos()
		case *ast.Field:
			nodeIdentPos = node.Names[0].Pos()
		case *ast.FuncDecl:
			nodeIdentPos = node.Name.Pos()
		case *ast.GenDecl:
			switch node := node.Specs[0].(type) {
			case *ast.TypeSpec:
				nodeIdentPos = node.Name.Pos()
			case *ast.ValueSpec:
				nodeIdentPos = node.Names[0].Pos()
			default:
				//position := pp.fileSet.Position(node.Pos())
				//fmt.Printf("file=%v,line=%v,column=%v", position.Filename, position.Line, position.Column)
				//return fmt.Errorf("parse package: don't support parse comment for [%#v]", node)
			}
		case *ast.File:
			nodeIdentPos = node.Name.Pos()
		case *ast.ReturnStmt, *ast.DeclStmt, *ast.BranchStmt, *ast.AssignStmt,
			*ast.IfStmt, *ast.ForStmt, *ast.Ident, *ast.ImportSpec, *ast.RangeStmt:
			continue
		default:
			//position := pp.fileSet.Position(node.Pos())
			//fmt.Printf("file=%v,line=%v,column=%v", position.Filename, position.Line, position.Column)
			//return fmt.Errorf("parse package: don't support parse comment for [%#v]", node)
		}
		commentLines := convertCommentGroupsToStrings(commentGroups)
		pp.posToComments.Store(nodeIdentPos, commentLines)
	}
	return nil
}

func convertCommentGroupsToStrings(commentGroups []*ast.CommentGroup) []string {
	var commentLines []string
	for _, commentGroup := range commentGroups {
		comments := commentGroup.List
		for _, comment := range comments {
			var commentLine string
			if strings.HasPrefix(comment.Text, "//") {
				commentLine = strings.TrimSpace(strings.TrimLeft(comment.Text, "//"))
			} else {
				singleCommentGroup := &ast.CommentGroup{List: []*ast.Comment{comment}}
				commentLine = strings.TrimSpace(singleCommentGroup.Text())
			}
			commentLines = append(commentLines, commentLine)
		}
	}
	return commentLines
}
