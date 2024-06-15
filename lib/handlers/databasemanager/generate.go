package databasemanager

import (
	"github.com/AgrafeModel/AuthProviderGO/config"
)

func ConfigToSQL(conf *config.Config) string {
	sql := ConfigToUserSQL(conf)
	return sql
}

func ConfigToUserSQL(conf *config.Config) string {
	sql := `CREATE TABLE users (
id INT PRIMARY KEY AUTO_INCREMENT,
`

	for _, field := range conf.Users.Fields {
		sql += field.Name + " " + getSQLTypes(field.Type, field.Params) + ",\n"
	}

	sql += `created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);`
	return sql
}
