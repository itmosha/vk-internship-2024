package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		Env string `env:"ENV"`
		HTTPServer
		DB
	}
	HTTPServer struct {
		RunPort     string `env:"RUN_PORT"`
		Timeout     time.Duration
		IdleTimeout time.Duration
	}
	DB struct {
		Address  string `env:"POSTGRES_ADDRESS"`
		User     string `env:"POSTGRES_USER"`
		Name     string `env:"POSTGRES_NAME"`
		Password string `env:"POSTGRES_PASSWORD"`
	}
)

// Create new app config.
func NewConfig() (cfg *Config) {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("could not load .env file: %s\n", err)
	}
	cfg = &Config{}

	cfg.Env = readEnvVar("ENV")
	cfg.HTTPServer.RunPort = readEnvVar("RUN_PORT")
	cfg.HTTPServer.Timeout = time.Second * 5
	cfg.HTTPServer.IdleTimeout = time.Second * 60
	cfg.DB.Address = readEnvVar("POSTGRES_ADDRESS")
	cfg.DB.User = readEnvVar("POSTGRES_USER")
	cfg.DB.Name = readEnvVar("POSTGRES_NAME")
	cfg.DB.Password = readEnvVar("POSTGRES_PASSWORD")
	return
}

// Read .env variable with provided name.
func readEnvVar(name string) (value string) {
	value = os.Getenv(name)
	if value == "" {
		log.Fatalf("env variable %s not found\n", name)
	}
	return
}
