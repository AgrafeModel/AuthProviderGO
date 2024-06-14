package main

import (
	"github.com/AgrafeModel/AuthProviderGO/config"
	"github.com/AgrafeModel/AuthProviderGO/lib/setup/templates"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//------- ENVIRONMENT VARIABLES ------- //

	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	_, err = config.SetupConfig()
	if err != nil {
		panic("Error loading configuration")
	}

	r := gin.Default()
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
		component := templates.HTMLPage(config.Conf, templates.Register(config.Conf))
		c.HTML(200, "register", component)
	})
	r.POST("/register", func(c *gin.Context) {}) // From the client

	r.GET("/authorize", func(c *gin.Context) {})
	r.POST("/authorize", func(c *gin.Context) {}) // From the client

	//------- TOKEN ENDPOINTS ------- //

	r.POST("/token", func(c *gin.Context) {})

	r.POST("/revoke", func(c *gin.Context) {})

	//------- USER ENDPOINTS ------- //

	r.GET("/user", func(c *gin.Context) {})

	r.Run()
}
