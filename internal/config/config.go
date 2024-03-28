package config

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env        string        `yaml:"env" env-required:"true"`
	SessionTTL time.Duration `yaml:"session_ttl"`
	Server     ServerConfig  `yaml:"server"`
}

type ServerConfig struct {
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}
	path := parseConfigPath()
	if path == "" {
		panic("config path not specified")
	}
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("config file does not exist")
	}
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("error while reading config: " + err.Error())
	}
	return &cfg
}

func parseConfigPath() string {
	var path string
	flag.StringVar(&path, "config_path", "", "path to config file")
	flag.Parse()
	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}
