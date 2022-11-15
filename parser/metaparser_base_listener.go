// Code generated from /home/kimloong/GolandProjects/gomelon/meta/parser/MetaParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MetaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseMetaParserListener is a complete listener for a parse tree produced by MetaParser.
type BaseMetaParserListener struct{}

var _ MetaParserListener = &BaseMetaParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseMetaParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseMetaParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseMetaParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseMetaParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseMetaParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseMetaParserListener) ExitRoot(ctx *RootContext) {}

// EnterSingleLine is called when production singleLine is entered.
func (s *BaseMetaParserListener) EnterSingleLine(ctx *SingleLineContext) {}

// ExitSingleLine is called when production singleLine is exited.
func (s *BaseMetaParserListener) ExitSingleLine(ctx *SingleLineContext) {}

// EnterMultipleLine is called when production multipleLine is entered.
func (s *BaseMetaParserListener) EnterMultipleLine(ctx *MultipleLineContext) {}

// ExitMultipleLine is called when production multipleLine is exited.
func (s *BaseMetaParserListener) ExitMultipleLine(ctx *MultipleLineContext) {}

// EnterMetaBody is called when production metaBody is entered.
func (s *BaseMetaParserListener) EnterMetaBody(ctx *MetaBodyContext) {}

// ExitMetaBody is called when production metaBody is exited.
func (s *BaseMetaParserListener) ExitMetaBody(ctx *MetaBodyContext) {}

// EnterMetaQualifyName is called when production metaQualifyName is entered.
func (s *BaseMetaParserListener) EnterMetaQualifyName(ctx *MetaQualifyNameContext) {}

// ExitMetaQualifyName is called when production metaQualifyName is exited.
func (s *BaseMetaParserListener) ExitMetaQualifyName(ctx *MetaQualifyNameContext) {}

// EnterFieldExpr is called when production fieldExpr is entered.
func (s *BaseMetaParserListener) EnterFieldExpr(ctx *FieldExprContext) {}

// ExitFieldExpr is called when production fieldExpr is exited.
func (s *BaseMetaParserListener) ExitFieldExpr(ctx *FieldExprContext) {}

// EnterFieldNameValueExpr is called when production fieldNameValueExpr is entered.
func (s *BaseMetaParserListener) EnterFieldNameValueExpr(ctx *FieldNameValueExprContext) {}

// ExitFieldNameValueExpr is called when production fieldNameValueExpr is exited.
func (s *BaseMetaParserListener) ExitFieldNameValueExpr(ctx *FieldNameValueExprContext) {}

// EnterFieldName is called when production fieldName is entered.
func (s *BaseMetaParserListener) EnterFieldName(ctx *FieldNameContext) {}

// ExitFieldName is called when production fieldName is exited.
func (s *BaseMetaParserListener) ExitFieldName(ctx *FieldNameContext) {}

// EnterValueExpr is called when production valueExpr is entered.
func (s *BaseMetaParserListener) EnterValueExpr(ctx *ValueExprContext) {}

// ExitValueExpr is called when production valueExpr is exited.
func (s *BaseMetaParserListener) ExitValueExpr(ctx *ValueExprContext) {}

// EnterBoolValue is called when production boolValue is entered.
func (s *BaseMetaParserListener) EnterBoolValue(ctx *BoolValueContext) {}

// ExitBoolValue is called when production boolValue is exited.
func (s *BaseMetaParserListener) ExitBoolValue(ctx *BoolValueContext) {}

// EnterStrValue is called when production strValue is entered.
func (s *BaseMetaParserListener) EnterStrValue(ctx *StrValueContext) {}

// ExitStrValue is called when production strValue is exited.
func (s *BaseMetaParserListener) ExitStrValue(ctx *StrValueContext) {}

// EnterFloatValue is called when production floatValue is entered.
func (s *BaseMetaParserListener) EnterFloatValue(ctx *FloatValueContext) {}

// ExitFloatValue is called when production floatValue is exited.
func (s *BaseMetaParserListener) ExitFloatValue(ctx *FloatValueContext) {}

// EnterIntegerValue is called when production integerValue is entered.
func (s *BaseMetaParserListener) EnterIntegerValue(ctx *IntegerValueContext) {}

// ExitIntegerValue is called when production integerValue is exited.
func (s *BaseMetaParserListener) ExitIntegerValue(ctx *IntegerValueContext) {}

// EnterFieldNameExpr is called when production fieldNameExpr is entered.
func (s *BaseMetaParserListener) EnterFieldNameExpr(ctx *FieldNameExprContext) {}

// ExitFieldNameExpr is called when production fieldNameExpr is exited.
func (s *BaseMetaParserListener) ExitFieldNameExpr(ctx *FieldNameExprContext) {}
