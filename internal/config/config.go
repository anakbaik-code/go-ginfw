package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	// App
	AppName    string `mapstructure:"APP_NAME"`
	AppEnv     string `mapstructure:"APP_ENV"`
	AppPort    string `mapstructure:"APP_PORT"`
	AppTimeout string `mapstructure:"APP_TIMEOUT"`
	ApiKey     string `mapstructure:"API_KEY"`
	// Database
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBSSLMode  string `mapstructure:"DB_SSL_MODE"`
	// DB Pool
	DBMaxOpenConns    int    `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleConns    int    `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBConnMaxLifetime string `mapstructure:"DB_CONN_MAX_LIFETIME"`
	// Migrate DB
	MigrateDatabaseURL string `mapstructure:"MIGRATE_DATABASE_URL"`
	MigratePath        string `mapstructure:"MIGRATE_PATH"`
	// JWT Config
	JwtSecret          string        `mapstructure:"JWT_SECRET"`
	JwtAccessTokenExp  time.Duration `mapstructure:"JWT_ACCESS_TOKEN_EXP"`
	JwtRefreshTokenExp time.Duration `mapstructure:"JWT_REFRESH_TOKEN_EXP"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}
	var config Config

	// unmarshal
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
