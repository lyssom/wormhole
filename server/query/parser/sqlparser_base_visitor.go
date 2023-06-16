// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseSqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSqlParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitDmlStatement(ctx *DmlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitDdlStatement(ctx *DdlStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitSystemStatement(ctx *SystemStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitCreateTable(ctx *CreateTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitTableSchemaClause(ctx *TableSchemaClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitTableElementExpr(ctx *TableElementExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitTableColumnDfnt(ctx *TableColumnDfntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnsClause(ctx *ColumnsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitInsertValuesSpec(ctx *InsertValuesSpecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitInsertMultiValue(ctx *InsertMultiValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitDataValue(ctx *DataValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitSelectStatement(ctx *SelectStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitJoinExpr(ctx *JoinExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitUseStatement(ctx *UseStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnTypeExpr(ctx *ColumnTypeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnExprList(ctx *ColumnExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnsExpr(ctx *ColumnsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnExpr(ctx *ColumnExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitColumnIdentifier(ctx *ColumnIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitNestedIdentifier(ctx *NestedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitRealLiteral(ctx *RealLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitTableExpr(ctx *TableExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitTableIdentifier(ctx *TableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitDatabaseIdentifier(ctx *DatabaseIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitEnumValue(ctx *EnumValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSqlParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) interface{} {
	return v.VisitChildren(ctx)
}
