package database_manager

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
	for _, field := range conf.Users.Fields {
		sql += field.Name + ", "
	}
	sql += "created_at, updated_at) VALUES ("
	for range conf.Users.Fields {
		sql += "?, "
	}
	sql += "?, ?);"
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

	args = append(args, "NOW()", "NOW()")

	_, err = stmt.Exec(args...)
	if err != nil {
		return err
	}
	return nil

}
