package meta

import (
	"github.com/google/shlex"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"reflect"
	"strings"
)

type MetaParser struct {
	packages                 *PackageParser
	pkg                      *packages.Package
	pkgPath                  string
	objectToParsedMetaGroups map[types.Object]map[string]MetaGroup
	metaNameToMeta           map[string]Meta
}

func NewMetaParser(packages *PackageParser, pkgPath string, metas []Meta) *MetaParser {
	metaNameToMeta := make(map[string]Meta, len(metas))
	for _, meta := range metas {
		metaNameToMeta[meta.Name()] = meta
	}
	return &MetaParser{
		packages:                 packages,
		pkg:                      packages.Package(pkgPath),
		pkgPath:                  pkgPath,
		objectToParsedMetaGroups: map[types.Object]map[string]MetaGroup{},
		metaNameToMeta:           metaNameToMeta,
	}
}

func (metaParser *MetaParser) FindByMetaName(metaNames ...string) map[types.Object]map[string]MetaGroup {

	objectToMetas := map[types.Object]map[string]MetaGroup{}
	scope := metaParser.pkg.Types.Scope()
	for _, typeName := range scope.Names() {
		object := scope.Lookup(typeName)
		var objectMetas map[string]MetaGroup
		objectMetas = metaParser.ObjectMetaGroups(object, metaNames...)
		objectToMetas[object] = objectMetas
	}
	return objectToMetas
}

func (metaParser *MetaParser) ObjectMetaGroups(object types.Object, metaNames ...string) (
	parsedMetaGroups map[string]MetaGroup) {

	parsedMetaGroups = map[string]MetaGroup{}
	for _, metaName := range metaNames {
		parsedMetaGroup := metaParser.ObjectMetaGroup(object, metaName)
		if len(parsedMetaGroup) == 0 {
			continue
		}
		parsedMetaGroups[metaName] = parsedMetaGroup
	}
	return
}

func (metaParser *MetaParser) ObjectMetaGroup(object types.Object, metaName string) (parsedMetaGroup MetaGroup) {
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
		metaNameToParsedMetaGroup = map[string]MetaGroup{}
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

func (metaParser *MetaParser) populateMetaFields(meta Meta, comment string) (parsedMeta Meta) {
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
			log.Fatalf("fail on set meta value, %v", err)
		}
	}
	parsedMeta = newMeta.Interface().(Meta)
	return
}
