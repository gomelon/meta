package meta

import (
	"fmt"
	"github.com/google/shlex"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"path"
	"reflect"
	"strings"
	"sync"
)

type Packages struct {
	pkgPathToPackage map[string]*packages.Package
	posToComments    sync.Map // key=types.Object.Pos(),value=comments []string
	pathToPkgPath    map[string]string
	fileSet          *token.FileSet
	workdir          string
}

func NewPackages(workdir string) *Packages {
	return &Packages{
		pkgPathToPackage: map[string]*packages.Package{},
		pathToPkgPath:    map[string]string{},
		fileSet:          token.NewFileSet(),
		workdir:          workdir,
	}
}

func (p *Packages) Load(paths ...string) (err error) {

	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedCompiledGoFiles |
			packages.NeedImports | packages.NeedTypes | packages.NeedSyntax,
		Dir:  p.workdir,
		Fset: p.fileSet,
		ParseFile: func(fileSet *token.FileSet, filename string, src []byte) (astFile *ast.File, err error) {
			mode := parser.ParseComments
			astFile, err = parser.ParseFile(fileSet, filename, src, mode)

			if err != nil {
				return
			}

			err = p.doPosToComments(astFile)
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
		p.pkgPathToPackage[pkg.PkgPath] = pkg
		goFile := pkg.GoFiles[0]
		goFilePath := path.Dir(goFile)
		p.pathToPkgPath[goFilePath] = pkg.PkgPath
	}
	return
}

func (p *Packages) FindByComment(commentPrefix string) []types.Object {
	var objects []types.Object
	for _, pkg := range p.pkgPathToPackage {
		scope := pkg.Types.Scope()
		for _, typeName := range scope.Names() {
			object := scope.Lookup(typeName)
			commentLines := p.filterComments(object.Pos(), commentPrefix)
			if len(commentLines) == 0 {
				continue
			}
			objects = append(objects, object)
		}
	}
	return objects
}

func (p *Packages) FindType(packagePath, typeName string) types.Object {
	pkg := p.Package(packagePath)
	if pkg == nil {
		return nil
	}
	return pkg.Types.Scope().Lookup(typeName)
}

func (p *Packages) Package(packagePath string) *packages.Package {
	return p.pkgPathToPackage[packagePath]
}

func (p *Packages) FindByMeta(metas ...Meta) (inputs []*Input, err error) {
	inputs = []*Input{}
	for _, pkg := range p.pkgPathToPackage {
		var input *Input
		input, err = p.findByMetaForPkg(pkg, metas)
		if err != nil {
			return
		}
		if input.IsEmpty() {
			continue
		}
		inputs = append(inputs, input)
	}
	return
}

func (p *Packages) findByMetaForPkg(pkg *packages.Package, metas Metas) (input *Input, err error) {
	scope := pkg.Types.Scope()
	input = NewInput(pkg, pkg.PkgPath)
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		var parsedMetas map[string][]Meta
		parsedMetas, err = p.parseMetas(object, metas)
		if err != nil {
			return
		}
		if len(parsedMetas) == 0 {
			continue
		}
		switch UnderlyingObj := object.Type().Underlying().(type) {
		case *types.Interface:
			input.Interfaces = append(input.Interfaces, &Interface{
				Interface: UnderlyingObj,
				Object:    object,
				Methods:   nil,
				Comments:  nil,
				Metas:     parsedMetas,
			})
		}
	}
	return
}

func (p *Packages) parseMetas(object types.Object, metas Metas) (parsedMetas map[string][]Meta, err error) {
	parsedMetas = map[string][]Meta{}
	for _, meta := range metas {
		var oneParsedMetas []Meta
		oneParsedMetas, err = p.parseMeta(object, meta)
		if err != nil {
			return
		}
		parsedMetas[meta.Name()] = append(parsedMetas[meta.Name()], oneParsedMetas...)
	}
	return
}

func (p *Packages) parseMeta(object types.Object, meta Meta) (parsedMetas []Meta, err error) {
	objectTarget := ObjectTarget(object)
	if meta.Target()&objectTarget == 0 {
		return
	}
	comments := p.filterComments(object.Pos(), meta.Name())
	if len(comments) == 0 {
		return
	}
	parsedMetas = []Meta{}
	var parsedMeta Meta
	for _, comment := range comments {
		parsedMeta, err = p.populateMetaFields(meta, comment)
		if err != nil {
			return
		}
		parsedMetas = append(parsedMetas, parsedMeta)
	}
	return
}

func (p *Packages) populateMetaFields(meta Meta, comment string) (parsedMeta Meta, err error) {
	strings.TrimLeft(comment, meta.Name())
	fieldAndValues, err := shlex.Split(comment)
	if err != nil {
		return
	}
	typeOfMeta := reflect.Indirect(reflect.ValueOf(meta)).Type()
	newMeta := reflect.New(typeOfMeta)
	valueOfMeta := reflect.Indirect(newMeta)

	//TODO 后面再来优化让元数据可以写得更简洁,当前实现默认都是A=B格式
	//first is meta prefix, must ignore
	for _, fieldAndValue := range fieldAndValues[1:] {
		//TODO 这里后面再处理非法格式吧
		parts := strings.SplitN(fieldAndValue, "=", 2)
		fieldValue := valueOfMeta.FieldByName(parts[0])
		err = SetValueFromString(fieldValue, parts[1])
		if err != nil {
			return
		}
	}
	parsedMeta = newMeta.Interface().(Meta)
	return
}

func (p *Packages) comments(pos token.Pos) []string {
	value, ok := p.posToComments.Load(pos)
	if !ok {
		return []string{}
	}
	return value.([]string)
}

func (p *Packages) filterComments(pos token.Pos, commentPrefix string) []string {
	var filteredComments []string
	comments := p.comments(pos)
	for _, comment := range comments {
		if strings.HasPrefix(comment, commentPrefix) {
			filteredComments = append(filteredComments, comment)
		}
	}
	return filteredComments
}

func (p *Packages) doPosToComments(astFile *ast.File) error {
	commentMap := ast.NewCommentMap(p.fileSet, astFile, astFile.Comments)
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
				position := p.fileSet.Position(node.Pos())
				log.Printf("file=%v,line=%v,column=%v", position.Filename, position.Line, position.Column)
				return fmt.Errorf("parse packages: don't support parse comment for [%#v]", node)
			}
		case *ast.File:
			nodeIdentPos = node.Name.Pos()
		case *ast.ReturnStmt, *ast.DeclStmt, *ast.BranchStmt, *ast.AssignStmt,
			*ast.IfStmt, *ast.ForStmt, *ast.Ident, *ast.ImportSpec, *ast.RangeStmt:
			continue
		default:
			position := p.fileSet.Position(node.Pos())
			log.Printf("file=%v,line=%v,column=%v", position.Filename, position.Line, position.Column)
			return fmt.Errorf("parse packages: don't support parse comment for [%#v]", node)
		}
		commentLines := convertCommentGroupsToStrings(commentGroups)
		p.posToComments.Store(nodeIdentPos, commentLines)
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

func KindOfObject(object types.Object) reflect.Kind {
	underlying := object.Type().Underlying()
	underlyingStr := underlying.String()
	switch {
	case strings.HasPrefix(underlyingStr, "interface{"):
		return reflect.Interface
	default:
		return reflect.Invalid
	}
}
