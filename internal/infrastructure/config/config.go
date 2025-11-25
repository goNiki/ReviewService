package config

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/goNiki/ReviewService/models/errorapp"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerConfig ServerConfig
	DBConfig     DBConfig
}

type ServerConfig struct {
	Env         string        `env:"SERVER_ENV" envDefault:"local"`
	Host        string        `env:"SERVER_HOST" envDefault:"localhost"`
	Port        string        `env:"SERVER_PORT" envDefault:"8080"`
	Timeout     time.Duration `env:"SERVER_TIMEOUT" envDefault:"4s"`
	IdleTimeout time.Duration `env:"SERVER_IDLE_TIMEOUT" envDefault:"60s"`
}

type DBConfig struct {
	Host              string        `env:"DB_HOST" envDefault:"localhost"`
	Port              string        `env:"DB_PORT" envDefault:"5432"`
	User              string        `env:"DB_USER" envDefault:"postgres"`
	Password          string        `env:"DB_PASSWORD" envDefault:"postgres"`
	Name              string        `env:"DB_NAME" envDefault:"postgres"`
	SslMode           string        `env:"DB_SSLMODE" envDefault:"disable"`
	MaxConns          int32         `env:"DB_MAXCONNS" envDefault:"20"`
	MinConns          int32         `env:"DB_MINCONNS" envDefault:"5"`
	MaxConnLifeTime   time.Duration `env:"DB_MAXCONNLIFETIME" envDefault:"30m"`
	MaxConnIdleTime   time.Duration `env:"DB_MAXCONNIDLETIME" envDefault:"5m"`
	HealthCheckPeriod time.Duration `env:"Db_HEALTHCHECKPERIOD" envDefault:"1m"`
}

func InitConfig() (*Config, error) {

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "./configs/.env"
	}

	if _, err := os.Stat(configPath); err == nil {
		if err := godotenv.Load(configPath); err != nil {
			return &Config{}, fmt.Errorf("%w: failed to load .env file from %s: %v", errorapp.ErrInitConfig, configPath, err)
		}
	} else {
		if _, err := os.Stat("./.env"); err == nil {
			_ = godotenv.Load("./.env")
		}
	}

	var server ServerConfig

	if err := env.Parse(&server); err != nil {
		return &Config{}, fmt.Errorf("%w: failed to parse server config: %v", errorapp.ErrInitConfig, err)
	}

	var db DBConfig

	if err := env.Parse(&db); err != nil {
		return &Config{}, fmt.Errorf("%w: failed to parse db config: %v", errorapp.ErrInitConfig, err)
	}

	return &Config{
		ServerConfig: server,
		DBConfig:     db,
	}, nil
}
