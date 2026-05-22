package interpreter

import (
	"fmt"
	"wormhole/server/sql/expression"
	"wormhole/server/sql/statement"
	"wormhole/server/storage"
)

func (i Interpreter) ExecuteUpdate(stmt statement.UpdateStatement) []byte {
	engine := storage.NewStorageEngine(i.GetDBabPath(), nil)

	// Read all rows
	cols, err := engine.SelectColumns(i.UsedDB, stmt.TableName, nil)
	if err != nil {
		return []byte(fmt.Sprintf("error: %v", err))
	}

	// Convert to rows
	var rowCount int
	for _, colData := range cols {
		if colData != nil {
			rowCount = len(colData.([]interface{}))
			break
		}
	}

	// Filter and update rows
	for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
		// Build row map
		row := make(expression.Row)
		for colName, colData := range cols {
			if colData != nil {
				arr := colData.([]interface{})
				if rowIdx < len(arr) {
					row[colName] = arr[rowIdx]
				}
			}
		}

		// Check WHERE condition
		if stmt.WhereCondition != nil {
			result, err := stmt.WhereCondition.(expression.Expression).Evaluate(row)
			if err != nil {
				continue
			}
			if result != true {
				continue
			}
		}

		// Apply SET clauses
		for _, setClause := range stmt.SetClauses {
			newVal, err := setClause.Value.(expression.Expression).Evaluate(row)
			if err != nil {
				continue
			}
			row[setClause.Column] = newVal
		}

		// Update in storage
		updateRow := make(map[string]interface{})
		for colName, val := range row {
			updateRow[colName] = val
		}
		if err := engine.UpdateRow(i.UsedDB, stmt.TableName, rowIdx, updateRow); err != nil {
			return []byte(fmt.Sprintf("error updating row: %v", err))
		}
	}

	return []byte(fmt.Sprintf("updated"))
}