package conf

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

// Config struct
type Config struct {
	Server struct {
		Port   string `default:":3000" envconfig:"SERVER_PORT"`
		Logger *logrus.Logger
	}
	Conn struct {
		Pool *pgxpool.Pool
	}
	Database struct {
		Host     string `default:"127.0.0.1" envconfig:"DB_HOST"`
		Name     string `default:"postgres" envconfig:"DB_NAME"`
		User     string `default:"postgres" envconfig:"DB_USER"`
		Password string `default:"postgres" envconfig:"DB_PASSWORD"`
		Port     string `default:"5434" envconfig:"DB_PORT"`
		Schema   string `default:"v3" envconfig:"DB_SCHEMA"`
		MaxConns int32  `default:"5" envconfig:"DB_MAX_CONNS"`
		Tmpl     string `default:"host=%s port=%d dbname=%s user=%s password=%s search_path=%s"`
	}
}

// ParseConfig ...
func ParseConfig(app string) (config *Config, err error) {
	if err := envconfig.Process(app, &config); err != nil {
		if err := envconfig.Usage(app, &config); err != nil {
			return config, err
		}
		return config, err
	}
	return config, nil
}
