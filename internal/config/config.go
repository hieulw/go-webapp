package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

const configPath = "configs/app.yml"

type Auth struct {
	ClientID          string `yaml:"client_id"          env:"OIDC_CLIENT_ID"`
	ClientSecret      string `yaml:"client_secret"      env:"OIDC_CLIENT_SECRET"`
	CallbackURL       string `yaml:"callback_url"       env:"OIDC_CALLBACK_URL"`
	DiscoveryEndpoint string `yaml:"discovery_endpoint" env:"OIDC_DISCOVERY_ENDPOINT"`
}
type Database struct {
	Driver string `yaml:"driver" env:"DB_DRIVER"`
	Name   string `yaml:"name"   env:"DB_DATABASE"`
	User   string `yaml:"user"   env:"DB_USERNAME"`
	Pass   string `yaml:"pass"   env:"DB_PASSWORD"`
	Host   string `yaml:"host"   env:"DB_HOST"`
	Port   int    `yaml:"port"   env:"DB_PORT"`
}
type Server struct {
	Host string `yaml:"host" env:"APP_HOST" env-default:"0.0.0.0"`
	Port int    `yaml:"port" env:"APP_PORT"`
}
type Config struct {
	Auth     Auth     `yaml:"auth"`
	Database Database `yaml:"database"`
	Server   Server   `yaml:"server"`
}

func New() Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func LoadConfig() (cfg Config, err error) {
	if err = cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return
	}
	return
}
