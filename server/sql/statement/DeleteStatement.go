package statement

type DeleteStatement struct {
	TableName      string
	WhereCondition interface{} // expression.Expression
	Kind           string
}

func (d DeleteStatement) GetKind() string {
	return d.Kind
}