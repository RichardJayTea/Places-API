package dbhelper

import (
	"fmt"
)

func BuildUpdateStatement(table string, data map[string]interface{}, where ...string) (string, []interface{}) {
	var setExpr string
	var params []interface{}

	for i, j := range data {
		setExpr += fmt.Sprintf("%s=?,", i)
		params = append(params, j)
	}
	setExpr = setExpr[:len(setExpr)-1]

	stmt := fmt.Sprintf("UPDATE %s SET %s", table, setExpr)

	if len(where) > 0 {
		stmt += fmt.Sprintf(" %s", where[0])
	}

	return stmt, params
}
