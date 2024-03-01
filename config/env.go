package config

import "github.com/spf13/viper"

type Env struct {
	DBDriver 		string `mapstructure:"DB_DRIVER"`
	DBSource 		string `mapstructure:"DB_SOURCE"`
	ServerAddress 	string `mapstructure:"SERVER_ADDRESS"`
}

// @param path string - path from .env file to the main.go file to load the environment variables to
// if main.go and .env file are in the same directory, use "."
func LoadEnv(path string) (env Env, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}