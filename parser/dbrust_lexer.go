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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 25, 158,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 6, 9, 87, 10, 9, 13, 9, 14, 9, 88, 3, 10,
	6, 10, 92, 10, 10, 13, 10, 14, 10, 93, 3, 10, 3, 10, 6, 10, 98, 10, 10,
	13, 10, 14, 10, 99, 3, 11, 3, 11, 7, 11, 104, 10, 11, 12, 11, 14, 11, 107,
	11, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 7, 13,
	117, 10, 13, 12, 13, 14, 13, 120, 11, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17,
	3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3,
	23, 3, 23, 3, 24, 6, 24, 150, 10, 24, 13, 24, 14, 24, 151, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 2, 2, 26, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15,
	9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33,
	18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 2, 3, 2,
	9, 3, 2, 50, 59, 3, 2, 36, 36, 3, 2, 41, 41, 5, 2, 67, 92, 97, 97, 99,
	124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 5, 2, 11, 12, 15, 15, 34, 34,
	9, 2, 34, 35, 37, 37, 45, 45, 47, 48, 60, 60, 66, 66, 93, 95, 2, 162, 2,
	3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2,
	11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2,
	2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2,
	2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2,
	2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3,
	2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 3, 51,
	3, 2, 2, 2, 5, 55, 3, 2, 2, 2, 7, 59, 3, 2, 2, 2, 9, 63, 3, 2, 2, 2, 11,
	68, 3, 2, 2, 2, 13, 73, 3, 2, 2, 2, 15, 78, 3, 2, 2, 2, 17, 86, 3, 2, 2,
	2, 19, 91, 3, 2, 2, 2, 21, 101, 3, 2, 2, 2, 23, 110, 3, 2, 2, 2, 25, 114,
	3, 2, 2, 2, 27, 121, 3, 2, 2, 2, 29, 127, 3, 2, 2, 2, 31, 132, 3, 2, 2,
	2, 33, 134, 3, 2, 2, 2, 35, 136, 3, 2, 2, 2, 37, 138, 3, 2, 2, 2, 39, 140,
	3, 2, 2, 2, 41, 142, 3, 2, 2, 2, 43, 144, 3, 2, 2, 2, 45, 146, 3, 2, 2,
	2, 47, 149, 3, 2, 2, 2, 49, 155, 3, 2, 2, 2, 51, 52, 7, 110, 2, 2, 52,
	53, 7, 103, 2, 2, 53, 54, 7, 118, 2, 2, 54, 4, 3, 2, 2, 2, 55, 56, 7, 107,
	2, 2, 56, 57, 7, 56, 2, 2, 57, 58, 7, 54, 2, 2, 58, 6, 3, 2, 2, 2, 59,
	60, 7, 104, 2, 2, 60, 61, 7, 56, 2, 2, 61, 62, 7, 54, 2, 2, 62, 8, 3, 2,
	2, 2, 63, 64, 7, 100, 2, 2, 64, 65, 7, 113, 2, 2, 65, 66, 7, 113, 2, 2,
	66, 67, 7, 110, 2, 2, 67, 10, 3, 2, 2, 2, 68, 69, 7, 101, 2, 2, 69, 70,
	7, 106, 2, 2, 70, 71, 7, 99, 2, 2, 71, 72, 7, 116, 2, 2, 72, 12, 3, 2,
	2, 2, 73, 74, 7, 40, 2, 2, 74, 75, 7, 117, 2, 2, 75, 76, 7, 118, 2, 2,
	76, 77, 7, 116, 2, 2, 77, 14, 3, 2, 2, 2, 78, 79, 7, 85, 2, 2, 79, 80,
	7, 118, 2, 2, 80, 81, 7, 116, 2, 2, 81, 82, 7, 107, 2, 2, 82, 83, 7, 112,
	2, 2, 83, 84, 7, 105, 2, 2, 84, 16, 3, 2, 2, 2, 85, 87, 9, 2, 2, 2, 86,
	85, 3, 2, 2, 2, 87, 88, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2,
	2, 89, 18, 3, 2, 2, 2, 90, 92, 9, 2, 2, 2, 91, 90, 3, 2, 2, 2, 92, 93,
	3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 94, 3, 2, 2, 2, 94, 95, 3, 2, 2, 2,
	95, 97, 7, 48, 2, 2, 96, 98, 9, 2, 2, 2, 97, 96, 3, 2, 2, 2, 98, 99, 3,
	2, 2, 2, 99, 97, 3, 2, 2, 2, 99, 100, 3, 2, 2, 2, 100, 20, 3, 2, 2, 2,
	101, 105, 7, 36, 2, 2, 102, 104, 10, 3, 2, 2, 103, 102, 3, 2, 2, 2, 104,
	107, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 108,
	3, 2, 2, 2, 107, 105, 3, 2, 2, 2, 108, 109, 7, 36, 2, 2, 109, 22, 3, 2,
	2, 2, 110, 111, 7, 41, 2, 2, 111, 112, 10, 4, 2, 2, 112, 113, 7, 41, 2,
	2, 113, 24, 3, 2, 2, 2, 114, 118, 9, 5, 2, 2, 115, 117, 9, 6, 2, 2, 116,
	115, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 26, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 122, 7, 104,
	2, 2, 122, 123, 7, 99, 2, 2, 123, 124, 7, 110, 2, 2, 124, 125, 7, 117,
	2, 2, 125, 126, 7, 103, 2, 2, 126, 28, 3, 2, 2, 2, 127, 128, 7, 118, 2,
	2, 128, 129, 7, 116, 2, 2, 129, 130, 7, 119, 2, 2, 130, 131, 7, 103, 2,
	2, 131, 30, 3, 2, 2, 2, 132, 133, 7, 60, 2, 2, 133, 32, 3, 2, 2, 2, 134,
	135, 7, 61, 2, 2, 135, 34, 3, 2, 2, 2, 136, 137, 7, 63, 2, 2, 137, 36,
	3, 2, 2, 2, 138, 139, 7, 44, 2, 2, 139, 38, 3, 2, 2, 2, 140, 141, 7, 49,
	2, 2, 141, 40, 3, 2, 2, 2, 142, 143, 7, 39, 2, 2, 143, 42, 3, 2, 2, 2,
	144, 145, 7, 45, 2, 2, 145, 44, 3, 2, 2, 2, 146, 147, 7, 47, 2, 2, 147,
	46, 3, 2, 2, 2, 148, 150, 9, 7, 2, 2, 149, 148, 3, 2, 2, 2, 150, 151, 3,
	2, 2, 2, 151, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 153, 3, 2, 2,
	2, 153, 154, 8, 24, 2, 2, 154, 48, 3, 2, 2, 2, 155, 156, 7, 94, 2, 2, 156,
	157, 9, 8, 2, 2, 157, 50, 3, 2, 2, 2, 9, 2, 88, 93, 99, 105, 118, 151,
	3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'let'", "'i64'", "'f64'", "'bool'", "'char'", "'&str'", "'String'",
	"", "", "", "", "", "'false'", "'true'", "':'", "';'", "'='", "'*'", "'/'",
	"'%'", "'+'", "'-'",
}

var lexerSymbolicNames = []string{
	"", "LET", "I64", "F64", "BOOL", "CHARTYPE", "STR", "STRCLASS", "NUMBER",
	"FLOAT", "STRING", "CHAR", "ID", "BFALSE", "BTRUE", "COLOM", "SEMI", "EQUALS",
	"MUL", "DIV", "MOD", "ADD", "SUB", "WHITESPACE",
}

var lexerRuleNames = []string{
	"LET", "I64", "F64", "BOOL", "CHARTYPE", "STR", "STRCLASS", "NUMBER", "FLOAT",
	"STRING", "CHAR", "ID", "BFALSE", "BTRUE", "COLOM", "SEMI", "EQUALS", "MUL",
	"DIV", "MOD", "ADD", "SUB", "WHITESPACE", "ESC_SEQ",
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
	DBRustLexerI64        = 2
	DBRustLexerF64        = 3
	DBRustLexerBOOL       = 4
	DBRustLexerCHARTYPE   = 5
	DBRustLexerSTR        = 6
	DBRustLexerSTRCLASS   = 7
	DBRustLexerNUMBER     = 8
	DBRustLexerFLOAT      = 9
	DBRustLexerSTRING     = 10
	DBRustLexerCHAR       = 11
	DBRustLexerID         = 12
	DBRustLexerBFALSE     = 13
	DBRustLexerBTRUE      = 14
	DBRustLexerCOLOM      = 15
	DBRustLexerSEMI       = 16
	DBRustLexerEQUALS     = 17
	DBRustLexerMUL        = 18
	DBRustLexerDIV        = 19
	DBRustLexerMOD        = 20
	DBRustLexerADD        = 21
	DBRustLexerSUB        = 22
	DBRustLexerWHITESPACE = 23
)
