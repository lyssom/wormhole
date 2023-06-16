// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SqlParser

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

type SqlParser struct {
	*antlr.BaseParser
}

var sqlparserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sqlparserParserInit() {
	staticData := &sqlparserParserStaticData
	staticData.literalNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "'->'", "'*'", "'`'", "'\\'", "':'", "','", "'||'", "'-'", "'.'",
		"'=='", "'='", "'>='", "'>'", "'{'", "'['", "'<='", "'('", "'<'", "",
		"'%'", "'+'", "'?'", "'\"'", "'''", "'}'", "']'", "')'", "';'", "'/'",
		"'_'",
	}
	staticData.symbolicNames = []string{
		"", "WS", "CREATE", "DATABASE", "DEFAULT", "EXISTS", "FROM", "IF", "INF",
		"INSERT", "INTO", "NAN_SQL", "NOT", "NULL_SQL", "SELECT", "SET", "TABLE",
		"USE", "VALUE", "VALUES", "SECOND", "MINUTE", "HOUR", "DAY", "WEEK",
		"MONTH", "QUARTER", "YEAR", "IDENTIFIER", "FLOATING_LITERAL", "OCTAL_LITERAL",
		"DECIMAL_LITERAL", "HEXADECIMAL_LITERAL", "EXPONENT_NUM_PART", "STRING_LITERAL",
		"ARROW", "ASTERISK", "BACKQUOTE", "BACKSLASH", "COLON", "COMMA", "CONCAT",
		"DASH", "DOT", "EQ_DOUBLE", "EQ_SINGLE", "GE", "GT", "LBRACE", "LBRACKET",
		"LE", "LPAREN", "LT", "NOT_EQ", "PERCENT", "PLUS", "QUERY", "QUOTE_DOUBLE",
		"QUOTE_SINGLE", "RBRACE", "RBRACKET", "RPAREN", "SEMICOLON", "SLASH",
		"UNDERSCORE", "MULTI_LINE_COMMENT", "SINGLE_LINE_COMMENT", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"statement", "dmlStatement", "ddlStatement", "systemStatement", "createDatabase",
		"createTable", "tableSchemaClause", "tableElementExpr", "tableColumnDfnt",
		"insertStatement", "columnsClause", "insertValuesSpec", "insertMultiValue",
		"dataValue", "selectStatement", "fromClause", "joinExpr", "useStatement",
		"columnTypeExpr", "columnExprList", "columnsExpr", "columnExpr", "columnIdentifier",
		"nestedIdentifier", "constant", "realLiteral", "tableExpr", "tableIdentifier",
		"databaseIdentifier", "floatingLiteral", "numberLiteral", "literal",
		"interval", "keyword", "identifier", "enumValue", "ifNotExists",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 340, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		1, 0, 1, 0, 1, 0, 3, 0, 78, 8, 0, 1, 1, 1, 1, 3, 1, 82, 8, 1, 1, 2, 1,
		2, 3, 2, 86, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 93, 8, 4, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 100, 8, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6,
		1, 6, 1, 6, 5, 6, 109, 8, 6, 10, 6, 12, 6, 112, 9, 6, 1, 6, 1, 6, 1, 7,
		1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 124, 8, 9, 1, 9, 1, 9,
		3, 9, 128, 8, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 137,
		8, 10, 10, 10, 12, 10, 140, 9, 10, 1, 10, 1, 10, 1, 11, 3, 11, 145, 8,
		11, 1, 11, 5, 11, 148, 8, 11, 10, 11, 12, 11, 151, 9, 11, 1, 12, 1, 12,
		1, 12, 3, 12, 156, 8, 12, 4, 12, 158, 8, 12, 11, 12, 12, 12, 159, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 4, 13, 169, 8, 13, 11, 13, 12,
		13, 170, 1, 13, 1, 13, 3, 13, 175, 8, 13, 1, 14, 1, 14, 1, 14, 3, 14, 180,
		8, 14, 1, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 199, 8, 18,
		10, 18, 12, 18, 202, 9, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		18, 5, 18, 211, 8, 18, 10, 18, 12, 18, 214, 9, 18, 1, 18, 1, 18, 1, 18,
		1, 18, 1, 18, 1, 18, 1, 18, 5, 18, 223, 8, 18, 10, 18, 12, 18, 226, 9,
		18, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 3, 18, 233, 8, 18, 1, 18, 1, 18,
		3, 18, 237, 8, 18, 1, 19, 1, 19, 1, 19, 5, 19, 242, 8, 19, 10, 19, 12,
		19, 245, 9, 19, 1, 20, 1, 20, 1, 21, 1, 21, 3, 21, 251, 8, 21, 1, 22, 1,
		22, 1, 22, 3, 22, 256, 8, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 3, 23,
		263, 8, 23, 1, 24, 3, 24, 266, 8, 24, 1, 24, 1, 24, 3, 24, 270, 8, 24,
		1, 24, 1, 24, 3, 24, 274, 8, 24, 1, 25, 1, 25, 1, 25, 3, 25, 279, 8, 25,
		1, 25, 1, 25, 1, 25, 3, 25, 284, 8, 25, 1, 26, 1, 26, 1, 27, 1, 27, 1,
		27, 3, 27, 291, 8, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 1, 29, 1, 29, 3, 29, 303, 8, 29, 3, 29, 305, 8, 29, 1, 30, 3, 30,
		308, 8, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 1, 30, 3, 30, 316, 8, 30,
		1, 31, 1, 31, 1, 31, 3, 31, 321, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1,
		34, 1, 34, 1, 34, 3, 34, 330, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 0, 0, 37, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18,
		20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54,
		56, 58, 60, 62, 64, 66, 68, 70, 72, 0, 5, 2, 0, 42, 42, 55, 55, 2, 0, 31,
		31, 33, 33, 1, 0, 30, 31, 1, 0, 20, 27, 4, 0, 2, 7, 9, 10, 12, 16, 18,
		19, 352, 0, 77, 1, 0, 0, 0, 2, 81, 1, 0, 0, 0, 4, 85, 1, 0, 0, 0, 6, 87,
		1, 0, 0, 0, 8, 89, 1, 0, 0, 0, 10, 96, 1, 0, 0, 0, 12, 104, 1, 0, 0, 0,
		14, 115, 1, 0, 0, 0, 16, 117, 1, 0, 0, 0, 18, 120, 1, 0, 0, 0, 20, 132,
		1, 0, 0, 0, 22, 149, 1, 0, 0, 0, 24, 152, 1, 0, 0, 0, 26, 174, 1, 0, 0,
		0, 28, 176, 1, 0, 0, 0, 30, 181, 1, 0, 0, 0, 32, 184, 1, 0, 0, 0, 34, 186,
		1, 0, 0, 0, 36, 236, 1, 0, 0, 0, 38, 238, 1, 0, 0, 0, 40, 246, 1, 0, 0,
		0, 42, 250, 1, 0, 0, 0, 44, 255, 1, 0, 0, 0, 46, 259, 1, 0, 0, 0, 48, 273,
		1, 0, 0, 0, 50, 283, 1, 0, 0, 0, 52, 285, 1, 0, 0, 0, 54, 290, 1, 0, 0,
		0, 56, 294, 1, 0, 0, 0, 58, 304, 1, 0, 0, 0, 60, 307, 1, 0, 0, 0, 62, 320,
		1, 0, 0, 0, 64, 322, 1, 0, 0, 0, 66, 324, 1, 0, 0, 0, 68, 329, 1, 0, 0,
		0, 70, 331, 1, 0, 0, 0, 72, 335, 1, 0, 0, 0, 74, 78, 3, 4, 2, 0, 75, 78,
		3, 2, 1, 0, 76, 78, 3, 6, 3, 0, 77, 74, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0,
		77, 76, 1, 0, 0, 0, 78, 1, 1, 0, 0, 0, 79, 82, 3, 28, 14, 0, 80, 82, 3,
		18, 9, 0, 81, 79, 1, 0, 0, 0, 81, 80, 1, 0, 0, 0, 82, 3, 1, 0, 0, 0, 83,
		86, 3, 8, 4, 0, 84, 86, 3, 10, 5, 0, 85, 83, 1, 0, 0, 0, 85, 84, 1, 0,
		0, 0, 86, 5, 1, 0, 0, 0, 87, 88, 3, 34, 17, 0, 88, 7, 1, 0, 0, 0, 89, 90,
		5, 2, 0, 0, 90, 92, 5, 3, 0, 0, 91, 93, 3, 72, 36, 0, 92, 91, 1, 0, 0,
		0, 92, 93, 1, 0, 0, 0, 93, 94, 1, 0, 0, 0, 94, 95, 3, 56, 28, 0, 95, 9,
		1, 0, 0, 0, 96, 97, 5, 2, 0, 0, 97, 99, 5, 16, 0, 0, 98, 100, 3, 72, 36,
		0, 99, 98, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101,
		102, 3, 54, 27, 0, 102, 103, 3, 12, 6, 0, 103, 11, 1, 0, 0, 0, 104, 105,
		5, 51, 0, 0, 105, 110, 3, 14, 7, 0, 106, 107, 5, 40, 0, 0, 107, 109, 3,
		14, 7, 0, 108, 106, 1, 0, 0, 0, 109, 112, 1, 0, 0, 0, 110, 108, 1, 0, 0,
		0, 110, 111, 1, 0, 0, 0, 111, 113, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 113,
		114, 5, 61, 0, 0, 114, 13, 1, 0, 0, 0, 115, 116, 3, 16, 8, 0, 116, 15,
		1, 0, 0, 0, 117, 118, 3, 46, 23, 0, 118, 119, 3, 36, 18, 0, 119, 17, 1,
		0, 0, 0, 120, 121, 5, 9, 0, 0, 121, 123, 5, 10, 0, 0, 122, 124, 5, 16,
		0, 0, 123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0,
		125, 127, 3, 54, 27, 0, 126, 128, 3, 20, 10, 0, 127, 126, 1, 0, 0, 0, 127,
		128, 1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 130, 5, 19, 0, 0, 130, 131,
		3, 22, 11, 0, 131, 19, 1, 0, 0, 0, 132, 133, 5, 51, 0, 0, 133, 138, 3,
		46, 23, 0, 134, 135, 5, 40, 0, 0, 135, 137, 3, 46, 23, 0, 136, 134, 1,
		0, 0, 0, 137, 140, 1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 138, 139, 1, 0, 0,
		0, 139, 141, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 141, 142, 5, 61, 0, 0, 142,
		21, 1, 0, 0, 0, 143, 145, 5, 40, 0, 0, 144, 143, 1, 0, 0, 0, 144, 145,
		1, 0, 0, 0, 145, 146, 1, 0, 0, 0, 146, 148, 3, 24, 12, 0, 147, 144, 1,
		0, 0, 0, 148, 151, 1, 0, 0, 0, 149, 147, 1, 0, 0, 0, 149, 150, 1, 0, 0,
		0, 150, 23, 1, 0, 0, 0, 151, 149, 1, 0, 0, 0, 152, 157, 5, 51, 0, 0, 153,
		155, 3, 26, 13, 0, 154, 156, 5, 40, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156,
		1, 0, 0, 0, 156, 158, 1, 0, 0, 0, 157, 153, 1, 0, 0, 0, 158, 159, 1, 0,
		0, 0, 159, 157, 1, 0, 0, 0, 159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0,
		161, 162, 5, 61, 0, 0, 162, 25, 1, 0, 0, 0, 163, 175, 3, 48, 24, 0, 164,
		165, 5, 51, 0, 0, 165, 168, 3, 48, 24, 0, 166, 167, 5, 40, 0, 0, 167, 169,
		3, 48, 24, 0, 168, 166, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 168, 1,
		0, 0, 0, 170, 171, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 173, 5, 61, 0,
		0, 173, 175, 1, 0, 0, 0, 174, 163, 1, 0, 0, 0, 174, 164, 1, 0, 0, 0, 175,
		27, 1, 0, 0, 0, 176, 177, 5, 14, 0, 0, 177, 179, 3, 38, 19, 0, 178, 180,
		3, 30, 15, 0, 179, 178, 1, 0, 0, 0, 179, 180, 1, 0, 0, 0, 180, 29, 1, 0,
		0, 0, 181, 182, 5, 6, 0, 0, 182, 183, 3, 32, 16, 0, 183, 31, 1, 0, 0, 0,
		184, 185, 3, 52, 26, 0, 185, 33, 1, 0, 0, 0, 186, 187, 5, 17, 0, 0, 187,
		188, 3, 56, 28, 0, 188, 35, 1, 0, 0, 0, 189, 237, 3, 68, 34, 0, 190, 191,
		3, 68, 34, 0, 191, 192, 5, 51, 0, 0, 192, 193, 3, 68, 34, 0, 193, 200,
		3, 36, 18, 0, 194, 195, 5, 40, 0, 0, 195, 196, 3, 68, 34, 0, 196, 197,
		3, 36, 18, 0, 197, 199, 1, 0, 0, 0, 198, 194, 1, 0, 0, 0, 199, 202, 1,
		0, 0, 0, 200, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 203, 1, 0, 0,
		0, 202, 200, 1, 0, 0, 0, 203, 204, 5, 61, 0, 0, 204, 237, 1, 0, 0, 0, 205,
		206, 3, 68, 34, 0, 206, 207, 5, 51, 0, 0, 207, 212, 3, 70, 35, 0, 208,
		209, 5, 40, 0, 0, 209, 211, 3, 70, 35, 0, 210, 208, 1, 0, 0, 0, 211, 214,
		1, 0, 0, 0, 212, 210, 1, 0, 0, 0, 212, 213, 1, 0, 0, 0, 213, 215, 1, 0,
		0, 0, 214, 212, 1, 0, 0, 0, 215, 216, 5, 61, 0, 0, 216, 237, 1, 0, 0, 0,
		217, 218, 3, 68, 34, 0, 218, 219, 5, 51, 0, 0, 219, 224, 3, 36, 18, 0,
		220, 221, 5, 40, 0, 0, 221, 223, 3, 36, 18, 0, 222, 220, 1, 0, 0, 0, 223,
		226, 1, 0, 0, 0, 224, 222, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 227,
		1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 227, 228, 5, 61, 0, 0, 228, 237, 1, 0,
		0, 0, 229, 230, 3, 68, 34, 0, 230, 232, 5, 51, 0, 0, 231, 233, 3, 38, 19,
		0, 232, 231, 1, 0, 0, 0, 232, 233, 1, 0, 0, 0, 233, 234, 1, 0, 0, 0, 234,
		235, 5, 61, 0, 0, 235, 237, 1, 0, 0, 0, 236, 189, 1, 0, 0, 0, 236, 190,
		1, 0, 0, 0, 236, 205, 1, 0, 0, 0, 236, 217, 1, 0, 0, 0, 236, 229, 1, 0,
		0, 0, 237, 37, 1, 0, 0, 0, 238, 243, 3, 40, 20, 0, 239, 240, 5, 40, 0,
		0, 240, 242, 3, 40, 20, 0, 241, 239, 1, 0, 0, 0, 242, 245, 1, 0, 0, 0,
		243, 241, 1, 0, 0, 0, 243, 244, 1, 0, 0, 0, 244, 39, 1, 0, 0, 0, 245, 243,
		1, 0, 0, 0, 246, 247, 3, 42, 21, 0, 247, 41, 1, 0, 0, 0, 248, 251, 3, 62,
		31, 0, 249, 251, 3, 44, 22, 0, 250, 248, 1, 0, 0, 0, 250, 249, 1, 0, 0,
		0, 251, 43, 1, 0, 0, 0, 252, 253, 3, 54, 27, 0, 253, 254, 5, 43, 0, 0,
		254, 256, 1, 0, 0, 0, 255, 252, 1, 0, 0, 0, 255, 256, 1, 0, 0, 0, 256,
		257, 1, 0, 0, 0, 257, 258, 3, 46, 23, 0, 258, 45, 1, 0, 0, 0, 259, 262,
		3, 68, 34, 0, 260, 261, 5, 43, 0, 0, 261, 263, 3, 68, 34, 0, 262, 260,
		1, 0, 0, 0, 262, 263, 1, 0, 0, 0, 263, 47, 1, 0, 0, 0, 264, 266, 7, 0,
		0, 0, 265, 264, 1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 1, 0, 0, 0,
		267, 274, 3, 50, 25, 0, 268, 270, 7, 0, 0, 0, 269, 268, 1, 0, 0, 0, 269,
		270, 1, 0, 0, 0, 270, 271, 1, 0, 0, 0, 271, 274, 5, 31, 0, 0, 272, 274,
		5, 34, 0, 0, 273, 265, 1, 0, 0, 0, 273, 269, 1, 0, 0, 0, 273, 272, 1, 0,
		0, 0, 274, 49, 1, 0, 0, 0, 275, 276, 5, 31, 0, 0, 276, 278, 5, 43, 0, 0,
		277, 279, 7, 1, 0, 0, 278, 277, 1, 0, 0, 0, 278, 279, 1, 0, 0, 0, 279,
		284, 1, 0, 0, 0, 280, 281, 5, 43, 0, 0, 281, 284, 7, 1, 0, 0, 282, 284,
		5, 33, 0, 0, 283, 275, 1, 0, 0, 0, 283, 280, 1, 0, 0, 0, 283, 282, 1, 0,
		0, 0, 284, 51, 1, 0, 0, 0, 285, 286, 3, 54, 27, 0, 286, 53, 1, 0, 0, 0,
		287, 288, 3, 56, 28, 0, 288, 289, 5, 43, 0, 0, 289, 291, 1, 0, 0, 0, 290,
		287, 1, 0, 0, 0, 290, 291, 1, 0, 0, 0, 291, 292, 1, 0, 0, 0, 292, 293,
		3, 68, 34, 0, 293, 55, 1, 0, 0, 0, 294, 295, 3, 68, 34, 0, 295, 57, 1,
		0, 0, 0, 296, 305, 5, 29, 0, 0, 297, 298, 5, 43, 0, 0, 298, 305, 7, 2,
		0, 0, 299, 300, 5, 31, 0, 0, 300, 302, 5, 43, 0, 0, 301, 303, 7, 2, 0,
		0, 302, 301, 1, 0, 0, 0, 302, 303, 1, 0, 0, 0, 303, 305, 1, 0, 0, 0, 304,
		296, 1, 0, 0, 0, 304, 297, 1, 0, 0, 0, 304, 299, 1, 0, 0, 0, 305, 59, 1,
		0, 0, 0, 306, 308, 7, 0, 0, 0, 307, 306, 1, 0, 0, 0, 307, 308, 1, 0, 0,
		0, 308, 315, 1, 0, 0, 0, 309, 316, 3, 58, 29, 0, 310, 316, 5, 30, 0, 0,
		311, 316, 5, 31, 0, 0, 312, 316, 5, 32, 0, 0, 313, 316, 5, 8, 0, 0, 314,
		316, 5, 11, 0, 0, 315, 309, 1, 0, 0, 0, 315, 310, 1, 0, 0, 0, 315, 311,
		1, 0, 0, 0, 315, 312, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315, 314, 1, 0,
		0, 0, 316, 61, 1, 0, 0, 0, 317, 321, 3, 60, 30, 0, 318, 321, 5, 34, 0,
		0, 319, 321, 5, 13, 0, 0, 320, 317, 1, 0, 0, 0, 320, 318, 1, 0, 0, 0, 320,
		319, 1, 0, 0, 0, 321, 63, 1, 0, 0, 0, 322, 323, 7, 3, 0, 0, 323, 65, 1,
		0, 0, 0, 324, 325, 7, 4, 0, 0, 325, 67, 1, 0, 0, 0, 326, 330, 5, 28, 0,
		0, 327, 330, 3, 64, 32, 0, 328, 330, 3, 66, 33, 0, 329, 326, 1, 0, 0, 0,
		329, 327, 1, 0, 0, 0, 329, 328, 1, 0, 0, 0, 330, 69, 1, 0, 0, 0, 331, 332,
		5, 34, 0, 0, 332, 333, 5, 45, 0, 0, 333, 334, 3, 60, 30, 0, 334, 71, 1,
		0, 0, 0, 335, 336, 5, 7, 0, 0, 336, 337, 5, 12, 0, 0, 337, 338, 5, 5, 0,
		0, 338, 73, 1, 0, 0, 0, 37, 77, 81, 85, 92, 99, 110, 123, 127, 138, 144,
		149, 155, 159, 170, 174, 179, 200, 212, 224, 232, 236, 243, 250, 255, 262,
		265, 269, 273, 278, 283, 290, 302, 304, 307, 315, 320, 329,
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

// SqlParserInit initializes any static state used to implement SqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SqlParserInit() {
	staticData := &sqlparserParserStaticData
	staticData.once.Do(sqlparserParserInit)
}

// NewSqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSqlParser(input antlr.TokenStream) *SqlParser {
	SqlParserInit()
	this := new(SqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &sqlparserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "java-escape"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF                 = antlr.TokenEOF
	SqlParserWS                  = 1
	SqlParserCREATE              = 2
	SqlParserDATABASE            = 3
	SqlParserDEFAULT             = 4
	SqlParserEXISTS              = 5
	SqlParserFROM                = 6
	SqlParserIF                  = 7
	SqlParserINF                 = 8
	SqlParserINSERT              = 9
	SqlParserINTO                = 10
	SqlParserNAN_SQL             = 11
	SqlParserNOT                 = 12
	SqlParserNULL_SQL            = 13
	SqlParserSELECT              = 14
	SqlParserSET                 = 15
	SqlParserTABLE               = 16
	SqlParserUSE                 = 17
	SqlParserVALUE               = 18
	SqlParserVALUES              = 19
	SqlParserSECOND              = 20
	SqlParserMINUTE              = 21
	SqlParserHOUR                = 22
	SqlParserDAY                 = 23
	SqlParserWEEK                = 24
	SqlParserMONTH               = 25
	SqlParserQUARTER             = 26
	SqlParserYEAR                = 27
	SqlParserIDENTIFIER          = 28
	SqlParserFLOATING_LITERAL    = 29
	SqlParserOCTAL_LITERAL       = 30
	SqlParserDECIMAL_LITERAL     = 31
	SqlParserHEXADECIMAL_LITERAL = 32
	SqlParserEXPONENT_NUM_PART   = 33
	SqlParserSTRING_LITERAL      = 34
	SqlParserARROW               = 35
	SqlParserASTERISK            = 36
	SqlParserBACKQUOTE           = 37
	SqlParserBACKSLASH           = 38
	SqlParserCOLON               = 39
	SqlParserCOMMA               = 40
	SqlParserCONCAT              = 41
	SqlParserDASH                = 42
	SqlParserDOT                 = 43
	SqlParserEQ_DOUBLE           = 44
	SqlParserEQ_SINGLE           = 45
	SqlParserGE                  = 46
	SqlParserGT                  = 47
	SqlParserLBRACE              = 48
	SqlParserLBRACKET            = 49
	SqlParserLE                  = 50
	SqlParserLPAREN              = 51
	SqlParserLT                  = 52
	SqlParserNOT_EQ              = 53
	SqlParserPERCENT             = 54
	SqlParserPLUS                = 55
	SqlParserQUERY               = 56
	SqlParserQUOTE_DOUBLE        = 57
	SqlParserQUOTE_SINGLE        = 58
	SqlParserRBRACE              = 59
	SqlParserRBRACKET            = 60
	SqlParserRPAREN              = 61
	SqlParserSEMICOLON           = 62
	SqlParserSLASH               = 63
	SqlParserUNDERSCORE          = 64
	SqlParserMULTI_LINE_COMMENT  = 65
	SqlParserSINGLE_LINE_COMMENT = 66
	SqlParserWHITESPACE          = 67
)

// SqlParser rules.
const (
	SqlParserRULE_statement          = 0
	SqlParserRULE_dmlStatement       = 1
	SqlParserRULE_ddlStatement       = 2
	SqlParserRULE_systemStatement    = 3
	SqlParserRULE_createDatabase     = 4
	SqlParserRULE_createTable        = 5
	SqlParserRULE_tableSchemaClause  = 6
	SqlParserRULE_tableElementExpr   = 7
	SqlParserRULE_tableColumnDfnt    = 8
	SqlParserRULE_insertStatement    = 9
	SqlParserRULE_columnsClause      = 10
	SqlParserRULE_insertValuesSpec   = 11
	SqlParserRULE_insertMultiValue   = 12
	SqlParserRULE_dataValue          = 13
	SqlParserRULE_selectStatement    = 14
	SqlParserRULE_fromClause         = 15
	SqlParserRULE_joinExpr           = 16
	SqlParserRULE_useStatement       = 17
	SqlParserRULE_columnTypeExpr     = 18
	SqlParserRULE_columnExprList     = 19
	SqlParserRULE_columnsExpr        = 20
	SqlParserRULE_columnExpr         = 21
	SqlParserRULE_columnIdentifier   = 22
	SqlParserRULE_nestedIdentifier   = 23
	SqlParserRULE_constant           = 24
	SqlParserRULE_realLiteral        = 25
	SqlParserRULE_tableExpr          = 26
	SqlParserRULE_tableIdentifier    = 27
	SqlParserRULE_databaseIdentifier = 28
	SqlParserRULE_floatingLiteral    = 29
	SqlParserRULE_numberLiteral      = 30
	SqlParserRULE_literal            = 31
	SqlParserRULE_interval           = 32
	SqlParserRULE_keyword            = 33
	SqlParserRULE_identifier         = 34
	SqlParserRULE_enumValue          = 35
	SqlParserRULE_ifNotExists        = 36
)

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) DdlStatement() IDdlStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDdlStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDdlStatementContext)
}

func (s *StatementContext) DmlStatement() IDmlStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDmlStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDmlStatementContext)
}

func (s *StatementContext) SystemStatement() ISystemStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISystemStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISystemStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_statement)

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

	p.SetState(77)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserCREATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.DdlStatement()
		}

	case SqlParserINSERT, SqlParserSELECT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(75)
			p.DmlStatement()
		}

	case SqlParserUSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(76)
			p.SystemStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDmlStatementContext is an interface to support dynamic dispatch.
type IDmlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDmlStatementContext differentiates from other interfaces.
	IsDmlStatementContext()
}

type DmlStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDmlStatementContext() *DmlStatementContext {
	var p = new(DmlStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_dmlStatement
	return p
}

func (*DmlStatementContext) IsDmlStatementContext() {}

func NewDmlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DmlStatementContext {
	var p = new(DmlStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_dmlStatement

	return p
}

func (s *DmlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DmlStatementContext) SelectStatement() ISelectStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectStatementContext)
}

func (s *DmlStatementContext) InsertStatement() IInsertStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertStatementContext)
}

func (s *DmlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DmlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DmlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDmlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DmlStatement() (localctx IDmlStatementContext) {
	this := p
	_ = this

	localctx = NewDmlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_dmlStatement)

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

	p.SetState(81)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSELECT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(79)
			p.SelectStatement()
		}

	case SqlParserINSERT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(80)
			p.InsertStatement()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDdlStatementContext is an interface to support dynamic dispatch.
type IDdlStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDdlStatementContext differentiates from other interfaces.
	IsDdlStatementContext()
}

type DdlStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDdlStatementContext() *DdlStatementContext {
	var p = new(DdlStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_ddlStatement
	return p
}

func (*DdlStatementContext) IsDdlStatementContext() {}

func NewDdlStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DdlStatementContext {
	var p = new(DdlStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_ddlStatement

	return p
}

func (s *DdlStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DdlStatementContext) CreateDatabase() ICreateDatabaseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateDatabaseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateDatabaseContext)
}

func (s *DdlStatementContext) CreateTable() ICreateTableContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICreateTableContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICreateTableContext)
}

func (s *DdlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DdlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DdlStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDdlStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DdlStatement() (localctx IDdlStatementContext) {
	this := p
	_ = this

	localctx = NewDdlStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_ddlStatement)

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

	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(83)
			p.CreateDatabase()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(84)
			p.CreateTable()
		}

	}

	return localctx
}

// ISystemStatementContext is an interface to support dynamic dispatch.
type ISystemStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSystemStatementContext differentiates from other interfaces.
	IsSystemStatementContext()
}

type SystemStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySystemStatementContext() *SystemStatementContext {
	var p = new(SystemStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_systemStatement
	return p
}

func (*SystemStatementContext) IsSystemStatementContext() {}

func NewSystemStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SystemStatementContext {
	var p = new(SystemStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_systemStatement

	return p
}

func (s *SystemStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SystemStatementContext) UseStatement() IUseStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUseStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUseStatementContext)
}

func (s *SystemStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SystemStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SystemStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSystemStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SystemStatement() (localctx ISystemStatementContext) {
	this := p
	_ = this

	localctx = NewSystemStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_systemStatement)

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
		p.SetState(87)
		p.UseStatement()
	}

	return localctx
}

// ICreateDatabaseContext is an interface to support dynamic dispatch.
type ICreateDatabaseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateDatabaseContext differentiates from other interfaces.
	IsCreateDatabaseContext()
}

type CreateDatabaseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateDatabaseContext() *CreateDatabaseContext {
	var p = new(CreateDatabaseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_createDatabase
	return p
}

func (*CreateDatabaseContext) IsCreateDatabaseContext() {}

func NewCreateDatabaseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateDatabaseContext {
	var p = new(CreateDatabaseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_createDatabase

	return p
}

func (s *CreateDatabaseContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateDatabaseContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SqlParserCREATE, 0)
}

func (s *CreateDatabaseContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(SqlParserDATABASE, 0)
}

func (s *CreateDatabaseContext) DatabaseIdentifier() IDatabaseIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabaseIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabaseIdentifierContext)
}

func (s *CreateDatabaseContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateDatabaseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateDatabaseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateDatabaseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitCreateDatabase(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) CreateDatabase() (localctx ICreateDatabaseContext) {
	this := p
	_ = this

	localctx = NewCreateDatabaseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_createDatabase)

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
		p.SetState(89)
		p.Match(SqlParserCREATE)
	}
	{
		p.SetState(90)
		p.Match(SqlParserDATABASE)
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(91)
			p.IfNotExists()
		}

	}
	{
		p.SetState(94)
		p.DatabaseIdentifier()
	}

	return localctx
}

// ICreateTableContext is an interface to support dynamic dispatch.
type ICreateTableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCreateTableContext differentiates from other interfaces.
	IsCreateTableContext()
}

type CreateTableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCreateTableContext() *CreateTableContext {
	var p = new(CreateTableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_createTable
	return p
}

func (*CreateTableContext) IsCreateTableContext() {}

func NewCreateTableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreateTableContext {
	var p = new(CreateTableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_createTable

	return p
}

func (s *CreateTableContext) GetParser() antlr.Parser { return s.parser }

func (s *CreateTableContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SqlParserCREATE, 0)
}

func (s *CreateTableContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserTABLE, 0)
}

func (s *CreateTableContext) TableIdentifier() ITableIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableIdentifierContext)
}

func (s *CreateTableContext) TableSchemaClause() ITableSchemaClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableSchemaClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableSchemaClauseContext)
}

func (s *CreateTableContext) IfNotExists() IIfNotExistsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfNotExistsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfNotExistsContext)
}

func (s *CreateTableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreateTableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreateTableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitCreateTable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) CreateTable() (localctx ICreateTableContext) {
	this := p
	_ = this

	localctx = NewCreateTableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_createTable)

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
		p.SetState(96)
		p.Match(SqlParserCREATE)
	}
	{
		p.SetState(97)
		p.Match(SqlParserTABLE)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(98)
			p.IfNotExists()
		}

	}
	{
		p.SetState(101)
		p.TableIdentifier()
	}
	{
		p.SetState(102)
		p.TableSchemaClause()
	}

	return localctx
}

// ITableSchemaClauseContext is an interface to support dynamic dispatch.
type ITableSchemaClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableSchemaClauseContext differentiates from other interfaces.
	IsTableSchemaClauseContext()
}

type TableSchemaClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableSchemaClauseContext() *TableSchemaClauseContext {
	var p = new(TableSchemaClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableSchemaClause
	return p
}

func (*TableSchemaClauseContext) IsTableSchemaClauseContext() {}

func NewTableSchemaClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableSchemaClauseContext {
	var p = new(TableSchemaClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableSchemaClause

	return p
}

func (s *TableSchemaClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *TableSchemaClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserLPAREN, 0)
}

func (s *TableSchemaClauseContext) AllTableElementExpr() []ITableElementExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableElementExprContext); ok {
			len++
		}
	}

	tst := make([]ITableElementExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableElementExprContext); ok {
			tst[i] = t.(ITableElementExprContext)
			i++
		}
	}

	return tst
}

func (s *TableSchemaClauseContext) TableElementExpr(i int) ITableElementExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableElementExprContext); ok {
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

	return t.(ITableElementExprContext)
}

func (s *TableSchemaClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserRPAREN, 0)
}

func (s *TableSchemaClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *TableSchemaClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *TableSchemaClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableSchemaClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableSchemaClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitTableSchemaClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableSchemaClause() (localctx ITableSchemaClauseContext) {
	this := p
	_ = this

	localctx = NewTableSchemaClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_tableSchemaClause)
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
		p.SetState(104)
		p.Match(SqlParserLPAREN)
	}
	{
		p.SetState(105)
		p.TableElementExpr()
	}
	p.SetState(110)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(106)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(107)
			p.TableElementExpr()
		}

		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(113)
		p.Match(SqlParserRPAREN)
	}

	return localctx
}

// ITableElementExprContext is an interface to support dynamic dispatch.
type ITableElementExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableElementExprContext differentiates from other interfaces.
	IsTableElementExprContext()
}

type TableElementExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableElementExprContext() *TableElementExprContext {
	var p = new(TableElementExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableElementExpr
	return p
}

func (*TableElementExprContext) IsTableElementExprContext() {}

func NewTableElementExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableElementExprContext {
	var p = new(TableElementExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableElementExpr

	return p
}

func (s *TableElementExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TableElementExprContext) TableColumnDfnt() ITableColumnDfntContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableColumnDfntContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableColumnDfntContext)
}

func (s *TableElementExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableElementExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableElementExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitTableElementExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableElementExpr() (localctx ITableElementExprContext) {
	this := p
	_ = this

	localctx = NewTableElementExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_tableElementExpr)

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
		p.SetState(115)
		p.TableColumnDfnt()
	}

	return localctx
}

// ITableColumnDfntContext is an interface to support dynamic dispatch.
type ITableColumnDfntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableColumnDfntContext differentiates from other interfaces.
	IsTableColumnDfntContext()
}

type TableColumnDfntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableColumnDfntContext() *TableColumnDfntContext {
	var p = new(TableColumnDfntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableColumnDfnt
	return p
}

func (*TableColumnDfntContext) IsTableColumnDfntContext() {}

func NewTableColumnDfntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableColumnDfntContext {
	var p = new(TableColumnDfntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableColumnDfnt

	return p
}

func (s *TableColumnDfntContext) GetParser() antlr.Parser { return s.parser }

func (s *TableColumnDfntContext) NestedIdentifier() INestedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedIdentifierContext)
}

func (s *TableColumnDfntContext) ColumnTypeExpr() IColumnTypeExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnTypeExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnTypeExprContext)
}

func (s *TableColumnDfntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableColumnDfntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableColumnDfntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitTableColumnDfnt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableColumnDfnt() (localctx ITableColumnDfntContext) {
	this := p
	_ = this

	localctx = NewTableColumnDfntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_tableColumnDfnt)

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
		p.SetState(117)
		p.NestedIdentifier()
	}
	{
		p.SetState(118)
		p.ColumnTypeExpr()
	}

	return localctx
}

// IInsertStatementContext is an interface to support dynamic dispatch.
type IInsertStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertStatementContext differentiates from other interfaces.
	IsInsertStatementContext()
}

type InsertStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertStatementContext() *InsertStatementContext {
	var p = new(InsertStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_insertStatement
	return p
}

func (*InsertStatementContext) IsInsertStatementContext() {}

func NewInsertStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertStatementContext {
	var p = new(InsertStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_insertStatement

	return p
}

func (s *InsertStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertStatementContext) INSERT() antlr.TerminalNode {
	return s.GetToken(SqlParserINSERT, 0)
}

func (s *InsertStatementContext) INTO() antlr.TerminalNode {
	return s.GetToken(SqlParserINTO, 0)
}

func (s *InsertStatementContext) VALUES() antlr.TerminalNode {
	return s.GetToken(SqlParserVALUES, 0)
}

func (s *InsertStatementContext) InsertValuesSpec() IInsertValuesSpecContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertValuesSpecContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInsertValuesSpecContext)
}

func (s *InsertStatementContext) TableIdentifier() ITableIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableIdentifierContext)
}

func (s *InsertStatementContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserTABLE, 0)
}

func (s *InsertStatementContext) ColumnsClause() IColumnsClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnsClauseContext)
}

func (s *InsertStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitInsertStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) InsertStatement() (localctx IInsertStatementContext) {
	this := p
	_ = this

	localctx = NewInsertStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SqlParserRULE_insertStatement)
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
		p.SetState(120)
		p.Match(SqlParserINSERT)
	}
	{
		p.SetState(121)
		p.Match(SqlParserINTO)
	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(122)
			p.Match(SqlParserTABLE)
		}

	}

	{
		p.SetState(125)
		p.TableIdentifier()
	}

	p.SetState(127)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserLPAREN {
		{
			p.SetState(126)
			p.ColumnsClause()
		}

	}
	{
		p.SetState(129)
		p.Match(SqlParserVALUES)
	}
	{
		p.SetState(130)
		p.InsertValuesSpec()
	}

	return localctx
}

// IColumnsClauseContext is an interface to support dynamic dispatch.
type IColumnsClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsClauseContext differentiates from other interfaces.
	IsColumnsClauseContext()
}

type ColumnsClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsClauseContext() *ColumnsClauseContext {
	var p = new(ColumnsClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnsClause
	return p
}

func (*ColumnsClauseContext) IsColumnsClauseContext() {}

func NewColumnsClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsClauseContext {
	var p = new(ColumnsClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnsClause

	return p
}

func (s *ColumnsClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsClauseContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserLPAREN, 0)
}

func (s *ColumnsClauseContext) AllNestedIdentifier() []INestedIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INestedIdentifierContext); ok {
			len++
		}
	}

	tst := make([]INestedIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INestedIdentifierContext); ok {
			tst[i] = t.(INestedIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ColumnsClauseContext) NestedIdentifier(i int) INestedIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedIdentifierContext); ok {
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

	return t.(INestedIdentifierContext)
}

func (s *ColumnsClauseContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserRPAREN, 0)
}

func (s *ColumnsClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *ColumnsClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *ColumnsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnsClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnsClause() (localctx IColumnsClauseContext) {
	this := p
	_ = this

	localctx = NewColumnsClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SqlParserRULE_columnsClause)
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
		p.SetState(132)
		p.Match(SqlParserLPAREN)
	}
	{
		p.SetState(133)
		p.NestedIdentifier()
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(134)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(135)
			p.NestedIdentifier()
		}

		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(141)
		p.Match(SqlParserRPAREN)
	}

	return localctx
}

// IInsertValuesSpecContext is an interface to support dynamic dispatch.
type IInsertValuesSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertValuesSpecContext differentiates from other interfaces.
	IsInsertValuesSpecContext()
}

type InsertValuesSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertValuesSpecContext() *InsertValuesSpecContext {
	var p = new(InsertValuesSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_insertValuesSpec
	return p
}

func (*InsertValuesSpecContext) IsInsertValuesSpecContext() {}

func NewInsertValuesSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertValuesSpecContext {
	var p = new(InsertValuesSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_insertValuesSpec

	return p
}

func (s *InsertValuesSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertValuesSpecContext) AllInsertMultiValue() []IInsertMultiValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInsertMultiValueContext); ok {
			len++
		}
	}

	tst := make([]IInsertMultiValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInsertMultiValueContext); ok {
			tst[i] = t.(IInsertMultiValueContext)
			i++
		}
	}

	return tst
}

func (s *InsertValuesSpecContext) InsertMultiValue(i int) IInsertMultiValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInsertMultiValueContext); ok {
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

	return t.(IInsertMultiValueContext)
}

func (s *InsertValuesSpecContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *InsertValuesSpecContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *InsertValuesSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertValuesSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertValuesSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitInsertValuesSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) InsertValuesSpec() (localctx IInsertValuesSpecContext) {
	this := p
	_ = this

	localctx = NewInsertValuesSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SqlParserRULE_insertValuesSpec)
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
	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA || _la == SqlParserLPAREN {
		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserCOMMA {
			{
				p.SetState(143)
				p.Match(SqlParserCOMMA)
			}

		}
		{
			p.SetState(146)
			p.InsertMultiValue()
		}

		p.SetState(151)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInsertMultiValueContext is an interface to support dynamic dispatch.
type IInsertMultiValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertMultiValueContext differentiates from other interfaces.
	IsInsertMultiValueContext()
}

type InsertMultiValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertMultiValueContext() *InsertMultiValueContext {
	var p = new(InsertMultiValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_insertMultiValue
	return p
}

func (*InsertMultiValueContext) IsInsertMultiValueContext() {}

func NewInsertMultiValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertMultiValueContext {
	var p = new(InsertMultiValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_insertMultiValue

	return p
}

func (s *InsertMultiValueContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertMultiValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserLPAREN, 0)
}

func (s *InsertMultiValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserRPAREN, 0)
}

func (s *InsertMultiValueContext) AllDataValue() []IDataValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDataValueContext); ok {
			len++
		}
	}

	tst := make([]IDataValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDataValueContext); ok {
			tst[i] = t.(IDataValueContext)
			i++
		}
	}

	return tst
}

func (s *InsertMultiValueContext) DataValue(i int) IDataValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDataValueContext); ok {
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

	return t.(IDataValueContext)
}

func (s *InsertMultiValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *InsertMultiValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *InsertMultiValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertMultiValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertMultiValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitInsertMultiValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) InsertMultiValue() (localctx IInsertMultiValueContext) {
	this := p
	_ = this

	localctx = NewInsertMultiValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_insertMultiValue)
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
		p.SetState(152)
		p.Match(SqlParserLPAREN)
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&38293818889469952) != 0 {
		{
			p.SetState(153)
			p.DataValue()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserCOMMA {
			{
				p.SetState(154)
				p.Match(SqlParserCOMMA)
			}

		}

		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(161)
		p.Match(SqlParserRPAREN)
	}

	return localctx
}

// IDataValueContext is an interface to support dynamic dispatch.
type IDataValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDataValueContext differentiates from other interfaces.
	IsDataValueContext()
}

type DataValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDataValueContext() *DataValueContext {
	var p = new(DataValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_dataValue
	return p
}

func (*DataValueContext) IsDataValueContext() {}

func NewDataValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DataValueContext {
	var p = new(DataValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_dataValue

	return p
}

func (s *DataValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DataValueContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *DataValueContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
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

	return t.(IConstantContext)
}

func (s *DataValueContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserLPAREN, 0)
}

func (s *DataValueContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserRPAREN, 0)
}

func (s *DataValueContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *DataValueContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *DataValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DataValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DataValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDataValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DataValue() (localctx IDataValueContext) {
	this := p
	_ = this

	localctx = NewDataValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SqlParserRULE_dataValue)
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

	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserDECIMAL_LITERAL, SqlParserEXPONENT_NUM_PART, SqlParserSTRING_LITERAL, SqlParserDASH, SqlParserDOT, SqlParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.Constant()
		}

	case SqlParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(165)
			p.Constant()
		}
		p.SetState(168)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SqlParserCOMMA {
			{
				p.SetState(166)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(167)
				p.Constant()
			}

			p.SetState(170)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(172)
			p.Match(SqlParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *SelectStatementContext) ColumnExprList() IColumnExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnExprListContext)
}

func (s *SelectStatementContext) FromClause() IFromClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFromClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFromClauseContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectStatement() (localctx ISelectStatementContext) {
	this := p
	_ = this

	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SqlParserRULE_selectStatement)
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
		p.SetState(176)
		p.Match(SqlParserSELECT)
	}
	{
		p.SetState(177)
		p.ColumnExprList()
	}
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserFROM {
		{
			p.SetState(178)
			p.FromClause()
		}

	}

	return localctx
}

// IFromClauseContext is an interface to support dynamic dispatch.
type IFromClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFromClauseContext differentiates from other interfaces.
	IsFromClauseContext()
}

type FromClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFromClauseContext() *FromClauseContext {
	var p = new(FromClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_fromClause
	return p
}

func (*FromClauseContext) IsFromClauseContext() {}

func NewFromClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FromClauseContext {
	var p = new(FromClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_fromClause

	return p
}

func (s *FromClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *FromClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *FromClauseContext) JoinExpr() IJoinExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IJoinExprContext)
}

func (s *FromClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FromClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FromClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitFromClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FromClause() (localctx IFromClauseContext) {
	this := p
	_ = this

	localctx = NewFromClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SqlParserRULE_fromClause)

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
		p.SetState(181)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(182)
		p.JoinExpr()
	}

	return localctx
}

// IJoinExprContext is an interface to support dynamic dispatch.
type IJoinExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinExprContext differentiates from other interfaces.
	IsJoinExprContext()
}

type JoinExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinExprContext() *JoinExprContext {
	var p = new(JoinExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_joinExpr
	return p
}

func (*JoinExprContext) IsJoinExprContext() {}

func NewJoinExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinExprContext {
	var p = new(JoinExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_joinExpr

	return p
}

func (s *JoinExprContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinExprContext) TableExpr() ITableExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableExprContext)
}

func (s *JoinExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitJoinExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) JoinExpr() (localctx IJoinExprContext) {
	this := p
	_ = this

	localctx = NewJoinExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SqlParserRULE_joinExpr)

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
		p.SetState(184)
		p.TableExpr()
	}

	return localctx
}

// IUseStatementContext is an interface to support dynamic dispatch.
type IUseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUseStatementContext differentiates from other interfaces.
	IsUseStatementContext()
}

type UseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUseStatementContext() *UseStatementContext {
	var p = new(UseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_useStatement
	return p
}

func (*UseStatementContext) IsUseStatementContext() {}

func NewUseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UseStatementContext {
	var p = new(UseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_useStatement

	return p
}

func (s *UseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UseStatementContext) USE() antlr.TerminalNode {
	return s.GetToken(SqlParserUSE, 0)
}

func (s *UseStatementContext) DatabaseIdentifier() IDatabaseIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabaseIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabaseIdentifierContext)
}

func (s *UseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) UseStatement() (localctx IUseStatementContext) {
	this := p
	_ = this

	localctx = NewUseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SqlParserRULE_useStatement)

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
		p.SetState(186)
		p.Match(SqlParserUSE)
	}
	{
		p.SetState(187)
		p.DatabaseIdentifier()
	}

	return localctx
}

// IColumnTypeExprContext is an interface to support dynamic dispatch.
type IColumnTypeExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnTypeExprContext differentiates from other interfaces.
	IsColumnTypeExprContext()
}

type ColumnTypeExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnTypeExprContext() *ColumnTypeExprContext {
	var p = new(ColumnTypeExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnTypeExpr
	return p
}

func (*ColumnTypeExprContext) IsColumnTypeExprContext() {}

func NewColumnTypeExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnTypeExprContext {
	var p = new(ColumnTypeExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnTypeExpr

	return p
}

func (s *ColumnTypeExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnTypeExprContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *ColumnTypeExprContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *ColumnTypeExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserLPAREN, 0)
}

func (s *ColumnTypeExprContext) AllColumnTypeExpr() []IColumnTypeExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnTypeExprContext); ok {
			len++
		}
	}

	tst := make([]IColumnTypeExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnTypeExprContext); ok {
			tst[i] = t.(IColumnTypeExprContext)
			i++
		}
	}

	return tst
}

func (s *ColumnTypeExprContext) ColumnTypeExpr(i int) IColumnTypeExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnTypeExprContext); ok {
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

	return t.(IColumnTypeExprContext)
}

func (s *ColumnTypeExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserRPAREN, 0)
}

func (s *ColumnTypeExprContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *ColumnTypeExprContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *ColumnTypeExprContext) AllEnumValue() []IEnumValueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueContext); ok {
			tst[i] = t.(IEnumValueContext)
			i++
		}
	}

	return tst
}

func (s *ColumnTypeExprContext) EnumValue(i int) IEnumValueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
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

	return t.(IEnumValueContext)
}

func (s *ColumnTypeExprContext) ColumnExprList() IColumnExprListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnExprListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnExprListContext)
}

func (s *ColumnTypeExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnTypeExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnTypeExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnTypeExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnTypeExpr() (localctx IColumnTypeExprContext) {
	this := p
	_ = this

	localctx = NewColumnTypeExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SqlParserRULE_columnTypeExpr)
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

	p.SetState(236)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(190)
			p.Identifier()
		}
		{
			p.SetState(191)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(192)
			p.Identifier()
		}
		{
			p.SetState(193)
			p.ColumnTypeExpr()
		}
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(194)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(195)
				p.Identifier()
			}
			{
				p.SetState(196)
				p.ColumnTypeExpr()
			}

			p.SetState(202)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(203)
			p.Match(SqlParserRPAREN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(205)
			p.Identifier()
		}
		{
			p.SetState(206)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(207)
			p.EnumValue()
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(208)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(209)
				p.EnumValue()
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(215)
			p.Match(SqlParserRPAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(217)
			p.Identifier()
		}
		{
			p.SetState(218)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(219)
			p.ColumnTypeExpr()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(220)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(221)
				p.ColumnTypeExpr()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(227)
			p.Match(SqlParserRPAREN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(229)
			p.Identifier()
		}
		{
			p.SetState(230)
			p.Match(SqlParserLPAREN)
		}
		p.SetState(232)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&36042016928169980) != 0 {
			{
				p.SetState(231)
				p.ColumnExprList()
			}

		}
		{
			p.SetState(234)
			p.Match(SqlParserRPAREN)
		}

	}

	return localctx
}

// IColumnExprListContext is an interface to support dynamic dispatch.
type IColumnExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnExprListContext differentiates from other interfaces.
	IsColumnExprListContext()
}

type ColumnExprListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnExprListContext() *ColumnExprListContext {
	var p = new(ColumnExprListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnExprList
	return p
}

func (*ColumnExprListContext) IsColumnExprListContext() {}

func NewColumnExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnExprListContext {
	var p = new(ColumnExprListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnExprList

	return p
}

func (s *ColumnExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnExprListContext) AllColumnsExpr() []IColumnsExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnsExprContext); ok {
			len++
		}
	}

	tst := make([]IColumnsExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnsExprContext); ok {
			tst[i] = t.(IColumnsExprContext)
			i++
		}
	}

	return tst
}

func (s *ColumnExprListContext) ColumnsExpr(i int) IColumnsExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnsExprContext); ok {
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

	return t.(IColumnsExprContext)
}

func (s *ColumnExprListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *ColumnExprListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *ColumnExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnExprListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnExprList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnExprList() (localctx IColumnExprListContext) {
	this := p
	_ = this

	localctx = NewColumnExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SqlParserRULE_columnExprList)
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
		p.SetState(238)
		p.ColumnsExpr()
	}
	p.SetState(243)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(239)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(240)
			p.ColumnsExpr()
		}

		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IColumnsExprContext is an interface to support dynamic dispatch.
type IColumnsExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnsExprContext differentiates from other interfaces.
	IsColumnsExprContext()
}

type ColumnsExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnsExprContext() *ColumnsExprContext {
	var p = new(ColumnsExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnsExpr
	return p
}

func (*ColumnsExprContext) IsColumnsExprContext() {}

func NewColumnsExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnsExprContext {
	var p = new(ColumnsExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnsExpr

	return p
}

func (s *ColumnsExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnsExprContext) ColumnExpr() IColumnExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnExprContext)
}

func (s *ColumnsExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnsExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnsExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnsExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnsExpr() (localctx IColumnsExprContext) {
	this := p
	_ = this

	localctx = NewColumnsExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_columnsExpr)

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
		p.SetState(246)
		p.ColumnExpr()
	}

	return localctx
}

// IColumnExprContext is an interface to support dynamic dispatch.
type IColumnExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnExprContext differentiates from other interfaces.
	IsColumnExprContext()
}

type ColumnExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnExprContext() *ColumnExprContext {
	var p = new(ColumnExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnExpr
	return p
}

func (*ColumnExprContext) IsColumnExprContext() {}

func NewColumnExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnExprContext {
	var p = new(ColumnExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnExpr

	return p
}

func (s *ColumnExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnExprContext) Literal() ILiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ColumnExprContext) ColumnIdentifier() IColumnIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IColumnIdentifierContext)
}

func (s *ColumnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnExpr() (localctx IColumnExprContext) {
	this := p
	_ = this

	localctx = NewColumnExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SqlParserRULE_columnExpr)

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

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(248)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(249)
			p.ColumnIdentifier()
		}

	}

	return localctx
}

// IColumnIdentifierContext is an interface to support dynamic dispatch.
type IColumnIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsColumnIdentifierContext differentiates from other interfaces.
	IsColumnIdentifierContext()
}

type ColumnIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyColumnIdentifierContext() *ColumnIdentifierContext {
	var p = new(ColumnIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnIdentifier
	return p
}

func (*ColumnIdentifierContext) IsColumnIdentifierContext() {}

func NewColumnIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnIdentifierContext {
	var p = new(ColumnIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnIdentifier

	return p
}

func (s *ColumnIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnIdentifierContext) NestedIdentifier() INestedIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INestedIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INestedIdentifierContext)
}

func (s *ColumnIdentifierContext) TableIdentifier() ITableIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableIdentifierContext)
}

func (s *ColumnIdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT, 0)
}

func (s *ColumnIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnIdentifier() (localctx IColumnIdentifierContext) {
	this := p
	_ = this

	localctx = NewColumnIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SqlParserRULE_columnIdentifier)

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
	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(252)
			p.TableIdentifier()
		}
		{
			p.SetState(253)
			p.Match(SqlParserDOT)
		}

	}
	{
		p.SetState(257)
		p.NestedIdentifier()
	}

	return localctx
}

// INestedIdentifierContext is an interface to support dynamic dispatch.
type INestedIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNestedIdentifierContext differentiates from other interfaces.
	IsNestedIdentifierContext()
}

type NestedIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNestedIdentifierContext() *NestedIdentifierContext {
	var p = new(NestedIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_nestedIdentifier
	return p
}

func (*NestedIdentifierContext) IsNestedIdentifierContext() {}

func NewNestedIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NestedIdentifierContext {
	var p = new(NestedIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_nestedIdentifier

	return p
}

func (s *NestedIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *NestedIdentifierContext) AllIdentifier() []IIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentifierContext); ok {
			tst[i] = t.(IIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *NestedIdentifierContext) Identifier(i int) IIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
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

	return t.(IIdentifierContext)
}

func (s *NestedIdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT, 0)
}

func (s *NestedIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NestedIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NestedIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitNestedIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) NestedIdentifier() (localctx INestedIdentifierContext) {
	this := p
	_ = this

	localctx = NewNestedIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SqlParserRULE_nestedIdentifier)
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
		p.SetState(259)
		p.Identifier()
	}
	p.SetState(262)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserDOT {
		{
			p.SetState(260)
			p.Match(SqlParserDOT)
		}
		{
			p.SetState(261)
			p.Identifier()
		}

	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) RealLiteral() IRealLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRealLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRealLiteralContext)
}

func (s *ConstantContext) DASH() antlr.TerminalNode {
	return s.GetToken(SqlParserDASH, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SqlParserPLUS, 0)
}

func (s *ConstantContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *ConstantContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING_LITERAL, 0)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Constant() (localctx IConstantContext) {
	this := p
	_ = this

	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SqlParserRULE_constant)
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

	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserDASH || _la == SqlParserPLUS {
			{
				p.SetState(264)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SqlParserDASH || _la == SqlParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(267)
			p.RealLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(269)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserDASH || _la == SqlParserPLUS {
			{
				p.SetState(268)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SqlParserDASH || _la == SqlParserPLUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(271)
			p.Match(SqlParserDECIMAL_LITERAL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)
			p.Match(SqlParserSTRING_LITERAL)
		}

	}

	return localctx
}

// IRealLiteralContext is an interface to support dynamic dispatch.
type IRealLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRealLiteralContext differentiates from other interfaces.
	IsRealLiteralContext()
}

type RealLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRealLiteralContext() *RealLiteralContext {
	var p = new(RealLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_realLiteral
	return p
}

func (*RealLiteralContext) IsRealLiteralContext() {}

func NewRealLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RealLiteralContext {
	var p = new(RealLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_realLiteral

	return p
}

func (s *RealLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RealLiteralContext) AllDECIMAL_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(SqlParserDECIMAL_LITERAL)
}

func (s *RealLiteralContext) DECIMAL_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, i)
}

func (s *RealLiteralContext) DOT() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT, 0)
}

func (s *RealLiteralContext) EXPONENT_NUM_PART() antlr.TerminalNode {
	return s.GetToken(SqlParserEXPONENT_NUM_PART, 0)
}

func (s *RealLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RealLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RealLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitRealLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) RealLiteral() (localctx IRealLiteralContext) {
	this := p
	_ = this

	localctx = NewRealLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SqlParserRULE_realLiteral)
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

	p.SetState(283)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(275)
			p.Match(SqlParserDECIMAL_LITERAL)
		}
		{
			p.SetState(276)
			p.Match(SqlParserDOT)
		}
		p.SetState(278)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(277)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SqlParserDECIMAL_LITERAL || _la == SqlParserEXPONENT_NUM_PART) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	case SqlParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(280)
			p.Match(SqlParserDOT)
		}
		{
			p.SetState(281)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserDECIMAL_LITERAL || _la == SqlParserEXPONENT_NUM_PART) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case SqlParserEXPONENT_NUM_PART:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(282)
			p.Match(SqlParserEXPONENT_NUM_PART)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITableExprContext is an interface to support dynamic dispatch.
type ITableExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableExprContext differentiates from other interfaces.
	IsTableExprContext()
}

type TableExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableExprContext() *TableExprContext {
	var p = new(TableExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableExpr
	return p
}

func (*TableExprContext) IsTableExprContext() {}

func NewTableExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableExprContext {
	var p = new(TableExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableExpr

	return p
}

func (s *TableExprContext) GetParser() antlr.Parser { return s.parser }

func (s *TableExprContext) TableIdentifier() ITableIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITableIdentifierContext)
}

func (s *TableExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitTableExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableExpr() (localctx ITableExprContext) {
	this := p
	_ = this

	localctx = NewTableExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SqlParserRULE_tableExpr)

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
		p.SetState(285)
		p.TableIdentifier()
	}

	return localctx
}

// ITableIdentifierContext is an interface to support dynamic dispatch.
type ITableIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTableIdentifierContext differentiates from other interfaces.
	IsTableIdentifierContext()
}

type TableIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTableIdentifierContext() *TableIdentifierContext {
	var p = new(TableIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_tableIdentifier
	return p
}

func (*TableIdentifierContext) IsTableIdentifierContext() {}

func NewTableIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TableIdentifierContext {
	var p = new(TableIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_tableIdentifier

	return p
}

func (s *TableIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *TableIdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *TableIdentifierContext) DatabaseIdentifier() IDatabaseIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatabaseIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatabaseIdentifierContext)
}

func (s *TableIdentifierContext) DOT() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT, 0)
}

func (s *TableIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TableIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TableIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitTableIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) TableIdentifier() (localctx ITableIdentifierContext) {
	this := p
	_ = this

	localctx = NewTableIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SqlParserRULE_tableIdentifier)

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
	p.SetState(290)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(287)
			p.DatabaseIdentifier()
		}
		{
			p.SetState(288)
			p.Match(SqlParserDOT)
		}

	}
	{
		p.SetState(292)
		p.Identifier()
	}

	return localctx
}

// IDatabaseIdentifierContext is an interface to support dynamic dispatch.
type IDatabaseIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDatabaseIdentifierContext differentiates from other interfaces.
	IsDatabaseIdentifierContext()
}

type DatabaseIdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatabaseIdentifierContext() *DatabaseIdentifierContext {
	var p = new(DatabaseIdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_databaseIdentifier
	return p
}

func (*DatabaseIdentifierContext) IsDatabaseIdentifierContext() {}

func NewDatabaseIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatabaseIdentifierContext {
	var p = new(DatabaseIdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_databaseIdentifier

	return p
}

func (s *DatabaseIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *DatabaseIdentifierContext) Identifier() IIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierContext)
}

func (s *DatabaseIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatabaseIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatabaseIdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDatabaseIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DatabaseIdentifier() (localctx IDatabaseIdentifierContext) {
	this := p
	_ = this

	localctx = NewDatabaseIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SqlParserRULE_databaseIdentifier)

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
		p.SetState(294)
		p.Identifier()
	}

	return localctx
}

// IFloatingLiteralContext is an interface to support dynamic dispatch.
type IFloatingLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatingLiteralContext differentiates from other interfaces.
	IsFloatingLiteralContext()
}

type FloatingLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatingLiteralContext() *FloatingLiteralContext {
	var p = new(FloatingLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_floatingLiteral
	return p
}

func (*FloatingLiteralContext) IsFloatingLiteralContext() {}

func NewFloatingLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatingLiteralContext {
	var p = new(FloatingLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_floatingLiteral

	return p
}

func (s *FloatingLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatingLiteralContext) FLOATING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserFLOATING_LITERAL, 0)
}

func (s *FloatingLiteralContext) DOT() antlr.TerminalNode {
	return s.GetToken(SqlParserDOT, 0)
}

func (s *FloatingLiteralContext) AllDECIMAL_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(SqlParserDECIMAL_LITERAL)
}

func (s *FloatingLiteralContext) DECIMAL_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, i)
}

func (s *FloatingLiteralContext) OCTAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserOCTAL_LITERAL, 0)
}

func (s *FloatingLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatingLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatingLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitFloatingLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FloatingLiteral() (localctx IFloatingLiteralContext) {
	this := p
	_ = this

	localctx = NewFloatingLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SqlParserRULE_floatingLiteral)
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

	p.SetState(304)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserFLOATING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(296)
			p.Match(SqlParserFLOATING_LITERAL)
		}

	case SqlParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(297)
			p.Match(SqlParserDOT)
		}
		{
			p.SetState(298)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserOCTAL_LITERAL || _la == SqlParserDECIMAL_LITERAL) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case SqlParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(299)
			p.Match(SqlParserDECIMAL_LITERAL)
		}
		{
			p.SetState(300)
			p.Match(SqlParserDOT)
		}
		p.SetState(302)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserOCTAL_LITERAL || _la == SqlParserDECIMAL_LITERAL {
			{
				p.SetState(301)
				_la = p.GetTokenStream().LA(1)

				if !(_la == SqlParserOCTAL_LITERAL || _la == SqlParserDECIMAL_LITERAL) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// INumberLiteralContext is an interface to support dynamic dispatch.
type INumberLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNumberLiteralContext differentiates from other interfaces.
	IsNumberLiteralContext()
}

type NumberLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumberLiteralContext() *NumberLiteralContext {
	var p = new(NumberLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_numberLiteral
	return p
}

func (*NumberLiteralContext) IsNumberLiteralContext() {}

func NewNumberLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_numberLiteral

	return p
}

func (s *NumberLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumberLiteralContext) FloatingLiteral() IFloatingLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatingLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatingLiteralContext)
}

func (s *NumberLiteralContext) OCTAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserOCTAL_LITERAL, 0)
}

func (s *NumberLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *NumberLiteralContext) HEXADECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserHEXADECIMAL_LITERAL, 0)
}

func (s *NumberLiteralContext) INF() antlr.TerminalNode {
	return s.GetToken(SqlParserINF, 0)
}

func (s *NumberLiteralContext) NAN_SQL() antlr.TerminalNode {
	return s.GetToken(SqlParserNAN_SQL, 0)
}

func (s *NumberLiteralContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SqlParserPLUS, 0)
}

func (s *NumberLiteralContext) DASH() antlr.TerminalNode {
	return s.GetToken(SqlParserDASH, 0)
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumberLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitNumberLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) NumberLiteral() (localctx INumberLiteralContext) {
	this := p
	_ = this

	localctx = NewNumberLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SqlParserRULE_numberLiteral)
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
	p.SetState(307)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserDASH || _la == SqlParserPLUS {
		{
			p.SetState(306)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserDASH || _la == SqlParserPLUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(309)
			p.FloatingLiteral()
		}

	case 2:
		{
			p.SetState(310)
			p.Match(SqlParserOCTAL_LITERAL)
		}

	case 3:
		{
			p.SetState(311)
			p.Match(SqlParserDECIMAL_LITERAL)
		}

	case 4:
		{
			p.SetState(312)
			p.Match(SqlParserHEXADECIMAL_LITERAL)
		}

	case 5:
		{
			p.SetState(313)
			p.Match(SqlParserINF)
		}

	case 6:
		{
			p.SetState(314)
			p.Match(SqlParserNAN_SQL)
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) NumberLiteral() INumberLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *LiteralContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING_LITERAL, 0)
}

func (s *LiteralContext) NULL_SQL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL_SQL, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Literal() (localctx ILiteralContext) {
	this := p
	_ = this

	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SqlParserRULE_literal)

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

	p.SetState(320)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserINF, SqlParserNAN_SQL, SqlParserFLOATING_LITERAL, SqlParserOCTAL_LITERAL, SqlParserDECIMAL_LITERAL, SqlParserHEXADECIMAL_LITERAL, SqlParserDASH, SqlParserDOT, SqlParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(317)
			p.NumberLiteral()
		}

	case SqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(318)
			p.Match(SqlParserSTRING_LITERAL)
		}

	case SqlParserNULL_SQL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.Match(SqlParserNULL_SQL)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIntervalContext is an interface to support dynamic dispatch.
type IIntervalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntervalContext differentiates from other interfaces.
	IsIntervalContext()
}

type IntervalContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntervalContext() *IntervalContext {
	var p = new(IntervalContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_interval
	return p
}

func (*IntervalContext) IsIntervalContext() {}

func NewIntervalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntervalContext {
	var p = new(IntervalContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_interval

	return p
}

func (s *IntervalContext) GetParser() antlr.Parser { return s.parser }

func (s *IntervalContext) SECOND() antlr.TerminalNode {
	return s.GetToken(SqlParserSECOND, 0)
}

func (s *IntervalContext) MINUTE() antlr.TerminalNode {
	return s.GetToken(SqlParserMINUTE, 0)
}

func (s *IntervalContext) HOUR() antlr.TerminalNode {
	return s.GetToken(SqlParserHOUR, 0)
}

func (s *IntervalContext) DAY() antlr.TerminalNode {
	return s.GetToken(SqlParserDAY, 0)
}

func (s *IntervalContext) WEEK() antlr.TerminalNode {
	return s.GetToken(SqlParserWEEK, 0)
}

func (s *IntervalContext) MONTH() antlr.TerminalNode {
	return s.GetToken(SqlParserMONTH, 0)
}

func (s *IntervalContext) QUARTER() antlr.TerminalNode {
	return s.GetToken(SqlParserQUARTER, 0)
}

func (s *IntervalContext) YEAR() antlr.TerminalNode {
	return s.GetToken(SqlParserYEAR, 0)
}

func (s *IntervalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntervalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntervalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitInterval(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Interval() (localctx IIntervalContext) {
	this := p
	_ = this

	localctx = NewIntervalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SqlParserRULE_interval)
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
		p.SetState(322)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&267386880) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IKeywordContext is an interface to support dynamic dispatch.
type IKeywordContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsKeywordContext differentiates from other interfaces.
	IsKeywordContext()
}

type KeywordContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordContext() *KeywordContext {
	var p = new(KeywordContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_keyword
	return p
}

func (*KeywordContext) IsKeywordContext() {}

func NewKeywordContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordContext {
	var p = new(KeywordContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_keyword

	return p
}

func (s *KeywordContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordContext) CREATE() antlr.TerminalNode {
	return s.GetToken(SqlParserCREATE, 0)
}

func (s *KeywordContext) DATABASE() antlr.TerminalNode {
	return s.GetToken(SqlParserDATABASE, 0)
}

func (s *KeywordContext) DEFAULT() antlr.TerminalNode {
	return s.GetToken(SqlParserDEFAULT, 0)
}

func (s *KeywordContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SqlParserEXISTS, 0)
}

func (s *KeywordContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *KeywordContext) IF() antlr.TerminalNode {
	return s.GetToken(SqlParserIF, 0)
}

func (s *KeywordContext) INSERT() antlr.TerminalNode {
	return s.GetToken(SqlParserINSERT, 0)
}

func (s *KeywordContext) INTO() antlr.TerminalNode {
	return s.GetToken(SqlParserINTO, 0)
}

func (s *KeywordContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *KeywordContext) NULL_SQL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL_SQL, 0)
}

func (s *KeywordContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *KeywordContext) SET() antlr.TerminalNode {
	return s.GetToken(SqlParserSET, 0)
}

func (s *KeywordContext) TABLE() antlr.TerminalNode {
	return s.GetToken(SqlParserTABLE, 0)
}

func (s *KeywordContext) VALUE() antlr.TerminalNode {
	return s.GetToken(SqlParserVALUE, 0)
}

func (s *KeywordContext) VALUES() antlr.TerminalNode {
	return s.GetToken(SqlParserVALUES, 0)
}

func (s *KeywordContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitKeyword(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Keyword() (localctx IKeywordContext) {
	this := p
	_ = this

	localctx = NewKeywordContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SqlParserRULE_keyword)
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
		p.SetState(324)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&915196) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IIdentifierContext is an interface to support dynamic dispatch.
type IIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIdentifierContext differentiates from other interfaces.
	IsIdentifierContext()
}

type IdentifierContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierContext() *IdentifierContext {
	var p = new(IdentifierContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_identifier
	return p
}

func (*IdentifierContext) IsIdentifierContext() {}

func NewIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierContext {
	var p = new(IdentifierContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_identifier

	return p
}

func (s *IdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(SqlParserIDENTIFIER, 0)
}

func (s *IdentifierContext) Interval() IIntervalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntervalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntervalContext)
}

func (s *IdentifierContext) Keyword() IKeywordContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordContext)
}

func (s *IdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitIdentifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Identifier() (localctx IIdentifierContext) {
	this := p
	_ = this

	localctx = NewIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SqlParserRULE_identifier)

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

	p.SetState(329)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.Match(SqlParserIDENTIFIER)
		}

	case SqlParserSECOND, SqlParserMINUTE, SqlParserHOUR, SqlParserDAY, SqlParserWEEK, SqlParserMONTH, SqlParserQUARTER, SqlParserYEAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(327)
			p.Interval()
		}

	case SqlParserCREATE, SqlParserDATABASE, SqlParserDEFAULT, SqlParserEXISTS, SqlParserFROM, SqlParserIF, SqlParserINSERT, SqlParserINTO, SqlParserNOT, SqlParserNULL_SQL, SqlParserSELECT, SqlParserSET, SqlParserTABLE, SqlParserVALUE, SqlParserVALUES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(328)
			p.Keyword()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_enumValue
	return p
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING_LITERAL, 0)
}

func (s *EnumValueContext) EQ_SINGLE() antlr.TerminalNode {
	return s.GetToken(SqlParserEQ_SINGLE, 0)
}

func (s *EnumValueContext) NumberLiteral() INumberLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumberLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumberLiteralContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitEnumValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) EnumValue() (localctx IEnumValueContext) {
	this := p
	_ = this

	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SqlParserRULE_enumValue)

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
		p.SetState(331)
		p.Match(SqlParserSTRING_LITERAL)
	}
	{
		p.SetState(332)
		p.Match(SqlParserEQ_SINGLE)
	}
	{
		p.SetState(333)
		p.NumberLiteral()
	}

	return localctx
}

// IIfNotExistsContext is an interface to support dynamic dispatch.
type IIfNotExistsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfNotExistsContext differentiates from other interfaces.
	IsIfNotExistsContext()
}

type IfNotExistsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfNotExistsContext() *IfNotExistsContext {
	var p = new(IfNotExistsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_ifNotExists
	return p
}

func (*IfNotExistsContext) IsIfNotExistsContext() {}

func NewIfNotExistsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfNotExistsContext {
	var p = new(IfNotExistsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_ifNotExists

	return p
}

func (s *IfNotExistsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfNotExistsContext) IF() antlr.TerminalNode {
	return s.GetToken(SqlParserIF, 0)
}

func (s *IfNotExistsContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *IfNotExistsContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(SqlParserEXISTS, 0)
}

func (s *IfNotExistsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfNotExistsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfNotExistsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitIfNotExists(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) IfNotExists() (localctx IIfNotExistsContext) {
	this := p
	_ = this

	localctx = NewIfNotExistsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SqlParserRULE_ifNotExists)

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
		p.SetState(335)
		p.Match(SqlParserIF)
	}
	{
		p.SetState(336)
		p.Match(SqlParserNOT)
	}
	{
		p.SetState(337)
		p.Match(SqlParserEXISTS)
	}

	return localctx
}
