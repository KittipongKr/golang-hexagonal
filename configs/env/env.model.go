package configs_env

type EnvConfig struct {
	App    AppConfig
	Fiber  FiberConfig
	Mongo  MongoConfig
	Jwt    JwtConfig
	Tracer TracerConfig
}

type AppConfig struct {
	AppName     string
	Version     string
	Environment string
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

type JwtConfig struct {
	HMACSecret string
	AccessExp  int
	RefreshExp int
}

type TracerConfig struct {
	Host string
	Port string
}
