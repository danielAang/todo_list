package internal

import "github.com/spf13/viper"

type Constants struct {
	Port  string
	Mongo struct {
		Url    string
		DbName string
	}
}

func initViper() (Constants, error) {
	viper.SetConfigName("application.yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		return Constants{}, err
	}
	viper.SetDefault("PORT", "8080")
	var constants Constants
	err := viper.Unmarshal(&constants)
	return constants, err
}
