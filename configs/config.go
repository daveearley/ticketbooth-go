package configs

import (
	dotenv "github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	AppHost  string `json:"app_host"`
	AppPort  string `json:"app_port"`
	AppDebug bool

	DbName     string `json:"db_name"`
	DbDriver   string `json:"db_driver"`
	DbHost     string `json:"db_host"`
	DbPort     string `json:"db_port"`
	DbUser     string `json:"db_user"`
	DbPassword string `json:"db_password"`

	JwtSecret string `json:"jwt_secret"`
}

// LoadConfig reads config values from the environment and returns a Config struct. If there is a .env in the
// project root this will be loaded also. Environment variables get priority over .env values
func LoadConfig() *Config {

	dotenv.Load()

	return &Config{
		AppDebug:   castBool(getEnv("APP_DEBUG", true)),
		AppPort:    getEnv("APP_PORT", true),
		AppHost:    getEnv("APP_HOST", true),
		DbName:     getEnv("DB_NAME", true),
		DbPassword: getEnv("DB_PASSWORD", true),
		DbHost:     getEnv("DB_HOST", true),
		DbPort:     getEnv("DB_PORT", true),
		DbUser:     getEnv("DB_USER", true),
		JwtSecret:  getEnv("JWT_SECRET", true),
	}
}

func getEnv(envName string, required bool) string {
	env := os.Getenv(envName)

	if env == "" && required {
		log.Fatal(envName + " environment variable must be set")
	}

	return env
}

func castBool(bool string) bool {
	if strings.ToLower(bool) == "true" {
		return true
	}

	return false
}
