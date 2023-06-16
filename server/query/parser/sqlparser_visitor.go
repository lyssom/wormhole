// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by SqlParser.
type SqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SqlParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by SqlParser#dmlStatement.
	VisitDmlStatement(ctx *DmlStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#ddlStatement.
	VisitDdlStatement(ctx *DdlStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#systemStatement.
	VisitSystemStatement(ctx *SystemStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) interface{}

	// Visit a parse tree produced by SqlParser#createTable.
	VisitCreateTable(ctx *CreateTableContext) interface{}

	// Visit a parse tree produced by SqlParser#tableSchemaClause.
	VisitTableSchemaClause(ctx *TableSchemaClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#tableElementExpr.
	VisitTableElementExpr(ctx *TableElementExprContext) interface{}

	// Visit a parse tree produced by SqlParser#tableColumnDfnt.
	VisitTableColumnDfnt(ctx *TableColumnDfntContext) interface{}

	// Visit a parse tree produced by SqlParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#columnsClause.
	VisitColumnsClause(ctx *ColumnsClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#insertValuesSpec.
	VisitInsertValuesSpec(ctx *InsertValuesSpecContext) interface{}

	// Visit a parse tree produced by SqlParser#insertMultiValue.
	VisitInsertMultiValue(ctx *InsertMultiValueContext) interface{}

	// Visit a parse tree produced by SqlParser#dataValue.
	VisitDataValue(ctx *DataValueContext) interface{}

	// Visit a parse tree produced by SqlParser#selectStatement.
	VisitSelectStatement(ctx *SelectStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by SqlParser#joinExpr.
	VisitJoinExpr(ctx *JoinExprContext) interface{}

	// Visit a parse tree produced by SqlParser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) interface{}

	// Visit a parse tree produced by SqlParser#columnTypeExpr.
	VisitColumnTypeExpr(ctx *ColumnTypeExprContext) interface{}

	// Visit a parse tree produced by SqlParser#columnExprList.
	VisitColumnExprList(ctx *ColumnExprListContext) interface{}

	// Visit a parse tree produced by SqlParser#columnsExpr.
	VisitColumnsExpr(ctx *ColumnsExprContext) interface{}

	// Visit a parse tree produced by SqlParser#columnExpr.
	VisitColumnExpr(ctx *ColumnExprContext) interface{}

	// Visit a parse tree produced by SqlParser#columnIdentifier.
	VisitColumnIdentifier(ctx *ColumnIdentifierContext) interface{}

	// Visit a parse tree produced by SqlParser#nestedIdentifier.
	VisitNestedIdentifier(ctx *NestedIdentifierContext) interface{}

	// Visit a parse tree produced by SqlParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by SqlParser#realLiteral.
	VisitRealLiteral(ctx *RealLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#tableExpr.
	VisitTableExpr(ctx *TableExprContext) interface{}

	// Visit a parse tree produced by SqlParser#tableIdentifier.
	VisitTableIdentifier(ctx *TableIdentifierContext) interface{}

	// Visit a parse tree produced by SqlParser#databaseIdentifier.
	VisitDatabaseIdentifier(ctx *DatabaseIdentifierContext) interface{}

	// Visit a parse tree produced by SqlParser#floatingLiteral.
	VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by SqlParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by SqlParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by SqlParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by SqlParser#enumValue.
	VisitEnumValue(ctx *EnumValueContext) interface{}

	// Visit a parse tree produced by SqlParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) interface{}
}
