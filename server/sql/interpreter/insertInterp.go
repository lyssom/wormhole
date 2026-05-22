package interpreter

import (
	"fmt"
	"wormhole/server/sql/statement"
	"wormhole/server/storage"
)

func (i Interpreter) ExecuteInsert(stmt statement.InsertStatement) []byte {
	rows := make([]map[string]interface{}, len(stmt.Values))
	for rowIdx, row := range stmt.Values {
		rowMap := make(map[string]interface{})
		for colIdx, colName := range stmt.Columns {
			rowMap[colName] = row[colIdx]
		}
		rows[rowIdx] = rowMap
	}

	engine := storage.NewStorageEngine(i.GetDBabPath(), nil)
	err := engine.Insert(i.UsedDB, stmt.TableName, rows)
	if err != nil {
		return []byte(fmt.Sprintf("error: %v", err))
	}

	return []byte("ok")
}