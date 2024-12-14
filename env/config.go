package env

import "github.com/caarlos0/env/v11"

// Config contain the values to be collected from the environment
type Config struct {
	Production              bool   `env:"PRODUCTION"`
	DatabaseHost            string `env:"DATABASE_HOST" envDefault:"localhost"`
	DatabaseName            string `env:"DATABASE_NAME" envDefault:"test"`
	DatabaseUser            string `env:"DATABASE_USER" envDefault:"test"`
	DatabasePort            string `env:"DATABASE_PORT" envDefault:"5432"`
	DatabasePassword        string `json:"-" env:"DATABASE_PASSWORD" envDefault:"password"`
	DatabaseMaxConn         int    `env:"DATABASE_MAX_CONN" envDefault:"10"`
	DatabaseConnMaxLifetime int    `env:"DATABASE_CONN_MAX_LIFETIME" envDefault:"300"`
}

// LoadOSEnv parses values from environment variables and populates the struct
func LoadOSEnv() (Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	return cfg, err
}
