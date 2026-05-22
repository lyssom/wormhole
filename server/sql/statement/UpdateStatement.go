package statement

type UpdateStatement struct {
	TableName      string
	SetClauses     []SetClause
	WhereCondition interface{} // component.FilterComponent - but we store expression.Expression directly
	Kind           string
}

type SetClause struct {
	Column string
	Value  interface{} // expression.Expression
}

func (u UpdateStatement) GetKind() string {
	return u.Kind
}