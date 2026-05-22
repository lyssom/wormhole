package interpreter

import (
	"fmt"
	"wormhole/server/sql/expression"
	"wormhole/server/sql/statement"
	"wormhole/server/storage"
)

func (i Interpreter) ExecuteDelete(stmt statement.DeleteStatement) []byte {
	engine := storage.NewStorageEngine(i.GetDBabPath(), nil)

	// Read all rows
	cols, err := engine.SelectColumns(i.UsedDB, stmt.TableName, nil)
	if err != nil {
		return []byte(fmt.Sprintf("error: %v", err))
	}

	// Determine row count
	var rowCount int
	for _, colData := range cols {
		if colData != nil {
			rowCount = len(colData.([]interface{}))
			break
		}
	}

	// Find rows to delete
	var rowsToDelete []int
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
		} else {
			// No WHERE - delete all rows
		}

		rowsToDelete = append(rowsToDelete, rowIdx)
	}

	// Delete rows (in reverse order to maintain indices)
	for _, idx := range rowsToDelete {
		if err := engine.DeleteRow(i.UsedDB, stmt.TableName, idx); err != nil {
			return []byte(fmt.Sprintf("error deleting row: %v", err))
		}
	}

	return []byte(fmt.Sprintf("deleted %d rows", len(rowsToDelete)))
}