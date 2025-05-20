package config

import (
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	APPname    string `mapstructure:"APP_NAME"`
	APPport    string `mapstructure:"APP_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
}

func LoadConfig() (*Config, error) {
	// Get the absolute path to the root directory
	absPath, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}

	viper.AddConfigPath(absPath) // Look in root directory
	viper.SetConfigName(".env")  // Exactly match file name
	viper.SetConfigType("env")   // For .env files
	viper.AutomaticEnv()         // Also read from environment variables

	// Set defaults for optional values
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_USER", "postgres")

	// Try to read the config file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No .env file found, using environment variables only")
		} else {
			return nil, err
		}
	}

	// Validate required variables
	required := []string{"DB_HOST", "DB_NAME", "DB_USER", "JWT_SECRET"}
	for _, key := range required {
		if viper.GetString(key) == "" {
			log.Fatalf("Required environment variable %s is not set", key)
		}
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
