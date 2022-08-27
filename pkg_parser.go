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

var defaultPkgParser = NewPkgParser()

type PkgParser struct {
	pkgPathToPkg         map[string]*packages.Package
	importedPkgPathToPkg map[string]*packages.Package
	posToComments        sync.Map // key=types.Object.Pos(),value=Comments []string
	pathToPkgPath        map[string]string
	pkgPathToPath        map[string]string
	objectToPlace        sync.Map // key=types.Object,value=Place
	fileSet              *token.FileSet
	filenameToAstFile    sync.Map //key=string,value=*ast.File
	fileParsingLock      sync.Map //key=string,value=*sync.Mutex
	anonymousAssign      map[string]map[types.Type]bool
	anonymousAssignTo    map[string]map[types.Type]bool
}

func NewPkgParser() *PkgParser {
	return &PkgParser{
		pkgPathToPkg:         map[string]*packages.Package{},
		importedPkgPathToPkg: map[string]*packages.Package{},
		pathToPkgPath:        map[string]string{},
		pkgPathToPath:        map[string]string{},
		fileSet:              token.NewFileSet(),
		anonymousAssign:      make(map[string]map[types.Type]bool, 64),
		anonymousAssignTo:    make(map[string]map[types.Type]bool, 64),
	}
}

//Load path, path may be relative path/absolute path/package path
func (pp *PkgParser) Load(paths ...string) (err error) {

	unloadedPaths := make([]string, 0, len(paths))
	for _, inputPath := range paths {
		if pp.pkgPathToPath[inputPath] != "" {
			continue
		}
		if pp.pathToPkgPath[inputPath] != "" {
			continue
		}
		absPath, _ := filepath.Abs(inputPath)
		if pp.pathToPkgPath[absPath] != "" {
			continue
		}
		unloadedPaths = append(unloadedPaths, inputPath)
	}

	if len(unloadedPaths) == 0 {
		return nil
	}

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles | packages.NeedImports |
			packages.NeedTypes | packages.NeedSyntax | packages.NeedTypesInfo,
		//BuildFlags: []string{"-tags", GeneratedBuildTag},
		Fset:      pp.fileSet,
		ParseFile: pp.parseFile,
		Tests:     false,
	}

	packageList, err := packages.Load(cfg, unloadedPaths...)
	if err != nil {
		return
	}
	for _, pkg := range packageList {
		if len(pkg.CompiledGoFiles) == 0 {
			continue
		}
		pp.pkgPathToPkg[pkg.PkgPath] = pkg
		goFile := pkg.GoFiles[0]
		goFilePath := path.Dir(goFile)
		pp.pathToPkgPath[goFilePath] = pkg.PkgPath
		pp.pkgPathToPath[pkg.PkgPath] = goFilePath

		for pkgPath, importedPkg := range pkg.Imports {
			pp.importedPkgPathToPkg[pkgPath] = importedPkg
		}

		pp.parseAnonymousAssign(pkg)
	}
	return
}

func (pp *PkgParser) Package(pkgPath string) *packages.Package {
	pkg := pp.pkgPathToPkg[pkgPath]
	if pkg == nil {
		pkg = pp.importedPkgPathToPkg[pkgPath]
	}
	return pkg
}

func (pp *PkgParser) Path(pkgPath string) string {
	return pp.pkgPathToPath[pkgPath]
}

func (pp *PkgParser) PkgPath(path string) string {
	return pp.pathToPkgPath[path]
}

func (pp *PkgParser) Structs(pkgPath string) []types.Object {
	return pp.filterByPlace(pkgPath, PlaceStruct)
}

func (pp *PkgParser) Interfaces(pkgPath string) []types.Object {
	return pp.filterByPlace(pkgPath, PlaceInterface)
}

func (pp *PkgParser) Functions(pkgPath string) []types.Object {
	return pp.filterByPlace(pkgPath, PlaceFunc)
}

func (pp *PkgParser) Methods(object types.Object) []types.Object {
	switch pp.ObjectPlace(object) {
	case PlaceInterface:
		return pp.InterfaceMethods(object)
	case PlaceStruct:
		return pp.StructMethods(object)
	default:
		return []types.Object{}
	}
}

func (pp *PkgParser) InterfaceMethods(object types.Object) []types.Object {
	iface := object.Type().Underlying().(*types.Interface)
	numMethods := iface.NumMethods()
	methods := make([]types.Object, 0, numMethods)
	for i := 0; i < numMethods; i++ {
		methods = append(methods, iface.Method(i))
	}
	return methods
}

func (pp *PkgParser) StructMethods(object types.Object) []types.Object {
	namedObject := object.Type().(*types.Named)
	numMethods := namedObject.NumMethods()
	methods := make([]types.Object, 0, numMethods)
	for i := 0; i < numMethods; i++ {
		methods = append(methods, namedObject.Method(i))
	}
	return methods
}

func (pp *PkgParser) Params(methodOrFunc types.Object) []types.Object {
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

func (pp *PkgParser) FirstParam(methodOrFunc types.Object) types.Object {
	params := pp.Params(methodOrFunc)
	if len(params) == 0 {
		return nil
	}
	return params[0]
}

func (pp *PkgParser) Results(methodOrFunc types.Object) []types.Object {
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

func (pp *PkgParser) FirstResult(methodOrFunc types.Object) types.Object {
	results := pp.Results(methodOrFunc)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (pp *PkgParser) LastResult(methodOrFunc types.Object) types.Object {
	results := pp.Results(methodOrFunc)
	if len(results) == 0 {
		return nil
	}
	return results[len(results)-1]
}

func (pp *PkgParser) HasErrorResult(methodOrFunc types.Object) bool {
	lastResult := pp.LastResult(methodOrFunc)
	if lastResult == nil {
		return false
	}
	return lastResult.Type().String() == "error"
}

func (pp *PkgParser) Object(pkgPath, typeName string) types.Object {
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

func (pp *PkgParser) ObjectComments(object types.Object) []string {
	pos := object.Pos()
	return pp.Comments(pos)
}

// AssignableToCtx reports whether a value of type V is assignable to context.Context.
// The behavior of AssignableTo is undefined if V or T is an uninstantiated generic type.
func (pp *PkgParser) AssignableToCtx(v types.Type) bool {
	if v.String() == "context.Context" {
		return true
	}
	ctxObject := pp.Object("context", "Context")
	return pp.AssignableTo(v, ctxObject.Type())
}

// AssignableTo reports whether a value of type V is assignable to a variable of type T.
// The behavior of AssignableTo is undefined if V or T is an uninstantiated generic type.
func (pp *PkgParser) AssignableTo(v, t types.Type) bool {
	if pp.TypeName(t) == TypeNameNamed {
		t = t.Underlying()
	}
	namedV, ok := v.(*types.Named)
	if !ok {
		return v.String() == t.String() || types.AssignableTo(v, t)
	}
	namedVObj := namedV.Obj()
	namedVObjPkgPath := namedVObj.Pkg().Path()
	namedVObjName := namedVObj.Name()
	vObject := pp.Object(namedVObjPkgPath, namedVObjName)
	return v.String() == t.String() || types.AssignableTo(vObject.Type(), t)
}

//AnonymousAssign return anonymous assign some value to t
//var _ Foo = FooImpl{} Foo is the t,FooImpl is the result
func (pp *PkgParser) AnonymousAssign(t types.Type) (values []types.Type) {
	assigns := pp.anonymousAssign[t.String()]
	for k, _ := range assigns {
		values = append(values, k)
	}
	return
}

//AnonymousAssignTo return anonymous assign some value to t
//var _ Foo = FooImpl{} Foo is the result,FooImpl is the v
func (pp *PkgParser) AnonymousAssignTo(v types.Type) (types []types.Type) {
	assigns := pp.anonymousAssignTo[v.String()]
	for k, _ := range assigns {
		types = append(types, k)
	}
	return
}

func (pp *PkgParser) Indirect(typ types.Type) types.Type {
	pointer := typ.(*types.Pointer)
	return pointer.Elem().(types.Type)
}

func (pp *PkgParser) ObjectPlace(object types.Object) (objectPlace Place) {
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

func (pp *PkgParser) TypeName(typ types.Type) string {
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

func (pp *PkgParser) UnderlyingType(typ types.Type) types.Type {
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

func (pp *PkgParser) Comments(pos token.Pos) []string {
	value, ok := pp.posToComments.Load(pos)
	if !ok {
		return []string{}
	}
	return value.([]string)
}

func (pp *PkgParser) parseFile(fileSet *token.FileSet, filename string, src []byte) (astFile *ast.File, err error) {
	lockerObj, _ := pp.fileParsingLock.LoadOrStore(filename, &sync.Mutex{})
	defer pp.fileParsingLock.Delete(filename)
	locker := lockerObj.(*sync.Mutex)
	locked := locker.TryLock()
	defer locker.Unlock()
	if !locked {
		locker.Lock()
		return
	}
	astFileObj, ok := pp.filenameToAstFile.Load(filename)
	if ok {
		astFile = astFileObj.(*ast.File)
		return
	}

	mode := parser.ParseComments
	astFile, err = parser.ParseFile(fileSet, filename, src, mode)

	if err != nil {
		return
	}

	err = pp.doPosToComments(astFile)

	pp.filenameToAstFile.Store(filename, astFile)
	return
}

func (pp *PkgParser) parseAnonymousAssign(pkg *packages.Package) {

	for _, syntax := range pkg.Syntax {
		ast.Inspect(syntax, func(node ast.Node) bool {
			switch node := node.(type) {
			case *ast.ValueSpec:
				if len(node.Names) == 0 || node.Names[0].Name != "_" {
					return false
				}
				typesInfo := pkg.TypesInfo
				typ := typesInfo.TypeOf(node.Type)
				value := typesInfo.TypeOf(node.Values[0])

				typeStr := typ.String()
				assigns := pp.anonymousAssign[typeStr]
				if assigns == nil {
					assigns = make(map[types.Type]bool, 4)
					pp.anonymousAssign[typeStr] = assigns
				}
				assigns[value] = true

				valueStr := value.String()
				assignTos := pp.anonymousAssignTo[valueStr]
				if assignTos == nil {
					assignTos = make(map[types.Type]bool, 4)
					pp.anonymousAssignTo[valueStr] = assignTos
				}
				assignTos[typ] = true
				return false
			case *ast.File, *ast.GenDecl:
				return true
			default:
				return false
			}
		})
	}
}

func (pp *PkgParser) doPosToComments(astFile *ast.File) error {
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
				continue
			}
		case *ast.File:
			nodeIdentPos = node.Name.Pos()
		case *ast.ReturnStmt, *ast.DeclStmt, *ast.BranchStmt, *ast.AssignStmt,
			*ast.IfStmt, *ast.ForStmt, *ast.Ident, *ast.ImportSpec, *ast.RangeStmt:
			continue
		default:
			continue
		}
		commentLines := convertCommentGroupsToStrings(commentGroups)
		pp.posToComments.Store(nodeIdentPos, commentLines)
	}
	return nil
}

func (pp *PkgParser) filterByPlace(pkgPath string, place Place) []types.Object {
	pkg := pp.Package(pkgPath)
	scope := pkg.Types.Scope()
	var result []types.Object
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		objectPlace := pp.ObjectPlace(object)
		if objectPlace&place > 0 {
			result = append(result, object)
		}
	}
	return result
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
				commentLine = strings.ReplaceAll(commentLine, "\n\t", "\n")
			}
			commentLines = append(commentLines, commentLine)
		}
	}
	return commentLines
}
