package meta

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path"
	"strings"
	"sync"
)

type PackageParser struct {
	pkgPathToPackage map[string]*packages.Package
	posToComments    sync.Map // key=types.Object.Pos(),value=Comments []string
	pathToPkgPath    map[string]string
	pkgPathToPath    map[string]string
	objectToType     sync.Map // key=types.Object,value=Type
	fileSet          *token.FileSet
}

func NewPackageParser() *PackageParser {
	return &PackageParser{
		pkgPathToPackage: map[string]*packages.Package{},
		pathToPkgPath:    map[string]string{},
		pkgPathToPath:    map[string]string{},
		fileSet:          token.NewFileSet(),
	}
}

func (packageParser *PackageParser) Load(paths ...string) (err error) {

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		//TODO 如果其它地方的代码有依赖生成的文件,但又不加入解析,是否有问题?
		BuildFlags: []string{"-tags", GeneratedBuildTag},
		Fset:       packageParser.fileSet,
		ParseFile: func(fileSet *token.FileSet, filename string, src []byte) (astFile *ast.File, err error) {
			mode := parser.ParseComments
			astFile, err = parser.ParseFile(fileSet, filename, src, mode)

			if err != nil {
				return
			}

			err = packageParser.doPosToComments(astFile)
			return
		},
		Tests: false,
	}

	packageList, err := packages.Load(cfg, paths...)
	if err != nil {
		return
	}
	for _, pkg := range packageList {
		if len(pkg.GoFiles) == 0 {
			return fmt.Errorf("missing go file in %s", pkg.PkgPath)
		}
		packageParser.pkgPathToPackage[pkg.PkgPath] = pkg
		goFile := pkg.GoFiles[0]
		goFilePath := path.Dir(goFile)
		packageParser.pathToPkgPath[goFilePath] = pkg.PkgPath
		packageParser.pkgPathToPath[pkg.PkgPath] = goFilePath
	}
	return
}

func (packageParser *PackageParser) Package(packagePath string) *packages.Package {
	return packageParser.pkgPathToPackage[packagePath]
}

func (packageParser *PackageParser) Path(pkgPath string) string {
	return packageParser.pkgPathToPath[pkgPath]
}

func (packageParser *PackageParser) PkgPath(path string) string {
	return packageParser.pathToPkgPath[path]
}

func (packageParser *PackageParser) TypeByPkgPathAndName(packagePath, typeName string) types.Object {
	pkg := packageParser.Package(packagePath)
	if pkg == nil {
		return nil
	}
	return pkg.Types.Scope().Lookup(typeName)
}

func (packageParser *PackageParser) Methods(object types.Object) []types.Object {
	switch packageParser.ObjectType(object) {
	case TypeInterface:
		return packageParser.InterfaceMethods(object)
	case TypeStruct:
		return packageParser.StructMethods(object)
	default:
		return []types.Object{}
	}
}

func (packageParser *PackageParser) InterfaceMethods(object types.Object) []types.Object {
	itf := object.Type().Underlying().(*types.Interface)
	numMethods := itf.NumMethods()
	methods := make([]types.Object, 0, numMethods)
	for i := 0; i < numMethods; i++ {
		methods = append(methods, itf.Method(i))
	}
	return methods
}

func (packageParser *PackageParser) StructMethods(object types.Object) []types.Object {
	namedObject := object.Type().(*types.Named)
	numMethods := namedObject.NumMethods()
	methods := make([]types.Object, 0, numMethods)
	for i := 0; i < numMethods; i++ {
		methods = append(methods, namedObject.Method(i))
	}
	return methods
}

func (packageParser *PackageParser) Params(methodOrFunc types.Object) []types.Object {
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

func (packageParser *PackageParser) FirstParam(methodOrFunc types.Object) types.Object {
	params := packageParser.Params(methodOrFunc)
	if len(params) == 0 {
		return nil
	}
	return params[0]
}

func (packageParser *PackageParser) Results(methodOrFunc types.Object) []types.Object {
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

func (packageParser *PackageParser) FirstResult(methodOrFunc types.Object) types.Object {
	results := packageParser.Results(methodOrFunc)
	if len(results) == 0 {
		return nil
	}
	return results[0]
}

func (packageParser *PackageParser) LastResult(methodOrFunc types.Object) types.Object {
	results := packageParser.Results(methodOrFunc)
	if len(results) == 0 {
		return nil
	}
	return results[len(results)-1]
}

func (packageParser *PackageParser) HasErrorResult(methodOrFunc types.Object) bool {
	lastResult := packageParser.LastResult(methodOrFunc)
	if lastResult == nil {
		return false
	}
	return lastResult.Type().String() == "error"
}

func (packageParser *PackageParser) IndirectObject(object types.Object) types.Type {
	pointer := object.Type().(*types.Pointer)
	return pointer.Elem().(types.Type)
}

func (packageParser *PackageParser) ObjectType(object types.Object) (objectType Type) {
	objectTypeValue, ok := packageParser.objectToType.Load(object)
	if ok {
		objectType = objectTypeValue.(Type)
		return
	}
	switch object := object.(type) {
	case *types.Const:
		objectType = TypeConst
	case *types.Var:
		if object.IsField() {
			objectType = TypeField
		} else if parent := object.Parent(); parent != nil {
			if strings.HasPrefix(parent.String(), "package") {
				objectType = TypeVar
			} else if strings.HasPrefix(parent.String(), "function") {
				objectType = TypeFuncVar
			}
		} else {
			objectType = TypeFuncVar
		}
	case *types.Func:
		signature := object.Type().(*types.Signature)
		receiver := signature.Recv()
		if receiver == nil {
			objectType = TypeFunc
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
				objectType = TypeStructMethod
			case *types.Interface:
				objectType = TypeInterfaceMethod
			default:
				objectType = TypeUnknown
			}
		}
	case *types.TypeName:
		switch object.Type().Underlying().(type) {
		case *types.Interface:
			objectType = TypeInterface
		case *types.Struct:
			objectType = TypeStruct
		default:
			objectType = TypeUnknown
		}
	default:
		objectType = TypeUnknown
	}
	packageParser.objectToType.Store(object, objectType)
	return
}

func (packageParser *PackageParser) Comments(pos token.Pos) []string {
	value, ok := packageParser.posToComments.Load(pos)
	if !ok {
		return []string{}
	}
	return value.([]string)
}

func (packageParser *PackageParser) doPosToComments(astFile *ast.File) error {
	commentMap := ast.NewCommentMap(packageParser.fileSet, astFile, astFile.Comments)
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
				position := packageParser.fileSet.Position(node.Pos())
				fmt.Printf("file=%v,line=%v,column=%v", position.Filename, position.Line, position.Column)
				return fmt.Errorf("parse packageParser: don't support parse comment for [%#v]", node)
			}
		case *ast.File:
			nodeIdentPos = node.Name.Pos()
		case *ast.ReturnStmt, *ast.DeclStmt, *ast.BranchStmt, *ast.AssignStmt,
			*ast.IfStmt, *ast.ForStmt, *ast.Ident, *ast.ImportSpec, *ast.RangeStmt:
			continue
		default:
			position := packageParser.fileSet.Position(node.Pos())
			fmt.Printf("file=%v,line=%v,column=%v", position.Filename, position.Line, position.Column)
			return fmt.Errorf("parse packageParser: don't support parse comment for [%#v]", node)
		}
		commentLines := convertCommentGroupsToStrings(commentGroups)
		packageParser.posToComments.Store(nodeIdentPos, commentLines)
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
