package query

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"

	"wormhole/server/sql/expression"
	"wormhole/server/sql/parser"
	"wormhole/server/sql/statement"
	"wormhole/server/sql/statement/component"
)

type MyVisitor struct {
	parser.BaseSqlParserVisitor
}

func NewVisitor() *MyVisitor {
	visitor := new(MyVisitor)
	return visitor
}

func (mv *MyVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(mv)
}

func (mv *MyVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var state interface{}
	for _, child := range node.GetChildren() {
		state = child.(antlr.ParseTree).Accept(mv)
	}
	return state
}

func (mv *MyVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return nil
}

func (mv *MyVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

func (mv *MyVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitDdlStatement(ctx *parser.DdlStatementContext) interface{} {
	return mv.VisitChildren(ctx)

}

func (mv *MyVisitor) VisitDmlStatement(ctx *parser.DmlStatementContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitSystemStatement(ctx *parser.SystemStatementContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitCreateDatabase(ctx *parser.CreateDatabaseContext) interface{} {
	db := mv.Visit(ctx.DatabaseIdentifier()).(string)
	return statement.CreateDatabaseStatement{
		Database: db,
		Kind:     "create database",
	}
}

func (mv *MyVisitor) VisitDatabaseIdentifier(ctx *parser.DatabaseIdentifierContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitCreateTable(ctx *parser.CreateTableContext) interface{} {
	tableNameResult := mv.Visit(ctx.TableIdentifier())
	var tableName string
	if tableNameResult != nil {
		if ids, ok := tableNameResult.([]string); ok {
			tableName = ids[len(ids)-1] // Use the last identifier (actual table name)
		}
	}
	columnsResult := mv.Visit(ctx.TableSchemaClause())
	var columns []statement.Column
	if columnsResult != nil {
		columns = columnsResult.([]statement.Column)
	}

	ct := statement.CreateTableStatement{
		Kind:      "create table",
		TableName: tableName,
		Columns:   columns,
	}

	return ct
}

func (mv *MyVisitor) VisitTableIdentifier(ctx *parser.TableIdentifierContext) interface{} {
	var ids []string
	if ctx.DatabaseIdentifier() != nil {
		if dbResult := mv.Visit(ctx.DatabaseIdentifier()); dbResult != nil {
			ids = append(ids, dbResult.(string))
		}
	}
	if ctx.Identifier() != nil {
		ids = append(ids, ctx.Identifier().GetText())
	}
	return ids
}

func (mv *MyVisitor) VisitTableSchemaClause(ctx *parser.TableSchemaClauseContext) interface{} {
	var columns []statement.Column
	for _, eleExpr := range ctx.AllTableElementExpr() {
		columns = append(columns, mv.Visit(eleExpr).(statement.Column))
	}
	return columns
}

func (mv *MyVisitor) VisitTableElementExpr(ctx *parser.TableElementExprContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitTableColumnDfnt(ctx *parser.TableColumnDfntContext) interface{} {
	dataType := mv.Visit(ctx.ColumnTypeExpr()).([]string)[0]
	column := mv.Visit(ctx.NestedIdentifier()).([]string)[0]
	return statement.Column{
		Name:     column,
		DataType: dataType,
	}
}

func (mv *MyVisitor) VisitColumnTypeExpr(ctx *parser.ColumnTypeExprContext) interface{} {
	var idfs []string
	for _, idf := range ctx.AllIdentifier() {
		idfs = append(idfs, mv.Visit(idf).(string))
	}
	return idfs
}

func (mv *MyVisitor) VisitNestedIdentifier(ctx *parser.NestedIdentifierContext) interface{} {

	var idfs []string
	for _, idf := range ctx.AllIdentifier() {
		idfs = append(idfs, mv.Visit(idf).(string))
	}
	return idfs
}

func (mv *MyVisitor) VisitSelectStatement(ctx *parser.SelectStatementContext) interface{} {
	selectClause := mv.Visit(ctx.ColumnExprList())

	var fromClause string
	if ctx.FromClause() != nil {
		if fc := mv.Visit(ctx.FromClause()); fc != nil {
			fromClause = fc.(string)
		}
	}

	var whereCondition component.FilterComponent
	if ctx.WhereClause() != nil {
		whereCondition = component.FilterComponent{
			Condition: mv.Visit(ctx.WhereClause()),
		}
	}

	var groupBy component.GroupByComponent
	if ctx.GroupByClause() != nil {
		gbResult := mv.Visit(ctx.GroupByClause())
		if gbResult != nil {
			groupBy = component.GroupByComponent{
				Columns: gbResult.([]string),
			}
		}
	}

	var having component.HavingComponent
	if ctx.HavingClause() != nil {
		having = component.HavingComponent{
			Condition: mv.Visit(ctx.HavingClause()),
		}
	}

	var orderBy component.OrderByComponent
	if ctx.OrderByClause() != nil {
		obResult := mv.Visit(ctx.OrderByClause())
		if obResult != nil {
			orderBy = obResult.(component.OrderByComponent)
		}
	}

	rowLimit := 0
	rowOffset := 0
	if ctx.LimitClause() != nil {
		limitText := ctx.LimitClause().GetText()
		fmt.Sscanf(limitText, "LIMIT %d", &rowLimit)
	}
	if ctx.OffsetClause() != nil {
		offsetText := ctx.OffsetClause().GetText()
		fmt.Sscanf(offsetText, "OFFSET %d", &rowOffset)
	}

	qs := statement.QueryStatement{
		SelectComponent: component.SelectComponent{
			ResultColumns: selectClause.([]string),
		},
		FromComponent:    component.FromComponent{FromClause: fromClause},
		WhereCondition:   whereCondition,
		GroupByComponent: groupBy,
		HavingCondition:  having,
		OrderByComponent: orderBy,
		RowLimit:         rowLimit,
		RowOffset:        rowOffset,
		Kind:             "select",
	}

	return qs
}

func (mv *MyVisitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	// FromClause has a joinExpr child
	if ctx.JoinExpr() != nil {
		joinResult := mv.Visit(ctx.JoinExpr())
		if joinResult != nil {
			if tableIds, ok := joinResult.([]string); ok && len(tableIds) > 0 {
				return tableIds[len(tableIds)-1]
			}
		}
	}
	return ""
}

func (mv *MyVisitor) VisitColumnExprList(ctx *parser.ColumnExprListContext) interface{} {
	var columns []string
	for _, expr := range ctx.AllColumnsExpr() {
		columns = append(columns, mv.Visit(expr).(string))
	}
	return columns
}

func (mv *MyVisitor) VisitColumnsExpr(ctx *parser.ColumnsExprContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitColumnExpr(ctx *parser.ColumnExprContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitInsertStatement(ctx *parser.InsertStatementContext) interface{} {
	tableNameResult := mv.Visit(ctx.TableIdentifier())
	var tableName string
	if tableNameResult != nil {
		if ids, ok := tableNameResult.([]string); ok {
			tableName = ids[len(ids)-1]
		}
	}

	insertColumnsResult := mv.Visit(ctx.ColumnsClause())
	var insertColumns []string
	if insertColumnsResult != nil {
		insertColumns = insertColumnsResult.([]string)
	}

	insertValuesResult := mv.Visit(ctx.InsertValuesSpec())
	var insertValues [][]interface{}
	if insertValuesResult != nil {
		insertValues = insertValuesResult.([][]interface{})
	}

	return statement.InsertStatement{
		Columns:   insertColumns,
		Values:    insertValues,
		TableName: tableName,
		Kind:      "insert",
	}
}

func (mv *MyVisitor) VisitInsertValuesSpec(ctx *parser.InsertValuesSpecContext) interface{} {

	var values [][]interface{}
	for _, mVal := range ctx.AllInsertMultiValue() {
		if mvResult := mv.Visit(mVal); mvResult != nil {
			values = append(values, mvResult.([]interface{}))
		}
	}
	return values
}

func (mv *MyVisitor) VisitInsertMultiValue(ctx *parser.InsertMultiValueContext) interface{} {
	var values []interface{}
	for _, dVal := range ctx.AllDataValue() {
		if dValResult := mv.Visit(dVal); dValResult != nil {
			values = append(values, dValResult)
		}
	}
	return values
}

func (mv *MyVisitor) VisitJoinExpr(ctx *parser.JoinExprContext) interface{} {
	// Visit the first table expression
	for _, te := range ctx.AllTableExpr() {
		if teResult := mv.Visit(te); teResult != nil {
			return teResult
		}
	}
	return nil
}

func (mv *MyVisitor) VisitTableExpr(ctx *parser.TableExprContext) interface{} {
	if ctx.TableIdentifier() != nil {
		return mv.Visit(ctx.TableIdentifier())
	}
	return nil
}

func (mv *MyVisitor) VisitDataValue(ctx *parser.DataValueContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitColumnsClause(ctx *parser.ColumnsClauseContext) interface{} {
	var columns []string
	for _, nidf := range ctx.AllNestedIdentifier() {
		columns = append(columns, mv.Visit(nidf).([]string)[0])
	}
	return columns
}

func (mv *MyVisitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	return ctx.GetText()
}

func (mv *MyVisitor) VisitUseStatement(ctx *parser.UseStatementContext) interface{} {
	return statement.UseStatement{
		Database: mv.VisitChildren(ctx).(string),
		Kind:     "use",
	}
}

func (mv *MyVisitor) VisitGroupByClause(ctx *parser.GroupByClauseContext) interface{} {
	var columns []string
	for _, col := range ctx.AllColumnIdentifier() {
		colID := mv.Visit(col).([]string)
		columns = append(columns, colID[len(colID)-1])
	}
	return columns
}

func (mv *MyVisitor) VisitHavingClause(ctx *parser.HavingClauseContext) interface{} {
	return mv.Visit(ctx.SearchCondition())
}

func (mv *MyVisitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	var exprs []component.OrderByExpression
	for _, ob := range ctx.AllOrderByExpr() {
		exprs = append(exprs, mv.Visit(ob).(component.OrderByExpression))
	}
	return component.OrderByComponent{Expressions: exprs}
}

func (mv *MyVisitor) VisitOrderByExpr(ctx *parser.OrderByExprContext) interface{} {
	colID := mv.Visit(ctx.ColumnIdentifier()).([]string)
	direction := "ASC"
	if ctx.DESC() != nil {
		direction = "DESC"
	}
	return component.OrderByExpression{
		Column:    colID[len(colID)-1],
		Direction: direction,
	}
}

func (mv *MyVisitor) VisitUpdateStatement(ctx *parser.UpdateStatementContext) interface{} {
	tableNameResult := mv.Visit(ctx.TableIdentifier())
	var tableName string
	if tableNameResult != nil {
		if ids, ok := tableNameResult.([]string); ok {
			tableName = ids[len(ids)-1]
		}
	}

	var setClauses []statement.SetClause
	for _, set := range ctx.AllSetClause() {
		if setResult := mv.Visit(set); setResult != nil {
			setClauses = append(setClauses, setResult.(statement.SetClause))
		}
	}

	var whereCondition interface{}
	if ctx.WhereClause() != nil {
		whereCondition = mv.Visit(ctx.WhereClause())
	}

	return statement.UpdateStatement{
		TableName:      tableName,
		SetClauses:     setClauses,
		WhereCondition: whereCondition,
		Kind:           "update",
	}
}

func (mv *MyVisitor) VisitColumnIdentifier(ctx *parser.ColumnIdentifierContext) interface{} {
	var ids []string
	if ctx.TableIdentifier() != nil {
		if tiResult := mv.Visit(ctx.TableIdentifier()); tiResult != nil {
			if ti, ok := tiResult.([]string); ok {
				ids = append(ids, ti...)
			}
		}
	}
	if ctx.NestedIdentifier() != nil {
		if niResult := mv.Visit(ctx.NestedIdentifier()); niResult != nil {
			if ni, ok := niResult.([]string); ok {
				ids = append(ids, ni...)
			}
		}
	}
	return ids
}

func (mv *MyVisitor) VisitSetClause(ctx *parser.SetClauseContext) interface{} {
	colIDResult := mv.Visit(ctx.ColumnIdentifier())
	var colName string
	if colIDResult != nil {
		if colID, ok := colIDResult.([]string); ok {
			colName = colID[len(colID)-1]
		}
	}
	return statement.SetClause{
		Column: colName,
		Value:  mv.Visit(ctx.Expression()),
	}
}

func (mv *MyVisitor) VisitDeleteStatement(ctx *parser.DeleteStatementContext) interface{} {
	tableNameResult := mv.Visit(ctx.TableIdentifier())
	var tableName string
	if tableNameResult != nil {
		if ids, ok := tableNameResult.([]string); ok {
			tableName = ids[len(ids)-1]
		}
	}

	var whereCondition interface{}
	if ctx.WhereClause() != nil {
		whereCondition = mv.Visit(ctx.WhereClause())
	}

	return statement.DeleteStatement{
		TableName:      tableName,
		WhereCondition: whereCondition,
		Kind:           "delete",
	}
}

func (mv *MyVisitor) VisitAggregateFunction(ctx *parser.AggregateFunctionContext) interface{} {
	funcName := ""
	if ctx.COUNT() != nil {
		funcName = "COUNT"
	} else if ctx.SUM() != nil {
		funcName = "SUM"
	} else if ctx.AVG() != nil {
		funcName = "AVG"
	} else if ctx.MIN() != nil {
		funcName = "MIN"
	} else if ctx.MAX() != nil {
		funcName = "MAX"
	}

	var arg expression.Expression
	if ctx.ASTERISK() != nil {
		arg = nil // COUNT(*)
	} else if ctx.Expression() != nil {
		arg = mv.Visit(ctx.Expression()).(expression.Expression)
	}

	return &expression.AggregateFunction{
		FuncName: funcName,
		Arg:      arg,
	}
}

func (mv *MyVisitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return mv.Visit(ctx.SearchCondition())
}

func (mv *MyVisitor) VisitSearchCondition(ctx *parser.SearchConditionContext) interface{} {
	if ctx.NOT() != nil {
		return &expression.LogicalNot{
			Operand: mv.Visit(ctx.SearchCondition(0)).(expression.Expression),
		}
	}

	if ctx.AND() != nil {
		return &expression.BooleanExpression{
			Left:     mv.Visit(ctx.SearchCondition(0)).(expression.Expression),
			Operator: "AND",
			Right:    mv.Visit(ctx.SearchCondition(1)).(expression.Expression),
		}
	}

	if ctx.OR() != nil {
		return &expression.BooleanExpression{
			Left:     mv.Visit(ctx.SearchCondition(0)).(expression.Expression),
			Operator: "OR",
			Right:    mv.Visit(ctx.SearchCondition(1)).(expression.Expression),
		}
	}

	return mv.Visit(ctx.Predicate())
}

func (mv *MyVisitor) VisitPredicate(ctx *parser.PredicateContext) interface{} {
	leftResult := mv.Visit(ctx.Expression(0))
	rightResult := mv.Visit(ctx.Expression(1))
	if leftResult == nil || rightResult == nil {
		return nil
	}
	left := leftResult.(expression.Expression)
	right := rightResult.(expression.Expression)

	if ctx.EQ_SINGLE() != nil {
		return &expression.ComparisonExpression{Left: left, Operator: "=", Right: right}
	}
	if ctx.NOT_EQ() != nil {
		return &expression.ComparisonExpression{Left: left, Operator: "!=", Right: right}
	}
	if ctx.LT() != nil {
		return &expression.ComparisonExpression{Left: left, Operator: "<", Right: right}
	}
	if ctx.LE() != nil {
		return &expression.ComparisonExpression{Left: left, Operator: "<=", Right: right}
	}
	if ctx.GT() != nil {
		return &expression.ComparisonExpression{Left: left, Operator: ">", Right: right}
	}
	if ctx.GE() != nil {
		return &expression.ComparisonExpression{Left: left, Operator: ">=", Right: right}
	}

	return nil
}

func (mv *MyVisitor) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	if ctx.ColumnIdentifier() != nil {
		colIDResult := mv.Visit(ctx.ColumnIdentifier())
		if colIDResult == nil {
			return nil
		}
		colID := colIDResult.([]string)
		if len(colID) == 2 {
			return &expression.ColumnRef{Table: colID[0], Column: colID[1]}
		}
		return &expression.ColumnRef{Column: colID[0]}
	}

	if ctx.Literal() != nil {
		litText := ctx.GetText()
		// Try to parse as number
		var val interface{}
		if _, err := fmt.Sscanf(litText, "%f", &val); err == nil {
			// Try int
			var i int64
			if _, err := fmt.Sscanf(litText, "%d", &i); err == nil {
				val = i
			} else {
				val = float64(val.(float64))
			}
		} else {
			// String literal - remove quotes
			if len(litText) >= 2 {
				val = litText[1 : len(litText)-1]
			} else {
				val = litText
			}
		}
		return &expression.Literal{Value: val}
	}

	// Binary arithmetic expression
	if len(ctx.AllExpression()) == 2 {
		leftResult := mv.Visit(ctx.Expression(0))
		rightResult := mv.Visit(ctx.Expression(1))
		if leftResult == nil || rightResult == nil {
			return nil
		}
		left := leftResult.(expression.Expression)
		right := rightResult.(expression.Expression)
		var op string
		if ctx.ASTERISK() != nil {
			op = "*"
		} else if ctx.SLASH() != nil {
			op = "/"
		} else if ctx.PERCENT() != nil {
			op = "%"
		} else if ctx.PLUS() != nil {
			op = "+"
		} else if ctx.DASH() != nil {
			op = "-"
		}
		return &expression.BinaryArithmetic{Left: left, Operator: op, Right: right}
	}

	return nil
}
