// Code generated from /home/kimloong/GolandProjects/gomelon/meta/parser/MetaLexer.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type MetaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var metalexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func metalexerLexerInit() {
	staticData := &metalexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
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
		"META_QUALIFY_NAME", "BOOLEAN", "FLOAT", "INTEGER", "LINE_COMMENT",
		"BLOCK_COMMENT_START", "BLOCK_COMMENT_END", "IDENT", "WS", "STRING",
		"DIGIT", "UNDERSCORE", "ALPHA", "STRING_F", "PLUS", "ASSIGNMENT", "DOT",
		"ERRCHAR",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 131, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 42, 8,
		0, 10, 0, 12, 0, 45, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 58, 8, 1, 1, 2, 4, 2, 61, 8, 2, 11, 2, 12, 2,
		62, 1, 2, 1, 2, 4, 2, 67, 8, 2, 11, 2, 12, 2, 68, 1, 3, 4, 3, 72, 8, 3,
		11, 3, 12, 3, 73, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6,
		1, 7, 1, 7, 3, 7, 87, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 93, 8, 7, 10,
		7, 12, 7, 96, 9, 7, 1, 8, 4, 8, 99, 8, 8, 11, 8, 12, 8, 100, 1, 8, 1, 8,
		1, 9, 1, 9, 5, 9, 107, 8, 9, 10, 9, 12, 9, 110, 9, 9, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 0, 0, 18, 1, 1, 3, 2, 5,
		3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 0, 25, 0,
		27, 0, 29, 11, 31, 12, 33, 13, 35, 14, 1, 0, 4, 3, 0, 9, 10, 13, 13, 32,
		32, 4, 0, 10, 10, 13, 13, 34, 34, 58, 58, 1, 0, 48, 57, 2, 0, 65, 90, 97,
		122, 138, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0,
		31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0,
		3, 57, 1, 0, 0, 0, 5, 60, 1, 0, 0, 0, 7, 71, 1, 0, 0, 0, 9, 75, 1, 0, 0,
		0, 11, 78, 1, 0, 0, 0, 13, 81, 1, 0, 0, 0, 15, 86, 1, 0, 0, 0, 17, 98,
		1, 0, 0, 0, 19, 104, 1, 0, 0, 0, 21, 113, 1, 0, 0, 0, 23, 115, 1, 0, 0,
		0, 25, 117, 1, 0, 0, 0, 27, 119, 1, 0, 0, 0, 29, 121, 1, 0, 0, 0, 31, 123,
		1, 0, 0, 0, 33, 125, 1, 0, 0, 0, 35, 127, 1, 0, 0, 0, 37, 43, 3, 29, 14,
		0, 38, 39, 3, 15, 7, 0, 39, 40, 3, 33, 16, 0, 40, 42, 1, 0, 0, 0, 41, 38,
		1, 0, 0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0,
		44, 46, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 3, 15, 7, 0, 47, 2, 1,
		0, 0, 0, 48, 49, 5, 116, 0, 0, 49, 50, 5, 114, 0, 0, 50, 51, 5, 117, 0,
		0, 51, 58, 5, 101, 0, 0, 52, 53, 5, 102, 0, 0, 53, 54, 5, 97, 0, 0, 54,
		55, 5, 108, 0, 0, 55, 56, 5, 115, 0, 0, 56, 58, 5, 101, 0, 0, 57, 48, 1,
		0, 0, 0, 57, 52, 1, 0, 0, 0, 58, 4, 1, 0, 0, 0, 59, 61, 3, 21, 10, 0, 60,
		59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0,
		0, 63, 64, 1, 0, 0, 0, 64, 66, 3, 33, 16, 0, 65, 67, 3, 21, 10, 0, 66,
		65, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 66, 1, 0, 0, 0, 68, 69, 1, 0, 0,
		0, 69, 6, 1, 0, 0, 0, 70, 72, 3, 21, 10, 0, 71, 70, 1, 0, 0, 0, 72, 73,
		1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0, 74, 8, 1, 0, 0, 0,
		75, 76, 5, 47, 0, 0, 76, 77, 5, 47, 0, 0, 77, 10, 1, 0, 0, 0, 78, 79, 5,
		47, 0, 0, 79, 80, 5, 42, 0, 0, 80, 12, 1, 0, 0, 0, 81, 82, 5, 42, 0, 0,
		82, 83, 5, 47, 0, 0, 83, 14, 1, 0, 0, 0, 84, 87, 3, 25, 12, 0, 85, 87,
		3, 23, 11, 0, 86, 84, 1, 0, 0, 0, 86, 85, 1, 0, 0, 0, 87, 94, 1, 0, 0,
		0, 88, 93, 3, 25, 12, 0, 89, 93, 3, 21, 10, 0, 90, 93, 3, 23, 11, 0, 91,
		93, 5, 45, 0, 0, 92, 88, 1, 0, 0, 0, 92, 89, 1, 0, 0, 0, 92, 90, 1, 0,
		0, 0, 92, 91, 1, 0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95,
		1, 0, 0, 0, 95, 16, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 99, 7, 0, 0, 0,
		98, 97, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 6, 8, 0, 0, 103, 18, 1, 0,
		0, 0, 104, 108, 3, 27, 13, 0, 105, 107, 8, 1, 0, 0, 106, 105, 1, 0, 0,
		0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109,
		111, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 111, 112, 3, 27, 13, 0, 112, 20,
		1, 0, 0, 0, 113, 114, 7, 2, 0, 0, 114, 22, 1, 0, 0, 0, 115, 116, 5, 95,
		0, 0, 116, 24, 1, 0, 0, 0, 117, 118, 7, 3, 0, 0, 118, 26, 1, 0, 0, 0, 119,
		120, 5, 34, 0, 0, 120, 28, 1, 0, 0, 0, 121, 122, 5, 43, 0, 0, 122, 30,
		1, 0, 0, 0, 123, 124, 5, 61, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 46,
		0, 0, 126, 34, 1, 0, 0, 0, 127, 128, 9, 0, 0, 0, 128, 129, 1, 0, 0, 0,
		129, 130, 6, 17, 0, 0, 130, 36, 1, 0, 0, 0, 11, 0, 43, 57, 62, 68, 73,
		86, 92, 94, 100, 108, 1, 0, 1, 0,
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

// MetaLexerInit initializes any static state used to implement MetaLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewMetaLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func MetaLexerInit() {
	staticData := &metalexerLexerStaticData
	staticData.once.Do(metalexerLexerInit)
}

// NewMetaLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewMetaLexer(input antlr.CharStream) *MetaLexer {
	MetaLexerInit()
	l := new(MetaLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &metalexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "MetaLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MetaLexer tokens.
const (
	MetaLexerMETA_QUALIFY_NAME   = 1
	MetaLexerBOOLEAN             = 2
	MetaLexerFLOAT               = 3
	MetaLexerINTEGER             = 4
	MetaLexerLINE_COMMENT        = 5
	MetaLexerBLOCK_COMMENT_START = 6
	MetaLexerBLOCK_COMMENT_END   = 7
	MetaLexerIDENT               = 8
	MetaLexerWS                  = 9
	MetaLexerSTRING              = 10
	MetaLexerPLUS                = 11
	MetaLexerASSIGNMENT          = 12
	MetaLexerDOT                 = 13
	MetaLexerERRCHAR             = 14
)
