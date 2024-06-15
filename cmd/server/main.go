package main

import (
	"fmt"

	"github.com/AgrafeModel/AuthProviderGO/config"
	"github.com/AgrafeModel/AuthProviderGO/lib/handlers"
	"github.com/AgrafeModel/AuthProviderGO/lib/handlers/databasemanager"
	"github.com/AgrafeModel/AuthProviderGO/lib/templates"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//------- ENVIRONMENT VARIABLES ------- //

	err := godotenv.Load(".env")
	if err != nil {
		panic(fmt.Sprintf("Error loading .env file: %s", err))
	}

	_, err = config.SetupConfig()
	if err != nil {
		panic(fmt.Sprintf("Error loading configuration: %s", err))
	}

	//------- DATABASE SETUP ------- //

	var db = databasemanager.DBManager{}
	err = db.SetupConnection()
	if err != nil {
		panic(fmt.Sprintf("Error setting up database connection: %s", err))
	}

	//------- GIN SETUP ------- //
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, //TODO: Make this configurable
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	ginHtmlRenderer := r.HTMLRender
	r.HTMLRender = &templates.HTMLTemplRenderer{
		FallbackHtmlRenderer: ginHtmlRenderer,
	}
	r.SetTrustedProxies(nil)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//------- CLIENT ENDPOINTS ------- //

	r.GET("/register", func(c *gin.Context) {
		component := templates.HTMLPage(config.Conf, templates.Register(config.Conf, nil))
		c.HTML(200, "register", component)
	})
	r.POST("/register", func(c *gin.Context) {
		err := handlers.HandleRegister(nil, c) //TODO: Implement the database manager
		if err != nil {
			component := templates.HTMLPage(config.Conf, templates.Register(config.Conf, err))
			c.HTML(200, "register", component)
		} else {
			c.Redirect(302, "/login")
		}
	}) // From the client

	r.GET("/authorize", func(c *gin.Context) {})
	r.POST("/authorize", func(c *gin.Context) {}) // From the client

	//------- TOKEN ENDPOINTS ------- //

	r.POST("/token", func(c *gin.Context) {})

	r.POST("/revoke", func(c *gin.Context) {})

	//------- USER ENDPOINTS ------- //

	r.GET("/user", func(c *gin.Context) {})

	r.Run()
}
