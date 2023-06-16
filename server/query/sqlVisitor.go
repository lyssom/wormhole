package query

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"

	"wormhole/server/query/parser"
	"wormhole/server/query/statement"
	"wormhole/server/query/statement/component"
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
	db := mv.VisitChildren(ctx)
	return statement.CreateDatabaseStatement{
		Database: db.(string),
		Kind:     "create database",
	}
}

func (mv *MyVisitor) VisitDatabaseIdentifier(ctx *parser.DatabaseIdentifierContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitCreateTable(ctx *parser.CreateTableContext) interface{} {
	tableName := mv.Visit(ctx.TableIdentifier())
	columns := mv.Visit(ctx.TableSchemaClause()).([]statement.Column)

	ct := statement.CreateTableStatement{
		Kind:      "create table",
		TableName: tableName.(string),
		Columns:   columns,
	}

	return ct
}

func (mv *MyVisitor) VisitTableIdentifier(ctx *parser.TableIdentifierContext) interface{} {
	return mv.VisitChildren(ctx)
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
	fromClause := mv.Visit(ctx.FromClause())

	qs := statement.QueryStatement{
		SelectComponent: component.SelectComponent{
			ResultColumns: selectClause.([]string),
		},
		FromComponent: component.FromComponent{FromClause: fromClause.(string)},
		RowLimit:      10,
		RowOffset:     0,
		Kind:          "select",
	}

	return qs
}

func (mv *MyVisitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitJoinExpr(ctx *parser.JoinExprContext) interface{} {
	return mv.VisitChildren(ctx)
}

func (mv *MyVisitor) VisitTableExpr(ctx *parser.TableExprContext) interface{} {
	return mv.VisitChildren(ctx)
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
	tableName := mv.Visit(ctx.TableIdentifier()).(string)
	insertColumns := mv.Visit(ctx.ColumnsClause()).([]string)
	insertValues := mv.Visit(ctx.InsertValuesSpec()).([][]interface{})
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
		values = append(values, mv.Visit(mVal).([]interface{}))
	}
	return values
}

func (mv *MyVisitor) VisitInsertMultiValue(ctx *parser.InsertMultiValueContext) interface{} {
	var values []interface{}
	for _, dVal := range ctx.AllDataValue() {
		values = append(values, mv.Visit(dVal))
	}
	return values
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
