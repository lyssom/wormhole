package interpreter

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"wormhole/server/query/statement"
)

type Interpreter struct {
	UsedDB string
}

func (i Interpreter) GetDBabPath() (dbpath string) {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	return path.Join(abPath, "../../../data", i.UsedDB)
}

func (i Interpreter) ExecuteCreateDatabase(dbName string) []byte {
	var abPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		abPath = path.Dir(filename)
	}
	os.Mkdir(path.Join(abPath, "../data", dbName), os.ModePerm)
	return []byte("ok")
}

func (i Interpreter) ExecuteCreateTable(table statement.CreateTableStatement) []byte {
	os.Mkdir(path.Join(i.GetDBabPath(), table.TableName), os.ModePerm)

	filePtr, err := os.Create(path.Join(i.GetDBabPath(), table.TableName, "table.json"))
	fmt.Println(err)
	defer filePtr.Close()
	// 创建Json编码器
	encoder := json.NewEncoder(filePtr)
	encoder.Encode(table)
	return []byte("ok")
}
