package statement

type UseStatement struct {
	Database string
	Kind     string
}

func (s UseStatement) GetKind() string {
	return s.Kind
}
