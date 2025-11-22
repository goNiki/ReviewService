package config

import (
	"fmt"
	"os"
	"time"

	"github.com/goNiki/ReviewService/models/errorapp"
	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerConfig *ServerConfig `yaml:"server"`
	DBConfig     *DBConfig     `yaml:"database"`
}

type ServerConfig struct {
	Env         string        `yaml:"env"`
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type DBConfig struct {
	Host              string        `yaml:"host"`
	Port              string        `yaml:"port"`
	User              string        `yaml:"user"`
	Password          string        `yaml:"password"`
	Name              string        `yaml:"name"`
	SslMode           string        `yaml:"sslmode"`
	MaxConns          int32         `yaml:"maxconns"`
	MinConns          int32         `yaml:"minconns"`
	MaxConnLifeTime   time.Duration `yaml:"maxconnlifetime"`
	MaxConnIdleTime   time.Duration `yaml:"maxconnidletime"`
	HealthCheckPeriod time.Duration `yaml:"healthcheckperiod"`
}

func InitConfig(configPath string) (*Config, error) {

	date, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{}, fmt.Errorf("%w: %v", errorapp.ErrInitConfig, err)
	}
	var cfg Config
	if err := yaml.Unmarshal(date, &cfg); err != nil {
		return &Config{}, fmt.Errorf("%w: %v", errorapp.ErrInitConfig, err)
	}

	return &cfg, nil
}
