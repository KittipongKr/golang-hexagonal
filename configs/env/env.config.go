package configs_env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var environments *EnvConfig

func ReadEnv(serverZone string) *EnvConfig {
	file := fmt.Sprintf(".env.%s", serverZone)
	if err := godotenv.Load(file); err != nil {
		log.Printf("File not found %s, ", file)
	} else {
		log.Printf("Success to read %s", file)
	}

	environments = &EnvConfig{
		App: AppConfig{
			AppName: os.Getenv("APP_NAME"),
		},
		Fiber: FiberConfig{
			Host:             os.Getenv("FIBER_HOST"),
			Port:             os.Getenv("FIBER_PORT"),
			AllowOrigins:     os.Getenv("FIBER_ALLOW_ORIGINS"),
			AllowHeaders:     os.Getenv("FIBER_ALLOW_HEADERS"),
			AllowMethods:     os.Getenv("FIBER_ALLOW_METHODS"),
			AllowCredentials: os.Getenv("FIBER_ALLOW_CREDENTIALS"),
		},
		Mongo: MongoConfig{
			UserName: os.Getenv("MONGO_USER"),
			Password: os.Getenv("MONGO_PASSWORD"),
			Host:     os.Getenv("MONGO_HOST"),
			Port:     os.Getenv("MONGO_PORT"),
			Database: os.Getenv("MONGO_BASE"),
		},
	}

	return environments
}
