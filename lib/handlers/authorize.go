package handlers

import (
	"github.com/AgrafeModel/AuthProviderGO/config"
	"github.com/AgrafeModel/AuthProviderGO/lib/handlers/databasemanager"
	"github.com/gin-gonic/gin"
)

func HandleAuthorize(db *databasemanager.DBManager, ctx *gin.Context) error {

	var data = make(map[string]string)
	for _, field := range config.Conf.Users.Fields {
		if field.Params["identifier"] == "true" || field.Name == "password" {
			data[field.Name] = ctx.PostForm(field.Name)
		}
	}

	//TODO: Authorize user with data

	return nil
}
