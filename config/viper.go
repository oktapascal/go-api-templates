package config

import "github.com/spf13/viper"

func InitConfig() {
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")

	err := viper.ReadInConfig()
	if err != nil {
		CreateLoggerFile().Fatal(err)
		return
	}
}
