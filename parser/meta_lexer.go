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
		"", "", "", "", "", "", "'//'", "'/*'", "'*/'", "", "", "'+'", "'='",
		"'.'",
	}
	staticData.symbolicNames = []string{
		"", "META_QUALIFY_NAME", "STRING", "BOOLEAN", "FLOAT", "INTEGER", "LINE_COMMENT",
		"BLOCK_COMMENT_START", "BLOCK_COMMENT_END", "IDENT", "WS", "PLUS", "ASSIGNMENT",
		"DOT", "ERRCHAR",
	}
	staticData.ruleNames = []string{
		"META_QUALIFY_NAME", "STRING", "BOOLEAN", "FLOAT", "INTEGER", "LINE_COMMENT",
		"BLOCK_COMMENT_START", "BLOCK_COMMENT_END", "IDENT", "WS", "DIGIT",
		"UNDERSCORE", "ALPHA", "STRING_F", "PLUS", "ASSIGNMENT", "DOT", "ERRCHAR",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 14, 131, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 42, 8,
		0, 10, 0, 12, 0, 45, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 5, 1, 51, 8, 1, 10,
		1, 12, 1, 54, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 1, 2, 1, 2, 3, 2, 67, 8, 2, 1, 3, 4, 3, 70, 8, 3, 11, 3, 12, 3, 71,
		1, 3, 1, 3, 4, 3, 76, 8, 3, 11, 3, 12, 3, 77, 1, 4, 4, 4, 81, 8, 4, 11,
		4, 12, 4, 82, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1,
		8, 1, 8, 3, 8, 96, 8, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 102, 8, 8, 10, 8,
		12, 8, 105, 9, 8, 1, 9, 4, 9, 108, 8, 9, 11, 9, 12, 9, 109, 1, 9, 1, 9,
		1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 17, 0, 0, 18, 1, 1, 3,
		2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 0,
		25, 0, 27, 0, 29, 11, 31, 12, 33, 13, 35, 14, 1, 0, 4, 1, 0, 34, 34, 3,
		0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 138, 0,
		1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0,
		9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0,
		0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0,
		0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 1, 37, 1, 0, 0, 0, 3, 48, 1, 0,
		0, 0, 5, 66, 1, 0, 0, 0, 7, 69, 1, 0, 0, 0, 9, 80, 1, 0, 0, 0, 11, 84,
		1, 0, 0, 0, 13, 87, 1, 0, 0, 0, 15, 90, 1, 0, 0, 0, 17, 95, 1, 0, 0, 0,
		19, 107, 1, 0, 0, 0, 21, 113, 1, 0, 0, 0, 23, 115, 1, 0, 0, 0, 25, 117,
		1, 0, 0, 0, 27, 119, 1, 0, 0, 0, 29, 121, 1, 0, 0, 0, 31, 123, 1, 0, 0,
		0, 33, 125, 1, 0, 0, 0, 35, 127, 1, 0, 0, 0, 37, 43, 3, 29, 14, 0, 38,
		39, 3, 17, 8, 0, 39, 40, 3, 33, 16, 0, 40, 42, 1, 0, 0, 0, 41, 38, 1, 0,
		0, 0, 42, 45, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0, 0, 44, 46,
		1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 46, 47, 3, 17, 8, 0, 47, 2, 1, 0, 0, 0,
		48, 52, 3, 27, 13, 0, 49, 51, 8, 0, 0, 0, 50, 49, 1, 0, 0, 0, 51, 54, 1,
		0, 0, 0, 52, 50, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 55, 1, 0, 0, 0, 54,
		52, 1, 0, 0, 0, 55, 56, 3, 27, 13, 0, 56, 4, 1, 0, 0, 0, 57, 58, 5, 116,
		0, 0, 58, 59, 5, 114, 0, 0, 59, 60, 5, 117, 0, 0, 60, 67, 5, 101, 0, 0,
		61, 62, 5, 102, 0, 0, 62, 63, 5, 97, 0, 0, 63, 64, 5, 108, 0, 0, 64, 65,
		5, 115, 0, 0, 65, 67, 5, 101, 0, 0, 66, 57, 1, 0, 0, 0, 66, 61, 1, 0, 0,
		0, 67, 6, 1, 0, 0, 0, 68, 70, 3, 21, 10, 0, 69, 68, 1, 0, 0, 0, 70, 71,
		1, 0, 0, 0, 71, 69, 1, 0, 0, 0, 71, 72, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0,
		73, 75, 3, 33, 16, 0, 74, 76, 3, 21, 10, 0, 75, 74, 1, 0, 0, 0, 76, 77,
		1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77, 78, 1, 0, 0, 0, 78, 8, 1, 0, 0, 0,
		79, 81, 3, 21, 10, 0, 80, 79, 1, 0, 0, 0, 81, 82, 1, 0, 0, 0, 82, 80, 1,
		0, 0, 0, 82, 83, 1, 0, 0, 0, 83, 10, 1, 0, 0, 0, 84, 85, 5, 47, 0, 0, 85,
		86, 5, 47, 0, 0, 86, 12, 1, 0, 0, 0, 87, 88, 5, 47, 0, 0, 88, 89, 5, 42,
		0, 0, 89, 14, 1, 0, 0, 0, 90, 91, 5, 42, 0, 0, 91, 92, 5, 47, 0, 0, 92,
		16, 1, 0, 0, 0, 93, 96, 3, 25, 12, 0, 94, 96, 3, 23, 11, 0, 95, 93, 1,
		0, 0, 0, 95, 94, 1, 0, 0, 0, 96, 103, 1, 0, 0, 0, 97, 102, 3, 25, 12, 0,
		98, 102, 3, 21, 10, 0, 99, 102, 3, 23, 11, 0, 100, 102, 5, 45, 0, 0, 101,
		97, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0, 0, 0, 101, 100, 1,
		0, 0, 0, 102, 105, 1, 0, 0, 0, 103, 101, 1, 0, 0, 0, 103, 104, 1, 0, 0,
		0, 104, 18, 1, 0, 0, 0, 105, 103, 1, 0, 0, 0, 106, 108, 7, 1, 0, 0, 107,
		106, 1, 0, 0, 0, 108, 109, 1, 0, 0, 0, 109, 107, 1, 0, 0, 0, 109, 110,
		1, 0, 0, 0, 110, 111, 1, 0, 0, 0, 111, 112, 6, 9, 0, 0, 112, 20, 1, 0,
		0, 0, 113, 114, 7, 2, 0, 0, 114, 22, 1, 0, 0, 0, 115, 116, 5, 95, 0, 0,
		116, 24, 1, 0, 0, 0, 117, 118, 7, 3, 0, 0, 118, 26, 1, 0, 0, 0, 119, 120,
		5, 34, 0, 0, 120, 28, 1, 0, 0, 0, 121, 122, 5, 43, 0, 0, 122, 30, 1, 0,
		0, 0, 123, 124, 5, 61, 0, 0, 124, 32, 1, 0, 0, 0, 125, 126, 5, 46, 0, 0,
		126, 34, 1, 0, 0, 0, 127, 128, 9, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130,
		6, 17, 0, 0, 130, 36, 1, 0, 0, 0, 11, 0, 43, 52, 66, 71, 77, 82, 95, 101,
		103, 109, 1, 0, 1, 0,
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
	MetaLexerSTRING              = 2
	MetaLexerBOOLEAN             = 3
	MetaLexerFLOAT               = 4
	MetaLexerINTEGER             = 5
	MetaLexerLINE_COMMENT        = 6
	MetaLexerBLOCK_COMMENT_START = 7
	MetaLexerBLOCK_COMMENT_END   = 8
	MetaLexerIDENT               = 9
	MetaLexerWS                  = 10
	MetaLexerPLUS                = 11
	MetaLexerASSIGNMENT          = 12
	MetaLexerDOT                 = 13
	MetaLexerERRCHAR             = 14
)
