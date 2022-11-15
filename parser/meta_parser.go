// Code generated from /home/kimloong/GolandProjects/gomelon/meta/parser/MetaParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MetaParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type MetaParser struct {
	*antlr.BaseParser
}

var metaparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func metaparserParserInit() {
	staticData := &metaparserParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "'//'", "'/*'", "'*/'", "", "", "", "'+'", "'='",
		"'.'",
	}
	staticData.symbolicNames = []string{
		"", "META_QUALIFY_NAME", "BOOLEAN", "FLOAT", "INTEGER", "LINE_COMMENT",
		"BLOCK_COMMENT_START", "BLOCK_COMMENT_END", "IDENT", "WS", "STRING",
		"PLUS", "ASSIGNMENT", "DOT", "ERRCHAR",
	}
	staticData.ruleNames = []string{
		"root", "singleLine", "multipleLine", "metaBody", "metaQualifyName",
		"fieldExpr", "fieldNameValueExpr", "fieldName", "valueExpr", "boolValue",
		"strValue", "floatValue", "integerValue", "fieldNameExpr",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 14, 78, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 1, 0, 1, 0, 1, 0, 1, 0, 3,
		0, 33, 8, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3,
		5, 3, 45, 8, 3, 10, 3, 12, 3, 48, 9, 3, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 54,
		8, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8,
		66, 8, 8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 13, 0, 0, 14, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 0, 0, 69, 0, 32, 1, 0, 0, 0, 2, 34, 1, 0, 0, 0, 4, 37, 1, 0, 0, 0,
		6, 41, 1, 0, 0, 0, 8, 49, 1, 0, 0, 0, 10, 53, 1, 0, 0, 0, 12, 55, 1, 0,
		0, 0, 14, 59, 1, 0, 0, 0, 16, 65, 1, 0, 0, 0, 18, 67, 1, 0, 0, 0, 20, 69,
		1, 0, 0, 0, 22, 71, 1, 0, 0, 0, 24, 73, 1, 0, 0, 0, 26, 75, 1, 0, 0, 0,
		28, 33, 3, 2, 1, 0, 29, 30, 3, 4, 2, 0, 30, 31, 5, 0, 0, 1, 31, 33, 1,
		0, 0, 0, 32, 28, 1, 0, 0, 0, 32, 29, 1, 0, 0, 0, 33, 1, 1, 0, 0, 0, 34,
		35, 5, 5, 0, 0, 35, 36, 3, 6, 3, 0, 36, 3, 1, 0, 0, 0, 37, 38, 5, 6, 0,
		0, 38, 39, 3, 6, 3, 0, 39, 40, 5, 7, 0, 0, 40, 5, 1, 0, 0, 0, 41, 42, 5,
		11, 0, 0, 42, 46, 3, 8, 4, 0, 43, 45, 3, 10, 5, 0, 44, 43, 1, 0, 0, 0,
		45, 48, 1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 7, 1, 0,
		0, 0, 48, 46, 1, 0, 0, 0, 49, 50, 5, 1, 0, 0, 50, 9, 1, 0, 0, 0, 51, 54,
		3, 12, 6, 0, 52, 54, 3, 26, 13, 0, 53, 51, 1, 0, 0, 0, 53, 52, 1, 0, 0,
		0, 54, 11, 1, 0, 0, 0, 55, 56, 3, 14, 7, 0, 56, 57, 5, 12, 0, 0, 57, 58,
		3, 16, 8, 0, 58, 13, 1, 0, 0, 0, 59, 60, 5, 8, 0, 0, 60, 15, 1, 0, 0, 0,
		61, 66, 3, 18, 9, 0, 62, 66, 3, 22, 11, 0, 63, 66, 3, 24, 12, 0, 64, 66,
		3, 20, 10, 0, 65, 61, 1, 0, 0, 0, 65, 62, 1, 0, 0, 0, 65, 63, 1, 0, 0,
		0, 65, 64, 1, 0, 0, 0, 66, 17, 1, 0, 0, 0, 67, 68, 5, 2, 0, 0, 68, 19,
		1, 0, 0, 0, 69, 70, 5, 10, 0, 0, 70, 21, 1, 0, 0, 0, 71, 72, 5, 3, 0, 0,
		72, 23, 1, 0, 0, 0, 73, 74, 5, 4, 0, 0, 74, 25, 1, 0, 0, 0, 75, 76, 3,
		14, 7, 0, 76, 27, 1, 0, 0, 0, 4, 32, 46, 53, 65,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// MetaParserInit initializes any static state used to implement MetaParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewMetaParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func MetaParserInit() {
	staticData := &metaparserParserStaticData
	staticData.once.Do(metaparserParserInit)
}

// NewMetaParser produces a new parser instance for the optional input antlr.TokenStream.
func NewMetaParser(input antlr.TokenStream) *MetaParser {
	MetaParserInit()
	this := new(MetaParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &metaparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "MetaParser.g4"

	return this
}

// MetaParser tokens.
const (
	MetaParserEOF                 = antlr.TokenEOF
	MetaParserMETA_QUALIFY_NAME   = 1
	MetaParserBOOLEAN             = 2
	MetaParserFLOAT               = 3
	MetaParserINTEGER             = 4
	MetaParserLINE_COMMENT        = 5
	MetaParserBLOCK_COMMENT_START = 6
	MetaParserBLOCK_COMMENT_END   = 7
	MetaParserIDENT               = 8
	MetaParserWS                  = 9
	MetaParserSTRING              = 10
	MetaParserPLUS                = 11
	MetaParserASSIGNMENT          = 12
	MetaParserDOT                 = 13
	MetaParserERRCHAR             = 14
)

// MetaParser rules.
const (
	MetaParserRULE_root               = 0
	MetaParserRULE_singleLine         = 1
	MetaParserRULE_multipleLine       = 2
	MetaParserRULE_metaBody           = 3
	MetaParserRULE_metaQualifyName    = 4
	MetaParserRULE_fieldExpr          = 5
	MetaParserRULE_fieldNameValueExpr = 6
	MetaParserRULE_fieldName          = 7
	MetaParserRULE_valueExpr          = 8
	MetaParserRULE_boolValue          = 9
	MetaParserRULE_strValue           = 10
	MetaParserRULE_floatValue         = 11
	MetaParserRULE_integerValue       = 12
	MetaParserRULE_fieldNameExpr      = 13
)

// IRootContext is an interface to support dynamic dispatch.
type IRootContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRootContext differentiates from other interfaces.
	IsRootContext()
}

type RootContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRootContext() *RootContext {
	var p = new(RootContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_root
	return p
}

func (*RootContext) IsRootContext() {}

func NewRootContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RootContext {
	var p = new(RootContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_root

	return p
}

func (s *RootContext) GetParser() antlr.Parser { return s.parser }

func (s *RootContext) SingleLine() ISingleLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISingleLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISingleLineContext)
}

func (s *RootContext) MultipleLine() IMultipleLineContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultipleLineContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultipleLineContext)
}

func (s *RootContext) EOF() antlr.TerminalNode {
	return s.GetToken(MetaParserEOF, 0)
}

func (s *RootContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RootContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RootContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterRoot(s)
	}
}

func (s *RootContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitRoot(s)
	}
}

func (p *MetaParser) Root() (localctx IRootContext) {
	this := p
	_ = this

	localctx = NewRootContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, MetaParserRULE_root)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MetaParserLINE_COMMENT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.SingleLine()
		}

	case MetaParserBLOCK_COMMENT_START:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(29)
			p.MultipleLine()
		}
		{
			p.SetState(30)
			p.Match(MetaParserEOF)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISingleLineContext is an interface to support dynamic dispatch.
type ISingleLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingleLineContext differentiates from other interfaces.
	IsSingleLineContext()
}

type SingleLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingleLineContext() *SingleLineContext {
	var p = new(SingleLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_singleLine
	return p
}

func (*SingleLineContext) IsSingleLineContext() {}

func NewSingleLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SingleLineContext {
	var p = new(SingleLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_singleLine

	return p
}

func (s *SingleLineContext) GetParser() antlr.Parser { return s.parser }

func (s *SingleLineContext) LINE_COMMENT() antlr.TerminalNode {
	return s.GetToken(MetaParserLINE_COMMENT, 0)
}

func (s *SingleLineContext) MetaBody() IMetaBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaBodyContext)
}

func (s *SingleLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SingleLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SingleLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterSingleLine(s)
	}
}

func (s *SingleLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitSingleLine(s)
	}
}

func (p *MetaParser) SingleLine() (localctx ISingleLineContext) {
	this := p
	_ = this

	localctx = NewSingleLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, MetaParserRULE_singleLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(MetaParserLINE_COMMENT)
	}
	{
		p.SetState(35)
		p.MetaBody()
	}

	return localctx
}

// IMultipleLineContext is an interface to support dynamic dispatch.
type IMultipleLineContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMultipleLineContext differentiates from other interfaces.
	IsMultipleLineContext()
}

type MultipleLineContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultipleLineContext() *MultipleLineContext {
	var p = new(MultipleLineContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_multipleLine
	return p
}

func (*MultipleLineContext) IsMultipleLineContext() {}

func NewMultipleLineContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultipleLineContext {
	var p = new(MultipleLineContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_multipleLine

	return p
}

func (s *MultipleLineContext) GetParser() antlr.Parser { return s.parser }

func (s *MultipleLineContext) BLOCK_COMMENT_START() antlr.TerminalNode {
	return s.GetToken(MetaParserBLOCK_COMMENT_START, 0)
}

func (s *MultipleLineContext) MetaBody() IMetaBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaBodyContext)
}

func (s *MultipleLineContext) BLOCK_COMMENT_END() antlr.TerminalNode {
	return s.GetToken(MetaParserBLOCK_COMMENT_END, 0)
}

func (s *MultipleLineContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultipleLineContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultipleLineContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterMultipleLine(s)
	}
}

func (s *MultipleLineContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitMultipleLine(s)
	}
}

func (p *MetaParser) MultipleLine() (localctx IMultipleLineContext) {
	this := p
	_ = this

	localctx = NewMultipleLineContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, MetaParserRULE_multipleLine)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(37)
		p.Match(MetaParserBLOCK_COMMENT_START)
	}
	{
		p.SetState(38)
		p.MetaBody()
	}
	{
		p.SetState(39)
		p.Match(MetaParserBLOCK_COMMENT_END)
	}

	return localctx
}

// IMetaBodyContext is an interface to support dynamic dispatch.
type IMetaBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetaBodyContext differentiates from other interfaces.
	IsMetaBodyContext()
}

type MetaBodyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaBodyContext() *MetaBodyContext {
	var p = new(MetaBodyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_metaBody
	return p
}

func (*MetaBodyContext) IsMetaBodyContext() {}

func NewMetaBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaBodyContext {
	var p = new(MetaBodyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_metaBody

	return p
}

func (s *MetaBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaBodyContext) PLUS() antlr.TerminalNode {
	return s.GetToken(MetaParserPLUS, 0)
}

func (s *MetaBodyContext) MetaQualifyName() IMetaQualifyNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMetaQualifyNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMetaQualifyNameContext)
}

func (s *MetaBodyContext) AllFieldExpr() []IFieldExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldExprContext); ok {
			len++
		}
	}

	tst := make([]IFieldExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldExprContext); ok {
			tst[i] = t.(IFieldExprContext)
			i++
		}
	}

	return tst
}

func (s *MetaBodyContext) FieldExpr(i int) IFieldExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldExprContext)
}

func (s *MetaBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaBodyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterMetaBody(s)
	}
}

func (s *MetaBodyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitMetaBody(s)
	}
}

func (p *MetaParser) MetaBody() (localctx IMetaBodyContext) {
	this := p
	_ = this

	localctx = NewMetaBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, MetaParserRULE_metaBody)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(41)
		p.Match(MetaParserPLUS)
	}
	{
		p.SetState(42)
		p.MetaQualifyName()
	}
	p.SetState(46)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == MetaParserIDENT {
		{
			p.SetState(43)
			p.FieldExpr()
		}

		p.SetState(48)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMetaQualifyNameContext is an interface to support dynamic dispatch.
type IMetaQualifyNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetaQualifyNameContext differentiates from other interfaces.
	IsMetaQualifyNameContext()
}

type MetaQualifyNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetaQualifyNameContext() *MetaQualifyNameContext {
	var p = new(MetaQualifyNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_metaQualifyName
	return p
}

func (*MetaQualifyNameContext) IsMetaQualifyNameContext() {}

func NewMetaQualifyNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetaQualifyNameContext {
	var p = new(MetaQualifyNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_metaQualifyName

	return p
}

func (s *MetaQualifyNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MetaQualifyNameContext) META_QUALIFY_NAME() antlr.TerminalNode {
	return s.GetToken(MetaParserMETA_QUALIFY_NAME, 0)
}

func (s *MetaQualifyNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaQualifyNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetaQualifyNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterMetaQualifyName(s)
	}
}

func (s *MetaQualifyNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitMetaQualifyName(s)
	}
}

func (p *MetaParser) MetaQualifyName() (localctx IMetaQualifyNameContext) {
	this := p
	_ = this

	localctx = NewMetaQualifyNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, MetaParserRULE_metaQualifyName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(MetaParserMETA_QUALIFY_NAME)
	}

	return localctx
}

// IFieldExprContext is an interface to support dynamic dispatch.
type IFieldExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldExprContext differentiates from other interfaces.
	IsFieldExprContext()
}

type FieldExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldExprContext() *FieldExprContext {
	var p = new(FieldExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_fieldExpr
	return p
}

func (*FieldExprContext) IsFieldExprContext() {}

func NewFieldExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldExprContext {
	var p = new(FieldExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_fieldExpr

	return p
}

func (s *FieldExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldExprContext) FieldNameValueExpr() IFieldNameValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameValueExprContext)
}

func (s *FieldExprContext) FieldNameExpr() IFieldNameExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameExprContext)
}

func (s *FieldExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterFieldExpr(s)
	}
}

func (s *FieldExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitFieldExpr(s)
	}
}

func (p *MetaParser) FieldExpr() (localctx IFieldExprContext) {
	this := p
	_ = this

	localctx = NewFieldExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, MetaParserRULE_fieldExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(51)
			p.FieldNameValueExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.FieldNameExpr()
		}

	}

	return localctx
}

// IFieldNameValueExprContext is an interface to support dynamic dispatch.
type IFieldNameValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameValueExprContext differentiates from other interfaces.
	IsFieldNameValueExprContext()
}

type FieldNameValueExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameValueExprContext() *FieldNameValueExprContext {
	var p = new(FieldNameValueExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_fieldNameValueExpr
	return p
}

func (*FieldNameValueExprContext) IsFieldNameValueExprContext() {}

func NewFieldNameValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameValueExprContext {
	var p = new(FieldNameValueExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_fieldNameValueExpr

	return p
}

func (s *FieldNameValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameValueExprContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldNameValueExprContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(MetaParserASSIGNMENT, 0)
}

func (s *FieldNameValueExprContext) ValueExpr() IValueExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueExprContext)
}

func (s *FieldNameValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterFieldNameValueExpr(s)
	}
}

func (s *FieldNameValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitFieldNameValueExpr(s)
	}
}

func (p *MetaParser) FieldNameValueExpr() (localctx IFieldNameValueExprContext) {
	this := p
	_ = this

	localctx = NewFieldNameValueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, MetaParserRULE_fieldNameValueExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(55)
		p.FieldName()
	}
	{
		p.SetState(56)
		p.Match(MetaParserASSIGNMENT)
	}
	{
		p.SetState(57)
		p.ValueExpr()
	}

	return localctx
}

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_fieldName
	return p
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) IDENT() antlr.TerminalNode {
	return s.GetToken(MetaParserIDENT, 0)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterFieldName(s)
	}
}

func (s *FieldNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitFieldName(s)
	}
}

func (p *MetaParser) FieldName() (localctx IFieldNameContext) {
	this := p
	_ = this

	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, MetaParserRULE_fieldName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(59)
		p.Match(MetaParserIDENT)
	}

	return localctx
}

// IValueExprContext is an interface to support dynamic dispatch.
type IValueExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueExprContext differentiates from other interfaces.
	IsValueExprContext()
}

type ValueExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueExprContext() *ValueExprContext {
	var p = new(ValueExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_valueExpr
	return p
}

func (*ValueExprContext) IsValueExprContext() {}

func NewValueExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueExprContext {
	var p = new(ValueExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_valueExpr

	return p
}

func (s *ValueExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueExprContext) BoolValue() IBoolValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *ValueExprContext) FloatValue() IFloatValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatValueContext)
}

func (s *ValueExprContext) IntegerValue() IIntegerValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntegerValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntegerValueContext)
}

func (s *ValueExprContext) StrValue() IStrValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrValueContext)
}

func (s *ValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterValueExpr(s)
	}
}

func (s *ValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitValueExpr(s)
	}
}

func (p *MetaParser) ValueExpr() (localctx IValueExprContext) {
	this := p
	_ = this

	localctx = NewValueExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, MetaParserRULE_valueExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(65)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case MetaParserBOOLEAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.BoolValue()
		}

	case MetaParserFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)
			p.FloatValue()
		}

	case MetaParserINTEGER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.IntegerValue()
		}

	case MetaParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)
			p.StrValue()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBoolValueContext is an interface to support dynamic dispatch.
type IBoolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoolValueContext differentiates from other interfaces.
	IsBoolValueContext()
}

type BoolValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolValueContext() *BoolValueContext {
	var p = new(BoolValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_boolValue
	return p
}

func (*BoolValueContext) IsBoolValueContext() {}

func NewBoolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolValueContext {
	var p = new(BoolValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_boolValue

	return p
}

func (s *BoolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolValueContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(MetaParserBOOLEAN, 0)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (p *MetaParser) BoolValue() (localctx IBoolValueContext) {
	this := p
	_ = this

	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, MetaParserRULE_boolValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(MetaParserBOOLEAN)
	}

	return localctx
}

// IStrValueContext is an interface to support dynamic dispatch.
type IStrValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStrValueContext differentiates from other interfaces.
	IsStrValueContext()
}

type StrValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrValueContext() *StrValueContext {
	var p = new(StrValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_strValue
	return p
}

func (*StrValueContext) IsStrValueContext() {}

func NewStrValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrValueContext {
	var p = new(StrValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_strValue

	return p
}

func (s *StrValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StrValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(MetaParserSTRING, 0)
}

func (s *StrValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterStrValue(s)
	}
}

func (s *StrValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitStrValue(s)
	}
}

func (p *MetaParser) StrValue() (localctx IStrValueContext) {
	this := p
	_ = this

	localctx = NewStrValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, MetaParserRULE_strValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(69)
		p.Match(MetaParserSTRING)
	}

	return localctx
}

// IFloatValueContext is an interface to support dynamic dispatch.
type IFloatValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatValueContext differentiates from other interfaces.
	IsFloatValueContext()
}

type FloatValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatValueContext() *FloatValueContext {
	var p = new(FloatValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_floatValue
	return p
}

func (*FloatValueContext) IsFloatValueContext() {}

func NewFloatValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatValueContext {
	var p = new(FloatValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_floatValue

	return p
}

func (s *FloatValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(MetaParserFLOAT, 0)
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (p *MetaParser) FloatValue() (localctx IFloatValueContext) {
	this := p
	_ = this

	localctx = NewFloatValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, MetaParserRULE_floatValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(71)
		p.Match(MetaParserFLOAT)
	}

	return localctx
}

// IIntegerValueContext is an interface to support dynamic dispatch.
type IIntegerValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerValueContext differentiates from other interfaces.
	IsIntegerValueContext()
}

type IntegerValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerValueContext() *IntegerValueContext {
	var p = new(IntegerValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_integerValue
	return p
}

func (*IntegerValueContext) IsIntegerValueContext() {}

func NewIntegerValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerValueContext {
	var p = new(IntegerValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_integerValue

	return p
}

func (s *IntegerValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerValueContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(MetaParserINTEGER, 0)
}

func (s *IntegerValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterIntegerValue(s)
	}
}

func (s *IntegerValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitIntegerValue(s)
	}
}

func (p *MetaParser) IntegerValue() (localctx IIntegerValueContext) {
	this := p
	_ = this

	localctx = NewIntegerValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, MetaParserRULE_integerValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(73)
		p.Match(MetaParserINTEGER)
	}

	return localctx
}

// IFieldNameExprContext is an interface to support dynamic dispatch.
type IFieldNameExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldNameExprContext differentiates from other interfaces.
	IsFieldNameExprContext()
}

type FieldNameExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameExprContext() *FieldNameExprContext {
	var p = new(FieldNameExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = MetaParserRULE_fieldNameExpr
	return p
}

func (*FieldNameExprContext) IsFieldNameExprContext() {}

func NewFieldNameExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameExprContext {
	var p = new(FieldNameExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = MetaParserRULE_fieldNameExpr

	return p
}

func (s *FieldNameExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameExprContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldNameExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.EnterFieldNameExpr(s)
	}
}

func (s *FieldNameExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(MetaParserListener); ok {
		listenerT.ExitFieldNameExpr(s)
	}
}

func (p *MetaParser) FieldNameExpr() (localctx IFieldNameExprContext) {
	this := p
	_ = this

	localctx = NewFieldNameExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, MetaParserRULE_fieldNameExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(75)
		p.FieldName()
	}

	return localctx
}
