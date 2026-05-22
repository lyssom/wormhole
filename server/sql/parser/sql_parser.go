// Code generated from SqlParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

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
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "'->'", "'*'", "'`'", "'\\'", "':'", "','",
		"'||'", "'-'", "'.'", "'=='", "'='", "'>='", "'>'", "'{'", "'['", "'<='",
		"'('", "'<'", "", "'%'", "'+'", "'?'", "'\"'", "'''", "'}'", "']'",
		"')'", "';'", "'/'", "'_'",
	}
	staticData.symbolicNames = []string{
		"", "WS", "AND", "ASC", "AVG", "BY", "COUNT", "CREATE", "DATABASE",
		"DEFAULT", "DELETE", "DESC", "EXISTS", "FROM", "GROUP", "HAVING", "IF",
		"INF", "INNER", "INSERT", "INTO", "JOIN", "LEFT", "LIMIT", "MAX", "MIN",
		"NAN_SQL", "NOT", "NULL_SQL", "OFFSET", "ON", "OR", "ORDER", "SELECT",
		"SET", "SUM", "TABLE", "UPDATE", "USE", "VALUE", "VALUES", "WHERE",
		"SECOND", "MINUTE", "HOUR", "DAY", "WEEK", "MONTH", "QUARTER", "YEAR",
		"IDENTIFIER", "FLOATING_LITERAL", "OCTAL_LITERAL", "DECIMAL_LITERAL",
		"HEXADECIMAL_LITERAL", "EXPONENT_NUM_PART", "STRING_LITERAL", "ARROW",
		"ASTERISK", "BACKQUOTE", "BACKSLASH", "COLON", "COMMA", "CONCAT", "DASH",
		"DOT", "EQ_DOUBLE", "EQ_SINGLE", "GE", "GT", "LBRACE", "LBRACKET", "LE",
		"LPAREN", "LT", "NOT_EQ", "PERCENT", "PLUS", "QUERY", "QUOTE_DOUBLE",
		"QUOTE_SINGLE", "RBRACE", "RBRACKET", "RPAREN", "SEMICOLON", "SLASH",
		"UNDERSCORE", "MULTI_LINE_COMMENT", "SINGLE_LINE_COMMENT", "WHITESPACE",
	}
	staticData.ruleNames = []string{
		"statement", "dmlStatement", "ddlStatement", "systemStatement", "createDatabase",
		"createTable", "tableSchemaClause", "tableElementExpr", "tableColumnDfnt",
		"insertStatement", "updateStatement", "deleteStatement", "setClause",
		"columnsClause", "insertValuesSpec", "insertMultiValue", "dataValue",
		"selectStatement", "fromClause", "whereClause", "groupByClause", "havingClause",
		"orderByClause", "orderByExpr", "limitClause", "offsetClause", "searchCondition",
		"predicate", "expression", "joinExpr", "joinOp", "useStatement", "columnTypeExpr",
		"columnExprList", "columnsExpr", "columnExpr", "aggregateFunction",
		"columnIdentifier", "nestedIdentifier", "constant", "realLiteral", "tableExpr",
		"tableIdentifier", "databaseIdentifier", "floatingLiteral", "numberLiteral",
		"literal", "interval", "keyword", "identifier", "enumValue", "ifNotExists",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 89, 565, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 1, 0, 1,
		0, 1, 0, 3, 0, 108, 8, 0, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 114, 8, 1, 1, 2,
		1, 2, 3, 2, 118, 8, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 3, 4, 125, 8, 4, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 132, 8, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1,
		6, 1, 6, 1, 6, 5, 6, 141, 8, 6, 10, 6, 12, 6, 144, 9, 6, 1, 6, 1, 6, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 3, 9, 156, 8, 9, 1, 9, 1,
		9, 3, 9, 160, 8, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 1, 10,
		1, 10, 5, 10, 171, 8, 10, 10, 10, 12, 10, 174, 9, 10, 1, 10, 3, 10, 177,
		8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 183, 8, 11, 1, 12, 1, 12, 1,
		12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 5, 13, 193, 8, 13, 10, 13, 12, 13,
		196, 9, 13, 1, 13, 1, 13, 1, 14, 3, 14, 201, 8, 14, 1, 14, 5, 14, 204,
		8, 14, 10, 14, 12, 14, 207, 9, 14, 1, 15, 1, 15, 1, 15, 3, 15, 212, 8,
		15, 4, 15, 214, 8, 15, 11, 15, 12, 15, 215, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 4, 16, 225, 8, 16, 11, 16, 12, 16, 226, 1, 16, 1,
		16, 3, 16, 231, 8, 16, 1, 17, 1, 17, 1, 17, 3, 17, 236, 8, 17, 1, 17, 3,
		17, 239, 8, 17, 1, 17, 3, 17, 242, 8, 17, 1, 17, 3, 17, 245, 8, 17, 1,
		17, 3, 17, 248, 8, 17, 1, 17, 3, 17, 251, 8, 17, 1, 17, 3, 17, 254, 8,
		17, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 20, 5, 20, 267, 8, 20, 10, 20, 12, 20, 270, 9, 20, 1, 21, 1, 21, 1,
		21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 22, 5, 22, 280, 8, 22, 10, 22, 12, 22,
		283, 9, 22, 1, 23, 1, 23, 3, 23, 287, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25,
		1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 3, 26, 299, 8, 26, 1, 26, 1,
		26, 1, 26, 1, 26, 1, 26, 1, 26, 5, 26, 307, 8, 26, 10, 26, 12, 26, 310,
		9, 26, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1,
		27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 3, 27, 336, 8, 27, 1, 28, 1, 28, 1, 28, 3,
		28, 341, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28,
		1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 5, 28, 358, 8, 28, 10,
		28, 12, 28, 361, 9, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 5, 29,
		369, 8, 29, 10, 29, 12, 29, 372, 9, 29, 1, 30, 3, 30, 375, 8, 30, 1, 30,
		1, 30, 1, 30, 3, 30, 380, 8, 30, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5, 32, 394, 8, 32, 10, 32,
		12, 32, 397, 9, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 5,
		32, 406, 8, 32, 10, 32, 12, 32, 409, 9, 32, 1, 32, 1, 32, 1, 32, 1, 32,
		1, 32, 1, 32, 1, 32, 5, 32, 418, 8, 32, 10, 32, 12, 32, 421, 9, 32, 1,
		32, 1, 32, 1, 32, 1, 32, 1, 32, 3, 32, 428, 8, 32, 1, 32, 1, 32, 3, 32,
		432, 8, 32, 1, 33, 1, 33, 1, 33, 5, 33, 437, 8, 33, 10, 33, 12, 33, 440,
		9, 33, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 3, 35, 447, 8, 35, 1, 36, 1,
		36, 1, 36, 1, 36, 3, 36, 453, 8, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 476, 8, 36, 1, 37, 1, 37,
		1, 37, 3, 37, 481, 8, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 3, 38, 488,
		8, 38, 1, 39, 3, 39, 491, 8, 39, 1, 39, 1, 39, 3, 39, 495, 8, 39, 1, 39,
		1, 39, 3, 39, 499, 8, 39, 1, 40, 1, 40, 1, 40, 3, 40, 504, 8, 40, 1, 40,
		1, 40, 1, 40, 3, 40, 509, 8, 40, 1, 41, 1, 41, 1, 42, 1, 42, 1, 42, 3,
		42, 516, 8, 42, 1, 42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 44, 1, 44,
		1, 44, 1, 44, 3, 44, 528, 8, 44, 3, 44, 530, 8, 44, 1, 45, 3, 45, 533,
		8, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 1, 45, 3, 45, 541, 8, 45, 1,
		46, 1, 46, 1, 46, 3, 46, 546, 8, 46, 1, 47, 1, 47, 1, 48, 1, 48, 1, 49,
		1, 49, 1, 49, 3, 49, 555, 8, 49, 1, 50, 1, 50, 1, 50, 1, 50, 1, 51, 1,
		51, 1, 51, 1, 51, 1, 51, 0, 2, 52, 56, 52, 0, 2, 4, 6, 8, 10, 12, 14, 16,
		18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52,
		54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88,
		90, 92, 94, 96, 98, 100, 102, 0, 6, 2, 0, 3, 3, 11, 11, 2, 0, 64, 64, 77,
		77, 2, 0, 53, 53, 55, 55, 1, 0, 52, 53, 1, 0, 42, 49, 8, 0, 7, 9, 12, 13,
		16, 16, 19, 20, 27, 28, 33, 34, 36, 36, 39, 40, 599, 0, 107, 1, 0, 0, 0,
		2, 113, 1, 0, 0, 0, 4, 117, 1, 0, 0, 0, 6, 119, 1, 0, 0, 0, 8, 121, 1,
		0, 0, 0, 10, 128, 1, 0, 0, 0, 12, 136, 1, 0, 0, 0, 14, 147, 1, 0, 0, 0,
		16, 149, 1, 0, 0, 0, 18, 152, 1, 0, 0, 0, 20, 164, 1, 0, 0, 0, 22, 178,
		1, 0, 0, 0, 24, 184, 1, 0, 0, 0, 26, 188, 1, 0, 0, 0, 28, 205, 1, 0, 0,
		0, 30, 208, 1, 0, 0, 0, 32, 230, 1, 0, 0, 0, 34, 232, 1, 0, 0, 0, 36, 255,
		1, 0, 0, 0, 38, 258, 1, 0, 0, 0, 40, 261, 1, 0, 0, 0, 42, 271, 1, 0, 0,
		0, 44, 274, 1, 0, 0, 0, 46, 284, 1, 0, 0, 0, 48, 288, 1, 0, 0, 0, 50, 291,
		1, 0, 0, 0, 52, 298, 1, 0, 0, 0, 54, 335, 1, 0, 0, 0, 56, 340, 1, 0, 0,
		0, 58, 362, 1, 0, 0, 0, 60, 379, 1, 0, 0, 0, 62, 381, 1, 0, 0, 0, 64, 431,
		1, 0, 0, 0, 66, 433, 1, 0, 0, 0, 68, 441, 1, 0, 0, 0, 70, 446, 1, 0, 0,
		0, 72, 475, 1, 0, 0, 0, 74, 480, 1, 0, 0, 0, 76, 484, 1, 0, 0, 0, 78, 498,
		1, 0, 0, 0, 80, 508, 1, 0, 0, 0, 82, 510, 1, 0, 0, 0, 84, 515, 1, 0, 0,
		0, 86, 519, 1, 0, 0, 0, 88, 529, 1, 0, 0, 0, 90, 532, 1, 0, 0, 0, 92, 545,
		1, 0, 0, 0, 94, 547, 1, 0, 0, 0, 96, 549, 1, 0, 0, 0, 98, 554, 1, 0, 0,
		0, 100, 556, 1, 0, 0, 0, 102, 560, 1, 0, 0, 0, 104, 108, 3, 4, 2, 0, 105,
		108, 3, 2, 1, 0, 106, 108, 3, 6, 3, 0, 107, 104, 1, 0, 0, 0, 107, 105,
		1, 0, 0, 0, 107, 106, 1, 0, 0, 0, 108, 1, 1, 0, 0, 0, 109, 114, 3, 34,
		17, 0, 110, 114, 3, 18, 9, 0, 111, 114, 3, 20, 10, 0, 112, 114, 3, 22,
		11, 0, 113, 109, 1, 0, 0, 0, 113, 110, 1, 0, 0, 0, 113, 111, 1, 0, 0, 0,
		113, 112, 1, 0, 0, 0, 114, 3, 1, 0, 0, 0, 115, 118, 3, 8, 4, 0, 116, 118,
		3, 10, 5, 0, 117, 115, 1, 0, 0, 0, 117, 116, 1, 0, 0, 0, 118, 5, 1, 0,
		0, 0, 119, 120, 3, 62, 31, 0, 120, 7, 1, 0, 0, 0, 121, 122, 5, 7, 0, 0,
		122, 124, 5, 8, 0, 0, 123, 125, 3, 102, 51, 0, 124, 123, 1, 0, 0, 0, 124,
		125, 1, 0, 0, 0, 125, 126, 1, 0, 0, 0, 126, 127, 3, 86, 43, 0, 127, 9,
		1, 0, 0, 0, 128, 129, 5, 7, 0, 0, 129, 131, 5, 36, 0, 0, 130, 132, 3, 102,
		51, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0,
		133, 134, 3, 84, 42, 0, 134, 135, 3, 12, 6, 0, 135, 11, 1, 0, 0, 0, 136,
		137, 5, 73, 0, 0, 137, 142, 3, 14, 7, 0, 138, 139, 5, 62, 0, 0, 139, 141,
		3, 14, 7, 0, 140, 138, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0,
		0, 0, 142, 143, 1, 0, 0, 0, 143, 145, 1, 0, 0, 0, 144, 142, 1, 0, 0, 0,
		145, 146, 5, 83, 0, 0, 146, 13, 1, 0, 0, 0, 147, 148, 3, 16, 8, 0, 148,
		15, 1, 0, 0, 0, 149, 150, 3, 76, 38, 0, 150, 151, 3, 64, 32, 0, 151, 17,
		1, 0, 0, 0, 152, 153, 5, 19, 0, 0, 153, 155, 5, 20, 0, 0, 154, 156, 5,
		36, 0, 0, 155, 154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0,
		0, 157, 159, 3, 84, 42, 0, 158, 160, 3, 26, 13, 0, 159, 158, 1, 0, 0, 0,
		159, 160, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 40, 0, 0, 162,
		163, 3, 28, 14, 0, 163, 19, 1, 0, 0, 0, 164, 165, 5, 37, 0, 0, 165, 166,
		3, 84, 42, 0, 166, 167, 5, 34, 0, 0, 167, 172, 3, 24, 12, 0, 168, 169,
		5, 62, 0, 0, 169, 171, 3, 24, 12, 0, 170, 168, 1, 0, 0, 0, 171, 174, 1,
		0, 0, 0, 172, 170, 1, 0, 0, 0, 172, 173, 1, 0, 0, 0, 173, 176, 1, 0, 0,
		0, 174, 172, 1, 0, 0, 0, 175, 177, 3, 38, 19, 0, 176, 175, 1, 0, 0, 0,
		176, 177, 1, 0, 0, 0, 177, 21, 1, 0, 0, 0, 178, 179, 5, 10, 0, 0, 179,
		180, 5, 13, 0, 0, 180, 182, 3, 84, 42, 0, 181, 183, 3, 38, 19, 0, 182,
		181, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183, 23, 1, 0, 0, 0, 184, 185, 3,
		74, 37, 0, 185, 186, 5, 67, 0, 0, 186, 187, 3, 56, 28, 0, 187, 25, 1, 0,
		0, 0, 188, 189, 5, 73, 0, 0, 189, 194, 3, 76, 38, 0, 190, 191, 5, 62, 0,
		0, 191, 193, 3, 76, 38, 0, 192, 190, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0,
		194, 192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196,
		194, 1, 0, 0, 0, 197, 198, 5, 83, 0, 0, 198, 27, 1, 0, 0, 0, 199, 201,
		5, 62, 0, 0, 200, 199, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0, 201, 202, 1, 0,
		0, 0, 202, 204, 3, 30, 15, 0, 203, 200, 1, 0, 0, 0, 204, 207, 1, 0, 0,
		0, 205, 203, 1, 0, 0, 0, 205, 206, 1, 0, 0, 0, 206, 29, 1, 0, 0, 0, 207,
		205, 1, 0, 0, 0, 208, 213, 5, 73, 0, 0, 209, 211, 3, 32, 16, 0, 210, 212,
		5, 62, 0, 0, 211, 210, 1, 0, 0, 0, 211, 212, 1, 0, 0, 0, 212, 214, 1, 0,
		0, 0, 213, 209, 1, 0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 213, 1, 0, 0, 0,
		215, 216, 1, 0, 0, 0, 216, 217, 1, 0, 0, 0, 217, 218, 5, 83, 0, 0, 218,
		31, 1, 0, 0, 0, 219, 231, 3, 78, 39, 0, 220, 221, 5, 73, 0, 0, 221, 224,
		3, 78, 39, 0, 222, 223, 5, 62, 0, 0, 223, 225, 3, 78, 39, 0, 224, 222,
		1, 0, 0, 0, 225, 226, 1, 0, 0, 0, 226, 224, 1, 0, 0, 0, 226, 227, 1, 0,
		0, 0, 227, 228, 1, 0, 0, 0, 228, 229, 5, 83, 0, 0, 229, 231, 1, 0, 0, 0,
		230, 219, 1, 0, 0, 0, 230, 220, 1, 0, 0, 0, 231, 33, 1, 0, 0, 0, 232, 233,
		5, 33, 0, 0, 233, 235, 3, 66, 33, 0, 234, 236, 3, 36, 18, 0, 235, 234,
		1, 0, 0, 0, 235, 236, 1, 0, 0, 0, 236, 238, 1, 0, 0, 0, 237, 239, 3, 38,
		19, 0, 238, 237, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 241, 1, 0, 0, 0,
		240, 242, 3, 40, 20, 0, 241, 240, 1, 0, 0, 0, 241, 242, 1, 0, 0, 0, 242,
		244, 1, 0, 0, 0, 243, 245, 3, 42, 21, 0, 244, 243, 1, 0, 0, 0, 244, 245,
		1, 0, 0, 0, 245, 247, 1, 0, 0, 0, 246, 248, 3, 44, 22, 0, 247, 246, 1,
		0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 250, 1, 0, 0, 0, 249, 251, 3, 48, 24,
		0, 250, 249, 1, 0, 0, 0, 250, 251, 1, 0, 0, 0, 251, 253, 1, 0, 0, 0, 252,
		254, 3, 50, 25, 0, 253, 252, 1, 0, 0, 0, 253, 254, 1, 0, 0, 0, 254, 35,
		1, 0, 0, 0, 255, 256, 5, 13, 0, 0, 256, 257, 3, 58, 29, 0, 257, 37, 1,
		0, 0, 0, 258, 259, 5, 41, 0, 0, 259, 260, 3, 52, 26, 0, 260, 39, 1, 0,
		0, 0, 261, 262, 5, 14, 0, 0, 262, 263, 5, 5, 0, 0, 263, 268, 3, 74, 37,
		0, 264, 265, 5, 62, 0, 0, 265, 267, 3, 74, 37, 0, 266, 264, 1, 0, 0, 0,
		267, 270, 1, 0, 0, 0, 268, 266, 1, 0, 0, 0, 268, 269, 1, 0, 0, 0, 269,
		41, 1, 0, 0, 0, 270, 268, 1, 0, 0, 0, 271, 272, 5, 15, 0, 0, 272, 273,
		3, 52, 26, 0, 273, 43, 1, 0, 0, 0, 274, 275, 5, 32, 0, 0, 275, 276, 5,
		5, 0, 0, 276, 281, 3, 46, 23, 0, 277, 278, 5, 62, 0, 0, 278, 280, 3, 46,
		23, 0, 279, 277, 1, 0, 0, 0, 280, 283, 1, 0, 0, 0, 281, 279, 1, 0, 0, 0,
		281, 282, 1, 0, 0, 0, 282, 45, 1, 0, 0, 0, 283, 281, 1, 0, 0, 0, 284, 286,
		3, 74, 37, 0, 285, 287, 7, 0, 0, 0, 286, 285, 1, 0, 0, 0, 286, 287, 1,
		0, 0, 0, 287, 47, 1, 0, 0, 0, 288, 289, 5, 23, 0, 0, 289, 290, 5, 53, 0,
		0, 290, 49, 1, 0, 0, 0, 291, 292, 5, 29, 0, 0, 292, 293, 5, 53, 0, 0, 293,
		51, 1, 0, 0, 0, 294, 295, 6, 26, -1, 0, 295, 296, 5, 27, 0, 0, 296, 299,
		3, 52, 26, 2, 297, 299, 3, 54, 27, 0, 298, 294, 1, 0, 0, 0, 298, 297, 1,
		0, 0, 0, 299, 308, 1, 0, 0, 0, 300, 301, 10, 4, 0, 0, 301, 302, 5, 2, 0,
		0, 302, 307, 3, 52, 26, 5, 303, 304, 10, 3, 0, 0, 304, 305, 5, 31, 0, 0,
		305, 307, 3, 52, 26, 4, 306, 300, 1, 0, 0, 0, 306, 303, 1, 0, 0, 0, 307,
		310, 1, 0, 0, 0, 308, 306, 1, 0, 0, 0, 308, 309, 1, 0, 0, 0, 309, 53, 1,
		0, 0, 0, 310, 308, 1, 0, 0, 0, 311, 312, 3, 56, 28, 0, 312, 313, 5, 67,
		0, 0, 313, 314, 3, 56, 28, 0, 314, 336, 1, 0, 0, 0, 315, 316, 3, 56, 28,
		0, 316, 317, 5, 75, 0, 0, 317, 318, 3, 56, 28, 0, 318, 336, 1, 0, 0, 0,
		319, 320, 3, 56, 28, 0, 320, 321, 5, 74, 0, 0, 321, 322, 3, 56, 28, 0,
		322, 336, 1, 0, 0, 0, 323, 324, 3, 56, 28, 0, 324, 325, 5, 72, 0, 0, 325,
		326, 3, 56, 28, 0, 326, 336, 1, 0, 0, 0, 327, 328, 3, 56, 28, 0, 328, 329,
		5, 69, 0, 0, 329, 330, 3, 56, 28, 0, 330, 336, 1, 0, 0, 0, 331, 332, 3,
		56, 28, 0, 332, 333, 5, 68, 0, 0, 333, 334, 3, 56, 28, 0, 334, 336, 1,
		0, 0, 0, 335, 311, 1, 0, 0, 0, 335, 315, 1, 0, 0, 0, 335, 319, 1, 0, 0,
		0, 335, 323, 1, 0, 0, 0, 335, 327, 1, 0, 0, 0, 335, 331, 1, 0, 0, 0, 336,
		55, 1, 0, 0, 0, 337, 338, 6, 28, -1, 0, 338, 341, 3, 92, 46, 0, 339, 341,
		3, 74, 37, 0, 340, 337, 1, 0, 0, 0, 340, 339, 1, 0, 0, 0, 341, 359, 1,
		0, 0, 0, 342, 343, 10, 5, 0, 0, 343, 344, 5, 58, 0, 0, 344, 358, 3, 56,
		28, 6, 345, 346, 10, 4, 0, 0, 346, 347, 5, 85, 0, 0, 347, 358, 3, 56, 28,
		5, 348, 349, 10, 3, 0, 0, 349, 350, 5, 76, 0, 0, 350, 358, 3, 56, 28, 4,
		351, 352, 10, 2, 0, 0, 352, 353, 5, 77, 0, 0, 353, 358, 3, 56, 28, 3, 354,
		355, 10, 1, 0, 0, 355, 356, 5, 64, 0, 0, 356, 358, 3, 56, 28, 2, 357, 342,
		1, 0, 0, 0, 357, 345, 1, 0, 0, 0, 357, 348, 1, 0, 0, 0, 357, 351, 1, 0,
		0, 0, 357, 354, 1, 0, 0, 0, 358, 361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0,
		359, 360, 1, 0, 0, 0, 360, 57, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 370,
		3, 82, 41, 0, 363, 364, 3, 60, 30, 0, 364, 365, 3, 82, 41, 0, 365, 366,
		5, 30, 0, 0, 366, 367, 3, 52, 26, 0, 367, 369, 1, 0, 0, 0, 368, 363, 1,
		0, 0, 0, 369, 372, 1, 0, 0, 0, 370, 368, 1, 0, 0, 0, 370, 371, 1, 0, 0,
		0, 371, 59, 1, 0, 0, 0, 372, 370, 1, 0, 0, 0, 373, 375, 5, 18, 0, 0, 374,
		373, 1, 0, 0, 0, 374, 375, 1, 0, 0, 0, 375, 376, 1, 0, 0, 0, 376, 380,
		5, 21, 0, 0, 377, 378, 5, 22, 0, 0, 378, 380, 5, 21, 0, 0, 379, 374, 1,
		0, 0, 0, 379, 377, 1, 0, 0, 0, 380, 61, 1, 0, 0, 0, 381, 382, 5, 38, 0,
		0, 382, 383, 3, 86, 43, 0, 383, 63, 1, 0, 0, 0, 384, 432, 3, 98, 49, 0,
		385, 386, 3, 98, 49, 0, 386, 387, 5, 73, 0, 0, 387, 388, 3, 98, 49, 0,
		388, 395, 3, 64, 32, 0, 389, 390, 5, 62, 0, 0, 390, 391, 3, 98, 49, 0,
		391, 392, 3, 64, 32, 0, 392, 394, 1, 0, 0, 0, 393, 389, 1, 0, 0, 0, 394,
		397, 1, 0, 0, 0, 395, 393, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396, 398,
		1, 0, 0, 0, 397, 395, 1, 0, 0, 0, 398, 399, 5, 83, 0, 0, 399, 432, 1, 0,
		0, 0, 400, 401, 3, 98, 49, 0, 401, 402, 5, 73, 0, 0, 402, 407, 3, 100,
		50, 0, 403, 404, 5, 62, 0, 0, 404, 406, 3, 100, 50, 0, 405, 403, 1, 0,
		0, 0, 406, 409, 1, 0, 0, 0, 407, 405, 1, 0, 0, 0, 407, 408, 1, 0, 0, 0,
		408, 410, 1, 0, 0, 0, 409, 407, 1, 0, 0, 0, 410, 411, 5, 83, 0, 0, 411,
		432, 1, 0, 0, 0, 412, 413, 3, 98, 49, 0, 413, 414, 5, 73, 0, 0, 414, 419,
		3, 64, 32, 0, 415, 416, 5, 62, 0, 0, 416, 418, 3, 64, 32, 0, 417, 415,
		1, 0, 0, 0, 418, 421, 1, 0, 0, 0, 419, 417, 1, 0, 0, 0, 419, 420, 1, 0,
		0, 0, 420, 422, 1, 0, 0, 0, 421, 419, 1, 0, 0, 0, 422, 423, 5, 83, 0, 0,
		423, 432, 1, 0, 0, 0, 424, 425, 3, 98, 49, 0, 425, 427, 5, 73, 0, 0, 426,
		428, 3, 66, 33, 0, 427, 426, 1, 0, 0, 0, 427, 428, 1, 0, 0, 0, 428, 429,
		1, 0, 0, 0, 429, 430, 5, 83, 0, 0, 430, 432, 1, 0, 0, 0, 431, 384, 1, 0,
		0, 0, 431, 385, 1, 0, 0, 0, 431, 400, 1, 0, 0, 0, 431, 412, 1, 0, 0, 0,
		431, 424, 1, 0, 0, 0, 432, 65, 1, 0, 0, 0, 433, 438, 3, 68, 34, 0, 434,
		435, 5, 62, 0, 0, 435, 437, 3, 68, 34, 0, 436, 434, 1, 0, 0, 0, 437, 440,
		1, 0, 0, 0, 438, 436, 1, 0, 0, 0, 438, 439, 1, 0, 0, 0, 439, 67, 1, 0,
		0, 0, 440, 438, 1, 0, 0, 0, 441, 442, 3, 70, 35, 0, 442, 69, 1, 0, 0, 0,
		443, 447, 3, 92, 46, 0, 444, 447, 3, 74, 37, 0, 445, 447, 3, 72, 36, 0,
		446, 443, 1, 0, 0, 0, 446, 444, 1, 0, 0, 0, 446, 445, 1, 0, 0, 0, 447,
		71, 1, 0, 0, 0, 448, 449, 5, 6, 0, 0, 449, 452, 5, 73, 0, 0, 450, 453,
		5, 58, 0, 0, 451, 453, 3, 56, 28, 0, 452, 450, 1, 0, 0, 0, 452, 451, 1,
		0, 0, 0, 453, 454, 1, 0, 0, 0, 454, 476, 5, 83, 0, 0, 455, 456, 5, 35,
		0, 0, 456, 457, 5, 73, 0, 0, 457, 458, 3, 56, 28, 0, 458, 459, 5, 83, 0,
		0, 459, 476, 1, 0, 0, 0, 460, 461, 5, 4, 0, 0, 461, 462, 5, 73, 0, 0, 462,
		463, 3, 56, 28, 0, 463, 464, 5, 83, 0, 0, 464, 476, 1, 0, 0, 0, 465, 466,
		5, 25, 0, 0, 466, 467, 5, 73, 0, 0, 467, 468, 3, 56, 28, 0, 468, 469, 5,
		83, 0, 0, 469, 476, 1, 0, 0, 0, 470, 471, 5, 24, 0, 0, 471, 472, 5, 73,
		0, 0, 472, 473, 3, 56, 28, 0, 473, 474, 5, 83, 0, 0, 474, 476, 1, 0, 0,
		0, 475, 448, 1, 0, 0, 0, 475, 455, 1, 0, 0, 0, 475, 460, 1, 0, 0, 0, 475,
		465, 1, 0, 0, 0, 475, 470, 1, 0, 0, 0, 476, 73, 1, 0, 0, 0, 477, 478, 3,
		84, 42, 0, 478, 479, 5, 65, 0, 0, 479, 481, 1, 0, 0, 0, 480, 477, 1, 0,
		0, 0, 480, 481, 1, 0, 0, 0, 481, 482, 1, 0, 0, 0, 482, 483, 3, 76, 38,
		0, 483, 75, 1, 0, 0, 0, 484, 487, 3, 98, 49, 0, 485, 486, 5, 65, 0, 0,
		486, 488, 3, 98, 49, 0, 487, 485, 1, 0, 0, 0, 487, 488, 1, 0, 0, 0, 488,
		77, 1, 0, 0, 0, 489, 491, 7, 1, 0, 0, 490, 489, 1, 0, 0, 0, 490, 491, 1,
		0, 0, 0, 491, 492, 1, 0, 0, 0, 492, 499, 3, 80, 40, 0, 493, 495, 7, 1,
		0, 0, 494, 493, 1, 0, 0, 0, 494, 495, 1, 0, 0, 0, 495, 496, 1, 0, 0, 0,
		496, 499, 5, 53, 0, 0, 497, 499, 5, 56, 0, 0, 498, 490, 1, 0, 0, 0, 498,
		494, 1, 0, 0, 0, 498, 497, 1, 0, 0, 0, 499, 79, 1, 0, 0, 0, 500, 501, 5,
		53, 0, 0, 501, 503, 5, 65, 0, 0, 502, 504, 7, 2, 0, 0, 503, 502, 1, 0,
		0, 0, 503, 504, 1, 0, 0, 0, 504, 509, 1, 0, 0, 0, 505, 506, 5, 65, 0, 0,
		506, 509, 7, 2, 0, 0, 507, 509, 5, 55, 0, 0, 508, 500, 1, 0, 0, 0, 508,
		505, 1, 0, 0, 0, 508, 507, 1, 0, 0, 0, 509, 81, 1, 0, 0, 0, 510, 511, 3,
		84, 42, 0, 511, 83, 1, 0, 0, 0, 512, 513, 3, 86, 43, 0, 513, 514, 5, 65,
		0, 0, 514, 516, 1, 0, 0, 0, 515, 512, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0,
		516, 517, 1, 0, 0, 0, 517, 518, 3, 98, 49, 0, 518, 85, 1, 0, 0, 0, 519,
		520, 3, 98, 49, 0, 520, 87, 1, 0, 0, 0, 521, 530, 5, 51, 0, 0, 522, 523,
		5, 65, 0, 0, 523, 530, 7, 3, 0, 0, 524, 525, 5, 53, 0, 0, 525, 527, 5,
		65, 0, 0, 526, 528, 7, 3, 0, 0, 527, 526, 1, 0, 0, 0, 527, 528, 1, 0, 0,
		0, 528, 530, 1, 0, 0, 0, 529, 521, 1, 0, 0, 0, 529, 522, 1, 0, 0, 0, 529,
		524, 1, 0, 0, 0, 530, 89, 1, 0, 0, 0, 531, 533, 7, 1, 0, 0, 532, 531, 1,
		0, 0, 0, 532, 533, 1, 0, 0, 0, 533, 540, 1, 0, 0, 0, 534, 541, 3, 88, 44,
		0, 535, 541, 5, 52, 0, 0, 536, 541, 5, 53, 0, 0, 537, 541, 5, 54, 0, 0,
		538, 541, 5, 17, 0, 0, 539, 541, 5, 26, 0, 0, 540, 534, 1, 0, 0, 0, 540,
		535, 1, 0, 0, 0, 540, 536, 1, 0, 0, 0, 540, 537, 1, 0, 0, 0, 540, 538,
		1, 0, 0, 0, 540, 539, 1, 0, 0, 0, 541, 91, 1, 0, 0, 0, 542, 546, 3, 90,
		45, 0, 543, 546, 5, 56, 0, 0, 544, 546, 5, 28, 0, 0, 545, 542, 1, 0, 0,
		0, 545, 543, 1, 0, 0, 0, 545, 544, 1, 0, 0, 0, 546, 93, 1, 0, 0, 0, 547,
		548, 7, 4, 0, 0, 548, 95, 1, 0, 0, 0, 549, 550, 7, 5, 0, 0, 550, 97, 1,
		0, 0, 0, 551, 555, 5, 50, 0, 0, 552, 555, 3, 94, 47, 0, 553, 555, 3, 96,
		48, 0, 554, 551, 1, 0, 0, 0, 554, 552, 1, 0, 0, 0, 554, 553, 1, 0, 0, 0,
		555, 99, 1, 0, 0, 0, 556, 557, 5, 56, 0, 0, 557, 558, 5, 67, 0, 0, 558,
		559, 3, 90, 45, 0, 559, 101, 1, 0, 0, 0, 560, 561, 5, 16, 0, 0, 561, 562,
		5, 27, 0, 0, 562, 563, 5, 12, 0, 0, 563, 103, 1, 0, 0, 0, 61, 107, 113,
		117, 124, 131, 142, 155, 159, 172, 176, 182, 194, 200, 205, 211, 215, 226,
		230, 235, 238, 241, 244, 247, 250, 253, 268, 281, 286, 298, 306, 308, 335,
		340, 357, 359, 370, 374, 379, 395, 407, 419, 427, 431, 438, 446, 452, 475,
		480, 487, 490, 494, 498, 503, 508, 515, 527, 529, 532, 540, 545, 554,
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
	this.GrammarFileName = "SqlParser.g4"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF                 = antlr.TokenEOF
	SqlParserWS                  = 1
	SqlParserAND                 = 2
	SqlParserASC                 = 3
	SqlParserAVG                 = 4
	SqlParserBY                  = 5
	SqlParserCOUNT               = 6
	SqlParserCREATE              = 7
	SqlParserDATABASE            = 8
	SqlParserDEFAULT             = 9
	SqlParserDELETE              = 10
	SqlParserDESC                = 11
	SqlParserEXISTS              = 12
	SqlParserFROM                = 13
	SqlParserGROUP               = 14
	SqlParserHAVING              = 15
	SqlParserIF                  = 16
	SqlParserINF                 = 17
	SqlParserINNER               = 18
	SqlParserINSERT              = 19
	SqlParserINTO                = 20
	SqlParserJOIN                = 21
	SqlParserLEFT                = 22
	SqlParserLIMIT               = 23
	SqlParserMAX                 = 24
	SqlParserMIN                 = 25
	SqlParserNAN_SQL             = 26
	SqlParserNOT                 = 27
	SqlParserNULL_SQL            = 28
	SqlParserOFFSET              = 29
	SqlParserON                  = 30
	SqlParserOR                  = 31
	SqlParserORDER               = 32
	SqlParserSELECT              = 33
	SqlParserSET                 = 34
	SqlParserSUM                 = 35
	SqlParserTABLE               = 36
	SqlParserUPDATE              = 37
	SqlParserUSE                 = 38
	SqlParserVALUE               = 39
	SqlParserVALUES              = 40
	SqlParserWHERE               = 41
	SqlParserSECOND              = 42
	SqlParserMINUTE              = 43
	SqlParserHOUR                = 44
	SqlParserDAY                 = 45
	SqlParserWEEK                = 46
	SqlParserMONTH               = 47
	SqlParserQUARTER             = 48
	SqlParserYEAR                = 49
	SqlParserIDENTIFIER          = 50
	SqlParserFLOATING_LITERAL    = 51
	SqlParserOCTAL_LITERAL       = 52
	SqlParserDECIMAL_LITERAL     = 53
	SqlParserHEXADECIMAL_LITERAL = 54
	SqlParserEXPONENT_NUM_PART   = 55
	SqlParserSTRING_LITERAL      = 56
	SqlParserARROW               = 57
	SqlParserASTERISK            = 58
	SqlParserBACKQUOTE           = 59
	SqlParserBACKSLASH           = 60
	SqlParserCOLON               = 61
	SqlParserCOMMA               = 62
	SqlParserCONCAT              = 63
	SqlParserDASH                = 64
	SqlParserDOT                 = 65
	SqlParserEQ_DOUBLE           = 66
	SqlParserEQ_SINGLE           = 67
	SqlParserGE                  = 68
	SqlParserGT                  = 69
	SqlParserLBRACE              = 70
	SqlParserLBRACKET            = 71
	SqlParserLE                  = 72
	SqlParserLPAREN              = 73
	SqlParserLT                  = 74
	SqlParserNOT_EQ              = 75
	SqlParserPERCENT             = 76
	SqlParserPLUS                = 77
	SqlParserQUERY               = 78
	SqlParserQUOTE_DOUBLE        = 79
	SqlParserQUOTE_SINGLE        = 80
	SqlParserRBRACE              = 81
	SqlParserRBRACKET            = 82
	SqlParserRPAREN              = 83
	SqlParserSEMICOLON           = 84
	SqlParserSLASH               = 85
	SqlParserUNDERSCORE          = 86
	SqlParserMULTI_LINE_COMMENT  = 87
	SqlParserSINGLE_LINE_COMMENT = 88
	SqlParserWHITESPACE          = 89
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
	SqlParserRULE_updateStatement    = 10
	SqlParserRULE_deleteStatement    = 11
	SqlParserRULE_setClause          = 12
	SqlParserRULE_columnsClause      = 13
	SqlParserRULE_insertValuesSpec   = 14
	SqlParserRULE_insertMultiValue   = 15
	SqlParserRULE_dataValue          = 16
	SqlParserRULE_selectStatement    = 17
	SqlParserRULE_fromClause         = 18
	SqlParserRULE_whereClause        = 19
	SqlParserRULE_groupByClause      = 20
	SqlParserRULE_havingClause       = 21
	SqlParserRULE_orderByClause      = 22
	SqlParserRULE_orderByExpr        = 23
	SqlParserRULE_limitClause        = 24
	SqlParserRULE_offsetClause       = 25
	SqlParserRULE_searchCondition    = 26
	SqlParserRULE_predicate          = 27
	SqlParserRULE_expression         = 28
	SqlParserRULE_joinExpr           = 29
	SqlParserRULE_joinOp             = 30
	SqlParserRULE_useStatement       = 31
	SqlParserRULE_columnTypeExpr     = 32
	SqlParserRULE_columnExprList     = 33
	SqlParserRULE_columnsExpr        = 34
	SqlParserRULE_columnExpr         = 35
	SqlParserRULE_aggregateFunction  = 36
	SqlParserRULE_columnIdentifier   = 37
	SqlParserRULE_nestedIdentifier   = 38
	SqlParserRULE_constant           = 39
	SqlParserRULE_realLiteral        = 40
	SqlParserRULE_tableExpr          = 41
	SqlParserRULE_tableIdentifier    = 42
	SqlParserRULE_databaseIdentifier = 43
	SqlParserRULE_floatingLiteral    = 44
	SqlParserRULE_numberLiteral      = 45
	SqlParserRULE_literal            = 46
	SqlParserRULE_interval           = 47
	SqlParserRULE_keyword            = 48
	SqlParserRULE_identifier         = 49
	SqlParserRULE_enumValue          = 50
	SqlParserRULE_ifNotExists        = 51
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

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitStatement(s)
	}
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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserCREATE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(104)
			p.DdlStatement()
		}

	case SqlParserDELETE, SqlParserINSERT, SqlParserSELECT, SqlParserUPDATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(105)
			p.DmlStatement()
		}

	case SqlParserUSE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(106)
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

func (s *DmlStatementContext) UpdateStatement() IUpdateStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUpdateStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUpdateStatementContext)
}

func (s *DmlStatementContext) DeleteStatement() IDeleteStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeleteStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeleteStatementContext)
}

func (s *DmlStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DmlStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DmlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDmlStatement(s)
	}
}

func (s *DmlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDmlStatement(s)
	}
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

	p.SetState(113)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSELECT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.SelectStatement()
		}

	case SqlParserINSERT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.InsertStatement()
		}

	case SqlParserUPDATE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.UpdateStatement()
		}

	case SqlParserDELETE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.DeleteStatement()
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

func (s *DdlStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDdlStatement(s)
	}
}

func (s *DdlStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDdlStatement(s)
	}
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

	p.SetState(117)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(115)
			p.CreateDatabase()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(116)
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

func (s *SystemStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSystemStatement(s)
	}
}

func (s *SystemStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSystemStatement(s)
	}
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
		p.SetState(119)
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

func (s *CreateDatabaseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterCreateDatabase(s)
	}
}

func (s *CreateDatabaseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitCreateDatabase(s)
	}
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
		p.SetState(121)
		p.Match(SqlParserCREATE)
	}
	{
		p.SetState(122)
		p.Match(SqlParserDATABASE)
	}
	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(123)
			p.IfNotExists()
		}

	}
	{
		p.SetState(126)
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

func (s *CreateTableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterCreateTable(s)
	}
}

func (s *CreateTableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitCreateTable(s)
	}
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
		p.SetState(128)
		p.Match(SqlParserCREATE)
	}
	{
		p.SetState(129)
		p.Match(SqlParserTABLE)
	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(130)
			p.IfNotExists()
		}

	}
	{
		p.SetState(133)
		p.TableIdentifier()
	}
	{
		p.SetState(134)
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

func (s *TableSchemaClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterTableSchemaClause(s)
	}
}

func (s *TableSchemaClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitTableSchemaClause(s)
	}
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
		p.SetState(136)
		p.Match(SqlParserLPAREN)
	}
	{
		p.SetState(137)
		p.TableElementExpr()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(138)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(139)
			p.TableElementExpr()
		}

		p.SetState(144)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(145)
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

func (s *TableElementExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterTableElementExpr(s)
	}
}

func (s *TableElementExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitTableElementExpr(s)
	}
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
		p.SetState(147)
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

func (s *TableColumnDfntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterTableColumnDfnt(s)
	}
}

func (s *TableColumnDfntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitTableColumnDfnt(s)
	}
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
		p.SetState(149)
		p.NestedIdentifier()
	}
	{
		p.SetState(150)
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

func (s *InsertStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterInsertStatement(s)
	}
}

func (s *InsertStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitInsertStatement(s)
	}
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
		p.SetState(152)
		p.Match(SqlParserINSERT)
	}
	{
		p.SetState(153)
		p.Match(SqlParserINTO)
	}
	p.SetState(155)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(154)
			p.Match(SqlParserTABLE)
		}

	}

	{
		p.SetState(157)
		p.TableIdentifier()
	}

	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserLPAREN {
		{
			p.SetState(158)
			p.ColumnsClause()
		}

	}
	{
		p.SetState(161)
		p.Match(SqlParserVALUES)
	}
	{
		p.SetState(162)
		p.InsertValuesSpec()
	}

	return localctx
}

// IUpdateStatementContext is an interface to support dynamic dispatch.
type IUpdateStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUpdateStatementContext differentiates from other interfaces.
	IsUpdateStatementContext()
}

type UpdateStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUpdateStatementContext() *UpdateStatementContext {
	var p = new(UpdateStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_updateStatement
	return p
}

func (*UpdateStatementContext) IsUpdateStatementContext() {}

func NewUpdateStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UpdateStatementContext {
	var p = new(UpdateStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_updateStatement

	return p
}

func (s *UpdateStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *UpdateStatementContext) UPDATE() antlr.TerminalNode {
	return s.GetToken(SqlParserUPDATE, 0)
}

func (s *UpdateStatementContext) TableIdentifier() ITableIdentifierContext {
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

func (s *UpdateStatementContext) SET() antlr.TerminalNode {
	return s.GetToken(SqlParserSET, 0)
}

func (s *UpdateStatementContext) AllSetClause() []ISetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISetClauseContext); ok {
			len++
		}
	}

	tst := make([]ISetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISetClauseContext); ok {
			tst[i] = t.(ISetClauseContext)
			i++
		}
	}

	return tst
}

func (s *UpdateStatementContext) SetClause(i int) ISetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISetClauseContext); ok {
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

	return t.(ISetClauseContext)
}

func (s *UpdateStatementContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *UpdateStatementContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *UpdateStatementContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *UpdateStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UpdateStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UpdateStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUpdateStatement(s)
	}
}

func (s *UpdateStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUpdateStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) UpdateStatement() (localctx IUpdateStatementContext) {
	this := p
	_ = this

	localctx = NewUpdateStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SqlParserRULE_updateStatement)
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
		p.SetState(164)
		p.Match(SqlParserUPDATE)
	}
	{
		p.SetState(165)
		p.TableIdentifier()
	}
	{
		p.SetState(166)
		p.Match(SqlParserSET)
	}
	{
		p.SetState(167)
		p.SetClause()
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(168)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(169)
			p.SetClause()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(176)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserWHERE {
		{
			p.SetState(175)
			p.WhereClause()
		}

	}

	return localctx
}

// IDeleteStatementContext is an interface to support dynamic dispatch.
type IDeleteStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeleteStatementContext differentiates from other interfaces.
	IsDeleteStatementContext()
}

type DeleteStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeleteStatementContext() *DeleteStatementContext {
	var p = new(DeleteStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_deleteStatement
	return p
}

func (*DeleteStatementContext) IsDeleteStatementContext() {}

func NewDeleteStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeleteStatementContext {
	var p = new(DeleteStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_deleteStatement

	return p
}

func (s *DeleteStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DeleteStatementContext) DELETE() antlr.TerminalNode {
	return s.GetToken(SqlParserDELETE, 0)
}

func (s *DeleteStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *DeleteStatementContext) TableIdentifier() ITableIdentifierContext {
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

func (s *DeleteStatementContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *DeleteStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeleteStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeleteStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDeleteStatement(s)
	}
}

func (s *DeleteStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDeleteStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DeleteStatement() (localctx IDeleteStatementContext) {
	this := p
	_ = this

	localctx = NewDeleteStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SqlParserRULE_deleteStatement)
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
		p.SetState(178)
		p.Match(SqlParserDELETE)
	}
	{
		p.SetState(179)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(180)
		p.TableIdentifier()
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserWHERE {
		{
			p.SetState(181)
			p.WhereClause()
		}

	}

	return localctx
}

// ISetClauseContext is an interface to support dynamic dispatch.
type ISetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSetClauseContext differentiates from other interfaces.
	IsSetClauseContext()
}

type SetClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySetClauseContext() *SetClauseContext {
	var p = new(SetClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_setClause
	return p
}

func (*SetClauseContext) IsSetClauseContext() {}

func NewSetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SetClauseContext {
	var p = new(SetClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_setClause

	return p
}

func (s *SetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SetClauseContext) ColumnIdentifier() IColumnIdentifierContext {
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

func (s *SetClauseContext) EQ_SINGLE() antlr.TerminalNode {
	return s.GetToken(SqlParserEQ_SINGLE, 0)
}

func (s *SetClauseContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSetClause(s)
	}
}

func (s *SetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSetClause(s)
	}
}

func (s *SetClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSetClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SetClause() (localctx ISetClauseContext) {
	this := p
	_ = this

	localctx = NewSetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_setClause)

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
		p.ColumnIdentifier()
	}
	{
		p.SetState(185)
		p.Match(SqlParserEQ_SINGLE)
	}
	{
		p.SetState(186)
		p.expression(0)
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

func (s *ColumnsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnsClause(s)
	}
}

func (s *ColumnsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnsClause(s)
	}
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
	p.EnterRule(localctx, 26, SqlParserRULE_columnsClause)
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
		p.SetState(188)
		p.Match(SqlParserLPAREN)
	}
	{
		p.SetState(189)
		p.NestedIdentifier()
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(190)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(191)
			p.NestedIdentifier()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(197)
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

func (s *InsertValuesSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterInsertValuesSpec(s)
	}
}

func (s *InsertValuesSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitInsertValuesSpec(s)
	}
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
	p.EnterRule(localctx, 28, SqlParserRULE_insertValuesSpec)
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
	p.SetState(205)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA || _la == SqlParserLPAREN {
		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserCOMMA {
			{
				p.SetState(199)
				p.Match(SqlParserCOMMA)
			}

		}
		{
			p.SetState(202)
			p.InsertMultiValue()
		}

		p.SetState(207)
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

func (s *InsertMultiValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterInsertMultiValue(s)
	}
}

func (s *InsertMultiValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitInsertMultiValue(s)
	}
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
	p.EnterRule(localctx, 30, SqlParserRULE_insertMultiValue)
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
		p.SetState(208)
		p.Match(SqlParserLPAREN)
	}
	p.SetState(213)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-53)&-(0x1f+1)) == 0 && ((1<<uint((_la-53)))&((1<<(SqlParserDECIMAL_LITERAL-53))|(1<<(SqlParserEXPONENT_NUM_PART-53))|(1<<(SqlParserSTRING_LITERAL-53))|(1<<(SqlParserDASH-53))|(1<<(SqlParserDOT-53))|(1<<(SqlParserLPAREN-53))|(1<<(SqlParserPLUS-53)))) != 0) {
		{
			p.SetState(209)
			p.DataValue()
		}
		p.SetState(211)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserCOMMA {
			{
				p.SetState(210)
				p.Match(SqlParserCOMMA)
			}

		}

		p.SetState(215)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(217)
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

func (s *DataValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDataValue(s)
	}
}

func (s *DataValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDataValue(s)
	}
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
	p.EnterRule(localctx, 32, SqlParserRULE_dataValue)
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

	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserDECIMAL_LITERAL, SqlParserEXPONENT_NUM_PART, SqlParserSTRING_LITERAL, SqlParserDASH, SqlParserDOT, SqlParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(219)
			p.Constant()
		}

	case SqlParserLPAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(220)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(221)
			p.Constant()
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == SqlParserCOMMA {
			{
				p.SetState(222)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(223)
				p.Constant()
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(228)
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

func (s *SelectStatementContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStatementContext) GroupByClause() IGroupByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupByClauseContext)
}

func (s *SelectStatementContext) HavingClause() IHavingClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingClauseContext)
}

func (s *SelectStatementContext) OrderByClause() IOrderByClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SelectStatementContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStatementContext) OffsetClause() IOffsetClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffsetClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffsetClauseContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectStatement(s)
	}
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
	p.EnterRule(localctx, 34, SqlParserRULE_selectStatement)
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
		p.SetState(232)
		p.Match(SqlParserSELECT)
	}
	{
		p.SetState(233)
		p.ColumnExprList()
	}
	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserFROM {
		{
			p.SetState(234)
			p.FromClause()
		}

	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserWHERE {
		{
			p.SetState(237)
			p.WhereClause()
		}

	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserGROUP {
		{
			p.SetState(240)
			p.GroupByClause()
		}

	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserHAVING {
		{
			p.SetState(243)
			p.HavingClause()
		}

	}
	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserORDER {
		{
			p.SetState(246)
			p.OrderByClause()
		}

	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserLIMIT {
		{
			p.SetState(249)
			p.LimitClause()
		}

	}
	p.SetState(253)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserOFFSET {
		{
			p.SetState(252)
			p.OffsetClause()
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

func (s *FromClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterFromClause(s)
	}
}

func (s *FromClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitFromClause(s)
	}
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
	p.EnterRule(localctx, 36, SqlParserRULE_fromClause)

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
		p.SetState(255)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(256)
		p.JoinExpr()
	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
}

func (s *WhereClauseContext) SearchCondition() ISearchConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchConditionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) WhereClause() (localctx IWhereClauseContext) {
	this := p
	_ = this

	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SqlParserRULE_whereClause)

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
		p.SetState(258)
		p.Match(SqlParserWHERE)
	}
	{
		p.SetState(259)
		p.searchCondition(0)
	}

	return localctx
}

// IGroupByClauseContext is an interface to support dynamic dispatch.
type IGroupByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupByClauseContext differentiates from other interfaces.
	IsGroupByClauseContext()
}

type GroupByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupByClauseContext() *GroupByClauseContext {
	var p = new(GroupByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupByClause
	return p
}

func (*GroupByClauseContext) IsGroupByClauseContext() {}

func NewGroupByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupByClauseContext {
	var p = new(GroupByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupByClause

	return p
}

func (s *GroupByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupByClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SqlParserGROUP, 0)
}

func (s *GroupByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *GroupByClauseContext) AllColumnIdentifier() []IColumnIdentifierContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IColumnIdentifierContext); ok {
			len++
		}
	}

	tst := make([]IColumnIdentifierContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IColumnIdentifierContext); ok {
			tst[i] = t.(IColumnIdentifierContext)
			i++
		}
	}

	return tst
}

func (s *GroupByClauseContext) ColumnIdentifier(i int) IColumnIdentifierContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IColumnIdentifierContext); ok {
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

	return t.(IColumnIdentifierContext)
}

func (s *GroupByClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *GroupByClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *GroupByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterGroupByClause(s)
	}
}

func (s *GroupByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitGroupByClause(s)
	}
}

func (s *GroupByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitGroupByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) GroupByClause() (localctx IGroupByClauseContext) {
	this := p
	_ = this

	localctx = NewGroupByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_groupByClause)
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
		p.SetState(261)
		p.Match(SqlParserGROUP)
	}
	{
		p.SetState(262)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(263)
		p.ColumnIdentifier()
	}
	p.SetState(268)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(264)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(265)
			p.ColumnIdentifier()
		}

		p.SetState(270)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHavingClauseContext is an interface to support dynamic dispatch.
type IHavingClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHavingClauseContext differentiates from other interfaces.
	IsHavingClauseContext()
}

type HavingClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingClauseContext() *HavingClauseContext {
	var p = new(HavingClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_havingClause
	return p
}

func (*HavingClauseContext) IsHavingClauseContext() {}

func NewHavingClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingClauseContext {
	var p = new(HavingClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_havingClause

	return p
}

func (s *HavingClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SqlParserHAVING, 0)
}

func (s *HavingClauseContext) SearchCondition() ISearchConditionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchConditionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISearchConditionContext)
}

func (s *HavingClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterHavingClause(s)
	}
}

func (s *HavingClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitHavingClause(s)
	}
}

func (s *HavingClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitHavingClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) HavingClause() (localctx IHavingClauseContext) {
	this := p
	_ = this

	localctx = NewHavingClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SqlParserRULE_havingClause)

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
		p.SetState(271)
		p.Match(SqlParserHAVING)
	}
	{
		p.SetState(272)
		p.searchCondition(0)
	}

	return localctx
}

// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *OrderByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *OrderByClauseContext) AllOrderByExpr() []IOrderByExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderByExprContext); ok {
			len++
		}
	}

	tst := make([]IOrderByExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderByExprContext); ok {
			tst[i] = t.(IOrderByExprContext)
			i++
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByExpr(i int) IOrderByExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderByExprContext); ok {
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

	return t.(IOrderByExprContext)
}

func (s *OrderByClauseContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SqlParserCOMMA)
}

func (s *OrderByClauseContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserCOMMA, i)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (s *OrderByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitOrderByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OrderByClause() (localctx IOrderByClauseContext) {
	this := p
	_ = this

	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SqlParserRULE_orderByClause)
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
		p.SetState(274)
		p.Match(SqlParserORDER)
	}
	{
		p.SetState(275)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(276)
		p.OrderByExpr()
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(277)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(278)
			p.OrderByExpr()
		}

		p.SetState(283)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByExprContext is an interface to support dynamic dispatch.
type IOrderByExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByExprContext differentiates from other interfaces.
	IsOrderByExprContext()
}

type OrderByExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByExprContext() *OrderByExprContext {
	var p = new(OrderByExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByExpr
	return p
}

func (*OrderByExprContext) IsOrderByExprContext() {}

func NewOrderByExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByExprContext {
	var p = new(OrderByExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByExpr

	return p
}

func (s *OrderByExprContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByExprContext) ColumnIdentifier() IColumnIdentifierContext {
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

func (s *OrderByExprContext) ASC() antlr.TerminalNode {
	return s.GetToken(SqlParserASC, 0)
}

func (s *OrderByExprContext) DESC() antlr.TerminalNode {
	return s.GetToken(SqlParserDESC, 0)
}

func (s *OrderByExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterOrderByExpr(s)
	}
}

func (s *OrderByExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitOrderByExpr(s)
	}
}

func (s *OrderByExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitOrderByExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OrderByExpr() (localctx IOrderByExprContext) {
	this := p
	_ = this

	localctx = NewOrderByExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SqlParserRULE_orderByExpr)
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
		p.SetState(284)
		p.ColumnIdentifier()
	}
	p.SetState(286)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserASC || _la == SqlParserDESC {
		{
			p.SetState(285)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserASC || _la == SqlParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
}

func (s *LimitClauseContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LimitClause() (localctx ILimitClauseContext) {
	this := p
	_ = this

	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SqlParserRULE_limitClause)

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
		p.SetState(288)
		p.Match(SqlParserLIMIT)
	}
	{
		p.SetState(289)
		p.Match(SqlParserDECIMAL_LITERAL)
	}

	return localctx
}

// IOffsetClauseContext is an interface to support dynamic dispatch.
type IOffsetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOffsetClauseContext differentiates from other interfaces.
	IsOffsetClauseContext()
}

type OffsetClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetClauseContext() *OffsetClauseContext {
	var p = new(OffsetClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_offsetClause
	return p
}

func (*OffsetClauseContext) IsOffsetClauseContext() {}

func NewOffsetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetClauseContext {
	var p = new(OffsetClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_offsetClause

	return p
}

func (s *OffsetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SqlParserOFFSET, 0)
}

func (s *OffsetClauseContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *OffsetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterOffsetClause(s)
	}
}

func (s *OffsetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitOffsetClause(s)
	}
}

func (s *OffsetClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitOffsetClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OffsetClause() (localctx IOffsetClauseContext) {
	this := p
	_ = this

	localctx = NewOffsetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SqlParserRULE_offsetClause)

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
		p.SetState(291)
		p.Match(SqlParserOFFSET)
	}
	{
		p.SetState(292)
		p.Match(SqlParserDECIMAL_LITERAL)
	}

	return localctx
}

// ISearchConditionContext is an interface to support dynamic dispatch.
type ISearchConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSearchConditionContext differentiates from other interfaces.
	IsSearchConditionContext()
}

type SearchConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySearchConditionContext() *SearchConditionContext {
	var p = new(SearchConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_searchCondition
	return p
}

func (*SearchConditionContext) IsSearchConditionContext() {}

func NewSearchConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SearchConditionContext {
	var p = new(SearchConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_searchCondition

	return p
}

func (s *SearchConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *SearchConditionContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *SearchConditionContext) AllSearchCondition() []ISearchConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchConditionContext); ok {
			len++
		}
	}

	tst := make([]ISearchConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchConditionContext); ok {
			tst[i] = t.(ISearchConditionContext)
			i++
		}
	}

	return tst
}

func (s *SearchConditionContext) SearchCondition(i int) ISearchConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchConditionContext); ok {
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

	return t.(ISearchConditionContext)
}

func (s *SearchConditionContext) Predicate() IPredicateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPredicateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *SearchConditionContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *SearchConditionContext) OR() antlr.TerminalNode {
	return s.GetToken(SqlParserOR, 0)
}

func (s *SearchConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SearchConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SearchConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSearchCondition(s)
	}
}

func (s *SearchConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSearchCondition(s)
	}
}

func (s *SearchConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSearchCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SearchCondition() (localctx ISearchConditionContext) {
	return p.searchCondition(0)
}

func (p *SqlParser) searchCondition(_p int) (localctx ISearchConditionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSearchConditionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISearchConditionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 52
	p.EnterRecursionRule(localctx, 52, SqlParserRULE_searchCondition, _p)

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
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(295)
			p.Match(SqlParserNOT)
		}
		{
			p.SetState(296)
			p.searchCondition(2)
		}

	case 2:
		{
			p.SetState(297)
			p.Predicate()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(308)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(306)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) {
			case 1:
				localctx = NewSearchConditionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_searchCondition)
				p.SetState(300)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(301)
					p.Match(SqlParserAND)
				}
				{
					p.SetState(302)
					p.searchCondition(5)
				}

			case 2:
				localctx = NewSearchConditionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_searchCondition)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(304)
					p.Match(SqlParserOR)
				}
				{
					p.SetState(305)
					p.searchCondition(4)
				}

			}

		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *PredicateContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *PredicateContext) EQ_SINGLE() antlr.TerminalNode {
	return s.GetToken(SqlParserEQ_SINGLE, 0)
}

func (s *PredicateContext) NOT_EQ() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT_EQ, 0)
}

func (s *PredicateContext) LT() antlr.TerminalNode {
	return s.GetToken(SqlParserLT, 0)
}

func (s *PredicateContext) LE() antlr.TerminalNode {
	return s.GetToken(SqlParserLE, 0)
}

func (s *PredicateContext) GT() antlr.TerminalNode {
	return s.GetToken(SqlParserGT, 0)
}

func (s *PredicateContext) GE() antlr.TerminalNode {
	return s.GetToken(SqlParserGE, 0)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterPredicate(s)
	}
}

func (s *PredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitPredicate(s)
	}
}

func (s *PredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Predicate() (localctx IPredicateContext) {
	this := p
	_ = this

	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SqlParserRULE_predicate)

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

	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(311)
			p.expression(0)
		}
		{
			p.SetState(312)
			p.Match(SqlParserEQ_SINGLE)
		}
		{
			p.SetState(313)
			p.expression(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(315)
			p.expression(0)
		}
		{
			p.SetState(316)
			p.Match(SqlParserNOT_EQ)
		}
		{
			p.SetState(317)
			p.expression(0)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(319)
			p.expression(0)
		}
		{
			p.SetState(320)
			p.Match(SqlParserLT)
		}
		{
			p.SetState(321)
			p.expression(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(323)
			p.expression(0)
		}
		{
			p.SetState(324)
			p.Match(SqlParserLE)
		}
		{
			p.SetState(325)
			p.expression(0)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(327)
			p.expression(0)
		}
		{
			p.SetState(328)
			p.Match(SqlParserGT)
		}
		{
			p.SetState(329)
			p.expression(0)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(331)
			p.expression(0)
		}
		{
			p.SetState(332)
			p.Match(SqlParserGE)
		}
		{
			p.SetState(333)
			p.expression(0)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Literal() ILiteralContext {
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

func (s *ExpressionContext) ColumnIdentifier() IColumnIdentifierContext {
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

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
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

	return t.(IExpressionContext)
}

func (s *ExpressionContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SqlParserASTERISK, 0)
}

func (s *ExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(SqlParserSLASH, 0)
}

func (s *ExpressionContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(SqlParserPERCENT, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SqlParserPLUS, 0)
}

func (s *ExpressionContext) DASH() antlr.TerminalNode {
	return s.GetToken(SqlParserDASH, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SqlParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, SqlParserRULE_expression, _p)

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
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(338)
			p.Literal()
		}

	case 2:
		{
			p.SetState(339)
			p.ColumnIdentifier()
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(357)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
				p.SetState(342)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(343)
					p.Match(SqlParserASTERISK)
				}
				{
					p.SetState(344)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
				p.SetState(345)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(346)
					p.Match(SqlParserSLASH)
				}
				{
					p.SetState(347)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
				p.SetState(348)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(349)
					p.Match(SqlParserPERCENT)
				}
				{
					p.SetState(350)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
				p.SetState(351)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(352)
					p.Match(SqlParserPLUS)
				}
				{
					p.SetState(353)
					p.expression(3)
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
				p.SetState(354)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(355)
					p.Match(SqlParserDASH)
				}
				{
					p.SetState(356)
					p.expression(2)
				}

			}

		}
		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
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

func (s *JoinExprContext) AllTableExpr() []ITableExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITableExprContext); ok {
			len++
		}
	}

	tst := make([]ITableExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITableExprContext); ok {
			tst[i] = t.(ITableExprContext)
			i++
		}
	}

	return tst
}

func (s *JoinExprContext) TableExpr(i int) ITableExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITableExprContext); ok {
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

	return t.(ITableExprContext)
}

func (s *JoinExprContext) AllJoinOp() []IJoinOpContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IJoinOpContext); ok {
			len++
		}
	}

	tst := make([]IJoinOpContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IJoinOpContext); ok {
			tst[i] = t.(IJoinOpContext)
			i++
		}
	}

	return tst
}

func (s *JoinExprContext) JoinOp(i int) IJoinOpContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IJoinOpContext); ok {
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

	return t.(IJoinOpContext)
}

func (s *JoinExprContext) AllON() []antlr.TerminalNode {
	return s.GetTokens(SqlParserON)
}

func (s *JoinExprContext) ON(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserON, i)
}

func (s *JoinExprContext) AllSearchCondition() []ISearchConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISearchConditionContext); ok {
			len++
		}
	}

	tst := make([]ISearchConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISearchConditionContext); ok {
			tst[i] = t.(ISearchConditionContext)
			i++
		}
	}

	return tst
}

func (s *JoinExprContext) SearchCondition(i int) ISearchConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISearchConditionContext); ok {
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

	return t.(ISearchConditionContext)
}

func (s *JoinExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterJoinExpr(s)
	}
}

func (s *JoinExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitJoinExpr(s)
	}
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
	p.EnterRule(localctx, 58, SqlParserRULE_joinExpr)
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
		p.SetState(362)
		p.TableExpr()
	}
	p.SetState(370)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SqlParserINNER)|(1<<SqlParserJOIN)|(1<<SqlParserLEFT))) != 0 {
		{
			p.SetState(363)
			p.JoinOp()
		}
		{
			p.SetState(364)
			p.TableExpr()
		}
		{
			p.SetState(365)
			p.Match(SqlParserON)
		}
		{
			p.SetState(366)
			p.searchCondition(0)
		}

		p.SetState(372)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IJoinOpContext is an interface to support dynamic dispatch.
type IJoinOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsJoinOpContext differentiates from other interfaces.
	IsJoinOpContext()
}

type JoinOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyJoinOpContext() *JoinOpContext {
	var p = new(JoinOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_joinOp
	return p
}

func (*JoinOpContext) IsJoinOpContext() {}

func NewJoinOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *JoinOpContext {
	var p = new(JoinOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_joinOp

	return p
}

func (s *JoinOpContext) GetParser() antlr.Parser { return s.parser }

func (s *JoinOpContext) JOIN() antlr.TerminalNode {
	return s.GetToken(SqlParserJOIN, 0)
}

func (s *JoinOpContext) INNER() antlr.TerminalNode {
	return s.GetToken(SqlParserINNER, 0)
}

func (s *JoinOpContext) LEFT() antlr.TerminalNode {
	return s.GetToken(SqlParserLEFT, 0)
}

func (s *JoinOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *JoinOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *JoinOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterJoinOp(s)
	}
}

func (s *JoinOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitJoinOp(s)
	}
}

func (s *JoinOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitJoinOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) JoinOp() (localctx IJoinOpContext) {
	this := p
	_ = this

	localctx = NewJoinOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SqlParserRULE_joinOp)
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

	p.SetState(379)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserINNER, SqlParserJOIN:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(374)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserINNER {
			{
				p.SetState(373)
				p.Match(SqlParserINNER)
			}

		}
		{
			p.SetState(376)
			p.Match(SqlParserJOIN)
		}

	case SqlParserLEFT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(377)
			p.Match(SqlParserLEFT)
		}
		{
			p.SetState(378)
			p.Match(SqlParserJOIN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *UseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUseStatement(s)
	}
}

func (s *UseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUseStatement(s)
	}
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
	p.EnterRule(localctx, 62, SqlParserRULE_useStatement)

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
		p.SetState(381)
		p.Match(SqlParserUSE)
	}
	{
		p.SetState(382)
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

func (s *ColumnTypeExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnTypeExpr(s)
	}
}

func (s *ColumnTypeExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnTypeExpr(s)
	}
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
	p.EnterRule(localctx, 64, SqlParserRULE_columnTypeExpr)
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

	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(384)
			p.Identifier()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(385)
			p.Identifier()
		}
		{
			p.SetState(386)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(387)
			p.Identifier()
		}
		{
			p.SetState(388)
			p.ColumnTypeExpr()
		}
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(389)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(390)
				p.Identifier()
			}
			{
				p.SetState(391)
				p.ColumnTypeExpr()
			}

			p.SetState(397)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(398)
			p.Match(SqlParserRPAREN)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(400)
			p.Identifier()
		}
		{
			p.SetState(401)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(402)
			p.EnumValue()
		}
		p.SetState(407)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(403)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(404)
				p.EnumValue()
			}

			p.SetState(409)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(410)
			p.Match(SqlParserRPAREN)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(412)
			p.Identifier()
		}
		{
			p.SetState(413)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(414)
			p.ColumnTypeExpr()
		}
		p.SetState(419)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(415)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(416)
				p.ColumnTypeExpr()
			}

			p.SetState(421)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(422)
			p.Match(SqlParserRPAREN)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(424)
			p.Identifier()
		}
		{
			p.SetState(425)
			p.Match(SqlParserLPAREN)
		}
		p.SetState(427)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SqlParserAVG)|(1<<SqlParserCOUNT)|(1<<SqlParserCREATE)|(1<<SqlParserDATABASE)|(1<<SqlParserDEFAULT)|(1<<SqlParserEXISTS)|(1<<SqlParserFROM)|(1<<SqlParserIF)|(1<<SqlParserINF)|(1<<SqlParserINSERT)|(1<<SqlParserINTO)|(1<<SqlParserMAX)|(1<<SqlParserMIN)|(1<<SqlParserNAN_SQL)|(1<<SqlParserNOT)|(1<<SqlParserNULL_SQL))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SqlParserSELECT-33))|(1<<(SqlParserSET-33))|(1<<(SqlParserSUM-33))|(1<<(SqlParserTABLE-33))|(1<<(SqlParserVALUE-33))|(1<<(SqlParserVALUES-33))|(1<<(SqlParserSECOND-33))|(1<<(SqlParserMINUTE-33))|(1<<(SqlParserHOUR-33))|(1<<(SqlParserDAY-33))|(1<<(SqlParserWEEK-33))|(1<<(SqlParserMONTH-33))|(1<<(SqlParserQUARTER-33))|(1<<(SqlParserYEAR-33))|(1<<(SqlParserIDENTIFIER-33))|(1<<(SqlParserFLOATING_LITERAL-33))|(1<<(SqlParserOCTAL_LITERAL-33))|(1<<(SqlParserDECIMAL_LITERAL-33))|(1<<(SqlParserHEXADECIMAL_LITERAL-33))|(1<<(SqlParserSTRING_LITERAL-33))|(1<<(SqlParserDASH-33)))) != 0) || _la == SqlParserDOT || _la == SqlParserPLUS {
			{
				p.SetState(426)
				p.ColumnExprList()
			}

		}
		{
			p.SetState(429)
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

func (s *ColumnExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnExprList(s)
	}
}

func (s *ColumnExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnExprList(s)
	}
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
	p.EnterRule(localctx, 66, SqlParserRULE_columnExprList)
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
		p.SetState(433)
		p.ColumnsExpr()
	}
	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(434)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(435)
			p.ColumnsExpr()
		}

		p.SetState(440)
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

func (s *ColumnsExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnsExpr(s)
	}
}

func (s *ColumnsExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnsExpr(s)
	}
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
	p.EnterRule(localctx, 68, SqlParserRULE_columnsExpr)

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
		p.SetState(441)
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

func (s *ColumnExprContext) AggregateFunction() IAggregateFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *ColumnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnExpr(s)
	}
}

func (s *ColumnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnExpr(s)
	}
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
	p.EnterRule(localctx, 70, SqlParserRULE_columnExpr)

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

	p.SetState(446)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(443)
			p.Literal()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)
			p.ColumnIdentifier()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(445)
			p.AggregateFunction()
		}

	}

	return localctx
}

// IAggregateFunctionContext is an interface to support dynamic dispatch.
type IAggregateFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggregateFunctionContext differentiates from other interfaces.
	IsAggregateFunctionContext()
}

type AggregateFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFunctionContext() *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_aggregateFunction
	return p
}

func (*AggregateFunctionContext) IsAggregateFunctionContext() {}

func NewAggregateFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_aggregateFunction

	return p
}

func (s *AggregateFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFunctionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOUNT, 0)
}

func (s *AggregateFunctionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserLPAREN, 0)
}

func (s *AggregateFunctionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(SqlParserRPAREN, 0)
}

func (s *AggregateFunctionContext) ASTERISK() antlr.TerminalNode {
	return s.GetToken(SqlParserASTERISK, 0)
}

func (s *AggregateFunctionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AggregateFunctionContext) SUM() antlr.TerminalNode {
	return s.GetToken(SqlParserSUM, 0)
}

func (s *AggregateFunctionContext) AVG() antlr.TerminalNode {
	return s.GetToken(SqlParserAVG, 0)
}

func (s *AggregateFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(SqlParserMIN, 0)
}

func (s *AggregateFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(SqlParserMAX, 0)
}

func (s *AggregateFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitAggregateFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) AggregateFunction() (localctx IAggregateFunctionContext) {
	this := p
	_ = this

	localctx = NewAggregateFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SqlParserRULE_aggregateFunction)

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

	p.SetState(475)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			p.Match(SqlParserCOUNT)
		}
		{
			p.SetState(449)
			p.Match(SqlParserLPAREN)
		}
		p.SetState(452)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case SqlParserASTERISK:
			{
				p.SetState(450)
				p.Match(SqlParserASTERISK)
			}

		case SqlParserCREATE, SqlParserDATABASE, SqlParserDEFAULT, SqlParserEXISTS, SqlParserFROM, SqlParserIF, SqlParserINF, SqlParserINSERT, SqlParserINTO, SqlParserNAN_SQL, SqlParserNOT, SqlParserNULL_SQL, SqlParserSELECT, SqlParserSET, SqlParserTABLE, SqlParserVALUE, SqlParserVALUES, SqlParserSECOND, SqlParserMINUTE, SqlParserHOUR, SqlParserDAY, SqlParserWEEK, SqlParserMONTH, SqlParserQUARTER, SqlParserYEAR, SqlParserIDENTIFIER, SqlParserFLOATING_LITERAL, SqlParserOCTAL_LITERAL, SqlParserDECIMAL_LITERAL, SqlParserHEXADECIMAL_LITERAL, SqlParserSTRING_LITERAL, SqlParserDASH, SqlParserDOT, SqlParserPLUS:
			{
				p.SetState(451)
				p.expression(0)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(454)
			p.Match(SqlParserRPAREN)
		}

	case SqlParserSUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(455)
			p.Match(SqlParserSUM)
		}
		{
			p.SetState(456)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(457)
			p.expression(0)
		}
		{
			p.SetState(458)
			p.Match(SqlParserRPAREN)
		}

	case SqlParserAVG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(460)
			p.Match(SqlParserAVG)
		}
		{
			p.SetState(461)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(462)
			p.expression(0)
		}
		{
			p.SetState(463)
			p.Match(SqlParserRPAREN)
		}

	case SqlParserMIN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(465)
			p.Match(SqlParserMIN)
		}
		{
			p.SetState(466)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(467)
			p.expression(0)
		}
		{
			p.SetState(468)
			p.Match(SqlParserRPAREN)
		}

	case SqlParserMAX:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(470)
			p.Match(SqlParserMAX)
		}
		{
			p.SetState(471)
			p.Match(SqlParserLPAREN)
		}
		{
			p.SetState(472)
			p.expression(0)
		}
		{
			p.SetState(473)
			p.Match(SqlParserRPAREN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

func (s *ColumnIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnIdentifier(s)
	}
}

func (s *ColumnIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnIdentifier(s)
	}
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
	p.EnterRule(localctx, 74, SqlParserRULE_columnIdentifier)

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
	p.SetState(480)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(477)
			p.TableIdentifier()
		}
		{
			p.SetState(478)
			p.Match(SqlParserDOT)
		}

	}
	{
		p.SetState(482)
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

func (s *NestedIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterNestedIdentifier(s)
	}
}

func (s *NestedIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitNestedIdentifier(s)
	}
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
	p.EnterRule(localctx, 76, SqlParserRULE_nestedIdentifier)

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
		p.SetState(484)
		p.Identifier()
	}
	p.SetState(487)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(485)
			p.Match(SqlParserDOT)
		}
		{
			p.SetState(486)
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

func (s *ConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterConstant(s)
	}
}

func (s *ConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitConstant(s)
	}
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
	p.EnterRule(localctx, 78, SqlParserRULE_constant)
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

	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 51, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(490)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserDASH || _la == SqlParserPLUS {
			{
				p.SetState(489)
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
			p.SetState(492)
			p.RealLiteral()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserDASH || _la == SqlParserPLUS {
			{
				p.SetState(493)
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
			p.SetState(496)
			p.Match(SqlParserDECIMAL_LITERAL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(497)
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

func (s *RealLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterRealLiteral(s)
	}
}

func (s *RealLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitRealLiteral(s)
	}
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
	p.EnterRule(localctx, 80, SqlParserRULE_realLiteral)
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

	p.SetState(508)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserDECIMAL_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(500)
			p.Match(SqlParserDECIMAL_LITERAL)
		}
		{
			p.SetState(501)
			p.Match(SqlParserDOT)
		}
		p.SetState(503)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 52, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(502)
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
			p.SetState(505)
			p.Match(SqlParserDOT)
		}
		{
			p.SetState(506)
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
			p.SetState(507)
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

func (s *TableExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterTableExpr(s)
	}
}

func (s *TableExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitTableExpr(s)
	}
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
	p.EnterRule(localctx, 82, SqlParserRULE_tableExpr)

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
		p.SetState(510)
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

func (s *TableIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterTableIdentifier(s)
	}
}

func (s *TableIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitTableIdentifier(s)
	}
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
	p.EnterRule(localctx, 84, SqlParserRULE_tableIdentifier)

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
	p.SetState(515)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 54, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(512)
			p.DatabaseIdentifier()
		}
		{
			p.SetState(513)
			p.Match(SqlParserDOT)
		}

	}
	{
		p.SetState(517)
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

func (s *DatabaseIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDatabaseIdentifier(s)
	}
}

func (s *DatabaseIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDatabaseIdentifier(s)
	}
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
	p.EnterRule(localctx, 86, SqlParserRULE_databaseIdentifier)

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
		p.SetState(519)
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

func (s *FloatingLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterFloatingLiteral(s)
	}
}

func (s *FloatingLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitFloatingLiteral(s)
	}
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
	p.EnterRule(localctx, 88, SqlParserRULE_floatingLiteral)
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

	p.SetState(529)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserFLOATING_LITERAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Match(SqlParserFLOATING_LITERAL)
		}

	case SqlParserDOT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(522)
			p.Match(SqlParserDOT)
		}
		{
			p.SetState(523)
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
			p.SetState(524)
			p.Match(SqlParserDECIMAL_LITERAL)
		}
		{
			p.SetState(525)
			p.Match(SqlParserDOT)
		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(526)
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

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
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
	p.EnterRule(localctx, 90, SqlParserRULE_numberLiteral)
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
	p.SetState(532)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserDASH || _la == SqlParserPLUS {
		{
			p.SetState(531)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserDASH || _la == SqlParserPLUS) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(540)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(534)
			p.FloatingLiteral()
		}

	case 2:
		{
			p.SetState(535)
			p.Match(SqlParserOCTAL_LITERAL)
		}

	case 3:
		{
			p.SetState(536)
			p.Match(SqlParserDECIMAL_LITERAL)
		}

	case 4:
		{
			p.SetState(537)
			p.Match(SqlParserHEXADECIMAL_LITERAL)
		}

	case 5:
		{
			p.SetState(538)
			p.Match(SqlParserINF)
		}

	case 6:
		{
			p.SetState(539)
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

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLiteral(s)
	}
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
	p.EnterRule(localctx, 92, SqlParserRULE_literal)

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

	p.SetState(545)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserINF, SqlParserNAN_SQL, SqlParserFLOATING_LITERAL, SqlParserOCTAL_LITERAL, SqlParserDECIMAL_LITERAL, SqlParserHEXADECIMAL_LITERAL, SqlParserDASH, SqlParserDOT, SqlParserPLUS:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(542)
			p.NumberLiteral()
		}

	case SqlParserSTRING_LITERAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(543)
			p.Match(SqlParserSTRING_LITERAL)
		}

	case SqlParserNULL_SQL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(544)
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

func (s *IntervalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterInterval(s)
	}
}

func (s *IntervalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitInterval(s)
	}
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
	p.EnterRule(localctx, 94, SqlParserRULE_interval)
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
		p.SetState(547)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(SqlParserSECOND-42))|(1<<(SqlParserMINUTE-42))|(1<<(SqlParserHOUR-42))|(1<<(SqlParserDAY-42))|(1<<(SqlParserWEEK-42))|(1<<(SqlParserMONTH-42))|(1<<(SqlParserQUARTER-42))|(1<<(SqlParserYEAR-42)))) != 0) {
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

func (s *KeywordContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterKeyword(s)
	}
}

func (s *KeywordContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitKeyword(s)
	}
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
	p.EnterRule(localctx, 96, SqlParserRULE_keyword)
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
		p.SetState(549)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<SqlParserCREATE)|(1<<SqlParserDATABASE)|(1<<SqlParserDEFAULT)|(1<<SqlParserEXISTS)|(1<<SqlParserFROM)|(1<<SqlParserIF)|(1<<SqlParserINSERT)|(1<<SqlParserINTO)|(1<<SqlParserNOT)|(1<<SqlParserNULL_SQL))) != 0) || (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(SqlParserSELECT-33))|(1<<(SqlParserSET-33))|(1<<(SqlParserTABLE-33))|(1<<(SqlParserVALUE-33))|(1<<(SqlParserVALUES-33)))) != 0)) {
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

func (s *IdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterIdentifier(s)
	}
}

func (s *IdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitIdentifier(s)
	}
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
	p.EnterRule(localctx, 98, SqlParserRULE_identifier)

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

	p.SetState(554)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(551)
			p.Match(SqlParserIDENTIFIER)
		}

	case SqlParserSECOND, SqlParserMINUTE, SqlParserHOUR, SqlParserDAY, SqlParserWEEK, SqlParserMONTH, SqlParserQUARTER, SqlParserYEAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(552)
			p.Interval()
		}

	case SqlParserCREATE, SqlParserDATABASE, SqlParserDEFAULT, SqlParserEXISTS, SqlParserFROM, SqlParserIF, SqlParserINSERT, SqlParserINTO, SqlParserNOT, SqlParserNULL_SQL, SqlParserSELECT, SqlParserSET, SqlParserTABLE, SqlParserVALUE, SqlParserVALUES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(553)
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

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitEnumValue(s)
	}
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
	p.EnterRule(localctx, 100, SqlParserRULE_enumValue)

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
		p.SetState(556)
		p.Match(SqlParserSTRING_LITERAL)
	}
	{
		p.SetState(557)
		p.Match(SqlParserEQ_SINGLE)
	}
	{
		p.SetState(558)
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

func (s *IfNotExistsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterIfNotExists(s)
	}
}

func (s *IfNotExistsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitIfNotExists(s)
	}
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
	p.EnterRule(localctx, 102, SqlParserRULE_ifNotExists)

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
		p.SetState(560)
		p.Match(SqlParserIF)
	}
	{
		p.SetState(561)
		p.Match(SqlParserNOT)
	}
	{
		p.SetState(562)
		p.Match(SqlParserEXISTS)
	}

	return localctx
}

func (p *SqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 26:
		var t *SearchConditionContext = nil
		if localctx != nil {
			t = localctx.(*SearchConditionContext)
		}
		return p.SearchCondition_Sempred(t, predIndex)

	case 28:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SqlParser) SearchCondition_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
