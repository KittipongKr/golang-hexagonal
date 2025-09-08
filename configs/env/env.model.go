package configs_env

type EnvConfig struct {
	App   AppConfig
	Fiber FiberConfig
	Mongo MongoConfig
}

type AppConfig struct {
	AppName string
}

type FiberConfig struct {
	Host             string
	Port             string
	AllowOrigins     string
	AllowHeaders     string
	AllowMethods     string
	AllowCredentials string
}

type MongoConfig struct {
	UserName string
	Password string
	Host     string
	Port     string
	Database string
}
