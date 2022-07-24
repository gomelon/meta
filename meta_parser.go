package meta

import (
	"encoding/json"
	"fmt"
	"github.com/google/shlex"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"reflect"
	"strings"
)

type Parser struct {
	packageParser            *PackageParser
	pkg                      *packages.Package
	pkgPath                  string
	objectToParsedMetaGroups map[types.Object]map[string]Group
	metaNameToMeta           map[string]Meta
}

func NewParser(packageParser *PackageParser, pkgPath string) *Parser {
	return &Parser{
		packageParser:            packageParser,
		pkg:                      packageParser.Package(pkgPath),
		pkgPath:                  pkgPath,
		objectToParsedMetaGroups: map[types.Object]map[string]Group{},
		metaNameToMeta:           make(map[string]Meta, 0),
	}
}

func (parser *Parser) AddMeta(metas ...Meta) {
	for _, meta := range metas {
		parser.metaNameToMeta[meta.Directive()] = meta
	}
}

func (parser *Parser) FindByMetaName(metaNames ...string) map[types.Object]map[string]Group {

	objectToMetas := map[types.Object]map[string]Group{}
	scope := parser.pkg.Types.Scope()
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		var objectMetas map[string]Group
		objectMetas = parser.ObjectMetaGroups(object, metaNames...)
		objectToMetas[object] = objectMetas
	}
	return objectToMetas
}

func (parser *Parser) ObjectMetaGroups(object types.Object, metaNames ...string) (
	parsedMetaGroups map[string]Group) {

	parsedMetaGroups = map[string]Group{}
	for _, metaName := range metaNames {
		parsedMetaGroup := parser.ObjectMetaGroup(object, metaName)
		if len(parsedMetaGroup) == 0 {
			continue
		}
		parsedMetaGroups[metaName] = parsedMetaGroup
	}
	return
}

func (parser *Parser) ObjectMetaGroup(object types.Object, metaName string) (parsedMetaGroup Group) {
	meta, ok := parser.metaNameToMeta[metaName]
	if !ok {
		return
	}

	objectPlace := parser.packageParser.ObjectPlace(object)
	if meta.PlaceAt()&objectPlace == 0 {
		return
	}

	metaNameToParsedMetaGroup, ok := parser.objectToParsedMetaGroups[object]
	if ok {
		parsedMetaGroup, ok = metaNameToParsedMetaGroup[metaName]
		if ok {
			return
		}
	}
	comments := parser.filterComments(object.Pos(), metaName)
	if len(comments) == 0 {
		return
	}

	for _, comment := range comments {
		parsedMeta := parser.populateMetaFields(meta, comment)
		parsedMetaGroup = append(parsedMetaGroup, parsedMeta)
	}
	if metaNameToParsedMetaGroup == nil {
		metaNameToParsedMetaGroup = map[string]Group{}
		parser.objectToParsedMetaGroups[object] = metaNameToParsedMetaGroup
	}
	metaNameToParsedMetaGroup[metaName] = parsedMetaGroup
	return
}

func (parser *Parser) filterComments(pos token.Pos, metaName string) []string {
	var filteredComments []string
	comments := parser.packageParser.Comments(pos)
	for _, comment := range comments {
		if strings.HasPrefix(comment, metaName) {
			filteredComments = append(filteredComments, comment)
		}
	}
	return filteredComments
}

func (parser *Parser) populateMetaFields(meta Meta, comment string) (parsedMeta Meta) {
	strings.TrimLeft(comment, meta.Directive())
	fieldAndValues, err := shlex.Split(comment)
	if err != nil {
		panic(fmt.Errorf("meta parse fail: %w", err))
	}
	typeOfMeta := reflect.Indirect(reflect.ValueOf(meta)).Type()
	newMeta := reflect.New(typeOfMeta)

	kv := make(map[string]any, typeOfMeta.NumField()*2)
	//first is meta prefix, must ignore
	for _, fieldAndValue := range fieldAndValues[1:] {
		parts := strings.SplitN(fieldAndValue, "=", 2)
		if len(parts) == 1 {
			kv[parts[0]] = "true"
		} else {
			kv[parts[0]] = parts[1]
		}
	}

	marshal, _ := json.Marshal(kv)
	err = json.Unmarshal(marshal, newMeta.Interface())
	if err != nil {
		panic(fmt.Errorf("meta parse fail: %w", err))
	}
	parsedMeta = newMeta.Interface().(Meta)
	return
}
