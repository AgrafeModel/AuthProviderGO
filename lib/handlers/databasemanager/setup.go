package databasemanager

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func (db *DBManager) SetupConnection() error {
	var err error

	db.DB, err = sql.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}
	db.DB.SetMaxIdleConns(0)

	// Ping connection to check if it's alive
	err = db.DB.Ping()
	if err != nil {
		return err
	}

	return nil
}
