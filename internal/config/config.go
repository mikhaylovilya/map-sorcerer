package config

import (
	"log" //nolint:depguard // logger is configured after config init
	"os"
	"path/filepath"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	EnvDebug string = "debug"
	EnvProd  string = "prod"
)

type Config struct {
	Env            string     `yaml:"env"`
	HTTPServer     HTTPServer `yaml:"http_server"`
	PostgresConfig PostgresConfig
}

type HTTPServer struct {
	Address     string        `yaml:"address"`
	Port        uint64        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type PostgresConfig struct {
	Host     string `env:"POSTGRES_HOST"`
	Name     string `env:"POSTGRES_NAME"`
	Port     uint64 `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	URI      string `env:"POSTGRES_URI"`
}

func Parse() *Config {
	configName := os.Getenv("CONFIG_PATH")
	if configName == "" {
		log.Fatal("No CONFIG_PATH found")
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Working directory was not found: %s", err)
	}

	configPath := filepath.Join(pwd, configName)
	if _, err = os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("No config file %s was found", configPath)
	}

	var config Config
	if err = cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatalf("Failed to parse config: %s", err)
	}

	envFile := ".env"
	envPath := filepath.Join(pwd, envFile)
	if _, err = os.Stat(envPath); os.IsNotExist(err) {
		log.Fatalf("No %s file was found", envFile)
	}

	if err = cleanenv.ReadConfig(envPath, &config); err != nil {
		log.Fatalf("Failed to parse env file: %s", err)
	}

	return &config
}
