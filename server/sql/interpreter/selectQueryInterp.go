package interpreter

import (
	"fmt"
	"sort"
	"wormhole/server/sql/expression"
	"wormhole/server/sql/statement"
	"wormhole/server/sql/statement/component"
	"wormhole/server/storage"
)

func (i Interpreter) ExecuteFetchColumns(stmt statement.QueryStatement) []byte {
	engine := storage.NewStorageEngine(i.GetDBabPath(), nil)

	// Get all tables involved in the query
	tables := stmt.FromComponent.Tables
	if len(tables) == 0 {
		// Simple case: single table
		if stmt.FromComponent.FromClause != "" {
			tables = []component.TableRef{{Name: stmt.FromComponent.FromClause}}
		}
	}

	var allRows []expression.Row

	// Read data from first table
	if len(tables) > 0 && tables[0].Name != "" {
		cols, err := engine.SelectColumns(i.UsedDB, tables[0].Name, nil)
		if err != nil {
			return []byte(fmt.Sprintf("error: %v", err))
		}
		allRows, err = i.columnsToRows(cols)
		if err != nil {
			return []byte(fmt.Sprintf("error: %v", err))
		}

		// Execute joins if needed
		if len(tables) > 1 {
			for idx := 1; idx < len(tables); idx++ {
				if tables[idx].Join == nil {
					continue
				}
				rightCols, err := engine.SelectColumns(i.UsedDB, tables[idx].Name, nil)
				if err != nil {
					return []byte(fmt.Sprintf("error: %v", err))
				}
				rightRows, err := i.columnsToRows(rightCols)
				if err != nil {
					return []byte(fmt.Sprintf("error: %v", err))
				}
				joinSpec := &component.JoinSpec{
					Type:      tables[idx].Join.Type,
					Table:     tables[idx].Name,
					Condition: tables[idx].Join.Condition,
				}
				allRows = i.performJoin(allRows, rightRows, joinSpec)
			}
		}
	}

	// Get columns to select
	resultColumns := stmt.SelectComponent.ResultColumns

	// If no data, return empty
	if len(allRows) == 0 {
		return formatResults(make(map[string]interface{}), resultColumns)
	}

	// 1. Apply WHERE filtering
	if stmt.WhereCondition.Condition != nil {
		allRows = i.filterRows(allRows, stmt.WhereCondition.Condition.(expression.Expression))
	}

	// 2. Apply ORDER BY
	if len(stmt.OrderByComponent.Expressions) > 0 {
		i.sortRows(allRows, stmt.OrderByComponent.Expressions)
	}

	// 3. Apply OFFSET
	if stmt.RowOffset > 0 && stmt.RowOffset < len(allRows) {
		allRows = allRows[stmt.RowOffset:]
	}

	// 4. Apply LIMIT
	if stmt.RowLimit > 0 && stmt.RowLimit < len(allRows) {
		allRows = allRows[:stmt.RowLimit]
	}

	// 5. Convert to columnar format for formatting
	resultCols := make(map[string]interface{})
	for _, colName := range resultColumns {
		resultCols[colName] = []interface{}{}
	}
	for _, row := range allRows {
		for colName, val := range row {
			if _, ok := resultCols[colName]; ok {
				resultCols[colName] = append(resultCols[colName].([]interface{}), val)
			}
		}
	}

	return formatResults(resultCols, resultColumns)
}

func (i Interpreter) columnsToRows(cols map[string]interface{}) ([]expression.Row, error) {
	var rowCount int
	for _, colData := range cols {
		if colData != nil {
			rowCount = len(colData.([]interface{}))
			break
		}
	}

	rows := make([]expression.Row, rowCount)
	for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
		row := make(expression.Row)
		for colName, colData := range cols {
			if colData != nil {
				arr := colData.([]interface{})
				if rowIdx < len(arr) {
					row[colName] = arr[rowIdx]
				}
			}
		}
		rows[rowIdx] = row
	}
	return rows, nil
}

func (i Interpreter) performJoin(left, right []expression.Row, joinSpec *component.JoinSpec) []expression.Row {
	var result []expression.Row

	for _, lRow := range left {
		matched := false
		for _, rRow := range right {
			// Merge rows
			merged := make(expression.Row)
			for k, v := range lRow {
				merged[k] = v
			}
			for k, v := range rRow {
				merged[k] = v
			}

			// Evaluate join condition
			if joinSpec.Condition != nil {
				condResult, err := joinSpec.Condition.(expression.Expression).Evaluate(merged)
				if err != nil {
					continue
				}
				if condResult != true {
					continue
				}
			}

			matched = true
			result = append(result, merged)
		}

		// LEFT JOIN: if no match, include left row with NULLs for right
		if !matched && joinSpec.Type == "LEFT" {
			merged := make(expression.Row)
			for k, v := range lRow {
				merged[k] = v
			}
			// Right side columns are nil
			for k := range right[0] {
				merged[k] = nil
			}
			result = append(result, merged)
		}
	}

	return result
}

func (i Interpreter) filterRows(rows []expression.Row, filter expression.Expression) []expression.Row {
	var result []expression.Row
	for _, row := range rows {
		evalResult, err := filter.Evaluate(row)
		if err != nil {
			continue
		}
		if evalResult == true {
			result = append(result, row)
		}
	}
	return result
}

func (i Interpreter) sortRows(rows []expression.Row, orderBys []component.OrderByExpression) {
	sort.SliceStable(rows, func(i, j int) bool {
		for _, ob := range orderBys {
			cmp := expression.Compare(rows[i][ob.Column], rows[j][ob.Column])
			if cmp == 0 {
				continue
			}
			if ob.Direction == "DESC" {
				return cmp > 0
			}
			return cmp < 0
		}
		return false
	})
}

func formatResults(cols map[string]interface{}, columns []string) []byte {
	result := []byte{}
	for _, colName := range columns {
		if colData, ok := cols[colName]; ok {
			result = append(result, []byte(fmt.Sprintf("%s: %v\n", colName, colData))...)
		}
	}
	return result
}