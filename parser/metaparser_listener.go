// Code generated from /home/kimloong/GolandProjects/gomelon/meta/parser/MetaParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // MetaParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// MetaParserListener is a complete listener for a parse tree produced by MetaParser.
type MetaParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSingleLine is called when entering the singleLine production.
	EnterSingleLine(c *SingleLineContext)

	// EnterMultipleLine is called when entering the multipleLine production.
	EnterMultipleLine(c *MultipleLineContext)

	// EnterMetaBody is called when entering the metaBody production.
	EnterMetaBody(c *MetaBodyContext)

	// EnterMetaQualifyName is called when entering the metaQualifyName production.
	EnterMetaQualifyName(c *MetaQualifyNameContext)

	// EnterFieldExpr is called when entering the fieldExpr production.
	EnterFieldExpr(c *FieldExprContext)

	// EnterFieldNameValueExpr is called when entering the fieldNameValueExpr production.
	EnterFieldNameValueExpr(c *FieldNameValueExprContext)

	// EnterFieldName is called when entering the fieldName production.
	EnterFieldName(c *FieldNameContext)

	// EnterValueExpr is called when entering the valueExpr production.
	EnterValueExpr(c *ValueExprContext)

	// EnterBoolValue is called when entering the boolValue production.
	EnterBoolValue(c *BoolValueContext)

	// EnterStrValue is called when entering the strValue production.
	EnterStrValue(c *StrValueContext)

	// EnterFloatValue is called when entering the floatValue production.
	EnterFloatValue(c *FloatValueContext)

	// EnterIntegerValue is called when entering the integerValue production.
	EnterIntegerValue(c *IntegerValueContext)

	// EnterFieldNameExpr is called when entering the fieldNameExpr production.
	EnterFieldNameExpr(c *FieldNameExprContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSingleLine is called when exiting the singleLine production.
	ExitSingleLine(c *SingleLineContext)

	// ExitMultipleLine is called when exiting the multipleLine production.
	ExitMultipleLine(c *MultipleLineContext)

	// ExitMetaBody is called when exiting the metaBody production.
	ExitMetaBody(c *MetaBodyContext)

	// ExitMetaQualifyName is called when exiting the metaQualifyName production.
	ExitMetaQualifyName(c *MetaQualifyNameContext)

	// ExitFieldExpr is called when exiting the fieldExpr production.
	ExitFieldExpr(c *FieldExprContext)

	// ExitFieldNameValueExpr is called when exiting the fieldNameValueExpr production.
	ExitFieldNameValueExpr(c *FieldNameValueExprContext)

	// ExitFieldName is called when exiting the fieldName production.
	ExitFieldName(c *FieldNameContext)

	// ExitValueExpr is called when exiting the valueExpr production.
	ExitValueExpr(c *ValueExprContext)

	// ExitBoolValue is called when exiting the boolValue production.
	ExitBoolValue(c *BoolValueContext)

	// ExitStrValue is called when exiting the strValue production.
	ExitStrValue(c *StrValueContext)

	// ExitFloatValue is called when exiting the floatValue production.
	ExitFloatValue(c *FloatValueContext)

	// ExitIntegerValue is called when exiting the integerValue production.
	ExitIntegerValue(c *IntegerValueContext)

	// ExitFieldNameExpr is called when exiting the fieldNameExpr production.
	ExitFieldNameExpr(c *FieldNameExprContext)
}
