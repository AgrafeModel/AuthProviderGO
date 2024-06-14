package database

import (
	"fmt"
	"strconv"

	"github.com/AgrafeModel/AuthProviderGO/config"
)

func ConfigToSQL(conf config.Config) string {
	sql := ConfigToUserSQL(conf.Users.Fields)
	return sql
}

func ConfigToUserSQL(conf []config.Field) string {
	sql := "CREATE TABLE users (\nid INT PRIMARY KEY AUTO_INCREMENT,\n"

	for _, field := range conf {
		sql += field.Name + " " + getSQLTypes(field.Type, field.Params) + ",\n"
	}

	sql += "created_at DATETIME,\nupdated_at DATETIME\n);"
	return sql
}

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
