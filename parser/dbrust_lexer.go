// Code generated from DBRustLexer.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 28, 177,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 6, 10, 102, 10, 10,
	13, 10, 14, 10, 103, 3, 11, 6, 11, 107, 10, 11, 13, 11, 14, 11, 108, 3,
	11, 3, 11, 6, 11, 113, 10, 11, 13, 11, 14, 11, 114, 3, 12, 3, 12, 7, 12,
	119, 10, 12, 12, 12, 14, 12, 122, 11, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 14, 3, 14, 7, 14, 132, 10, 14, 12, 14, 14, 14, 135, 11, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3,
	16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3,
	27, 6, 27, 169, 10, 27, 13, 27, 14, 27, 170, 3, 27, 3, 27, 3, 28, 3, 28,
	3, 28, 2, 2, 29, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19,
	11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37,
	20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55,
	2, 3, 2, 9, 3, 2, 50, 59, 3, 2, 36, 36, 3, 2, 41, 41, 5, 2, 67, 92, 97,
	97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15,
	34, 34, 9, 2, 34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2,
	181, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2,
	2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3,
	2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25,
	3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2,
	33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2,
	2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2,
	2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 3, 57, 3, 2,
	2, 2, 5, 61, 3, 2, 2, 2, 7, 70, 3, 2, 2, 2, 9, 74, 3, 2, 2, 2, 11, 78,
	3, 2, 2, 2, 13, 83, 3, 2, 2, 2, 15, 88, 3, 2, 2, 2, 17, 93, 3, 2, 2, 2,
	19, 101, 3, 2, 2, 2, 21, 106, 3, 2, 2, 2, 23, 116, 3, 2, 2, 2, 25, 125,
	3, 2, 2, 2, 27, 129, 3, 2, 2, 2, 29, 136, 3, 2, 2, 2, 31, 142, 3, 2, 2,
	2, 33, 147, 3, 2, 2, 2, 35, 149, 3, 2, 2, 2, 37, 151, 3, 2, 2, 2, 39, 153,
	3, 2, 2, 2, 41, 155, 3, 2, 2, 2, 43, 157, 3, 2, 2, 2, 45, 159, 3, 2, 2,
	2, 47, 161, 3, 2, 2, 2, 49, 163, 3, 2, 2, 2, 51, 165, 3, 2, 2, 2, 53, 168,
	3, 2, 2, 2, 55, 174, 3, 2, 2, 2, 57, 58, 7, 110, 2, 2, 58, 59, 7, 103,
	2, 2, 59, 60, 7, 118, 2, 2, 60, 4, 3, 2, 2, 2, 61, 62, 7, 114, 2, 2, 62,
	63, 7, 116, 2, 2, 63, 64, 7, 107, 2, 2, 64, 65, 7, 112, 2, 2, 65, 66, 7,
	118, 2, 2, 66, 67, 7, 110, 2, 2, 67, 68, 7, 112, 2, 2, 68, 69, 7, 35, 2,
	2, 69, 6, 3, 2, 2, 2, 70, 71, 7, 107, 2, 2, 71, 72, 7, 56, 2, 2, 72, 73,
	7, 54, 2, 2, 73, 8, 3, 2, 2, 2, 74, 75, 7, 104, 2, 2, 75, 76, 7, 56, 2,
	2, 76, 77, 7, 54, 2, 2, 77, 10, 3, 2, 2, 2, 78, 79, 7, 100, 2, 2, 79, 80,
	7, 113, 2, 2, 80, 81, 7, 113, 2, 2, 81, 82, 7, 110, 2, 2, 82, 12, 3, 2,
	2, 2, 83, 84, 7, 101, 2, 2, 84, 85, 7, 106, 2, 2, 85, 86, 7, 99, 2, 2,
	86, 87, 7, 116, 2, 2, 87, 14, 3, 2, 2, 2, 88, 89, 7, 40, 2, 2, 89, 90,
	7, 117, 2, 2, 90, 91, 7, 118, 2, 2, 91, 92, 7, 116, 2, 2, 92, 16, 3, 2,
	2, 2, 93, 94, 7, 85, 2, 2, 94, 95, 7, 118, 2, 2, 95, 96, 7, 116, 2, 2,
	96, 97, 7, 107, 2, 2, 97, 98, 7, 112, 2, 2, 98, 99, 7, 105, 2, 2, 99, 18,
	3, 2, 2, 2, 100, 102, 9, 2, 2, 2, 101, 100, 3, 2, 2, 2, 102, 103, 3, 2,
	2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 20, 3, 2, 2, 2,
	105, 107, 9, 2, 2, 2, 106, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108,
	106, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109, 110, 3, 2, 2, 2, 110, 112,
	7, 48, 2, 2, 111, 113, 9, 2, 2, 2, 112, 111, 3, 2, 2, 2, 113, 114, 3, 2,
	2, 2, 114, 112, 3, 2, 2, 2, 114, 115, 3, 2, 2, 2, 115, 22, 3, 2, 2, 2,
	116, 120, 7, 36, 2, 2, 117, 119, 10, 3, 2, 2, 118, 117, 3, 2, 2, 2, 119,
	122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 123,
	3, 2, 2, 2, 122, 120, 3, 2, 2, 2, 123, 124, 7, 36, 2, 2, 124, 24, 3, 2,
	2, 2, 125, 126, 7, 41, 2, 2, 126, 127, 10, 4, 2, 2, 127, 128, 7, 41, 2,
	2, 128, 26, 3, 2, 2, 2, 129, 133, 9, 5, 2, 2, 130, 132, 9, 6, 2, 2, 131,
	130, 3, 2, 2, 2, 132, 135, 3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 133, 134,
	3, 2, 2, 2, 134, 28, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 136, 137, 7, 104,
	2, 2, 137, 138, 7, 99, 2, 2, 138, 139, 7, 110, 2, 2, 139, 140, 7, 117,
	2, 2, 140, 141, 7, 103, 2, 2, 141, 30, 3, 2, 2, 2, 142, 143, 7, 118, 2,
	2, 143, 144, 7, 116, 2, 2, 144, 145, 7, 119, 2, 2, 145, 146, 7, 103, 2,
	2, 146, 32, 3, 2, 2, 2, 147, 148, 7, 42, 2, 2, 148, 34, 3, 2, 2, 2, 149,
	150, 7, 43, 2, 2, 150, 36, 3, 2, 2, 2, 151, 152, 7, 60, 2, 2, 152, 38,
	3, 2, 2, 2, 153, 154, 7, 61, 2, 2, 154, 40, 3, 2, 2, 2, 155, 156, 7, 63,
	2, 2, 156, 42, 3, 2, 2, 2, 157, 158, 7, 44, 2, 2, 158, 44, 3, 2, 2, 2,
	159, 160, 7, 49, 2, 2, 160, 46, 3, 2, 2, 2, 161, 162, 7, 39, 2, 2, 162,
	48, 3, 2, 2, 2, 163, 164, 7, 45, 2, 2, 164, 50, 3, 2, 2, 2, 165, 166, 7,
	47, 2, 2, 166, 52, 3, 2, 2, 2, 167, 169, 9, 7, 2, 2, 168, 167, 3, 2, 2,
	2, 169, 170, 3, 2, 2, 2, 170, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171,
	172, 3, 2, 2, 2, 172, 173, 8, 27, 2, 2, 173, 54, 3, 2, 2, 2, 174, 175,
	7, 94, 2, 2, 175, 176, 9, 8, 2, 2, 176, 56, 3, 2, 2, 2, 9, 2, 103, 108,
	114, 120, 133, 170, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'let'", "'println!'", "'i64'", "'f64'", "'bool'", "'char'", "'&str'",
	"'String'", "", "", "", "", "", "'false'", "'true'", "'('", "')'", "':'",
	"';'", "'='", "'*'", "'/'", "'%'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "LET", "PRINTLN", "I64", "F64", "BOOL", "CHARTYPE", "STR", "STRCLASS",
	"NUMBER", "FLOAT", "STRING", "CHAR", "ID", "BFALSE", "BTRUE", "OPENPAR",
	"CLOSEPAR", "COLOM", "SEMI", "EQUALS", "MUL", "DIV", "MOD", "ADD", "SUB",
	"WHITESPACE",
}

var lexerRuleNames = []string{
	"LET", "PRINTLN", "I64", "F64", "BOOL", "CHARTYPE", "STR", "STRCLASS",
	"NUMBER", "FLOAT", "STRING", "CHAR", "ID", "BFALSE", "BTRUE", "OPENPAR",
	"CLOSEPAR", "COLOM", "SEMI", "EQUALS", "MUL", "DIV", "MOD", "ADD", "SUB",
	"WHITESPACE", "ESC_SEQ",
}

type DBRustLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewDBRustLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *DBRustLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDBRustLexer(input antlr.CharStream) *DBRustLexer {
	l := new(DBRustLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "DBRustLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// DBRustLexer tokens.
const (
	DBRustLexerLET        = 1
	DBRustLexerPRINTLN    = 2
	DBRustLexerI64        = 3
	DBRustLexerF64        = 4
	DBRustLexerBOOL       = 5
	DBRustLexerCHARTYPE   = 6
	DBRustLexerSTR        = 7
	DBRustLexerSTRCLASS   = 8
	DBRustLexerNUMBER     = 9
	DBRustLexerFLOAT      = 10
	DBRustLexerSTRING     = 11
	DBRustLexerCHAR       = 12
	DBRustLexerID         = 13
	DBRustLexerBFALSE     = 14
	DBRustLexerBTRUE      = 15
	DBRustLexerOPENPAR    = 16
	DBRustLexerCLOSEPAR   = 17
	DBRustLexerCOLOM      = 18
	DBRustLexerSEMI       = 19
	DBRustLexerEQUALS     = 20
	DBRustLexerMUL        = 21
	DBRustLexerDIV        = 22
	DBRustLexerMOD        = 23
	DBRustLexerADD        = 24
	DBRustLexerSUB        = 25
	DBRustLexerWHITESPACE = 26
)
