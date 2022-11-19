parser grammar MetaParser;

options {
    tokenVocab = MetaLexer;
}

root:
    singleLine
    | multipleLine
    EOF;

singleLine:
    LINE_COMMENT metaBody;
multipleLine:
    BLOCK_COMMENT_START metaBody BLOCK_COMMENT_END ;

metaBody:
    metaQualifyName fieldExpr*;

metaQualifyName:
    META_QUALIFY_NAME;

fieldExpr:
    fieldNameValueExpr
    |fieldNameExpr;
fieldNameValueExpr:
    fieldName ASSIGNMENT valueExpr;
fieldName:
    IDENT;
valueExpr:
    boolValue
    | floatValue
    | integerValue
    | strValue;
boolValue:
    BOOLEAN;
strValue:
    STRING;
floatValue:
    FLOAT;
integerValue:
    INTEGER;
fieldNameExpr:
    fieldName;