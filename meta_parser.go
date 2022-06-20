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

type MetaParser struct {
	packages                 *PackageParser
	pkg                      *packages.Package
	pkgPath                  string
	objectToParsedMetaGroups map[types.Object]map[string]Group
	metaNameToMeta           map[string]Meta
}

func NewMetaParser(packages *PackageParser, pkgPath string, metas []Meta) *MetaParser {
	metaNameToMeta := make(map[string]Meta, len(metas))
	for _, meta := range metas {
		metaNameToMeta[meta.Directive()] = meta
	}
	return &MetaParser{
		packages:                 packages,
		pkg:                      packages.Package(pkgPath),
		pkgPath:                  pkgPath,
		objectToParsedMetaGroups: map[types.Object]map[string]Group{},
		metaNameToMeta:           metaNameToMeta,
	}
}

func (metaParser *MetaParser) FindByMetaName(metaNames ...string) map[types.Object]map[string]Group {

	objectToMetas := map[types.Object]map[string]Group{}
	scope := metaParser.pkg.Types.Scope()
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		var objectMetas map[string]Group
		objectMetas = metaParser.ObjectMetaGroups(object, metaNames...)
		objectToMetas[object] = objectMetas
	}
	return objectToMetas
}

func (metaParser *MetaParser) ObjectMetaGroups(object types.Object, metaNames ...string) (
	parsedMetaGroups map[string]Group) {

	parsedMetaGroups = map[string]Group{}
	for _, metaName := range metaNames {
		parsedMetaGroup := metaParser.ObjectMetaGroup(object, metaName)
		if len(parsedMetaGroup) == 0 {
			continue
		}
		parsedMetaGroups[metaName] = parsedMetaGroup
	}
	return
}

func (metaParser *MetaParser) ObjectMetaGroup(object types.Object, metaName string) (parsedMetaGroup Group) {
	meta, ok := metaParser.metaNameToMeta[metaName]
	if !ok {
		return
	}

	objectType := metaParser.packages.ObjectType(object)
	if meta.Target()&objectType == 0 {
		return
	}

	metaNameToParsedMetaGroup, ok := metaParser.objectToParsedMetaGroups[object]
	if ok {
		parsedMetaGroup, ok = metaNameToParsedMetaGroup[metaName]
		if ok {
			return
		}
	}
	comments := metaParser.filterComments(object.Pos(), metaName)
	if len(comments) == 0 {
		return
	}

	for _, comment := range comments {
		parsedMeta := metaParser.populateMetaFields(meta, comment)
		parsedMetaGroup = append(parsedMetaGroup, parsedMeta)
	}
	if metaNameToParsedMetaGroup == nil {
		metaNameToParsedMetaGroup = map[string]Group{}
		metaParser.objectToParsedMetaGroups[object] = metaNameToParsedMetaGroup
	}
	metaNameToParsedMetaGroup[metaName] = parsedMetaGroup
	return
}

func (metaParser *MetaParser) filterComments(pos token.Pos, metaName string) []string {
	var filteredComments []string
	comments := metaParser.packages.Comments(pos)
	for _, comment := range comments {
		if strings.HasPrefix(comment, metaName) {
			filteredComments = append(filteredComments, comment)
		}
	}
	return filteredComments
}

//func (metaParser *MetaParser) populateMetaFields(meta Meta, comment string) (parsedMeta Meta) {
//	strings.TrimLeft(comment, meta.Directive())
//	fieldAndValues, err := shlex.Split(comment)
//	if err != nil {
//		return
//	}
//	typeOfMeta := reflect.Indirect(reflect.ValueOf(meta)).Type()
//	newMeta := reflect.New(typeOfMeta)
//	valueOfMeta := reflect.Indirect(newMeta)
//
//	//first is meta prefix, must ignore
//	for _, fieldAndValue := range fieldAndValues[1:] {
//		parts := strings.SplitN(fieldAndValue, "=", 2)
//
//		field, found := typeOfMeta.FieldByNameFunc(func(name string) bool {
//			fieldName := xstrings.ToCamelCase(parts[0])
//			return fieldName == name || strings.EqualFold(fieldName, name)
//		})
//		if !found {
//			panic(fmt.Errorf("meta parse: field not found,[meta=%s,field=%s]",
//				meta.Directive(), parts[0]))
//		}
//		if !field.IsExported() {
//			panic(fmt.Errorf("meta parse: expected exported field but not,[meta=%s,field=%s]",
//				meta.Directive(), parts[0]))
//		}
//		fieldValue := valueOfMeta.FieldByName(field.Name)
//
//		if len(parts) == 1 {
//			if field.Type.Kind() != reflect.Bool {
//				panic(fmt.Errorf("meta parse: expected bool field but not,[meta=%s,field=%s]",
//					meta.Directive(), parts[0]))
//			}
//			fieldValue.SetBool(true)
//		} else {
//			err = setValueFromString(field, fieldValue, parts[1])
//		}
//		if err != nil {
//			log.Fatalf("fail on set meta value, %v", err)
//		}
//	}
//	parsedMeta = newMeta.Interface().(Meta)
//	return
//}

func (metaParser *MetaParser) populateMetaFields(meta Meta, comment string) (parsedMeta Meta) {
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
			kv[parts[0]] = true
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