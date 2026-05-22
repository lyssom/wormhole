package interpreter

import (
	"wormhole/server/sql/statement"
)

func (i Interpreter) ExecuteUse(stmt statement.UseStatement) []byte {
	return []byte(stmt.Database)
}
