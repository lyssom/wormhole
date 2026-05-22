// Code generated from SqlParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// SqlParserListener is a complete listener for a parse tree produced by SqlParser.
type SqlParserListener interface {
	antlr.ParseTreeListener

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterDmlStatement is called when entering the dmlStatement production.
	EnterDmlStatement(c *DmlStatementContext)

	// EnterDdlStatement is called when entering the ddlStatement production.
	EnterDdlStatement(c *DdlStatementContext)

	// EnterSystemStatement is called when entering the systemStatement production.
	EnterSystemStatement(c *SystemStatementContext)

	// EnterCreateDatabase is called when entering the createDatabase production.
	EnterCreateDatabase(c *CreateDatabaseContext)

	// EnterCreateTable is called when entering the createTable production.
	EnterCreateTable(c *CreateTableContext)

	// EnterTableSchemaClause is called when entering the tableSchemaClause production.
	EnterTableSchemaClause(c *TableSchemaClauseContext)

	// EnterTableElementExpr is called when entering the tableElementExpr production.
	EnterTableElementExpr(c *TableElementExprContext)

	// EnterTableColumnDfnt is called when entering the tableColumnDfnt production.
	EnterTableColumnDfnt(c *TableColumnDfntContext)

	// EnterInsertStatement is called when entering the insertStatement production.
	EnterInsertStatement(c *InsertStatementContext)

	// EnterUpdateStatement is called when entering the updateStatement production.
	EnterUpdateStatement(c *UpdateStatementContext)

	// EnterDeleteStatement is called when entering the deleteStatement production.
	EnterDeleteStatement(c *DeleteStatementContext)

	// EnterSetClause is called when entering the setClause production.
	EnterSetClause(c *SetClauseContext)

	// EnterColumnsClause is called when entering the columnsClause production.
	EnterColumnsClause(c *ColumnsClauseContext)

	// EnterInsertValuesSpec is called when entering the insertValuesSpec production.
	EnterInsertValuesSpec(c *InsertValuesSpecContext)

	// EnterInsertMultiValue is called when entering the insertMultiValue production.
	EnterInsertMultiValue(c *InsertMultiValueContext)

	// EnterDataValue is called when entering the dataValue production.
	EnterDataValue(c *DataValueContext)

	// EnterSelectStatement is called when entering the selectStatement production.
	EnterSelectStatement(c *SelectStatementContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterOrderByExpr is called when entering the orderByExpr production.
	EnterOrderByExpr(c *OrderByExprContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterOffsetClause is called when entering the offsetClause production.
	EnterOffsetClause(c *OffsetClauseContext)

	// EnterSearchCondition is called when entering the searchCondition production.
	EnterSearchCondition(c *SearchConditionContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterJoinExpr is called when entering the joinExpr production.
	EnterJoinExpr(c *JoinExprContext)

	// EnterJoinOp is called when entering the joinOp production.
	EnterJoinOp(c *JoinOpContext)

	// EnterUseStatement is called when entering the useStatement production.
	EnterUseStatement(c *UseStatementContext)

	// EnterColumnTypeExpr is called when entering the columnTypeExpr production.
	EnterColumnTypeExpr(c *ColumnTypeExprContext)

	// EnterColumnExprList is called when entering the columnExprList production.
	EnterColumnExprList(c *ColumnExprListContext)

	// EnterColumnsExpr is called when entering the columnsExpr production.
	EnterColumnsExpr(c *ColumnsExprContext)

	// EnterColumnExpr is called when entering the columnExpr production.
	EnterColumnExpr(c *ColumnExprContext)

	// EnterAggregateFunction is called when entering the aggregateFunction production.
	EnterAggregateFunction(c *AggregateFunctionContext)

	// EnterColumnIdentifier is called when entering the columnIdentifier production.
	EnterColumnIdentifier(c *ColumnIdentifierContext)

	// EnterNestedIdentifier is called when entering the nestedIdentifier production.
	EnterNestedIdentifier(c *NestedIdentifierContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterRealLiteral is called when entering the realLiteral production.
	EnterRealLiteral(c *RealLiteralContext)

	// EnterTableExpr is called when entering the tableExpr production.
	EnterTableExpr(c *TableExprContext)

	// EnterTableIdentifier is called when entering the tableIdentifier production.
	EnterTableIdentifier(c *TableIdentifierContext)

	// EnterDatabaseIdentifier is called when entering the databaseIdentifier production.
	EnterDatabaseIdentifier(c *DatabaseIdentifierContext)

	// EnterFloatingLiteral is called when entering the floatingLiteral production.
	EnterFloatingLiteral(c *FloatingLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// EnterIfNotExists is called when entering the ifNotExists production.
	EnterIfNotExists(c *IfNotExistsContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitDmlStatement is called when exiting the dmlStatement production.
	ExitDmlStatement(c *DmlStatementContext)

	// ExitDdlStatement is called when exiting the ddlStatement production.
	ExitDdlStatement(c *DdlStatementContext)

	// ExitSystemStatement is called when exiting the systemStatement production.
	ExitSystemStatement(c *SystemStatementContext)

	// ExitCreateDatabase is called when exiting the createDatabase production.
	ExitCreateDatabase(c *CreateDatabaseContext)

	// ExitCreateTable is called when exiting the createTable production.
	ExitCreateTable(c *CreateTableContext)

	// ExitTableSchemaClause is called when exiting the tableSchemaClause production.
	ExitTableSchemaClause(c *TableSchemaClauseContext)

	// ExitTableElementExpr is called when exiting the tableElementExpr production.
	ExitTableElementExpr(c *TableElementExprContext)

	// ExitTableColumnDfnt is called when exiting the tableColumnDfnt production.
	ExitTableColumnDfnt(c *TableColumnDfntContext)

	// ExitInsertStatement is called when exiting the insertStatement production.
	ExitInsertStatement(c *InsertStatementContext)

	// ExitUpdateStatement is called when exiting the updateStatement production.
	ExitUpdateStatement(c *UpdateStatementContext)

	// ExitDeleteStatement is called when exiting the deleteStatement production.
	ExitDeleteStatement(c *DeleteStatementContext)

	// ExitSetClause is called when exiting the setClause production.
	ExitSetClause(c *SetClauseContext)

	// ExitColumnsClause is called when exiting the columnsClause production.
	ExitColumnsClause(c *ColumnsClauseContext)

	// ExitInsertValuesSpec is called when exiting the insertValuesSpec production.
	ExitInsertValuesSpec(c *InsertValuesSpecContext)

	// ExitInsertMultiValue is called when exiting the insertMultiValue production.
	ExitInsertMultiValue(c *InsertMultiValueContext)

	// ExitDataValue is called when exiting the dataValue production.
	ExitDataValue(c *DataValueContext)

	// ExitSelectStatement is called when exiting the selectStatement production.
	ExitSelectStatement(c *SelectStatementContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitOrderByExpr is called when exiting the orderByExpr production.
	ExitOrderByExpr(c *OrderByExprContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitOffsetClause is called when exiting the offsetClause production.
	ExitOffsetClause(c *OffsetClauseContext)

	// ExitSearchCondition is called when exiting the searchCondition production.
	ExitSearchCondition(c *SearchConditionContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitJoinExpr is called when exiting the joinExpr production.
	ExitJoinExpr(c *JoinExprContext)

	// ExitJoinOp is called when exiting the joinOp production.
	ExitJoinOp(c *JoinOpContext)

	// ExitUseStatement is called when exiting the useStatement production.
	ExitUseStatement(c *UseStatementContext)

	// ExitColumnTypeExpr is called when exiting the columnTypeExpr production.
	ExitColumnTypeExpr(c *ColumnTypeExprContext)

	// ExitColumnExprList is called when exiting the columnExprList production.
	ExitColumnExprList(c *ColumnExprListContext)

	// ExitColumnsExpr is called when exiting the columnsExpr production.
	ExitColumnsExpr(c *ColumnsExprContext)

	// ExitColumnExpr is called when exiting the columnExpr production.
	ExitColumnExpr(c *ColumnExprContext)

	// ExitAggregateFunction is called when exiting the aggregateFunction production.
	ExitAggregateFunction(c *AggregateFunctionContext)

	// ExitColumnIdentifier is called when exiting the columnIdentifier production.
	ExitColumnIdentifier(c *ColumnIdentifierContext)

	// ExitNestedIdentifier is called when exiting the nestedIdentifier production.
	ExitNestedIdentifier(c *NestedIdentifierContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitRealLiteral is called when exiting the realLiteral production.
	ExitRealLiteral(c *RealLiteralContext)

	// ExitTableExpr is called when exiting the tableExpr production.
	ExitTableExpr(c *TableExprContext)

	// ExitTableIdentifier is called when exiting the tableIdentifier production.
	ExitTableIdentifier(c *TableIdentifierContext)

	// ExitDatabaseIdentifier is called when exiting the databaseIdentifier production.
	ExitDatabaseIdentifier(c *DatabaseIdentifierContext)

	// ExitFloatingLiteral is called when exiting the floatingLiteral production.
	ExitFloatingLiteral(c *FloatingLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)

	// ExitIfNotExists is called when exiting the ifNotExists production.
	ExitIfNotExists(c *IfNotExistsContext)
}
