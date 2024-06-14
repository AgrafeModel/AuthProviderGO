package main

import (
	"github.com/AgrafeModel/AuthProviderGO/config"
	"github.com/AgrafeModel/AuthProviderGO/lib/setup/database"
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

	// Generate the SQL for the users
	sql := database.ConfigToSQL(*conf)
	println(sql)
}
