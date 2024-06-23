package main

import (
	"fmt"

	"github.com/AgrafeModel/AuthProviderGO/config"
	"github.com/AgrafeModel/AuthProviderGO/lib/handlers/databasemanager"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	conf, err := config.SetupConfig()
	if err != nil {
		panic("Error loading configuration")
	}

	var db = databasemanager.DBManager{}
	err = db.SetupConnection()
	if err != nil {
		panic(fmt.Sprintf("Error setting up database connection: %s", err))
	}

	// Generate the SQL for the tables
	sql := databasemanager.ConfigToSQL(conf)
	res, err := db.DB.Exec(sql)
	if err != nil {
		panic(fmt.Sprintf("Error creating table: %s", err))
	}
	fmt.Println(res)

	println(sql)
}
