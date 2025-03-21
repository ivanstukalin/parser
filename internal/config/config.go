package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var requiredEnvs = []string{
	"SERVER_PORT",
	"DB_NAME",
	"DB_USER",
	"DB_PASSWORD",
	"DB_HOST",
	"DB_PORT",
}

type Config struct {
	ServerPort string     `envconfig:"SERVER_PORT" required:"true"`
	Database   *DBConfig  `envconfig:"DB" required:"true"`
	App        *AppConfig `envconfig:"APP" required:"true"`
}

type DBConfig struct {
	Name     string `envconfig:"DB_NAME" required:"true"`
	User     string `envconfig:"DB_USER" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Port     string `envconfig:"DB_PORT" required:"true"`
	SSLMode  string `envconfig:"SSL_MODE" required:"false"`

	MaxOpenConns    int           `envconfig:"MAX_OPEN_CONNS" default:"10"`
	MaxIdleConns    int           `envconfig:"MAX_IDLE_CONNS" default:"5"`
	ConnMaxLifetime time.Duration `envconfig:"CONN_MAX_LIFETIME" default:"1h"`
}

type AppConfig struct {
	SecretKey string `envconfig:"SECRET_KEY" required:"true"`
	URL       string `envconfig:"URL" required:"true"`
}

func LoadConfig() (*Config, error) {
	err := loadEnvFile()
	if err != nil {
		return nil, err
	}

	err = validateEnvVars()
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config: %w", err)
	}

	return cfg, nil
}

func loadEnvFile() error {
	if needUseLocalEnvFile() {
		err := godotenv.Load(".env")
		if err != nil {
			return err
		}
	}
	return nil
}

func needUseLocalEnvFile() bool {
	for _, arg := range os.Args {
		if arg == "--local-env" {
			return true
		}
	}
	return false
}

func validateEnvVars() error {
	missingEnvs := []string{}
	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			missingEnvs = append(missingEnvs, env)
		}
	}

	if len(missingEnvs) > 0 {
		return fmt.Errorf("missing required environment variables: %v", missingEnvs)
	}
	return nil
}
