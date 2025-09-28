package configs_env

import (
	"csat-servay/pkg/logs"
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

	expAccessStr := os.Getenv("JWT_ACCESS_EXP")
	expAccess, err := strconv.ParseInt(expAccessStr, 10, 64)
	if err != nil {
		logs.Error(err)
		expAccess = 900
	}

	expRefreshStr := os.Getenv("JWT_REFRESH_EXP")
	expRefresh, err := strconv.ParseInt(expRefreshStr, 10, 64)
	if err != nil {
		logs.Error(err)
		expRefresh = 3600
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
		One: OneConfig{
			URL:          os.Getenv("ONE_URL"),
			GrantType:    os.Getenv("ONE_GRANT_TYPE"),
			ClientID:     os.Getenv("ONE_CLIENT_ID"),
			ClientSecret: os.Getenv("ONE_CLIENT_SECRET"),
		},
		Jwt: JwtConfig{
			HMACSecret: os.Getenv("JWT_HMAC_SECRET"),
			AccessExp:  expAccess,
			RefreshExp: expRefresh,
		},
	}

	return environments
}
