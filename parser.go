package meta

import (
	"fmt"
	"github.com/google/shlex"
	"go/token"
	"go/types"
	"strings"
)

var defaultParser = NewParser(defaultPkgParser)

type Parser struct {
	pkgParser                *PkgParser
	objectToParsedMetaGroups map[types.Object]map[string]Group
}

func NewParser(pkgParser *PkgParser) *Parser {
	return &Parser{
		pkgParser:                pkgParser,
		objectToParsedMetaGroups: map[types.Object]map[string]Group{},
	}
}

func (parser *Parser) FilterByMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if parser.HasMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (parser *Parser) HasMeta(metaName string, object types.Object) bool {
	metas := parser.ObjectMetaGroups(object, metaName)
	return len(metas) > 0
}

func (parser *Parser) FilterByMethodContainsMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if parser.HasMethodContainsMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (parser *Parser) HasMethodContainsMeta(metaName string, object types.Object) bool {
	methods := parser.pkgParser.Methods(object)
	for _, method := range methods {
		if parser.HasMeta(metaName, method) {
			return true
		}
	}
	return false
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
		parsedMeta := parser.populateMetaFields(metaName, comment)
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
	comments := parser.pkgParser.Comments(pos)
	for _, comment := range comments {
		if strings.HasPrefix(comment, metaName) {
			filteredComments = append(filteredComments, comment)
		}
	}
	return filteredComments
}

func (parser *Parser) populateMetaFields(metaName, comment string) (parsedMeta *Meta) {
	propertiesStr := strings.TrimLeft(comment, metaName)
	fieldAndValues, err := shlex.Split(propertiesStr)
	if err != nil {
		panic(fmt.Errorf("meta parse fail: %w", err))
	}

	properties := make(map[string]string, len(fieldAndValues))
	for _, fieldAndValue := range fieldAndValues {
		parts := strings.SplitN(fieldAndValue, "=", 2)
		if len(parts) == 1 {
			properties[fieldAndValue] = fieldAndValue
		} else {
			properties[parts[0]] = parts[1]
		}
	}

	meta := New(metaName)
	meta.SetProperties(properties)
	return meta
}
