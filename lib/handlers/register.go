package handlers

import (
	"fmt"
	"strconv"

	"github.com/AgrafeModel/AuthProviderGO/config"
	"github.com/AgrafeModel/AuthProviderGO/lib/handlers/database_manager"
	"github.com/gin-gonic/gin"
)

func HandleRegister(db *database_manager.DBManager, ctx *gin.Context) error {

	var data = make(map[string]string)
	for _, field := range config.Conf.Users.Fields {
		val := ctx.PostForm(field.Name)
		if val == "" {

			return fmt.Errorf(field.Name + " is required")
		}

		//-- Size validation
		if field.Params["min"] != nil {
			min := field.Params["min"].(int)
			if len(val) < min {
				return fmt.Errorf(field.Name + " must be at least " + strconv.Itoa(min) + " characters")
			}
		}
		if field.Params["max"] != nil {
			max := field.Params["max"].(int)
			if len(val) > max {
				return fmt.Errorf(field.Name + " must be at most " + strconv.Itoa(max) + " characters")
			}
		}

		//-- Confirm validation
		if field.Params["confirm"] != nil && field.Params["confirm"].(bool) {
			val_confirm := ctx.PostForm(field.Name + "_confirm")
			if val_confirm == "" {
				return fmt.Errorf(field.Name + "_confirm is required")
			}
			if val != val_confirm {
				return fmt.Errorf(field.Name + " and " + field.Name + "_confirm must match")
			}
		}

		data[field.Name] = val
	}

	// Insert user
	err := db.InsertUserQuery(&data, config.Conf)
	if err != nil {
		return err
	}
	return nil
}
