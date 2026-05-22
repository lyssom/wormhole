package statement

import (
	"wormhole/server/query/statement/component"
)

type Statement interface {
	GetKind() string
}

type QueryStatement struct {
	SelectComponent component.SelectComponent
	FromComponent   component.FromComponent
	// whereCondition component.whereCondition
	RowLimit  int
	RowOffset int
	Kind      string
}

func (q QueryStatement) GetKind() string {
	return q.Kind
}
