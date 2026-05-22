package interpreter

import (
	"fmt"
	"wormhole/server/sql/statement"
	"wormhole/server/storage"
)

func (i Interpreter) ExecuteFetchColumns(stmt statement.QueryStatement) []byte {
	engine := storage.NewStorageEngine(i.GetDBabPath(), nil)
	cols, err := engine.SelectColumns(i.UsedDB, stmt.FromComponent.FromClause, stmt.SelectComponent.ResultColumns)
	if err != nil {
		return []byte(fmt.Sprintf("error: %v", err))
	}

	result := []byte{}
	for _, colName := range stmt.SelectComponent.ResultColumns {
		if colData, ok := cols[colName]; ok {
			result = append(result, []byte(fmt.Sprintf("%s: %v\n", colName, colData))...)
		}
	}
	return result
}