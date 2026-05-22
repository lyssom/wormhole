package statement

type CreateDatabaseStatement struct {
	Database string
	Kind     string
}

func (s CreateDatabaseStatement) GetKind() string {
	return s.Kind
}

type Column struct {
	Name     string // name of the column.
	DataType string // data type of the column.
}

type CreateTableStatement struct {
	TableName string
	Columns   []Column
	Kind      string
}

func (s CreateTableStatement) GetKind() string {
	return s.Kind
}
