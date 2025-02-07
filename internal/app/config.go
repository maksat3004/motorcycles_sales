package app

import (
	"github.com/spf13/viper"
	"log"
)

// Config структура для хранения конфигурации приложения
type Config struct {
	Port        string `mapstructure:"PORT"`
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	JwtSecret   string `mapstructure:"qwerty"`
}

// LoadConfig загружает конфигурацию из файла или переменных окружения
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("No config file found: %v", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
