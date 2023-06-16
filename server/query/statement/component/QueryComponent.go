package component

type SelectComponent struct {
	ResultColumns    []string
	AliasToColumnMap map[string]string // map from an alias to the column it points to.  used for join operations.  used for query generation.  used
}

type FromComponent struct {
	FromClause string
}
