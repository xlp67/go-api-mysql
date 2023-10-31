package config

import (
	"github.com/spf13/viper"
)
type config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBName string `mapstructure:"DB_NAME"`
	DBUser string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	ServerPort string `mapstructure:"SERVER_PORT"`
}
type APIConfig struct {
	
}
type DBConfig struct {
	
}

var cfg *config

func LoadConfig(path string) (*config, error) {
	viper.SetConfigName("app")
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {panic(err)}
	err = viper.Unmarshal(&cfg)
	if err != nil {panic(err)}	
	return cfg, nil
}