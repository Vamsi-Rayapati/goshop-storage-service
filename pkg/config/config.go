package config

import (
	"github.com/spf13/viper"
)

type ApplicationConfig struct {
	Port      int
	AWSApiKey string
	AWSSecret string
	AWSRegion string
}

var Config *ApplicationConfig

func LoadConfig() *ApplicationConfig {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	viper.ReadInConfig()

	Config = &ApplicationConfig{
		Port:      viper.GetInt("PORT"),
		AWSApiKey: viper.GetString("AWS_API_KEY"),
		AWSSecret: viper.GetString("AWS_SECRET"),
		AWSRegion: viper.GetString("AWS_REGION"),
	}

	return Config

}
