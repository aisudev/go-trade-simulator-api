package configs

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

func ViperSetup() {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}
