// Generated from SqlParser.g4 by ANTLR 4.7.2
package parser;
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link SqlParser}.
 */
public interface SqlParserListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link SqlParser#statement}.
	 * @param ctx the parse tree
	 */
	void enterStatement(SqlParser.StatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#statement}.
	 * @param ctx the parse tree
	 */
	void exitStatement(SqlParser.StatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#dmlStatement}.
	 * @param ctx the parse tree
	 */
	void enterDmlStatement(SqlParser.DmlStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#dmlStatement}.
	 * @param ctx the parse tree
	 */
	void exitDmlStatement(SqlParser.DmlStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#ddlStatement}.
	 * @param ctx the parse tree
	 */
	void enterDdlStatement(SqlParser.DdlStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#ddlStatement}.
	 * @param ctx the parse tree
	 */
	void exitDdlStatement(SqlParser.DdlStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#systemStatement}.
	 * @param ctx the parse tree
	 */
	void enterSystemStatement(SqlParser.SystemStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#systemStatement}.
	 * @param ctx the parse tree
	 */
	void exitSystemStatement(SqlParser.SystemStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#createDatabase}.
	 * @param ctx the parse tree
	 */
	void enterCreateDatabase(SqlParser.CreateDatabaseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#createDatabase}.
	 * @param ctx the parse tree
	 */
	void exitCreateDatabase(SqlParser.CreateDatabaseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#createTable}.
	 * @param ctx the parse tree
	 */
	void enterCreateTable(SqlParser.CreateTableContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#createTable}.
	 * @param ctx the parse tree
	 */
	void exitCreateTable(SqlParser.CreateTableContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#tableSchemaClause}.
	 * @param ctx the parse tree
	 */
	void enterTableSchemaClause(SqlParser.TableSchemaClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#tableSchemaClause}.
	 * @param ctx the parse tree
	 */
	void exitTableSchemaClause(SqlParser.TableSchemaClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#tableElementExpr}.
	 * @param ctx the parse tree
	 */
	void enterTableElementExpr(SqlParser.TableElementExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#tableElementExpr}.
	 * @param ctx the parse tree
	 */
	void exitTableElementExpr(SqlParser.TableElementExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#tableColumnDfnt}.
	 * @param ctx the parse tree
	 */
	void enterTableColumnDfnt(SqlParser.TableColumnDfntContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#tableColumnDfnt}.
	 * @param ctx the parse tree
	 */
	void exitTableColumnDfnt(SqlParser.TableColumnDfntContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#insertStatement}.
	 * @param ctx the parse tree
	 */
	void enterInsertStatement(SqlParser.InsertStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#insertStatement}.
	 * @param ctx the parse tree
	 */
	void exitInsertStatement(SqlParser.InsertStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#columnsClause}.
	 * @param ctx the parse tree
	 */
	void enterColumnsClause(SqlParser.ColumnsClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#columnsClause}.
	 * @param ctx the parse tree
	 */
	void exitColumnsClause(SqlParser.ColumnsClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#insertValuesSpec}.
	 * @param ctx the parse tree
	 */
	void enterInsertValuesSpec(SqlParser.InsertValuesSpecContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#insertValuesSpec}.
	 * @param ctx the parse tree
	 */
	void exitInsertValuesSpec(SqlParser.InsertValuesSpecContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#insertMultiValue}.
	 * @param ctx the parse tree
	 */
	void enterInsertMultiValue(SqlParser.InsertMultiValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#insertMultiValue}.
	 * @param ctx the parse tree
	 */
	void exitInsertMultiValue(SqlParser.InsertMultiValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#dataValue}.
	 * @param ctx the parse tree
	 */
	void enterDataValue(SqlParser.DataValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#dataValue}.
	 * @param ctx the parse tree
	 */
	void exitDataValue(SqlParser.DataValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#selectStatement}.
	 * @param ctx the parse tree
	 */
	void enterSelectStatement(SqlParser.SelectStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#selectStatement}.
	 * @param ctx the parse tree
	 */
	void exitSelectStatement(SqlParser.SelectStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#fromClause}.
	 * @param ctx the parse tree
	 */
	void enterFromClause(SqlParser.FromClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#fromClause}.
	 * @param ctx the parse tree
	 */
	void exitFromClause(SqlParser.FromClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void enterWhereClause(SqlParser.WhereClauseContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#whereClause}.
	 * @param ctx the parse tree
	 */
	void exitWhereClause(SqlParser.WhereClauseContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#searchCondition}.
	 * @param ctx the parse tree
	 */
	void enterSearchCondition(SqlParser.SearchConditionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#searchCondition}.
	 * @param ctx the parse tree
	 */
	void exitSearchCondition(SqlParser.SearchConditionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#predicate}.
	 * @param ctx the parse tree
	 */
	void enterPredicate(SqlParser.PredicateContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#predicate}.
	 * @param ctx the parse tree
	 */
	void exitPredicate(SqlParser.PredicateContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#expression}.
	 * @param ctx the parse tree
	 */
	void enterExpression(SqlParser.ExpressionContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#expression}.
	 * @param ctx the parse tree
	 */
	void exitExpression(SqlParser.ExpressionContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#joinExpr}.
	 * @param ctx the parse tree
	 */
	void enterJoinExpr(SqlParser.JoinExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#joinExpr}.
	 * @param ctx the parse tree
	 */
	void exitJoinExpr(SqlParser.JoinExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#useStatement}.
	 * @param ctx the parse tree
	 */
	void enterUseStatement(SqlParser.UseStatementContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#useStatement}.
	 * @param ctx the parse tree
	 */
	void exitUseStatement(SqlParser.UseStatementContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#columnTypeExpr}.
	 * @param ctx the parse tree
	 */
	void enterColumnTypeExpr(SqlParser.ColumnTypeExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#columnTypeExpr}.
	 * @param ctx the parse tree
	 */
	void exitColumnTypeExpr(SqlParser.ColumnTypeExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#columnExprList}.
	 * @param ctx the parse tree
	 */
	void enterColumnExprList(SqlParser.ColumnExprListContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#columnExprList}.
	 * @param ctx the parse tree
	 */
	void exitColumnExprList(SqlParser.ColumnExprListContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#columnsExpr}.
	 * @param ctx the parse tree
	 */
	void enterColumnsExpr(SqlParser.ColumnsExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#columnsExpr}.
	 * @param ctx the parse tree
	 */
	void exitColumnsExpr(SqlParser.ColumnsExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#columnExpr}.
	 * @param ctx the parse tree
	 */
	void enterColumnExpr(SqlParser.ColumnExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#columnExpr}.
	 * @param ctx the parse tree
	 */
	void exitColumnExpr(SqlParser.ColumnExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#columnIdentifier}.
	 * @param ctx the parse tree
	 */
	void enterColumnIdentifier(SqlParser.ColumnIdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#columnIdentifier}.
	 * @param ctx the parse tree
	 */
	void exitColumnIdentifier(SqlParser.ColumnIdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#nestedIdentifier}.
	 * @param ctx the parse tree
	 */
	void enterNestedIdentifier(SqlParser.NestedIdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#nestedIdentifier}.
	 * @param ctx the parse tree
	 */
	void exitNestedIdentifier(SqlParser.NestedIdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#constant}.
	 * @param ctx the parse tree
	 */
	void enterConstant(SqlParser.ConstantContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#constant}.
	 * @param ctx the parse tree
	 */
	void exitConstant(SqlParser.ConstantContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#realLiteral}.
	 * @param ctx the parse tree
	 */
	void enterRealLiteral(SqlParser.RealLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#realLiteral}.
	 * @param ctx the parse tree
	 */
	void exitRealLiteral(SqlParser.RealLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#tableExpr}.
	 * @param ctx the parse tree
	 */
	void enterTableExpr(SqlParser.TableExprContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#tableExpr}.
	 * @param ctx the parse tree
	 */
	void exitTableExpr(SqlParser.TableExprContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#tableIdentifier}.
	 * @param ctx the parse tree
	 */
	void enterTableIdentifier(SqlParser.TableIdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#tableIdentifier}.
	 * @param ctx the parse tree
	 */
	void exitTableIdentifier(SqlParser.TableIdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#databaseIdentifier}.
	 * @param ctx the parse tree
	 */
	void enterDatabaseIdentifier(SqlParser.DatabaseIdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#databaseIdentifier}.
	 * @param ctx the parse tree
	 */
	void exitDatabaseIdentifier(SqlParser.DatabaseIdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#floatingLiteral}.
	 * @param ctx the parse tree
	 */
	void enterFloatingLiteral(SqlParser.FloatingLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#floatingLiteral}.
	 * @param ctx the parse tree
	 */
	void exitFloatingLiteral(SqlParser.FloatingLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#numberLiteral}.
	 * @param ctx the parse tree
	 */
	void enterNumberLiteral(SqlParser.NumberLiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#numberLiteral}.
	 * @param ctx the parse tree
	 */
	void exitNumberLiteral(SqlParser.NumberLiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#literal}.
	 * @param ctx the parse tree
	 */
	void enterLiteral(SqlParser.LiteralContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#literal}.
	 * @param ctx the parse tree
	 */
	void exitLiteral(SqlParser.LiteralContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#interval}.
	 * @param ctx the parse tree
	 */
	void enterInterval(SqlParser.IntervalContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#interval}.
	 * @param ctx the parse tree
	 */
	void exitInterval(SqlParser.IntervalContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#keyword}.
	 * @param ctx the parse tree
	 */
	void enterKeyword(SqlParser.KeywordContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#keyword}.
	 * @param ctx the parse tree
	 */
	void exitKeyword(SqlParser.KeywordContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#identifier}.
	 * @param ctx the parse tree
	 */
	void enterIdentifier(SqlParser.IdentifierContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#identifier}.
	 * @param ctx the parse tree
	 */
	void exitIdentifier(SqlParser.IdentifierContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#enumValue}.
	 * @param ctx the parse tree
	 */
	void enterEnumValue(SqlParser.EnumValueContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#enumValue}.
	 * @param ctx the parse tree
	 */
	void exitEnumValue(SqlParser.EnumValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link SqlParser#ifNotExists}.
	 * @param ctx the parse tree
	 */
	void enterIfNotExists(SqlParser.IfNotExistsContext ctx);
	/**
	 * Exit a parse tree produced by {@link SqlParser#ifNotExists}.
	 * @param ctx the parse tree
	 */
	void exitIfNotExists(SqlParser.IfNotExistsContext ctx);
}