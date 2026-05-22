package component

type SelectComponent struct {
	ResultColumns    []string
	AliasToColumnMap map[string]string
}

type FromComponent struct {
	FromClause string
	Tables     []TableRef // for JOIN support
}

type TableRef struct {
	Name    string
	Alias   string
	Join    *JoinSpec
}

type JoinSpec struct {
	Type      string // "INNER", "LEFT"
	Table     string
	Condition interface{} // expression.Expression
}

type FilterComponent struct {
	Condition interface{} // expression.Expression
}

type GroupByComponent struct {
	Columns []string
}

type HavingComponent struct {
	Condition interface{} // expression.Expression
}

type OrderByComponent struct {
	Expressions []OrderByExpression
}

type OrderByExpression struct {
	Column    string
	Direction string // "ASC" or "DESC"
}
