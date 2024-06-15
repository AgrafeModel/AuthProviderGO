package databasemanager

import (
	"fmt"
	"strconv"
)

func getSQLTypes(fieldType string, params map[string]interface{}) string {
	switch fieldType {
	case "string":
		if max, ok := params["max"]; ok {
			return fmt.Sprintf("VARCHAR(%s)", strconv.Itoa(max.(int)))
		}
		return "VARCHAR(255)"
	case "int":
		return "INT"
	case "bool":
		return "BOOLEAN"
	case "date":
		return "DATE"
	case "datetime":
		return "DATETIME"
	default:
		return "VARCHAR(255)"
	}
}
