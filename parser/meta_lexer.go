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
		4, 0, 14, 129, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 4, 0, 41, 8, 0, 11,
		0, 12, 0, 42, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 56, 8, 1, 1, 2, 4, 2, 59, 8, 2, 11, 2, 12, 2, 60, 1, 2,
		1, 2, 4, 2, 65, 8, 2, 11, 2, 12, 2, 66, 1, 3, 4, 3, 70, 8, 3, 11, 3, 12,
		3, 71, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7,
		3, 7, 85, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 5, 7, 91, 8, 7, 10, 7, 12, 7, 94,
		9, 7, 1, 8, 4, 8, 97, 8, 8, 11, 8, 12, 8, 98, 1, 8, 1, 8, 1, 9, 1, 9, 5,
		9, 105, 8, 9, 10, 9, 12, 9, 108, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1,
		16, 1, 17, 1, 17, 1, 17, 1, 17, 0, 0, 18, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5,
		11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 0, 25, 0, 27, 0, 29, 11,
		31, 12, 33, 13, 35, 14, 1, 0, 4, 3, 0, 9, 10, 13, 13, 32, 32, 4, 0, 10,
		10, 13, 13, 34, 34, 58, 58, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 136, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 40, 1, 0, 0, 0, 3, 55, 1, 0,
		0, 0, 5, 58, 1, 0, 0, 0, 7, 69, 1, 0, 0, 0, 9, 73, 1, 0, 0, 0, 11, 76,
		1, 0, 0, 0, 13, 79, 1, 0, 0, 0, 15, 84, 1, 0, 0, 0, 17, 96, 1, 0, 0, 0,
		19, 102, 1, 0, 0, 0, 21, 111, 1, 0, 0, 0, 23, 113, 1, 0, 0, 0, 25, 115,
		1, 0, 0, 0, 27, 117, 1, 0, 0, 0, 29, 119, 1, 0, 0, 0, 31, 121, 1, 0, 0,
		0, 33, 123, 1, 0, 0, 0, 35, 125, 1, 0, 0, 0, 37, 38, 3, 15, 7, 0, 38, 39,
		3, 33, 16, 0, 39, 41, 1, 0, 0, 0, 40, 37, 1, 0, 0, 0, 41, 42, 1, 0, 0,
		0, 42, 40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 45,
		3, 15, 7, 0, 45, 2, 1, 0, 0, 0, 46, 47, 5, 116, 0, 0, 47, 48, 5, 114, 0,
		0, 48, 49, 5, 117, 0, 0, 49, 56, 5, 101, 0, 0, 50, 51, 5, 102, 0, 0, 51,
		52, 5, 97, 0, 0, 52, 53, 5, 108, 0, 0, 53, 54, 5, 115, 0, 0, 54, 56, 5,
		101, 0, 0, 55, 46, 1, 0, 0, 0, 55, 50, 1, 0, 0, 0, 56, 4, 1, 0, 0, 0, 57,
		59, 3, 21, 10, 0, 58, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 58, 1, 0,
		0, 0, 60, 61, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 64, 3, 33, 16, 0, 63,
		65, 3, 21, 10, 0, 64, 63, 1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 64, 1, 0,
		0, 0, 66, 67, 1, 0, 0, 0, 67, 6, 1, 0, 0, 0, 68, 70, 3, 21, 10, 0, 69,
		68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0,
		0, 72, 8, 1, 0, 0, 0, 73, 74, 5, 47, 0, 0, 74, 75, 5, 47, 0, 0, 75, 10,
		1, 0, 0, 0, 76, 77, 5, 47, 0, 0, 77, 78, 5, 42, 0, 0, 78, 12, 1, 0, 0,
		0, 79, 80, 5, 42, 0, 0, 80, 81, 5, 47, 0, 0, 81, 14, 1, 0, 0, 0, 82, 85,
		3, 25, 12, 0, 83, 85, 3, 23, 11, 0, 84, 82, 1, 0, 0, 0, 84, 83, 1, 0, 0,
		0, 85, 92, 1, 0, 0, 0, 86, 91, 3, 25, 12, 0, 87, 91, 3, 21, 10, 0, 88,
		91, 3, 23, 11, 0, 89, 91, 5, 45, 0, 0, 90, 86, 1, 0, 0, 0, 90, 87, 1, 0,
		0, 0, 90, 88, 1, 0, 0, 0, 90, 89, 1, 0, 0, 0, 91, 94, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 92, 93, 1, 0, 0, 0, 93, 16, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0,
		95, 97, 7, 0, 0, 0, 96, 95, 1, 0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1,
		0, 0, 0, 98, 99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 6, 8, 0, 0,
		101, 18, 1, 0, 0, 0, 102, 106, 3, 27, 13, 0, 103, 105, 8, 1, 0, 0, 104,
		103, 1, 0, 0, 0, 105, 108, 1, 0, 0, 0, 106, 104, 1, 0, 0, 0, 106, 107,
		1, 0, 0, 0, 107, 109, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 109, 110, 3, 27,
		13, 0, 110, 20, 1, 0, 0, 0, 111, 112, 7, 2, 0, 0, 112, 22, 1, 0, 0, 0,
		113, 114, 5, 95, 0, 0, 114, 24, 1, 0, 0, 0, 115, 116, 7, 3, 0, 0, 116,
		26, 1, 0, 0, 0, 117, 118, 5, 34, 0, 0, 118, 28, 1, 0, 0, 0, 119, 120, 5,
		43, 0, 0, 120, 30, 1, 0, 0, 0, 121, 122, 5, 61, 0, 0, 122, 32, 1, 0, 0,
		0, 123, 124, 5, 46, 0, 0, 124, 34, 1, 0, 0, 0, 125, 126, 9, 0, 0, 0, 126,
		127, 1, 0, 0, 0, 127, 128, 6, 17, 0, 0, 128, 36, 1, 0, 0, 0, 11, 0, 42,
		55, 60, 66, 71, 84, 90, 92, 98, 106, 1, 0, 1, 0,
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
