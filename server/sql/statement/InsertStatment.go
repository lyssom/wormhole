package statement

type InsertStatement struct {
	TableName string
	Columns   []string
	Values    [][]interface{}
	Kind      string
}

func (s InsertStatement) GetKind() string {
	return s.Kind
}
