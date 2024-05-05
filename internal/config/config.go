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
	Env           string              `yaml:"env" env-required:"true"`
	SessionTTL    time.Duration       `yaml:"session_ttl"`
	Main          Main                `yaml:"main"`
	Auth          Auth                `yaml:"auth"`
	Profile       Profile             `yaml:"profile"`
	CSAT          CSAT                `yaml:"csat"`
	Review        Review              `yaml:"review"`
	Minio         MinioConfig         `yaml:"minio"`
	Postgres      PostgresConfig      `yaml:"postgres"`
	Elasticsearch ElasticsearchConfig `yaml:"elasticsearch"`
}

// Main app

type Main struct {
	Server  ServerConfig  `yaml:"server"`
	Clients ClientsConfig `yaml:"clients"`
}

type ServerConfig struct {
	Port        string        `yaml:"port"`
	Timeout     time.Duration `yaml:"timeout"`
	IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type ClientsConfig struct {
	AuthClient    ClientConfig `yaml:"auth"`
	ProfileClient ClientConfig `yaml:"profile"`
	CSATClient    ClientConfig `yaml:"csat"`
	ReviewClient  ClientConfig `yaml:"review"`
}

type ClientConfig struct {
	Addr         string        `yaml:"addr"`
	Retries      int           `yaml:"retries"`
	RetryTimeout time.Duration `yaml:"retry_timeout"`
}

// Auth service

type Auth struct {
	Server GRPCServer  `yaml:"server"`
	Redis  RedisConfig `yaml:"redis"`
}

type GRPCServer struct {
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
}

// Profile service

type Profile struct {
	Server GRPCServer `yaml:"server"`
}

// CSAT service

type CSAT struct {
	Server   GRPCServer     `yaml:"server"`
	Postgres PostgresConfig `yaml:"postgres"`
}

// Review service

type Review struct {
	Server GRPCServer `yaml:"server"`
}

// Data storage

type MinioConfig struct {
	Addr           string `yaml:"addr"`
	MinioUser      string `yaml:"minio_user" env:"MINIO_USER"`
	MinioPassword  string `yaml:"minio_password" env:"MINIO_PASSWORD"`
	MinioAccessKey string `yaml:"minio_access_key" env:"MINIO_ACCESS_KEY"`
	MinioSecretKey string `yaml:"minio_secret_key" env:"MINIO_SECRET_KEY"`
	AvatarBucket   string `yaml:"avatar_bucket"`
}

type PostgresConfig struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user" env:"POSTGRES_USER"`
	Password     string `yaml:"password" env:"POSTGRES_PASSWORD"`
	Database     string `yaml:"database"`
	Sslmode      string `yaml:"sslmode"`
	SearchPath   string `yaml:"search_path"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleTime  int    `yaml:"max_idle_time"`
}

type ElasticsearchConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user" env:"ELASTIC_USER"`
	Password string `yaml:"password" env:"ELASTIC_PASSWORD"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file reading error:", err.Error())
	}
	path := parseConfigPath()
	if path == "" {
		panic("config path not specified")
	}
	if _, err = os.Stat(path); errors.Is(err, os.ErrNotExist) {
		panic("config file does not exist")
	}
	var cfg Config
	if err = cleanenv.ReadConfig(path, &cfg); err != nil {
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
