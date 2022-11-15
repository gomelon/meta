package meta

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/gomelon/meta/parser"
	"github.com/google/shlex"
	"go/token"
	"go/types"
	"strconv"
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

func (p *Parser) FilterByMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if p.HasMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (p *Parser) HasMeta(metaName string, object types.Object) bool {
	metas := p.ObjectMetaGroups(object, metaName)
	return len(metas) > 0
}

func (p *Parser) FilterByMethodHasMeta(metaName string, objects []types.Object) []types.Object {
	filteredObjects := make([]types.Object, 0, 8)
	for _, object := range objects {
		if p.HasMethodHasMeta(metaName, object) {
			filteredObjects = append(filteredObjects, object)
		}
	}
	return filteredObjects
}

func (p *Parser) HasMethodHasMeta(metaName string, object types.Object) bool {
	methods := p.pkgParser.Methods(object)
	for _, method := range methods {
		if p.HasMeta(metaName, method) {
			return true
		}
	}
	return false
}

func (p *Parser) ObjectMetaGroups(object types.Object, metaNames ...string) (
	parsedMetaGroups map[string]Group) {

	parsedMetaGroups = map[string]Group{}
	for _, metaName := range metaNames {
		parsedMetaGroup := p.ObjectMetaGroup(object, metaName)
		if len(parsedMetaGroup) == 0 {
			continue
		}
		parsedMetaGroups[metaName] = parsedMetaGroup
	}
	return
}

func (p *Parser) ObjectMetaGroup(object types.Object, metaName string) (parsedMetaGroup Group) {
	metaNameToParsedMetaGroup, ok := p.objectToParsedMetaGroups[object]
	if ok {
		parsedMetaGroup, ok = metaNameToParsedMetaGroup[metaName]
		if ok {
			return
		}
	}

	comments := p.pkgParser.Comments(object.Pos())
	for _, comment := range comments {
		parsedMeta, parsed := parse(metaName, comment)
		if !parsed {
			continue
		}
		parsedMetaGroup = append(parsedMetaGroup, parsedMeta)
	}
	if metaNameToParsedMetaGroup == nil {
		metaNameToParsedMetaGroup = map[string]Group{}
		p.objectToParsedMetaGroups[object] = metaNameToParsedMetaGroup
	}
	metaNameToParsedMetaGroup[metaName] = parsedMetaGroup
	return
}

func (p *Parser) ObjectMeta(object types.Object, metaName string) (m *Meta) {
	group := p.ObjectMetaGroup(object, metaName)
	if group != nil && len(group) > 0 {
		return group[0]
	} else {
		return nil
	}
}

func (p *Parser) filterComments(pos token.Pos, metaName string) []string {
	var filteredComments []string
	comments := p.pkgParser.Comments(pos)
	for _, comment := range comments {
		if strings.HasPrefix(comment, metaName) {
			filteredComments = append(filteredComments, comment)
		}
	}
	return filteredComments
}

func (p *Parser) populateMetaFields(metaName, comment string) (parsedMeta *Meta) {
	antlr.NewInputStream(comment)

	propertiesStr := strings.TrimLeft(comment, metaName)
	fieldAndValues, err := shlex.Split(propertiesStr)
	if err != nil {
		panic(fmt.Errorf("meta parse fail: %w", err))
	}

	properties := make(map[string]any, len(fieldAndValues))
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

func parse(qualifyName, comment string) (*Meta, bool) {
	if strings.Index(comment, qualifyName) < 1 {

	}
	is := antlr.NewInputStream(comment)
	lexer := parser.NewMetaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewMetaParser(stream)
	listener := newSpecMetaParserListener(qualifyName)
	p.AddParseListener(listener)
	p.MetaBody()
	parsedMeta := listener.parsedMeta
	return parsedMeta, parsedMeta != nil
}

type specMetaParserListener struct {
	*parser.BaseMetaParserListener
	specQualifyName  string
	parsedMeta       *Meta
	currentFieldName string
}

func newSpecMetaParserListener(specQualifyName string) *specMetaParserListener {
	return &specMetaParserListener{
		BaseMetaParserListener: &parser.BaseMetaParserListener{},
		specQualifyName:        specQualifyName,
	}
}

func (listener *specMetaParserListener) ExitMetaQualifyName(ctx *parser.MetaQualifyNameContext) {
	qualifyName := ctx.GetText()
	if listener.specQualifyName != qualifyName {
		return
	}
	listener.parsedMeta = New(qualifyName)
}

func (listener *specMetaParserListener) ExitFieldName(ctx *parser.FieldNameContext) {
	if listener.parsedMeta == nil {
		return
	}
	listener.currentFieldName = ctx.GetText()
}

func (listener *specMetaParserListener) ExitFieldNameExpr(ctx *parser.FieldNameExprContext) {
	if listener.parsedMeta == nil {
		return
	}
	listener.parsedMeta.properties[listener.currentFieldName] = true
}

func (listener *specMetaParserListener) ExitBoolValue(ctx *parser.BoolValueContext) {
	if listener.parsedMeta == nil {
		return
	}
	listener.parsedMeta.properties[listener.currentFieldName] = ctx.GetText() == "true"
}

func (listener *specMetaParserListener) ExitStrValue(ctx *parser.StrValueContext) {
	if listener.parsedMeta == nil {
		return
	}
	text := ctx.GetText()
	listener.parsedMeta.properties[listener.currentFieldName] = text[1 : len(text)-1]
}

func (listener *specMetaParserListener) ExitFloatValue(ctx *parser.FloatValueContext) {
	if listener.parsedMeta == nil {
		return
	}
	text := ctx.GetText()
	parseFloat, _ := strconv.ParseFloat(text, 64)
	listener.parsedMeta.properties[listener.currentFieldName] = parseFloat
}

func (listener *specMetaParserListener) ExitIntegerValue(ctx *parser.IntegerValueContext) {
	if listener.parsedMeta == nil {
		return
	}
	text := ctx.GetText()
	parseInt, _ := strconv.ParseInt(text, 10, 64)
	listener.parsedMeta.properties[listener.currentFieldName] = parseInt
}
