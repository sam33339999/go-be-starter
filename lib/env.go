package lib

import (
	"log"

	"github.com/spf13/viper"
)

// Env has environment stored
type Env struct {
	// general
	AppName     string `mapstructure:"APP_NAME"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	Environment string `mapstructure:"ENV"`

	// log
	LogOutput string `mapstructure:"LOG_OUTPUT"`
	LogLevel  string `mapstructure:"LOG_LEVEL"`

	// database
	DBUsername string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASS"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`

	// jwt
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	JWTTtl           int64  `mapstructure:"JWT_TTL_HOURS"`
	RefreshJWTSecret string `mapstructure:"REFRESH_JWT_SECRET"`
	RefreshJWTTtl    int64  `mapstructure:"REFRESH_JWT_TTL_HOURS"`
}

// NewEnv creates a new environment
func NewEnv() Env {

	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("☠️ cannot read configuration")
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("☠️ environment can't be loaded: ", err)
	}

	return env
}
