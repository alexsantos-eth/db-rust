// Code generated from DBRust.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // DBRust

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import I "main/interfaces"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 296,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	7, 4, 48, 10, 4, 12, 4, 14, 4, 51, 11, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 78, 10, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 110, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 7, 8, 126,
	10, 8, 12, 8, 14, 8, 129, 11, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 144, 10, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 7, 9, 151, 10, 9, 12, 9, 14, 9, 154, 11, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10,
	178, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 5, 11, 192, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 214, 10, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 226, 10,
	13, 3, 14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 246, 10,
	16, 12, 16, 14, 16, 249, 11, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 5, 18, 290, 10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 2,
	5, 14, 16, 30, 20, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 2, 2, 2, 318, 2, 38, 3, 2, 2, 2, 4, 41, 3, 2, 2, 2, 6,
	49, 3, 2, 2, 2, 8, 77, 3, 2, 2, 2, 10, 109, 3, 2, 2, 2, 12, 111, 3, 2,
	2, 2, 14, 116, 3, 2, 2, 2, 16, 143, 3, 2, 2, 2, 18, 177, 3, 2, 2, 2, 20,
	191, 3, 2, 2, 2, 22, 213, 3, 2, 2, 2, 24, 225, 3, 2, 2, 2, 26, 227, 3,
	2, 2, 2, 28, 230, 3, 2, 2, 2, 30, 236, 3, 2, 2, 2, 32, 250, 3, 2, 2, 2,
	34, 289, 3, 2, 2, 2, 36, 291, 3, 2, 2, 2, 38, 39, 5, 6, 4, 2, 39, 40, 8,
	2, 1, 2, 40, 3, 3, 2, 2, 2, 41, 42, 7, 23, 2, 2, 42, 43, 5, 6, 4, 2, 43,
	44, 7, 24, 2, 2, 44, 45, 8, 3, 1, 2, 45, 5, 3, 2, 2, 2, 46, 48, 5, 8, 5,
	2, 47, 46, 3, 2, 2, 2, 48, 51, 3, 2, 2, 2, 49, 47, 3, 2, 2, 2, 49, 50,
	3, 2, 2, 2, 50, 52, 3, 2, 2, 2, 51, 49, 3, 2, 2, 2, 52, 53, 8, 4, 1, 2,
	53, 7, 3, 2, 2, 2, 54, 55, 5, 10, 6, 2, 55, 56, 7, 28, 2, 2, 56, 57, 8,
	5, 1, 2, 57, 78, 3, 2, 2, 2, 58, 59, 5, 24, 13, 2, 59, 60, 7, 28, 2, 2,
	60, 61, 8, 5, 1, 2, 61, 78, 3, 2, 2, 2, 62, 63, 5, 12, 7, 2, 63, 64, 7,
	28, 2, 2, 64, 65, 8, 5, 1, 2, 65, 78, 3, 2, 2, 2, 66, 67, 5, 26, 14, 2,
	67, 68, 7, 28, 2, 2, 68, 69, 8, 5, 1, 2, 69, 78, 3, 2, 2, 2, 70, 71, 5,
	34, 18, 2, 71, 72, 8, 5, 1, 2, 72, 78, 3, 2, 2, 2, 73, 74, 5, 36, 19, 2,
	74, 75, 7, 28, 2, 2, 75, 76, 8, 5, 1, 2, 76, 78, 3, 2, 2, 2, 77, 54, 3,
	2, 2, 2, 77, 58, 3, 2, 2, 2, 77, 62, 3, 2, 2, 2, 77, 66, 3, 2, 2, 2, 77,
	70, 3, 2, 2, 2, 77, 73, 3, 2, 2, 2, 78, 9, 3, 2, 2, 2, 79, 80, 7, 3, 2,
	2, 80, 81, 7, 20, 2, 2, 81, 82, 7, 27, 2, 2, 82, 83, 5, 20, 11, 2, 83,
	84, 7, 37, 2, 2, 84, 85, 5, 16, 9, 2, 85, 86, 8, 6, 1, 2, 86, 110, 3, 2,
	2, 2, 87, 88, 7, 3, 2, 2, 88, 89, 7, 4, 2, 2, 89, 90, 7, 20, 2, 2, 90,
	91, 7, 27, 2, 2, 91, 92, 5, 20, 11, 2, 92, 93, 7, 37, 2, 2, 93, 94, 5,
	16, 9, 2, 94, 95, 8, 6, 1, 2, 95, 110, 3, 2, 2, 2, 96, 97, 7, 3, 2, 2,
	97, 98, 7, 4, 2, 2, 98, 99, 7, 20, 2, 2, 99, 100, 7, 37, 2, 2, 100, 101,
	5, 16, 9, 2, 101, 102, 8, 6, 1, 2, 102, 110, 3, 2, 2, 2, 103, 104, 7, 3,
	2, 2, 104, 105, 7, 20, 2, 2, 105, 106, 7, 37, 2, 2, 106, 107, 5, 16, 9,
	2, 107, 108, 8, 6, 1, 2, 108, 110, 3, 2, 2, 2, 109, 79, 3, 2, 2, 2, 109,
	87, 3, 2, 2, 2, 109, 96, 3, 2, 2, 2, 109, 103, 3, 2, 2, 2, 110, 11, 3,
	2, 2, 2, 111, 112, 7, 20, 2, 2, 112, 113, 7, 37, 2, 2, 113, 114, 5, 16,
	9, 2, 114, 115, 8, 7, 1, 2, 115, 13, 3, 2, 2, 2, 116, 117, 8, 8, 1, 2,
	117, 118, 5, 16, 9, 2, 118, 119, 8, 8, 1, 2, 119, 127, 3, 2, 2, 2, 120,
	121, 12, 4, 2, 2, 121, 122, 7, 29, 2, 2, 122, 123, 5, 16, 9, 2, 123, 124,
	8, 8, 1, 2, 124, 126, 3, 2, 2, 2, 125, 120, 3, 2, 2, 2, 126, 129, 3, 2,
	2, 2, 127, 125, 3, 2, 2, 2, 127, 128, 3, 2, 2, 2, 128, 15, 3, 2, 2, 2,
	129, 127, 3, 2, 2, 2, 130, 131, 8, 9, 1, 2, 131, 132, 7, 36, 2, 2, 132,
	133, 5, 16, 9, 5, 133, 134, 8, 9, 1, 2, 134, 144, 3, 2, 2, 2, 135, 136,
	7, 21, 2, 2, 136, 137, 5, 16, 9, 2, 137, 138, 7, 22, 2, 2, 138, 139, 8,
	9, 1, 2, 139, 144, 3, 2, 2, 2, 140, 141, 5, 22, 12, 2, 141, 142, 8, 9,
	1, 2, 142, 144, 3, 2, 2, 2, 143, 130, 3, 2, 2, 2, 143, 135, 3, 2, 2, 2,
	143, 140, 3, 2, 2, 2, 144, 152, 3, 2, 2, 2, 145, 146, 12, 6, 2, 2, 146,
	147, 5, 18, 10, 2, 147, 148, 5, 16, 9, 7, 148, 149, 8, 9, 1, 2, 149, 151,
	3, 2, 2, 2, 150, 145, 3, 2, 2, 2, 151, 154, 3, 2, 2, 2, 152, 150, 3, 2,
	2, 2, 152, 153, 3, 2, 2, 2, 153, 17, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2,
	155, 156, 7, 40, 2, 2, 156, 178, 8, 10, 1, 2, 157, 158, 7, 41, 2, 2, 158,
	178, 8, 10, 1, 2, 159, 160, 7, 42, 2, 2, 160, 178, 8, 10, 1, 2, 161, 162,
	7, 43, 2, 2, 162, 178, 8, 10, 1, 2, 163, 164, 7, 44, 2, 2, 164, 178, 8,
	10, 1, 2, 165, 166, 7, 32, 2, 2, 166, 178, 8, 10, 1, 2, 167, 168, 7, 34,
	2, 2, 168, 178, 8, 10, 1, 2, 169, 170, 7, 35, 2, 2, 170, 178, 8, 10, 1,
	2, 171, 172, 7, 33, 2, 2, 172, 178, 8, 10, 1, 2, 173, 174, 7, 38, 2, 2,
	174, 178, 8, 10, 1, 2, 175, 176, 7, 39, 2, 2, 176, 178, 8, 10, 1, 2, 177,
	155, 3, 2, 2, 2, 177, 157, 3, 2, 2, 2, 177, 159, 3, 2, 2, 2, 177, 161,
	3, 2, 2, 2, 177, 163, 3, 2, 2, 2, 177, 165, 3, 2, 2, 2, 177, 167, 3, 2,
	2, 2, 177, 169, 3, 2, 2, 2, 177, 171, 3, 2, 2, 2, 177, 173, 3, 2, 2, 2,
	177, 175, 3, 2, 2, 2, 178, 19, 3, 2, 2, 2, 179, 180, 7, 8, 2, 2, 180, 192,
	8, 11, 1, 2, 181, 182, 7, 9, 2, 2, 182, 192, 8, 11, 1, 2, 183, 184, 7,
	10, 2, 2, 184, 192, 8, 11, 1, 2, 185, 186, 7, 11, 2, 2, 186, 192, 8, 11,
	1, 2, 187, 188, 7, 12, 2, 2, 188, 192, 8, 11, 1, 2, 189, 190, 7, 13, 2,
	2, 190, 192, 8, 11, 1, 2, 191, 179, 3, 2, 2, 2, 191, 181, 3, 2, 2, 2, 191,
	183, 3, 2, 2, 2, 191, 185, 3, 2, 2, 2, 191, 187, 3, 2, 2, 2, 191, 189,
	3, 2, 2, 2, 192, 21, 3, 2, 2, 2, 193, 194, 7, 16, 2, 2, 194, 214, 8, 12,
	1, 2, 195, 196, 7, 17, 2, 2, 196, 214, 8, 12, 1, 2, 197, 198, 7, 18, 2,
	2, 198, 214, 8, 12, 1, 2, 199, 200, 7, 19, 2, 2, 200, 214, 8, 12, 1, 2,
	201, 202, 7, 14, 2, 2, 202, 214, 8, 12, 1, 2, 203, 204, 7, 15, 2, 2, 204,
	214, 8, 12, 1, 2, 205, 206, 7, 20, 2, 2, 206, 214, 8, 12, 1, 2, 207, 208,
	5, 26, 14, 2, 208, 209, 8, 12, 1, 2, 209, 214, 3, 2, 2, 2, 210, 211, 5,
	24, 13, 2, 211, 212, 8, 12, 1, 2, 212, 214, 3, 2, 2, 2, 213, 193, 3, 2,
	2, 2, 213, 195, 3, 2, 2, 2, 213, 197, 3, 2, 2, 2, 213, 199, 3, 2, 2, 2,
	213, 201, 3, 2, 2, 2, 213, 203, 3, 2, 2, 2, 213, 205, 3, 2, 2, 2, 213,
	207, 3, 2, 2, 2, 213, 210, 3, 2, 2, 2, 214, 23, 3, 2, 2, 2, 215, 216, 7,
	20, 2, 2, 216, 217, 7, 21, 2, 2, 217, 218, 5, 14, 8, 2, 218, 219, 7, 22,
	2, 2, 219, 220, 8, 13, 1, 2, 220, 226, 3, 2, 2, 2, 221, 222, 7, 20, 2,
	2, 222, 223, 7, 21, 2, 2, 223, 224, 7, 22, 2, 2, 224, 226, 8, 13, 1, 2,
	225, 215, 3, 2, 2, 2, 225, 221, 3, 2, 2, 2, 226, 25, 3, 2, 2, 2, 227, 228,
	5, 28, 15, 2, 228, 229, 8, 14, 1, 2, 229, 27, 3, 2, 2, 2, 230, 231, 7,
	5, 2, 2, 231, 232, 7, 21, 2, 2, 232, 233, 5, 14, 8, 2, 233, 234, 7, 22,
	2, 2, 234, 235, 8, 15, 1, 2, 235, 29, 3, 2, 2, 2, 236, 237, 8, 16, 1, 2,
	237, 238, 5, 32, 17, 2, 238, 239, 8, 16, 1, 2, 239, 247, 3, 2, 2, 2, 240,
	241, 12, 4, 2, 2, 241, 242, 7, 29, 2, 2, 242, 243, 5, 32, 17, 2, 243, 244,
	8, 16, 1, 2, 244, 246, 3, 2, 2, 2, 245, 240, 3, 2, 2, 2, 246, 249, 3, 2,
	2, 2, 247, 245, 3, 2, 2, 2, 247, 248, 3, 2, 2, 2, 248, 31, 3, 2, 2, 2,
	249, 247, 3, 2, 2, 2, 250, 251, 7, 20, 2, 2, 251, 252, 7, 27, 2, 2, 252,
	253, 5, 20, 11, 2, 253, 254, 8, 17, 1, 2, 254, 33, 3, 2, 2, 2, 255, 256,
	7, 6, 2, 2, 256, 257, 7, 20, 2, 2, 257, 258, 7, 21, 2, 2, 258, 259, 5,
	30, 16, 2, 259, 260, 7, 22, 2, 2, 260, 261, 5, 4, 3, 2, 261, 262, 8, 18,
	1, 2, 262, 290, 3, 2, 2, 2, 263, 264, 7, 6, 2, 2, 264, 265, 7, 20, 2, 2,
	265, 266, 7, 21, 2, 2, 266, 267, 7, 22, 2, 2, 267, 268, 5, 4, 3, 2, 268,
	269, 8, 18, 1, 2, 269, 290, 3, 2, 2, 2, 270, 271, 7, 6, 2, 2, 271, 272,
	7, 20, 2, 2, 272, 273, 7, 21, 2, 2, 273, 274, 5, 30, 16, 2, 274, 275, 7,
	22, 2, 2, 275, 276, 7, 25, 2, 2, 276, 277, 5, 20, 11, 2, 277, 278, 5, 4,
	3, 2, 278, 279, 8, 18, 1, 2, 279, 290, 3, 2, 2, 2, 280, 281, 7, 6, 2, 2,
	281, 282, 7, 20, 2, 2, 282, 283, 7, 21, 2, 2, 283, 284, 7, 22, 2, 2, 284,
	285, 7, 25, 2, 2, 285, 286, 5, 20, 11, 2, 286, 287, 5, 4, 3, 2, 287, 288,
	8, 18, 1, 2, 288, 290, 3, 2, 2, 2, 289, 255, 3, 2, 2, 2, 289, 263, 3, 2,
	2, 2, 289, 270, 3, 2, 2, 2, 289, 280, 3, 2, 2, 2, 290, 35, 3, 2, 2, 2,
	291, 292, 7, 7, 2, 2, 292, 293, 5, 16, 9, 2, 293, 294, 8, 19, 1, 2, 294,
	37, 3, 2, 2, 2, 14, 49, 77, 109, 127, 143, 152, 177, 191, 213, 225, 247,
	289,
}
var literalNames = []string{
	"", "'let'", "'mut'", "'println!'", "'fn'", "'return'", "'i64'", "'f64'",
	"'bool'", "'char'", "'&str'", "'String'", "'false'", "'true'", "", "",
	"", "", "", "'('", "')'", "'{'", "'}'", "'->'", "'.'", "':'", "';'", "','",
	"'&&'", "'||'", "'!='", "'=='", "'>='", "'<='", "'!'", "'='", "'>'", "'<'",
	"'*'", "'/'", "'%'", "'+'", "'-'",
}
var symbolicNames = []string{
	"", "LET", "MUT", "PRINTLN", "FN", "RETURN", "I64", "F64", "BOOL", "CHARTYPE",
	"STR", "STRCLASS", "BFALSE", "BTRUE", "NUMBER", "FLOAT", "STRING", "CHAR",
	"ID", "OPENPAR", "CLOSEPAR", "OPENBRACKET", "CLOSEBRACKET", "ARROW", "DOT",
	"COLOM", "SEMI", "COMMA", "AND", "OR", "NOTEQUALS", "EQUALSEQUALS", "MOREOREQUALS",
	"LESSOREQUALS", "NOT", "EQUALS", "MAJOR", "MINOR", "MUL", "DIV", "MOD",
	"ADD", "SUB", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instructionsBlock", "instructions", "instruction", "declaration",
	"assignment", "expList", "expression", "expOp", "valueType", "value", "functionCall",
	"methods", "printlnCall", "paramList", "param", "function", "returnValue",
}

type DBRustParser struct {
	*antlr.BaseParser
}

// NewDBRustParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *DBRustParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewDBRustParser(input antlr.TokenStream) *DBRustParser {
	this := new(DBRustParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "DBRust.g4"

	return this
}

// DBRustParser tokens.
const (
	DBRustParserEOF          = antlr.TokenEOF
	DBRustParserLET          = 1
	DBRustParserMUT          = 2
	DBRustParserPRINTLN      = 3
	DBRustParserFN           = 4
	DBRustParserRETURN       = 5
	DBRustParserI64          = 6
	DBRustParserF64          = 7
	DBRustParserBOOL         = 8
	DBRustParserCHARTYPE     = 9
	DBRustParserSTR          = 10
	DBRustParserSTRCLASS     = 11
	DBRustParserBFALSE       = 12
	DBRustParserBTRUE        = 13
	DBRustParserNUMBER       = 14
	DBRustParserFLOAT        = 15
	DBRustParserSTRING       = 16
	DBRustParserCHAR         = 17
	DBRustParserID           = 18
	DBRustParserOPENPAR      = 19
	DBRustParserCLOSEPAR     = 20
	DBRustParserOPENBRACKET  = 21
	DBRustParserCLOSEBRACKET = 22
	DBRustParserARROW        = 23
	DBRustParserDOT          = 24
	DBRustParserCOLOM        = 25
	DBRustParserSEMI         = 26
	DBRustParserCOMMA        = 27
	DBRustParserAND          = 28
	DBRustParserOR           = 29
	DBRustParserNOTEQUALS    = 30
	DBRustParserEQUALSEQUALS = 31
	DBRustParserMOREOREQUALS = 32
	DBRustParserLESSOREQUALS = 33
	DBRustParserNOT          = 34
	DBRustParserEQUALS       = 35
	DBRustParserMAJOR        = 36
	DBRustParserMINOR        = 37
	DBRustParserMUL          = 38
	DBRustParserDIV          = 39
	DBRustParserMOD          = 40
	DBRustParserADD          = 41
	DBRustParserSUB          = 42
	DBRustParserWHITESPACE   = 43
)

// DBRustParser rules.
const (
	DBRustParserRULE_start             = 0
	DBRustParserRULE_instructionsBlock = 1
	DBRustParserRULE_instructions      = 2
	DBRustParserRULE_instruction       = 3
	DBRustParserRULE_declaration       = 4
	DBRustParserRULE_assignment        = 5
	DBRustParserRULE_expList           = 6
	DBRustParserRULE_expression        = 7
	DBRustParserRULE_expOp             = 8
	DBRustParserRULE_valueType         = 9
	DBRustParserRULE_value             = 10
	DBRustParserRULE_functionCall      = 11
	DBRustParserRULE_methods           = 12
	DBRustParserRULE_printlnCall       = 13
	DBRustParserRULE_paramList         = 14
	DBRustParserRULE_param             = 15
	DBRustParserRULE_function          = 16
	DBRustParserRULE_returnValue       = 17
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetList returns the list attribute.
	GetList() *arrayList.List

	// SetList sets the list attribute.
	SetList(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	list          *arrayList.List
	_instructions IInstructionsContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *StartContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *StartContext) GetList() *arrayList.List { return s.list }

func (s *StartContext) SetList(v *arrayList.List) { s.list = v }

func (s *StartContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *DBRustParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, DBRustParserRULE_start)

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
		p.SetState(36)

		var _x = p.Instructions()

		localctx.(*StartContext)._instructions = _x
	}
	localctx.(*StartContext).list = localctx.(*StartContext).Get_instructions().GetL()

	return localctx
}

// IInstructionsBlockContext is an interface to support dynamic dispatch.
type IInstructionsBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instructions returns the _instructions rule contexts.
	Get_instructions() IInstructionsContext

	// Set_instructions sets the _instructions rule contexts.
	Set_instructions(IInstructionsContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstructionsBlockContext differentiates from other interfaces.
	IsInstructionsBlockContext()
}

type InstructionsBlockContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	l             *arrayList.List
	_instructions IInstructionsContext
}

func NewEmptyInstructionsBlockContext() *InstructionsBlockContext {
	var p = new(InstructionsBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_instructionsBlock
	return p
}

func (*InstructionsBlockContext) IsInstructionsBlockContext() {}

func NewInstructionsBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsBlockContext {
	var p = new(InstructionsBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_instructionsBlock

	return p
}

func (s *InstructionsBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsBlockContext) Get_instructions() IInstructionsContext { return s._instructions }

func (s *InstructionsBlockContext) Set_instructions(v IInstructionsContext) { s._instructions = v }

func (s *InstructionsBlockContext) GetL() *arrayList.List { return s.l }

func (s *InstructionsBlockContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstructionsBlockContext) OPENBRACKET() antlr.TerminalNode {
	return s.GetToken(DBRustParserOPENBRACKET, 0)
}

func (s *InstructionsBlockContext) Instructions() IInstructionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsContext)
}

func (s *InstructionsBlockContext) CLOSEBRACKET() antlr.TerminalNode {
	return s.GetToken(DBRustParserCLOSEBRACKET, 0)
}

func (s *InstructionsBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterInstructionsBlock(s)
	}
}

func (s *InstructionsBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitInstructionsBlock(s)
	}
}

func (p *DBRustParser) InstructionsBlock() (localctx IInstructionsBlockContext) {
	localctx = NewInstructionsBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, DBRustParserRULE_instructionsBlock)

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
		p.SetState(39)
		p.Match(DBRustParserOPENBRACKET)
	}
	{
		p.SetState(40)

		var _x = p.Instructions()

		localctx.(*InstructionsBlockContext)._instructions = _x
	}
	{
		p.SetState(41)
		p.Match(DBRustParserCLOSEBRACKET)
	}

	localctx.(*InstructionsBlockContext).l = localctx.(*InstructionsBlockContext).Get_instructions().GetL()

	return localctx
}

// IInstructionsContext is an interface to support dynamic dispatch.
type IInstructionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruction returns the _instruction rule contexts.
	Get_instruction() IInstructionContext

	// Set_instruction sets the _instruction rule contexts.
	Set_instruction(IInstructionContext)

	// GetE returns the e rule context list.
	GetE() []IInstructionContext

	// SetE sets the e rule context list.
	SetE([]IInstructionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstructionsContext differentiates from other interfaces.
	IsInstructionsContext()
}

type InstructionsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruction IInstructionContext
	e            []IInstructionContext
}

func NewEmptyInstructionsContext() *InstructionsContext {
	var p = new(InstructionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_instructions
	return p
}

func (*InstructionsContext) IsInstructionsContext() {}

func NewInstructionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionsContext {
	var p = new(InstructionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_instructions

	return p
}

func (s *InstructionsContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionsContext) Get_instruction() IInstructionContext { return s._instruction }

func (s *InstructionsContext) Set_instruction(v IInstructionContext) { s._instruction = v }

func (s *InstructionsContext) GetE() []IInstructionContext { return s.e }

func (s *InstructionsContext) SetE(v []IInstructionContext) { s.e = v }

func (s *InstructionsContext) GetL() *arrayList.List { return s.l }

func (s *InstructionsContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstructionsContext) AllInstruction() []IInstructionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstructionContext)(nil)).Elem())
	var tst = make([]IInstructionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstructionContext)
		}
	}

	return tst
}

func (s *InstructionsContext) Instruction(i int) IInstructionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *InstructionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterInstructions(s)
	}
}

func (s *InstructionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitInstructions(s)
	}
}

func (p *DBRustParser) Instructions() (localctx IInstructionsContext) {
	localctx = NewInstructionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, DBRustParserRULE_instructions)

	localctx.(*InstructionsContext).l = arrayList.New()

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
	p.SetState(47)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<DBRustParserLET)|(1<<DBRustParserPRINTLN)|(1<<DBRustParserFN)|(1<<DBRustParserRETURN)|(1<<DBRustParserID))) != 0 {
		{
			p.SetState(44)

			var _x = p.Instruction()

			localctx.(*InstructionsContext)._instruction = _x
		}
		localctx.(*InstructionsContext).e = append(localctx.(*InstructionsContext).e, localctx.(*InstructionsContext)._instruction)

		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstructionsContext).GetE()
	for _, e := range listInt {
		localctx.(*InstructionsContext).l.Add(e.GetState())
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetDecltn returns the decltn rule contexts.
	GetDecltn() IDeclarationContext

	// GetCalls returns the calls rule contexts.
	GetCalls() IFunctionCallContext

	// GetAssign returns the assign rule contexts.
	GetAssign() IAssignmentContext

	// GetMth returns the mth rule contexts.
	GetMth() IMethodsContext

	// GetFn returns the fn rule contexts.
	GetFn() IFunctionContext

	// GetRtn returns the rtn rule contexts.
	GetRtn() IReturnValueContext

	// SetDecltn sets the decltn rule contexts.
	SetDecltn(IDeclarationContext)

	// SetCalls sets the calls rule contexts.
	SetCalls(IFunctionCallContext)

	// SetAssign sets the assign rule contexts.
	SetAssign(IAssignmentContext)

	// SetMth sets the mth rule contexts.
	SetMth(IMethodsContext)

	// SetFn sets the fn rule contexts.
	SetFn(IFunctionContext)

	// SetRtn sets the rtn rule contexts.
	SetRtn(IReturnValueContext)

	// GetState returns the state attribute.
	GetState() I.IInstruction

	// SetState sets the state attribute.
	SetState(I.IInstruction)

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	state  I.IInstruction
	decltn IDeclarationContext
	calls  IFunctionCallContext
	assign IAssignmentContext
	mth    IMethodsContext
	fn     IFunctionContext
	rtn    IReturnValueContext
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) GetDecltn() IDeclarationContext { return s.decltn }

func (s *InstructionContext) GetCalls() IFunctionCallContext { return s.calls }

func (s *InstructionContext) GetAssign() IAssignmentContext { return s.assign }

func (s *InstructionContext) GetMth() IMethodsContext { return s.mth }

func (s *InstructionContext) GetFn() IFunctionContext { return s.fn }

func (s *InstructionContext) GetRtn() IReturnValueContext { return s.rtn }

func (s *InstructionContext) SetDecltn(v IDeclarationContext) { s.decltn = v }

func (s *InstructionContext) SetCalls(v IFunctionCallContext) { s.calls = v }

func (s *InstructionContext) SetAssign(v IAssignmentContext) { s.assign = v }

func (s *InstructionContext) SetMth(v IMethodsContext) { s.mth = v }

func (s *InstructionContext) SetFn(v IFunctionContext) { s.fn = v }

func (s *InstructionContext) SetRtn(v IReturnValueContext) { s.rtn = v }

func (s *InstructionContext) GetState() I.IInstruction { return s.state }

func (s *InstructionContext) SetState(v I.IInstruction) { s.state = v }

func (s *InstructionContext) SEMI() antlr.TerminalNode {
	return s.GetToken(DBRustParserSEMI, 0)
}

func (s *InstructionContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *InstructionContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *InstructionContext) Assignment() IAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAssignmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAssignmentContext)
}

func (s *InstructionContext) Methods() IMethodsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodsContext)
}

func (s *InstructionContext) Function() IFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *InstructionContext) ReturnValue() IReturnValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReturnValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReturnValueContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterInstruction(s)
	}
}

func (s *InstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitInstruction(s)
	}
}

func (p *DBRustParser) Instruction() (localctx IInstructionContext) {
	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, DBRustParserRULE_instruction)

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

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)

			var _x = p.Declaration()

			localctx.(*InstructionContext).decltn = _x
		}
		{
			p.SetState(53)
			p.Match(DBRustParserSEMI)
		}
		localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetDecltn().GetState()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)

			var _x = p.FunctionCall()

			localctx.(*InstructionContext).calls = _x
		}
		{
			p.SetState(57)
			p.Match(DBRustParserSEMI)
		}
		localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetCalls().GetState()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(60)

			var _x = p.Assignment()

			localctx.(*InstructionContext).assign = _x
		}
		{
			p.SetState(61)
			p.Match(DBRustParserSEMI)
		}
		localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetAssign().GetState()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(64)

			var _x = p.Methods()

			localctx.(*InstructionContext).mth = _x
		}
		{
			p.SetState(65)
			p.Match(DBRustParserSEMI)
		}
		localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetMth().GetState()

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(68)

			var _x = p.Function()

			localctx.(*InstructionContext).fn = _x
		}
		localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetFn().GetState()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(71)

			var _x = p.ReturnValue()

			localctx.(*InstructionContext).rtn = _x
		}
		{
			p.SetState(72)
			p.Match(DBRustParserSEMI)
		}
		localctx.(*InstructionContext).state = localctx.(*InstructionContext).GetRtn().GetState()

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_valueType returns the _valueType rule contexts.
	Get_valueType() IValueTypeContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_valueType sets the _valueType rule contexts.
	Set_valueType(IValueTypeContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetState returns the state attribute.
	GetState() I.Declaration

	// SetState sets the state attribute.
	SetState(I.Declaration)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	state       I.Declaration
	_ID         antlr.Token
	_valueType  IValueTypeContext
	_expression IExpressionContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) Get_ID() antlr.Token { return s._ID }

func (s *DeclarationContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *DeclarationContext) Get_valueType() IValueTypeContext { return s._valueType }

func (s *DeclarationContext) Get_expression() IExpressionContext { return s._expression }

func (s *DeclarationContext) Set_valueType(v IValueTypeContext) { s._valueType = v }

func (s *DeclarationContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DeclarationContext) GetState() I.Declaration { return s.state }

func (s *DeclarationContext) SetState(v I.Declaration) { s.state = v }

func (s *DeclarationContext) LET() antlr.TerminalNode {
	return s.GetToken(DBRustParserLET, 0)
}

func (s *DeclarationContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *DeclarationContext) COLOM() antlr.TerminalNode {
	return s.GetToken(DBRustParserCOLOM, 0)
}

func (s *DeclarationContext) ValueType() IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *DeclarationContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserEQUALS, 0)
}

func (s *DeclarationContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DeclarationContext) MUT() antlr.TerminalNode {
	return s.GetToken(DBRustParserMUT, 0)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (p *DBRustParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, DBRustParserRULE_declaration)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(77)
			p.Match(DBRustParserLET)
		}
		{
			p.SetState(78)

			var _m = p.Match(DBRustParserID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(79)
			p.Match(DBRustParserCOLOM)
		}
		{
			p.SetState(80)

			var _x = p.ValueType()

			localctx.(*DeclarationContext)._valueType = _x
		}
		{
			p.SetState(81)
			p.Match(DBRustParserEQUALS)
		}
		{
			p.SetState(82)

			var _x = p.expression(0)

			localctx.(*DeclarationContext)._expression = _x
		}

		expPoint := localctx.(*DeclarationContext).Get_expression().GetState()
		localctx.(*DeclarationContext).state = I.Declaration{
			Instruction: I.Instruction{"Declaration"},
			Mut:         false,
			Type:        localctx.(*DeclarationContext).Get_valueType().GetState(),
			Id: (func() string {
				if localctx.(*DeclarationContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*DeclarationContext).Get_ID().GetText()
				}
			}()),
			Expression: &expPoint,
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(85)
			p.Match(DBRustParserLET)
		}
		{
			p.SetState(86)
			p.Match(DBRustParserMUT)
		}
		{
			p.SetState(87)

			var _m = p.Match(DBRustParserID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(88)
			p.Match(DBRustParserCOLOM)
		}
		{
			p.SetState(89)

			var _x = p.ValueType()

			localctx.(*DeclarationContext)._valueType = _x
		}
		{
			p.SetState(90)
			p.Match(DBRustParserEQUALS)
		}
		{
			p.SetState(91)

			var _x = p.expression(0)

			localctx.(*DeclarationContext)._expression = _x
		}

		expPoint := localctx.(*DeclarationContext).Get_expression().GetState()
		localctx.(*DeclarationContext).state = I.Declaration{
			Instruction: I.Instruction{"Declaration"},
			Mut:         true,
			Type:        localctx.(*DeclarationContext).Get_valueType().GetState(),
			Id: (func() string {
				if localctx.(*DeclarationContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*DeclarationContext).Get_ID().GetText()
				}
			}()),
			Expression: &expPoint,
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(94)
			p.Match(DBRustParserLET)
		}
		{
			p.SetState(95)
			p.Match(DBRustParserMUT)
		}
		{
			p.SetState(96)

			var _m = p.Match(DBRustParserID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(97)
			p.Match(DBRustParserEQUALS)
		}
		{
			p.SetState(98)

			var _x = p.expression(0)

			localctx.(*DeclarationContext)._expression = _x
		}

		expPoint := localctx.(*DeclarationContext).Get_expression().GetState()
		localctx.(*DeclarationContext).state = I.Declaration{
			Instruction: I.Instruction{"Declaration"},
			Mut:         true,
			Type:        I.UNDEF,
			Id: (func() string {
				if localctx.(*DeclarationContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*DeclarationContext).Get_ID().GetText()
				}
			}()),
			Expression: &expPoint,
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(101)
			p.Match(DBRustParserLET)
		}
		{
			p.SetState(102)

			var _m = p.Match(DBRustParserID)

			localctx.(*DeclarationContext)._ID = _m
		}
		{
			p.SetState(103)
			p.Match(DBRustParserEQUALS)
		}
		{
			p.SetState(104)

			var _x = p.expression(0)

			localctx.(*DeclarationContext)._expression = _x
		}

		expPoint := localctx.(*DeclarationContext).Get_expression().GetState()
		localctx.(*DeclarationContext).state = I.Declaration{
			Instruction: I.Instruction{"Declaration"},
			Mut:         true,
			Type:        I.UNDEF,
			Id: (func() string {
				if localctx.(*DeclarationContext).Get_ID() == nil {
					return ""
				} else {
					return localctx.(*DeclarationContext).Get_ID().GetText()
				}
			}()),
			Expression: &expPoint,
		}

	}

	return localctx
}

// IAssignmentContext is an interface to support dynamic dispatch.
type IAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetState returns the state attribute.
	GetState() I.Assignment

	// SetState sets the state attribute.
	SetState(I.Assignment)

	// IsAssignmentContext differentiates from other interfaces.
	IsAssignmentContext()
}

type AssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	state       I.Assignment
	_ID         antlr.Token
	_expression IExpressionContext
}

func NewEmptyAssignmentContext() *AssignmentContext {
	var p = new(AssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_assignment
	return p
}

func (*AssignmentContext) IsAssignmentContext() {}

func NewAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentContext {
	var p = new(AssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_assignment

	return p
}

func (s *AssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentContext) Get_ID() antlr.Token { return s._ID }

func (s *AssignmentContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *AssignmentContext) Get_expression() IExpressionContext { return s._expression }

func (s *AssignmentContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *AssignmentContext) GetState() I.Assignment { return s.state }

func (s *AssignmentContext) SetState(v I.Assignment) { s.state = v }

func (s *AssignmentContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *AssignmentContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserEQUALS, 0)
}

func (s *AssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterAssignment(s)
	}
}

func (s *AssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitAssignment(s)
	}
}

func (p *DBRustParser) Assignment() (localctx IAssignmentContext) {
	localctx = NewAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, DBRustParserRULE_assignment)

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
		p.SetState(109)

		var _m = p.Match(DBRustParserID)

		localctx.(*AssignmentContext)._ID = _m
	}
	{
		p.SetState(110)
		p.Match(DBRustParserEQUALS)
	}
	{
		p.SetState(111)

		var _x = p.expression(0)

		localctx.(*AssignmentContext)._expression = _x
	}

	expPoint := localctx.(*AssignmentContext).Get_expression().GetState()
	localctx.(*AssignmentContext).state = I.Assignment{
		Instruction: I.Instruction{"Assignment"},
		Id: (func() string {
			if localctx.(*AssignmentContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*AssignmentContext).Get_ID().GetText()
			}
		}()),
		Expression: &expPoint,
	}

	return localctx
}

// IExpListContext is an interface to support dynamic dispatch.
type IExpListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IExpListContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IExpListContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsExpListContext differentiates from other interfaces.
	IsExpListContext()
}

type ExpListContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	list        IExpListContext
	_expression IExpressionContext
}

func NewEmptyExpListContext() *ExpListContext {
	var p = new(ExpListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_expList
	return p
}

func (*ExpListContext) IsExpListContext() {}

func NewExpListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpListContext {
	var p = new(ExpListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_expList

	return p
}

func (s *ExpListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpListContext) GetList() IExpListContext { return s.list }

func (s *ExpListContext) Get_expression() IExpressionContext { return s._expression }

func (s *ExpListContext) SetList(v IExpListContext) { s.list = v }

func (s *ExpListContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ExpListContext) GetL() *arrayList.List { return s.l }

func (s *ExpListContext) SetL(v *arrayList.List) { s.l = v }

func (s *ExpListContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(DBRustParserCOMMA, 0)
}

func (s *ExpListContext) ExpList() IExpListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpListContext)
}

func (s *ExpListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterExpList(s)
	}
}

func (s *ExpListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitExpList(s)
	}
}

func (p *DBRustParser) ExpList() (localctx IExpListContext) {
	return p.expList(0)
}

func (p *DBRustParser) expList(_p int) (localctx IExpListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 12
	p.EnterRecursionRule(localctx, 12, DBRustParserRULE_expList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)

		var _x = p.expression(0)

		localctx.(*ExpListContext)._expression = _x
	}

	localctx.(*ExpListContext).l = arrayList.New()
	localctx.(*ExpListContext).l.Add(localctx.(*ExpListContext).Get_expression().GetState())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(125)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpListContext(p, _parentctx, _parentState)
			localctx.(*ExpListContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, DBRustParserRULE_expList)
			p.SetState(118)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(119)
				p.Match(DBRustParserCOMMA)
			}
			{
				p.SetState(120)

				var _x = p.expression(0)

				localctx.(*ExpListContext)._expression = _x
			}

			localctx.(*ExpListContext).GetList().GetL().Add(localctx.(*ExpListContext).Get_expression().GetState())
			localctx.(*ExpListContext).l = localctx.(*ExpListContext).GetList().GetL()

		}
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLeftExp returns the leftExp rule contexts.
	GetLeftExp() IExpressionContext

	// GetExp returns the exp rule contexts.
	GetExp() IExpressionContext

	// Get_value returns the _value rule contexts.
	Get_value() IValueContext

	// Get_expOp returns the _expOp rule contexts.
	Get_expOp() IExpOpContext

	// GetRightExp returns the rightExp rule contexts.
	GetRightExp() IExpressionContext

	// SetLeftExp sets the leftExp rule contexts.
	SetLeftExp(IExpressionContext)

	// SetExp sets the exp rule contexts.
	SetExp(IExpressionContext)

	// Set_value sets the _value rule contexts.
	Set_value(IValueContext)

	// Set_expOp sets the _expOp rule contexts.
	Set_expOp(IExpOpContext)

	// SetRightExp sets the rightExp rule contexts.
	SetRightExp(IExpressionContext)

	// GetState returns the state attribute.
	GetState() I.Expression

	// SetState sets the state attribute.
	SetState(I.Expression)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	state    I.Expression
	leftExp  IExpressionContext
	exp      IExpressionContext
	_value   IValueContext
	_expOp   IExpOpContext
	rightExp IExpressionContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) GetLeftExp() IExpressionContext { return s.leftExp }

func (s *ExpressionContext) GetExp() IExpressionContext { return s.exp }

func (s *ExpressionContext) Get_value() IValueContext { return s._value }

func (s *ExpressionContext) Get_expOp() IExpOpContext { return s._expOp }

func (s *ExpressionContext) GetRightExp() IExpressionContext { return s.rightExp }

func (s *ExpressionContext) SetLeftExp(v IExpressionContext) { s.leftExp = v }

func (s *ExpressionContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *ExpressionContext) Set_value(v IValueContext) { s._value = v }

func (s *ExpressionContext) Set_expOp(v IExpOpContext) { s._expOp = v }

func (s *ExpressionContext) SetRightExp(v IExpressionContext) { s.rightExp = v }

func (s *ExpressionContext) GetState() I.Expression { return s.state }

func (s *ExpressionContext) SetState(v I.Expression) { s.state = v }

func (s *ExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(DBRustParserNOT, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) OPENPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserOPENPAR, 0)
}

func (s *ExpressionContext) CLOSEPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserCLOSEPAR, 0)
}

func (s *ExpressionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ExpressionContext) ExpOp() IExpOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpOpContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *DBRustParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *DBRustParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, DBRustParserRULE_expression, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DBRustParserNOT:
		{
			p.SetState(129)
			p.Match(DBRustParserNOT)
		}
		{
			p.SetState(130)

			var _x = p.expression(3)

			localctx.(*ExpressionContext).exp = _x
		}

		exp := localctx.(*ExpressionContext).GetExp().GetState()
		localctx.(*ExpressionContext).state = I.Expression{
			Value:     nil,
			Left:      &exp,
			Right:     nil,
			Operation: I.UNOT,
		}

	case DBRustParserOPENPAR:
		{
			p.SetState(133)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(134)

			var _x = p.expression(0)

			localctx.(*ExpressionContext).exp = _x
		}
		{
			p.SetState(135)
			p.Match(DBRustParserCLOSEPAR)
		}

		exp := localctx.(*ExpressionContext).GetExp().GetState()
		localctx.(*ExpressionContext).state = I.Expression{
			Value:     nil,
			Left:      &exp,
			Right:     nil,
			Operation: I.NOOP,
		}

	case DBRustParserPRINTLN, DBRustParserBFALSE, DBRustParserBTRUE, DBRustParserNUMBER, DBRustParserFLOAT, DBRustParserSTRING, DBRustParserCHAR, DBRustParserID:
		{
			p.SetState(138)

			var _x = p.Value()

			localctx.(*ExpressionContext)._value = _x
		}

		sym := localctx.(*ExpressionContext).Get_value().GetState()
		localctx.(*ExpressionContext).state = I.Expression{
			Value:     &sym,
			Left:      nil,
			Right:     nil,
			Operation: I.NOOP,
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(150)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			localctx.(*ExpressionContext).leftExp = _prevctx
			p.PushNewRecursionContext(localctx, _startState, DBRustParserRULE_expression)
			p.SetState(143)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(144)

				var _x = p.ExpOp()

				localctx.(*ExpressionContext)._expOp = _x
			}
			{
				p.SetState(145)

				var _x = p.expression(5)

				localctx.(*ExpressionContext).rightExp = _x
			}

			left, right := localctx.(*ExpressionContext).GetLeftExp().GetState(), localctx.(*ExpressionContext).GetRightExp().GetState()
			localctx.(*ExpressionContext).state = I.Expression{
				Value:     nil,
				Left:      &left,
				Right:     &right,
				Operation: localctx.(*ExpressionContext).Get_expOp().GetState(),
			}

		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

// IExpOpContext is an interface to support dynamic dispatch.
type IExpOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetState returns the state attribute.
	GetState() I.Operation

	// SetState sets the state attribute.
	SetState(I.Operation)

	// IsExpOpContext differentiates from other interfaces.
	IsExpOpContext()
}

type ExpOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	state  I.Operation
}

func NewEmptyExpOpContext() *ExpOpContext {
	var p = new(ExpOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_expOp
	return p
}

func (*ExpOpContext) IsExpOpContext() {}

func NewExpOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpOpContext {
	var p = new(ExpOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_expOp

	return p
}

func (s *ExpOpContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpOpContext) GetState() I.Operation { return s.state }

func (s *ExpOpContext) SetState(v I.Operation) { s.state = v }

func (s *ExpOpContext) MUL() antlr.TerminalNode {
	return s.GetToken(DBRustParserMUL, 0)
}

func (s *ExpOpContext) DIV() antlr.TerminalNode {
	return s.GetToken(DBRustParserDIV, 0)
}

func (s *ExpOpContext) MOD() antlr.TerminalNode {
	return s.GetToken(DBRustParserMOD, 0)
}

func (s *ExpOpContext) ADD() antlr.TerminalNode {
	return s.GetToken(DBRustParserADD, 0)
}

func (s *ExpOpContext) SUB() antlr.TerminalNode {
	return s.GetToken(DBRustParserSUB, 0)
}

func (s *ExpOpContext) NOTEQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserNOTEQUALS, 0)
}

func (s *ExpOpContext) MOREOREQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserMOREOREQUALS, 0)
}

func (s *ExpOpContext) LESSOREQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserLESSOREQUALS, 0)
}

func (s *ExpOpContext) EQUALSEQUALS() antlr.TerminalNode {
	return s.GetToken(DBRustParserEQUALSEQUALS, 0)
}

func (s *ExpOpContext) MAJOR() antlr.TerminalNode {
	return s.GetToken(DBRustParserMAJOR, 0)
}

func (s *ExpOpContext) MINOR() antlr.TerminalNode {
	return s.GetToken(DBRustParserMINOR, 0)
}

func (s *ExpOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterExpOp(s)
	}
}

func (s *ExpOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitExpOp(s)
	}
}

func (p *DBRustParser) ExpOp() (localctx IExpOpContext) {
	localctx = NewExpOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, DBRustParserRULE_expOp)

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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DBRustParserMUL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(DBRustParserMUL)
		}

		localctx.(*ExpOpContext).state = I.MUL

	case DBRustParserDIV:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(155)
			p.Match(DBRustParserDIV)
		}

		localctx.(*ExpOpContext).state = I.DIV

	case DBRustParserMOD:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.Match(DBRustParserMOD)
		}

		localctx.(*ExpOpContext).state = I.MOD

	case DBRustParserADD:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(159)
			p.Match(DBRustParserADD)
		}

		localctx.(*ExpOpContext).state = I.ADD

	case DBRustParserSUB:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.Match(DBRustParserSUB)
		}

		localctx.(*ExpOpContext).state = I.SUB

	case DBRustParserNOTEQUALS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(163)
			p.Match(DBRustParserNOTEQUALS)
		}

		localctx.(*ExpOpContext).state = I.NOTEQUALS

	case DBRustParserMOREOREQUALS:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(165)
			p.Match(DBRustParserMOREOREQUALS)
		}

		localctx.(*ExpOpContext).state = I.MOREOREQUALS

	case DBRustParserLESSOREQUALS:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(167)
			p.Match(DBRustParserLESSOREQUALS)
		}

		localctx.(*ExpOpContext).state = I.LESSOREQUALS

	case DBRustParserEQUALSEQUALS:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.Match(DBRustParserEQUALSEQUALS)
		}

		localctx.(*ExpOpContext).state = I.EQUALSEQUALS

	case DBRustParserMAJOR:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(171)
			p.Match(DBRustParserMAJOR)
		}

		localctx.(*ExpOpContext).state = I.MAJOR

	case DBRustParserMINOR:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(173)
			p.Match(DBRustParserMINOR)
		}

		localctx.(*ExpOpContext).state = I.MINOR

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueTypeContext is an interface to support dynamic dispatch.
type IValueTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetState returns the state attribute.
	GetState() I.ValueType

	// SetState sets the state attribute.
	SetState(I.ValueType)

	// IsValueTypeContext differentiates from other interfaces.
	IsValueTypeContext()
}

type ValueTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	state  I.ValueType
}

func NewEmptyValueTypeContext() *ValueTypeContext {
	var p = new(ValueTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_valueType
	return p
}

func (*ValueTypeContext) IsValueTypeContext() {}

func NewValueTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueTypeContext {
	var p = new(ValueTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_valueType

	return p
}

func (s *ValueTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueTypeContext) GetState() I.ValueType { return s.state }

func (s *ValueTypeContext) SetState(v I.ValueType) { s.state = v }

func (s *ValueTypeContext) I64() antlr.TerminalNode {
	return s.GetToken(DBRustParserI64, 0)
}

func (s *ValueTypeContext) F64() antlr.TerminalNode {
	return s.GetToken(DBRustParserF64, 0)
}

func (s *ValueTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(DBRustParserBOOL, 0)
}

func (s *ValueTypeContext) CHARTYPE() antlr.TerminalNode {
	return s.GetToken(DBRustParserCHARTYPE, 0)
}

func (s *ValueTypeContext) STR() antlr.TerminalNode {
	return s.GetToken(DBRustParserSTR, 0)
}

func (s *ValueTypeContext) STRCLASS() antlr.TerminalNode {
	return s.GetToken(DBRustParserSTRCLASS, 0)
}

func (s *ValueTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterValueType(s)
	}
}

func (s *ValueTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitValueType(s)
	}
}

func (p *DBRustParser) ValueType() (localctx IValueTypeContext) {
	localctx = NewValueTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, DBRustParserRULE_valueType)

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

	p.SetState(189)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case DBRustParserI64:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(DBRustParserI64)
		}

		localctx.(*ValueTypeContext).state = I.INTEGER

	case DBRustParserF64:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.Match(DBRustParserF64)
		}

		localctx.(*ValueTypeContext).state = I.FLOAT

	case DBRustParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(181)
			p.Match(DBRustParserBOOL)
		}

		localctx.(*ValueTypeContext).state = I.BOOL

	case DBRustParserCHARTYPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(183)
			p.Match(DBRustParserCHARTYPE)
		}

		localctx.(*ValueTypeContext).state = I.CHAR

	case DBRustParserSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(185)
			p.Match(DBRustParserSTR)
		}

		localctx.(*ValueTypeContext).state = I.STR

	case DBRustParserSTRCLASS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(187)
			p.Match(DBRustParserSTRCLASS)
		}

		localctx.(*ValueTypeContext).state = I.STRING

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_FLOAT returns the _FLOAT token.
	Get_FLOAT() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_CHAR returns the _CHAR token.
	Get_CHAR() antlr.Token

	// Get_BFALSE returns the _BFALSE token.
	Get_BFALSE() antlr.Token

	// Get_BTRUE returns the _BTRUE token.
	Get_BTRUE() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_FLOAT sets the _FLOAT token.
	Set_FLOAT(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_CHAR sets the _CHAR token.
	Set_CHAR(antlr.Token)

	// Set_BFALSE sets the _BFALSE token.
	Set_BFALSE(antlr.Token)

	// Set_BTRUE sets the _BTRUE token.
	Set_BTRUE(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_methods returns the _methods rule contexts.
	Get_methods() IMethodsContext

	// Get_functionCall returns the _functionCall rule contexts.
	Get_functionCall() IFunctionCallContext

	// Set_methods sets the _methods rule contexts.
	Set_methods(IMethodsContext)

	// Set_functionCall sets the _functionCall rule contexts.
	Set_functionCall(IFunctionCallContext)

	// GetState returns the state attribute.
	GetState() I.IValue

	// SetState sets the state attribute.
	SetState(I.IValue)

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser        antlr.Parser
	state         I.IValue
	_NUMBER       antlr.Token
	_FLOAT        antlr.Token
	_STRING       antlr.Token
	_CHAR         antlr.Token
	_BFALSE       antlr.Token
	_BTRUE        antlr.Token
	_ID           antlr.Token
	_methods      IMethodsContext
	_functionCall IFunctionCallContext
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *ValueContext) Get_FLOAT() antlr.Token { return s._FLOAT }

func (s *ValueContext) Get_STRING() antlr.Token { return s._STRING }

func (s *ValueContext) Get_CHAR() antlr.Token { return s._CHAR }

func (s *ValueContext) Get_BFALSE() antlr.Token { return s._BFALSE }

func (s *ValueContext) Get_BTRUE() antlr.Token { return s._BTRUE }

func (s *ValueContext) Get_ID() antlr.Token { return s._ID }

func (s *ValueContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *ValueContext) Set_FLOAT(v antlr.Token) { s._FLOAT = v }

func (s *ValueContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *ValueContext) Set_CHAR(v antlr.Token) { s._CHAR = v }

func (s *ValueContext) Set_BFALSE(v antlr.Token) { s._BFALSE = v }

func (s *ValueContext) Set_BTRUE(v antlr.Token) { s._BTRUE = v }

func (s *ValueContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ValueContext) Get_methods() IMethodsContext { return s._methods }

func (s *ValueContext) Get_functionCall() IFunctionCallContext { return s._functionCall }

func (s *ValueContext) Set_methods(v IMethodsContext) { s._methods = v }

func (s *ValueContext) Set_functionCall(v IFunctionCallContext) { s._functionCall = v }

func (s *ValueContext) GetState() I.IValue { return s.state }

func (s *ValueContext) SetState(v I.IValue) { s.state = v }

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(DBRustParserNUMBER, 0)
}

func (s *ValueContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(DBRustParserFLOAT, 0)
}

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(DBRustParserSTRING, 0)
}

func (s *ValueContext) CHAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserCHAR, 0)
}

func (s *ValueContext) BFALSE() antlr.TerminalNode {
	return s.GetToken(DBRustParserBFALSE, 0)
}

func (s *ValueContext) BTRUE() antlr.TerminalNode {
	return s.GetToken(DBRustParserBTRUE, 0)
}

func (s *ValueContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *ValueContext) Methods() IMethodsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMethodsContext)
}

func (s *ValueContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitValue(s)
	}
}

func (p *DBRustParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, DBRustParserRULE_value)

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

	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(191)

			var _m = p.Match(DBRustParserNUMBER)

			localctx.(*ValueContext)._NUMBER = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"NUMBER", localctx.(*ValueContext).Get_NUMBER().GetLine(), localctx.(*ValueContext).Get_NUMBER().GetColumn()}, (func() string {
			if localctx.(*ValueContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_NUMBER().GetText()
			}
		}()), I.INTEGER}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(193)

			var _m = p.Match(DBRustParserFLOAT)

			localctx.(*ValueContext)._FLOAT = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"FLOAT", localctx.(*ValueContext).Get_FLOAT().GetLine(), localctx.(*ValueContext).Get_FLOAT().GetColumn()}, (func() string {
			if localctx.(*ValueContext).Get_FLOAT() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_FLOAT().GetText()
			}
		}()), I.FLOAT}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(195)

			var _m = p.Match(DBRustParserSTRING)

			localctx.(*ValueContext)._STRING = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"SRING", localctx.(*ValueContext).Get_STRING().GetLine(), localctx.(*ValueContext).Get_STRING().GetColumn()}, (func() string {
			if localctx.(*ValueContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ValueContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_STRING().GetText()
			}
		}()))-1], I.STR}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(197)

			var _m = p.Match(DBRustParserCHAR)

			localctx.(*ValueContext)._CHAR = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"CHAR", localctx.(*ValueContext).Get_CHAR().GetLine(), localctx.(*ValueContext).Get_CHAR().GetColumn()}, (func() string {
			if localctx.(*ValueContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_CHAR().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*ValueContext).Get_CHAR() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_CHAR().GetText()
			}
		}()))-1], I.CHAR}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(199)

			var _m = p.Match(DBRustParserBFALSE)

			localctx.(*ValueContext)._BFALSE = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"BFALSE", localctx.(*ValueContext).Get_BFALSE().GetLine(), localctx.(*ValueContext).Get_BFALSE().GetColumn()}, false, I.BOOL}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(201)

			var _m = p.Match(DBRustParserBTRUE)

			localctx.(*ValueContext)._BTRUE = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"BTRUE", localctx.(*ValueContext).Get_BTRUE().GetLine(), localctx.(*ValueContext).Get_BTRUE().GetColumn()}, false, I.BOOL}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(203)

			var _m = p.Match(DBRustParserID)

			localctx.(*ValueContext)._ID = _m
		}

		localctx.(*ValueContext).state = I.Value{I.Token{"ID", localctx.(*ValueContext).Get_ID().GetLine(), localctx.(*ValueContext).Get_ID().GetColumn()}, (func() string {
			if localctx.(*ValueContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*ValueContext).Get_ID().GetText()
			}
		}()), I.ID}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(205)

			var _x = p.Methods()

			localctx.(*ValueContext)._methods = _x
		}

		localctx.(*ValueContext).SetState(localctx.(*ValueContext).Get_methods().GetState())

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(208)

			var _x = p.FunctionCall()

			localctx.(*ValueContext)._functionCall = _x
		}

		localctx.(*ValueContext).SetState(localctx.(*ValueContext).Get_functionCall().GetState())

	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_expList returns the _expList rule contexts.
	Get_expList() IExpListContext

	// Set_expList sets the _expList rule contexts.
	Set_expList(IExpListContext)

	// GetState returns the state attribute.
	GetState() I.IFunctionCall

	// SetState sets the state attribute.
	SetState(I.IFunctionCall)

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	state    I.IFunctionCall
	_ID      antlr.Token
	_expList IExpListContext
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionCallContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionCallContext) Get_expList() IExpListContext { return s._expList }

func (s *FunctionCallContext) Set_expList(v IExpListContext) { s._expList = v }

func (s *FunctionCallContext) GetState() I.IFunctionCall { return s.state }

func (s *FunctionCallContext) SetState(v I.IFunctionCall) { s.state = v }

func (s *FunctionCallContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *FunctionCallContext) OPENPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserOPENPAR, 0)
}

func (s *FunctionCallContext) ExpList() IExpListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpListContext)
}

func (s *FunctionCallContext) CLOSEPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserCLOSEPAR, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *DBRustParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, DBRustParserRULE_functionCall)

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

	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(213)

			var _m = p.Match(DBRustParserID)

			localctx.(*FunctionCallContext)._ID = _m
		}
		{
			p.SetState(214)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(215)

			var _x = p.expList(0)

			localctx.(*FunctionCallContext)._expList = _x
		}
		{
			p.SetState(216)
			p.Match(DBRustParserCLOSEPAR)
		}

		localctx.(*FunctionCallContext).state = I.FunctionCall{I.Instruction{"FunctionCall"}, I.Value{I.Token{"FunctionCall", localctx.(*FunctionCallContext).Get_ID().GetLine(), localctx.(*FunctionCallContext).Get_ID().GetColumn()}, (func() string {
			if localctx.(*FunctionCallContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionCallContext).Get_ID().GetText()
			}
		}()), I.VOID}, (func() string {
			if localctx.(*FunctionCallContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionCallContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionCallContext).Get_expList().GetL().ToArray()}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)

			var _m = p.Match(DBRustParserID)

			localctx.(*FunctionCallContext)._ID = _m
		}
		{
			p.SetState(220)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(221)
			p.Match(DBRustParserCLOSEPAR)
		}

		localctx.(*FunctionCallContext).state = I.FunctionCall{I.Instruction{"FunctionCall"}, I.Value{I.Token{"FunctionCall", localctx.(*FunctionCallContext).Get_ID().GetLine(), localctx.(*FunctionCallContext).Get_ID().GetColumn()}, (func() string {
			if localctx.(*FunctionCallContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionCallContext).Get_ID().GetText()
			}
		}()), I.VOID}, (func() string {
			if localctx.(*FunctionCallContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionCallContext).Get_ID().GetText()
			}
		}()), make([]interface{}, 0)}

	}

	return localctx
}

// IMethodsContext is an interface to support dynamic dispatch.
type IMethodsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_printlnCall returns the _printlnCall rule contexts.
	Get_printlnCall() IPrintlnCallContext

	// Set_printlnCall sets the _printlnCall rule contexts.
	Set_printlnCall(IPrintlnCallContext)

	// GetState returns the state attribute.
	GetState() I.IFunctionCall

	// SetState sets the state attribute.
	SetState(I.IFunctionCall)

	// IsMethodsContext differentiates from other interfaces.
	IsMethodsContext()
}

type MethodsContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	state        I.IFunctionCall
	_printlnCall IPrintlnCallContext
}

func NewEmptyMethodsContext() *MethodsContext {
	var p = new(MethodsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_methods
	return p
}

func (*MethodsContext) IsMethodsContext() {}

func NewMethodsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodsContext {
	var p = new(MethodsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_methods

	return p
}

func (s *MethodsContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodsContext) Get_printlnCall() IPrintlnCallContext { return s._printlnCall }

func (s *MethodsContext) Set_printlnCall(v IPrintlnCallContext) { s._printlnCall = v }

func (s *MethodsContext) GetState() I.IFunctionCall { return s.state }

func (s *MethodsContext) SetState(v I.IFunctionCall) { s.state = v }

func (s *MethodsContext) PrintlnCall() IPrintlnCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrintlnCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrintlnCallContext)
}

func (s *MethodsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterMethods(s)
	}
}

func (s *MethodsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitMethods(s)
	}
}

func (p *DBRustParser) Methods() (localctx IMethodsContext) {
	localctx = NewMethodsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, DBRustParserRULE_methods)

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
		p.SetState(225)

		var _x = p.PrintlnCall()

		localctx.(*MethodsContext)._printlnCall = _x
	}
	localctx.(*MethodsContext).state = localctx.(*MethodsContext).Get_printlnCall().GetState()

	return localctx
}

// IPrintlnCallContext is an interface to support dynamic dispatch.
type IPrintlnCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PRINTLN returns the _PRINTLN token.
	Get_PRINTLN() antlr.Token

	// Set_PRINTLN sets the _PRINTLN token.
	Set_PRINTLN(antlr.Token)

	// Get_expList returns the _expList rule contexts.
	Get_expList() IExpListContext

	// Set_expList sets the _expList rule contexts.
	Set_expList(IExpListContext)

	// GetState returns the state attribute.
	GetState() I.PrintlnCall

	// SetState sets the state attribute.
	SetState(I.PrintlnCall)

	// IsPrintlnCallContext differentiates from other interfaces.
	IsPrintlnCallContext()
}

type PrintlnCallContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	state    I.PrintlnCall
	_PRINTLN antlr.Token
	_expList IExpListContext
}

func NewEmptyPrintlnCallContext() *PrintlnCallContext {
	var p = new(PrintlnCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_printlnCall
	return p
}

func (*PrintlnCallContext) IsPrintlnCallContext() {}

func NewPrintlnCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnCallContext {
	var p = new(PrintlnCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_printlnCall

	return p
}

func (s *PrintlnCallContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnCallContext) Get_PRINTLN() antlr.Token { return s._PRINTLN }

func (s *PrintlnCallContext) Set_PRINTLN(v antlr.Token) { s._PRINTLN = v }

func (s *PrintlnCallContext) Get_expList() IExpListContext { return s._expList }

func (s *PrintlnCallContext) Set_expList(v IExpListContext) { s._expList = v }

func (s *PrintlnCallContext) GetState() I.PrintlnCall { return s.state }

func (s *PrintlnCallContext) SetState(v I.PrintlnCall) { s.state = v }

func (s *PrintlnCallContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(DBRustParserPRINTLN, 0)
}

func (s *PrintlnCallContext) OPENPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserOPENPAR, 0)
}

func (s *PrintlnCallContext) ExpList() IExpListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpListContext)
}

func (s *PrintlnCallContext) CLOSEPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserCLOSEPAR, 0)
}

func (s *PrintlnCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterPrintlnCall(s)
	}
}

func (s *PrintlnCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitPrintlnCall(s)
	}
}

func (p *DBRustParser) PrintlnCall() (localctx IPrintlnCallContext) {
	localctx = NewPrintlnCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, DBRustParserRULE_printlnCall)

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
		p.SetState(228)

		var _m = p.Match(DBRustParserPRINTLN)

		localctx.(*PrintlnCallContext)._PRINTLN = _m
	}
	{
		p.SetState(229)
		p.Match(DBRustParserOPENPAR)
	}
	{
		p.SetState(230)

		var _x = p.expList(0)

		localctx.(*PrintlnCallContext)._expList = _x
	}
	{
		p.SetState(231)
		p.Match(DBRustParserCLOSEPAR)
	}

	localctx.(*PrintlnCallContext).state = I.PrintlnCall{I.FunctionCall{I.Instruction{"FunctionCall"}, I.Value{I.Token{"Println", localctx.(*PrintlnCallContext).Get_PRINTLN().GetLine(), localctx.(*PrintlnCallContext).Get_PRINTLN().GetColumn()}, "Println", I.VOID}, "Println", localctx.(*PrintlnCallContext).Get_expList().GetL().ToArray()}}

	return localctx
}

// IParamListContext is an interface to support dynamic dispatch.
type IParamListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IParamListContext

	// Get_param returns the _param rule contexts.
	Get_param() IParamContext

	// SetList sets the list rule contexts.
	SetList(IParamListContext)

	// Set_param sets the _param rule contexts.
	Set_param(IParamContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsParamListContext differentiates from other interfaces.
	IsParamListContext()
}

type ParamListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	l      *arrayList.List
	list   IParamListContext
	_param IParamContext
}

func NewEmptyParamListContext() *ParamListContext {
	var p = new(ParamListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_paramList
	return p
}

func (*ParamListContext) IsParamListContext() {}

func NewParamListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamListContext {
	var p = new(ParamListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_paramList

	return p
}

func (s *ParamListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamListContext) GetList() IParamListContext { return s.list }

func (s *ParamListContext) Get_param() IParamContext { return s._param }

func (s *ParamListContext) SetList(v IParamListContext) { s.list = v }

func (s *ParamListContext) Set_param(v IParamContext) { s._param = v }

func (s *ParamListContext) GetL() *arrayList.List { return s.l }

func (s *ParamListContext) SetL(v *arrayList.List) { s.l = v }

func (s *ParamListContext) Param() IParamContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamContext)
}

func (s *ParamListContext) COMMA() antlr.TerminalNode {
	return s.GetToken(DBRustParserCOMMA, 0)
}

func (s *ParamListContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *ParamListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterParamList(s)
	}
}

func (s *ParamListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitParamList(s)
	}
}

func (p *DBRustParser) ParamList() (localctx IParamListContext) {
	return p.paramList(0)
}

func (p *DBRustParser) paramList(_p int) (localctx IParamListContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewParamListContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IParamListContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 28
	p.EnterRecursionRule(localctx, 28, DBRustParserRULE_paramList, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(235)

		var _x = p.Param()

		localctx.(*ParamListContext)._param = _x
	}

	localctx.(*ParamListContext).l = arrayList.New()
	localctx.(*ParamListContext).l.Add(localctx.(*ParamListContext).Get_param().GetState())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewParamListContext(p, _parentctx, _parentState)
			localctx.(*ParamListContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, DBRustParserRULE_paramList)
			p.SetState(238)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(239)
				p.Match(DBRustParserCOMMA)
			}
			{
				p.SetState(240)

				var _x = p.Param()

				localctx.(*ParamListContext)._param = _x
			}

			localctx.(*ParamListContext).GetList().GetL().Add(localctx.(*ParamListContext).Get_param().GetState())
			localctx.(*ParamListContext).l = localctx.(*ParamListContext).GetList().GetL()

		}
		p.SetState(247)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IParamContext is an interface to support dynamic dispatch.
type IParamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_valueType returns the _valueType rule contexts.
	Get_valueType() IValueTypeContext

	// Set_valueType sets the _valueType rule contexts.
	Set_valueType(IValueTypeContext)

	// GetState returns the state attribute.
	GetState() I.FunctionParam

	// SetState sets the state attribute.
	SetState(I.FunctionParam)

	// IsParamContext differentiates from other interfaces.
	IsParamContext()
}

type ParamContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	state      I.FunctionParam
	_ID        antlr.Token
	_valueType IValueTypeContext
}

func NewEmptyParamContext() *ParamContext {
	var p = new(ParamContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_param
	return p
}

func (*ParamContext) IsParamContext() {}

func NewParamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamContext {
	var p = new(ParamContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_param

	return p
}

func (s *ParamContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamContext) Get_ID() antlr.Token { return s._ID }

func (s *ParamContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *ParamContext) Get_valueType() IValueTypeContext { return s._valueType }

func (s *ParamContext) Set_valueType(v IValueTypeContext) { s._valueType = v }

func (s *ParamContext) GetState() I.FunctionParam { return s.state }

func (s *ParamContext) SetState(v I.FunctionParam) { s.state = v }

func (s *ParamContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *ParamContext) COLOM() antlr.TerminalNode {
	return s.GetToken(DBRustParserCOLOM, 0)
}

func (s *ParamContext) ValueType() IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *ParamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterParam(s)
	}
}

func (s *ParamContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitParam(s)
	}
}

func (p *DBRustParser) Param() (localctx IParamContext) {
	localctx = NewParamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, DBRustParserRULE_param)

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
		p.SetState(248)

		var _m = p.Match(DBRustParserID)

		localctx.(*ParamContext)._ID = _m
	}
	{
		p.SetState(249)
		p.Match(DBRustParserCOLOM)
	}
	{
		p.SetState(250)

		var _x = p.ValueType()

		localctx.(*ParamContext)._valueType = _x
	}

	localctx.(*ParamContext).SetState(I.FunctionParam{(func() string {
		if localctx.(*ParamContext).Get_ID() == nil {
			return ""
		} else {
			return localctx.(*ParamContext).Get_ID().GetText()
		}
	}()), localctx.(*ParamContext).Get_valueType().GetState()})

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_FN returns the _FN token.
	Get_FN() antlr.Token

	// Get_ID returns the _ID token.
	Get_ID() antlr.Token

	// Set_FN sets the _FN token.
	Set_FN(antlr.Token)

	// Set_ID sets the _ID token.
	Set_ID(antlr.Token)

	// Get_paramList returns the _paramList rule contexts.
	Get_paramList() IParamListContext

	// Get_instructionsBlock returns the _instructionsBlock rule contexts.
	Get_instructionsBlock() IInstructionsBlockContext

	// Get_valueType returns the _valueType rule contexts.
	Get_valueType() IValueTypeContext

	// Set_paramList sets the _paramList rule contexts.
	Set_paramList(IParamListContext)

	// Set_instructionsBlock sets the _instructionsBlock rule contexts.
	Set_instructionsBlock(IInstructionsBlockContext)

	// Set_valueType sets the _valueType rule contexts.
	Set_valueType(IValueTypeContext)

	// GetState returns the state attribute.
	GetState() I.Function

	// SetState sets the state attribute.
	SetState(I.Function)

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	state              I.Function
	_FN                antlr.Token
	_ID                antlr.Token
	_paramList         IParamListContext
	_instructionsBlock IInstructionsBlockContext
	_valueType         IValueTypeContext
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) Get_FN() antlr.Token { return s._FN }

func (s *FunctionContext) Get_ID() antlr.Token { return s._ID }

func (s *FunctionContext) Set_FN(v antlr.Token) { s._FN = v }

func (s *FunctionContext) Set_ID(v antlr.Token) { s._ID = v }

func (s *FunctionContext) Get_paramList() IParamListContext { return s._paramList }

func (s *FunctionContext) Get_instructionsBlock() IInstructionsBlockContext {
	return s._instructionsBlock
}

func (s *FunctionContext) Get_valueType() IValueTypeContext { return s._valueType }

func (s *FunctionContext) Set_paramList(v IParamListContext) { s._paramList = v }

func (s *FunctionContext) Set_instructionsBlock(v IInstructionsBlockContext) {
	s._instructionsBlock = v
}

func (s *FunctionContext) Set_valueType(v IValueTypeContext) { s._valueType = v }

func (s *FunctionContext) GetState() I.Function { return s.state }

func (s *FunctionContext) SetState(v I.Function) { s.state = v }

func (s *FunctionContext) FN() antlr.TerminalNode {
	return s.GetToken(DBRustParserFN, 0)
}

func (s *FunctionContext) ID() antlr.TerminalNode {
	return s.GetToken(DBRustParserID, 0)
}

func (s *FunctionContext) OPENPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserOPENPAR, 0)
}

func (s *FunctionContext) ParamList() IParamListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParamListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParamListContext)
}

func (s *FunctionContext) CLOSEPAR() antlr.TerminalNode {
	return s.GetToken(DBRustParserCLOSEPAR, 0)
}

func (s *FunctionContext) InstructionsBlock() IInstructionsBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstructionsBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstructionsBlockContext)
}

func (s *FunctionContext) ARROW() antlr.TerminalNode {
	return s.GetToken(DBRustParserARROW, 0)
}

func (s *FunctionContext) ValueType() IValueTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueTypeContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterFunction(s)
	}
}

func (s *FunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitFunction(s)
	}
}

func (p *DBRustParser) Function() (localctx IFunctionContext) {
	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, DBRustParserRULE_function)

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

	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)

			var _m = p.Match(DBRustParserFN)

			localctx.(*FunctionContext)._FN = _m
		}
		{
			p.SetState(254)

			var _m = p.Match(DBRustParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(255)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(256)

			var _x = p.paramList(0)

			localctx.(*FunctionContext)._paramList = _x
		}
		{
			p.SetState(257)
			p.Match(DBRustParserCLOSEPAR)
		}
		{
			p.SetState(258)

			var _x = p.InstructionsBlock()

			localctx.(*FunctionContext)._instructionsBlock = _x
		}

		localctx.(*FunctionContext).SetState(I.Function{I.Instruction{"Function"}, I.Token{"Function", localctx.(*FunctionContext).Get_FN().GetLine(), localctx.(*FunctionContext).Get_FN().GetColumn()}, (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_paramList().GetL().ToArray(), localctx.(*FunctionContext).Get_instructionsBlock().GetL().ToArray(), I.VOID, nil})

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(261)

			var _m = p.Match(DBRustParserFN)

			localctx.(*FunctionContext)._FN = _m
		}
		{
			p.SetState(262)

			var _m = p.Match(DBRustParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(263)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(264)
			p.Match(DBRustParserCLOSEPAR)
		}
		{
			p.SetState(265)

			var _x = p.InstructionsBlock()

			localctx.(*FunctionContext)._instructionsBlock = _x
		}

		localctx.(*FunctionContext).SetState(I.Function{I.Instruction{"Function"}, I.Token{"Function", localctx.(*FunctionContext).Get_FN().GetLine(), localctx.(*FunctionContext).Get_FN().GetColumn()}, (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), make([]interface{}, 0), localctx.(*FunctionContext).Get_instructionsBlock().GetL().ToArray(), I.VOID, nil})

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(268)

			var _m = p.Match(DBRustParserFN)

			localctx.(*FunctionContext)._FN = _m
		}
		{
			p.SetState(269)

			var _m = p.Match(DBRustParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(270)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(271)

			var _x = p.paramList(0)

			localctx.(*FunctionContext)._paramList = _x
		}
		{
			p.SetState(272)
			p.Match(DBRustParserCLOSEPAR)
		}
		{
			p.SetState(273)
			p.Match(DBRustParserARROW)
		}
		{
			p.SetState(274)

			var _x = p.ValueType()

			localctx.(*FunctionContext)._valueType = _x
		}
		{
			p.SetState(275)

			var _x = p.InstructionsBlock()

			localctx.(*FunctionContext)._instructionsBlock = _x
		}

		localctx.(*FunctionContext).SetState(I.Function{I.Instruction{"Function"}, I.Token{"Function", localctx.(*FunctionContext).Get_FN().GetLine(), localctx.(*FunctionContext).Get_FN().GetColumn()}, (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), localctx.(*FunctionContext).Get_paramList().GetL().ToArray(), localctx.(*FunctionContext).Get_instructionsBlock().GetL().ToArray(), localctx.(*FunctionContext).Get_valueType().GetState(), nil})

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(278)

			var _m = p.Match(DBRustParserFN)

			localctx.(*FunctionContext)._FN = _m
		}
		{
			p.SetState(279)

			var _m = p.Match(DBRustParserID)

			localctx.(*FunctionContext)._ID = _m
		}
		{
			p.SetState(280)
			p.Match(DBRustParserOPENPAR)
		}
		{
			p.SetState(281)
			p.Match(DBRustParserCLOSEPAR)
		}
		{
			p.SetState(282)
			p.Match(DBRustParserARROW)
		}
		{
			p.SetState(283)

			var _x = p.ValueType()

			localctx.(*FunctionContext)._valueType = _x
		}
		{
			p.SetState(284)

			var _x = p.InstructionsBlock()

			localctx.(*FunctionContext)._instructionsBlock = _x
		}

		localctx.(*FunctionContext).SetState(I.Function{I.Instruction{"Function"}, I.Token{"Function", localctx.(*FunctionContext).Get_FN().GetLine(), localctx.(*FunctionContext).Get_FN().GetColumn()}, (func() string {
			if localctx.(*FunctionContext).Get_ID() == nil {
				return ""
			} else {
				return localctx.(*FunctionContext).Get_ID().GetText()
			}
		}()), make([]interface{}, 0), localctx.(*FunctionContext).Get_instructionsBlock().GetL().ToArray(), localctx.(*FunctionContext).Get_valueType().GetState(), nil})

	}

	return localctx
}

// IReturnValueContext is an interface to support dynamic dispatch.
type IReturnValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_RETURN returns the _RETURN token.
	Get_RETURN() antlr.Token

	// Set_RETURN sets the _RETURN token.
	Set_RETURN(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetState returns the state attribute.
	GetState() I.ReturnValue

	// SetState sets the state attribute.
	SetState(I.ReturnValue)

	// IsReturnValueContext differentiates from other interfaces.
	IsReturnValueContext()
}

type ReturnValueContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	state       I.ReturnValue
	_RETURN     antlr.Token
	_expression IExpressionContext
}

func NewEmptyReturnValueContext() *ReturnValueContext {
	var p = new(ReturnValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = DBRustParserRULE_returnValue
	return p
}

func (*ReturnValueContext) IsReturnValueContext() {}

func NewReturnValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnValueContext {
	var p = new(ReturnValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = DBRustParserRULE_returnValue

	return p
}

func (s *ReturnValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnValueContext) Get_RETURN() antlr.Token { return s._RETURN }

func (s *ReturnValueContext) Set_RETURN(v antlr.Token) { s._RETURN = v }

func (s *ReturnValueContext) Get_expression() IExpressionContext { return s._expression }

func (s *ReturnValueContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ReturnValueContext) GetState() I.ReturnValue { return s.state }

func (s *ReturnValueContext) SetState(v I.ReturnValue) { s.state = v }

func (s *ReturnValueContext) RETURN() antlr.TerminalNode {
	return s.GetToken(DBRustParserRETURN, 0)
}

func (s *ReturnValueContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.EnterReturnValue(s)
	}
}

func (s *ReturnValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(DBRustListener); ok {
		listenerT.ExitReturnValue(s)
	}
}

func (p *DBRustParser) ReturnValue() (localctx IReturnValueContext) {
	localctx = NewReturnValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, DBRustParserRULE_returnValue)

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
		p.SetState(289)

		var _m = p.Match(DBRustParserRETURN)

		localctx.(*ReturnValueContext)._RETURN = _m
	}
	{
		p.SetState(290)

		var _x = p.expression(0)

		localctx.(*ReturnValueContext)._expression = _x
	}

	localctx.(*ReturnValueContext).state = I.ReturnValue{I.Instruction{"Return"}, I.Token{"Return", localctx.(*ReturnValueContext).Get_RETURN().GetLine(), localctx.(*ReturnValueContext).Get_RETURN().GetColumn()}, localctx.(*ReturnValueContext).Get_expression().GetState()}

	return localctx
}

func (p *DBRustParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 6:
		var t *ExpListContext = nil
		if localctx != nil {
			t = localctx.(*ExpListContext)
		}
		return p.ExpList_Sempred(t, predIndex)

	case 7:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 14:
		var t *ParamListContext = nil
		if localctx != nil {
			t = localctx.(*ParamListContext)
		}
		return p.ParamList_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *DBRustParser) ExpList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *DBRustParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *DBRustParser) ParamList_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
