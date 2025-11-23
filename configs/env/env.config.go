package configs_env

import (
	"fmt"
	"log"
	"os"
	"strconv"

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

	jwtAccessExp, err := strconv.Atoi(os.Getenv("JWT_ACCESS_EXP"))
	if err != nil {
		log.Printf("Error parsing JWT_ACCESS_EXP: %v", err)
	}

	jwtRefreshExp, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXP"))
	if err != nil {
		log.Printf("Error parsing JWT_REFRESH_EXP: %v", err)
	}

	environments = &EnvConfig{
		App: AppConfig{
			AppName:     os.Getenv("APP_NAME"),
			Version:     os.Getenv("APP_VERSION"),
			Environment: os.Getenv("APP_ENV"),
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
		Jwt: JwtConfig{
			HMACSecret: os.Getenv("JWT_HMAC_SECRET"),
			AccessExp:  jwtAccessExp,
			RefreshExp: jwtRefreshExp,
		},
		Tracer: TracerConfig{
			Host: os.Getenv("TRACING_HOST"),
			Port: os.Getenv("TRACING_PORT"),
		},
	}

	return environments
}
