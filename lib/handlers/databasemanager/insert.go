package databasemanager

import (
	"database/sql"
	"log"

	"github.com/AgrafeModel/AuthProviderGO/config"
)

type DBManager struct {
	DB *sql.DB
}

func InsertUserSQL(conf *config.Config) string {
	sql := "INSERT INTO users ("
	for i, field := range conf.Users.Fields {
		sql += field.Name
		if i < len(conf.Users.Fields)-1 {
			sql += ", "
		}
	}
	sql += ") VALUES ("
	for i := range conf.Users.Fields {
		sql += "?"
		if i < len(conf.Users.Fields)-1 {
			sql += ", "
		}
	}
	sql += ");"
	log.Println(sql)
	return sql
}

func (db *DBManager) InsertUserQuery(data *map[string]string, conf *config.Config) error {
	sql := InsertUserSQL(conf)
	stmt, err := db.DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()

	args := make([]interface{}, 0)
	for _, field := range conf.Users.Fields {
		args = append(args, (*data)[field.Name])
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil

}
