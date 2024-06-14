// THIS IS THE CONFIG HANDLER
package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Conf is a global variable that holds the configuration for the application.
var Conf *Config

// SetupConfig sets up the configuration for the application.
// It reads the environment variable and sets up the configuration file.
// It returns the configuration file and an error if there is one.
func SetupConfig() (*Config, error) {
	vConf := viper.New()

	vConf.SetEnvPrefix("app")

	vConf.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "_", "__")) // APP_DATA__BASE_PASS -> app.data_base.pass
	vConf.AutomaticEnv()

	vConf.SetConfigType("yaml")
	vConf.AddConfigPath("./config/edit/")
	if err := vConf.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	if err := vConf.Unmarshal(&Conf); err != nil {
		log.Fatalf("Error unmarshalling config file, %s", err)
	}

	// Validate the required fields for user
	validateConfig()

	return Conf, nil
}
