package interpreter

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"wormhole/server/engine/parquet"
	"wormhole/server/engine/parquet/file"
	"wormhole/server/engine/parquet/schema"
	"wormhole/server/query/statement"
)

// type RowType struct {
// 	FirstName string `parquet:"first_name"`
// 	LastName  string `parquet:"last_name"`
// }

func getField(col map[string]interface{}) *schema.PrimitiveNode {
	var s *schema.PrimitiveNode
	columnName := col["Name"].(string)
	switch col["DataType"].(string) {
	case "Int":
		s = schema.NewInt32Node(columnName, parquet.Repetitions.Required, -1)
	case "String", "TEXT":
		s = schema.NewByteArrayNode(columnName, parquet.Repetitions.Required, -1)
		// case "create table":
		// case "insert":
	}
	return s
}

func transpose(matrix [][]interface{}) [][]interface{} {
	m := len(matrix)
	n := len(matrix[0])
	transposed := make([][]interface{}, n)
	for i := 0; i < n; i++ {
		transposed[i] = make([]interface{}, m)
		for j := 0; j < m; j++ {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}

func (i Interpreter) ExecuteInsert(stmt statement.InsertStatement) []byte {
	const pageSize = 15

	var (
		props = parquet.NewWriterProperties(parquet.WithDictionaryDefault(false), parquet.WithDataPageSize(pageSize))
	)

	tableMsg := path.Join(i.GetDBabPath(), stmt.TableName, "table.json")
	fmt.Println(tableMsg)

	jsonFile, _ := os.Open(tableMsg)
	defer jsonFile.Close()
	tablemap := make(map[string]interface{})
	tablejson, _ := io.ReadAll(jsonFile)
	json.Unmarshal([]byte(tablejson), &tablemap)

	var fieldList schema.FieldList
	fmt.Println(1111231)
	fmt.Println(tablemap)
	for _, col := range tablemap["Columns"].([]interface{}) {
		res := getField(col.(map[string]interface{}))
		fieldList = append(fieldList, res)
	}

	sc, _ := schema.NewGroupNode("schema", parquet.Repetitions.Required, fieldList, -1)

	sink, _ := os.Create(path.Join(i.GetDBabPath(), stmt.TableName, "1.parquet"))

	writer := file.NewParquetWriter(sink, sc, file.WithWriterProps(props))
	rgWriter := writer.AppendBufferedRowGroup()

	values := transpose(stmt.Values)
	for i, col := range tablemap["Columns"].([]interface{}) {
		cwr, _ := rgWriter.Column(i)
		switch col.(map[string]interface{})["DataType"].(string) {
		case "Int":
			cw := cwr.(*file.Int32ColumnChunkWriter)
			var formatVal []int32
			for _, val := range values[i] {
				intVal, _ := strconv.ParseInt(val.(string), 10, 32)
				formatVal = append(formatVal, int32(intVal))
			}
			cw.WriteBatch(formatVal, nil, nil)
		case "String", "TEXT":
			cw := cwr.(*file.ByteArrayColumnChunkWriter)
			var formatVal []parquet.ByteArray
			for _, val := range values[i] {
				formatVal = append(formatVal, []byte(val.(string)))
			}
			cw.WriteBatch(formatVal, nil, nil)

		}
	}
	rgWriter.Close()
	writer.Close()

	return []byte("ok")
}
