// Code generated from SqlParser.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // SqlParser

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseSqlParserListener is a complete listener for a parse tree produced by SqlParser.
type BaseSqlParserListener struct{}

var _ SqlParserListener = &BaseSqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseSqlParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseSqlParserListener) ExitStatement(ctx *StatementContext) {}

// EnterDmlStatement is called when production dmlStatement is entered.
func (s *BaseSqlParserListener) EnterDmlStatement(ctx *DmlStatementContext) {}

// ExitDmlStatement is called when production dmlStatement is exited.
func (s *BaseSqlParserListener) ExitDmlStatement(ctx *DmlStatementContext) {}

// EnterDdlStatement is called when production ddlStatement is entered.
func (s *BaseSqlParserListener) EnterDdlStatement(ctx *DdlStatementContext) {}

// ExitDdlStatement is called when production ddlStatement is exited.
func (s *BaseSqlParserListener) ExitDdlStatement(ctx *DdlStatementContext) {}

// EnterSystemStatement is called when production systemStatement is entered.
func (s *BaseSqlParserListener) EnterSystemStatement(ctx *SystemStatementContext) {}

// ExitSystemStatement is called when production systemStatement is exited.
func (s *BaseSqlParserListener) ExitSystemStatement(ctx *SystemStatementContext) {}

// EnterCreateDatabase is called when production createDatabase is entered.
func (s *BaseSqlParserListener) EnterCreateDatabase(ctx *CreateDatabaseContext) {}

// ExitCreateDatabase is called when production createDatabase is exited.
func (s *BaseSqlParserListener) ExitCreateDatabase(ctx *CreateDatabaseContext) {}

// EnterCreateTable is called when production createTable is entered.
func (s *BaseSqlParserListener) EnterCreateTable(ctx *CreateTableContext) {}

// ExitCreateTable is called when production createTable is exited.
func (s *BaseSqlParserListener) ExitCreateTable(ctx *CreateTableContext) {}

// EnterTableSchemaClause is called when production tableSchemaClause is entered.
func (s *BaseSqlParserListener) EnterTableSchemaClause(ctx *TableSchemaClauseContext) {}

// ExitTableSchemaClause is called when production tableSchemaClause is exited.
func (s *BaseSqlParserListener) ExitTableSchemaClause(ctx *TableSchemaClauseContext) {}

// EnterTableElementExpr is called when production tableElementExpr is entered.
func (s *BaseSqlParserListener) EnterTableElementExpr(ctx *TableElementExprContext) {}

// ExitTableElementExpr is called when production tableElementExpr is exited.
func (s *BaseSqlParserListener) ExitTableElementExpr(ctx *TableElementExprContext) {}

// EnterTableColumnDfnt is called when production tableColumnDfnt is entered.
func (s *BaseSqlParserListener) EnterTableColumnDfnt(ctx *TableColumnDfntContext) {}

// ExitTableColumnDfnt is called when production tableColumnDfnt is exited.
func (s *BaseSqlParserListener) ExitTableColumnDfnt(ctx *TableColumnDfntContext) {}

// EnterInsertStatement is called when production insertStatement is entered.
func (s *BaseSqlParserListener) EnterInsertStatement(ctx *InsertStatementContext) {}

// ExitInsertStatement is called when production insertStatement is exited.
func (s *BaseSqlParserListener) ExitInsertStatement(ctx *InsertStatementContext) {}

// EnterUpdateStatement is called when production updateStatement is entered.
func (s *BaseSqlParserListener) EnterUpdateStatement(ctx *UpdateStatementContext) {}

// ExitUpdateStatement is called when production updateStatement is exited.
func (s *BaseSqlParserListener) ExitUpdateStatement(ctx *UpdateStatementContext) {}

// EnterDeleteStatement is called when production deleteStatement is entered.
func (s *BaseSqlParserListener) EnterDeleteStatement(ctx *DeleteStatementContext) {}

// ExitDeleteStatement is called when production deleteStatement is exited.
func (s *BaseSqlParserListener) ExitDeleteStatement(ctx *DeleteStatementContext) {}

// EnterSetClause is called when production setClause is entered.
func (s *BaseSqlParserListener) EnterSetClause(ctx *SetClauseContext) {}

// ExitSetClause is called when production setClause is exited.
func (s *BaseSqlParserListener) ExitSetClause(ctx *SetClauseContext) {}

// EnterColumnsClause is called when production columnsClause is entered.
func (s *BaseSqlParserListener) EnterColumnsClause(ctx *ColumnsClauseContext) {}

// ExitColumnsClause is called when production columnsClause is exited.
func (s *BaseSqlParserListener) ExitColumnsClause(ctx *ColumnsClauseContext) {}

// EnterInsertValuesSpec is called when production insertValuesSpec is entered.
func (s *BaseSqlParserListener) EnterInsertValuesSpec(ctx *InsertValuesSpecContext) {}

// ExitInsertValuesSpec is called when production insertValuesSpec is exited.
func (s *BaseSqlParserListener) ExitInsertValuesSpec(ctx *InsertValuesSpecContext) {}

// EnterInsertMultiValue is called when production insertMultiValue is entered.
func (s *BaseSqlParserListener) EnterInsertMultiValue(ctx *InsertMultiValueContext) {}

// ExitInsertMultiValue is called when production insertMultiValue is exited.
func (s *BaseSqlParserListener) ExitInsertMultiValue(ctx *InsertMultiValueContext) {}

// EnterDataValue is called when production dataValue is entered.
func (s *BaseSqlParserListener) EnterDataValue(ctx *DataValueContext) {}

// ExitDataValue is called when production dataValue is exited.
func (s *BaseSqlParserListener) ExitDataValue(ctx *DataValueContext) {}

// EnterSelectStatement is called when production selectStatement is entered.
func (s *BaseSqlParserListener) EnterSelectStatement(ctx *SelectStatementContext) {}

// ExitSelectStatement is called when production selectStatement is exited.
func (s *BaseSqlParserListener) ExitSelectStatement(ctx *SelectStatementContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseSqlParserListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseSqlParserListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSqlParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSqlParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseSqlParserListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseSqlParserListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSqlParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSqlParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseSqlParserListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseSqlParserListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterOrderByExpr is called when production orderByExpr is entered.
func (s *BaseSqlParserListener) EnterOrderByExpr(ctx *OrderByExprContext) {}

// ExitOrderByExpr is called when production orderByExpr is exited.
func (s *BaseSqlParserListener) ExitOrderByExpr(ctx *OrderByExprContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BaseSqlParserListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BaseSqlParserListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterSearchCondition is called when production searchCondition is entered.
func (s *BaseSqlParserListener) EnterSearchCondition(ctx *SearchConditionContext) {}

// ExitSearchCondition is called when production searchCondition is exited.
func (s *BaseSqlParserListener) ExitSearchCondition(ctx *SearchConditionContext) {}

// EnterPredicate is called when production predicate is entered.
func (s *BaseSqlParserListener) EnterPredicate(ctx *PredicateContext) {}

// ExitPredicate is called when production predicate is exited.
func (s *BaseSqlParserListener) ExitPredicate(ctx *PredicateContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterJoinExpr is called when production joinExpr is entered.
func (s *BaseSqlParserListener) EnterJoinExpr(ctx *JoinExprContext) {}

// ExitJoinExpr is called when production joinExpr is exited.
func (s *BaseSqlParserListener) ExitJoinExpr(ctx *JoinExprContext) {}

// EnterJoinOp is called when production joinOp is entered.
func (s *BaseSqlParserListener) EnterJoinOp(ctx *JoinOpContext) {}

// ExitJoinOp is called when production joinOp is exited.
func (s *BaseSqlParserListener) ExitJoinOp(ctx *JoinOpContext) {}

// EnterUseStatement is called when production useStatement is entered.
func (s *BaseSqlParserListener) EnterUseStatement(ctx *UseStatementContext) {}

// ExitUseStatement is called when production useStatement is exited.
func (s *BaseSqlParserListener) ExitUseStatement(ctx *UseStatementContext) {}

// EnterColumnTypeExpr is called when production columnTypeExpr is entered.
func (s *BaseSqlParserListener) EnterColumnTypeExpr(ctx *ColumnTypeExprContext) {}

// ExitColumnTypeExpr is called when production columnTypeExpr is exited.
func (s *BaseSqlParserListener) ExitColumnTypeExpr(ctx *ColumnTypeExprContext) {}

// EnterColumnExprList is called when production columnExprList is entered.
func (s *BaseSqlParserListener) EnterColumnExprList(ctx *ColumnExprListContext) {}

// ExitColumnExprList is called when production columnExprList is exited.
func (s *BaseSqlParserListener) ExitColumnExprList(ctx *ColumnExprListContext) {}

// EnterColumnsExpr is called when production columnsExpr is entered.
func (s *BaseSqlParserListener) EnterColumnsExpr(ctx *ColumnsExprContext) {}

// ExitColumnsExpr is called when production columnsExpr is exited.
func (s *BaseSqlParserListener) ExitColumnsExpr(ctx *ColumnsExprContext) {}

// EnterColumnExpr is called when production columnExpr is entered.
func (s *BaseSqlParserListener) EnterColumnExpr(ctx *ColumnExprContext) {}

// ExitColumnExpr is called when production columnExpr is exited.
func (s *BaseSqlParserListener) ExitColumnExpr(ctx *ColumnExprContext) {}

// EnterAggregateFunction is called when production aggregateFunction is entered.
func (s *BaseSqlParserListener) EnterAggregateFunction(ctx *AggregateFunctionContext) {}

// ExitAggregateFunction is called when production aggregateFunction is exited.
func (s *BaseSqlParserListener) ExitAggregateFunction(ctx *AggregateFunctionContext) {}

// EnterColumnIdentifier is called when production columnIdentifier is entered.
func (s *BaseSqlParserListener) EnterColumnIdentifier(ctx *ColumnIdentifierContext) {}

// ExitColumnIdentifier is called when production columnIdentifier is exited.
func (s *BaseSqlParserListener) ExitColumnIdentifier(ctx *ColumnIdentifierContext) {}

// EnterNestedIdentifier is called when production nestedIdentifier is entered.
func (s *BaseSqlParserListener) EnterNestedIdentifier(ctx *NestedIdentifierContext) {}

// ExitNestedIdentifier is called when production nestedIdentifier is exited.
func (s *BaseSqlParserListener) ExitNestedIdentifier(ctx *NestedIdentifierContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseSqlParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseSqlParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterRealLiteral is called when production realLiteral is entered.
func (s *BaseSqlParserListener) EnterRealLiteral(ctx *RealLiteralContext) {}

// ExitRealLiteral is called when production realLiteral is exited.
func (s *BaseSqlParserListener) ExitRealLiteral(ctx *RealLiteralContext) {}

// EnterTableExpr is called when production tableExpr is entered.
func (s *BaseSqlParserListener) EnterTableExpr(ctx *TableExprContext) {}

// ExitTableExpr is called when production tableExpr is exited.
func (s *BaseSqlParserListener) ExitTableExpr(ctx *TableExprContext) {}

// EnterTableIdentifier is called when production tableIdentifier is entered.
func (s *BaseSqlParserListener) EnterTableIdentifier(ctx *TableIdentifierContext) {}

// ExitTableIdentifier is called when production tableIdentifier is exited.
func (s *BaseSqlParserListener) ExitTableIdentifier(ctx *TableIdentifierContext) {}

// EnterDatabaseIdentifier is called when production databaseIdentifier is entered.
func (s *BaseSqlParserListener) EnterDatabaseIdentifier(ctx *DatabaseIdentifierContext) {}

// ExitDatabaseIdentifier is called when production databaseIdentifier is exited.
func (s *BaseSqlParserListener) ExitDatabaseIdentifier(ctx *DatabaseIdentifierContext) {}

// EnterFloatingLiteral is called when production floatingLiteral is entered.
func (s *BaseSqlParserListener) EnterFloatingLiteral(ctx *FloatingLiteralContext) {}

// ExitFloatingLiteral is called when production floatingLiteral is exited.
func (s *BaseSqlParserListener) ExitFloatingLiteral(ctx *FloatingLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseSqlParserListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseSqlParserListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseSqlParserListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseSqlParserListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseSqlParserListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseSqlParserListener) ExitInterval(ctx *IntervalContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseSqlParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseSqlParserListener) ExitKeyword(ctx *KeywordContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseSqlParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseSqlParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseSqlParserListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseSqlParserListener) ExitEnumValue(ctx *EnumValueContext) {}

// EnterIfNotExists is called when production ifNotExists is entered.
func (s *BaseSqlParserListener) EnterIfNotExists(ctx *IfNotExistsContext) {}

// ExitIfNotExists is called when production ifNotExists is exited.
func (s *BaseSqlParserListener) ExitIfNotExists(ctx *IfNotExistsContext) {}
