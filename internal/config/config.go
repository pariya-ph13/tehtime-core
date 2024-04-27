package config

type Config struct {
	Database   Database   `mapstructure:"database"`
	HttpServer HttpServer `mapstructure:"http_server"`
	Debug      bool       `mapstructure:"debug" default:"${DEBUG | true }"`
}

type Database struct {
	Host     string `mapstructure:"host"  default:"${DATABASE_HOST | }"`
	User     string `mapstructure:"user"  default:"${DATABASE_USER | }"`
	Password string `mapstructure:"password"  default:"${DATABASE_PASSWORD | }"`
	DBName   string `mapstructure:"dbname"  default:"${DATABASE_DBNAME | }"`
	Port     string `mapstructure:"port"  default:"${DATABASE_PORT | }"`
}

type HttpServer struct {
	Address string `mapstructure:"address" default:"${HTTP_ENDPOINT | }"`
}

type Sentry struct {
	Dsn string `mapstructure:"dsn"`
}

type SMap struct {
	Url     string `mapstructure:"url"`
	Token   string `mapstructure:"token"`
	Timeout int    `mapstructure:"timeout"`
}

func (c Config) GetConfig() Config {
	return c
}
