package statement

import (
	"wormhole/server/sql/statement/component"
)

type Statement interface {
	GetKind() string
}

type QueryStatement struct {
	SelectComponent component.SelectComponent
	FromComponent   component.FromComponent
	WhereCondition  component.FilterComponent
	GroupByComponent component.GroupByComponent
	HavingCondition component.HavingComponent
	OrderByComponent component.OrderByComponent
	RowLimit        int
	RowOffset       int
	Kind            string
}

func (q QueryStatement) GetKind() string {
	return q.Kind
}
