package interpreter

import (
	"wormhole/server/query/statement"
)

func (i Interpreter) ExecuteUse(stmt statement.UseStatement) []byte {
	return []byte(stmt.Database)
}
