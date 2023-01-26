package bootstrap

import (
	"errors"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	HTTPPort       string `env:"HTTP_PORT" envDefault:"80"`
	BasePath       string `env:"HTTP_BASE_PATH" envDefault:""`
	AllowedOrigins string `env:"HTTP_ALLOWED_ORIGINS" envDefault:"*"`
	Host           string `env:"HTTP_HOST"  envDefault:""`
	LogLevel       string `env:"LOG_LEVEL"`
	DB             DBConfig
	FirebaseConfig FirebaseConfig
	RabbitMQConfig RabbitMQConfig
}

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	Port     string `env:"DB_PORT"`
	Database string `env:"DB_DATABASE"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
}

type FirebaseConfig struct {
	FirebaseKeyFilePath string `env:"FB_KEYFILE_PATH"`
}

type RabbitMQConfig struct {
	UseTLS    bool   `env:"RABBIT_USE_TLS" envDefault:"false"`
	Port      string `env:"RABBIT_SERVICE_PORT" envDefault:"5672"`
	Url       string `env:"RABBIT_SERVICE" envDefault:"localhost"`
	CertsPath string `env:"RABBIT_CERTS_PATH" `
	User      string `env:"RABBIT_SERVICE_USER" envDefault:"localhost"`
	Password  string `env:"RABBIT_SERVICE_PASSWORD" envDefault:"localhost"`
}

func NewConfig(configFiles ...string) (*Config, error) {
	var c Config
	err := godotenv.Load(configFiles...)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	return &c, env.Parse(&c, env.Options{
		RequiredIfNoDef: true,
	})

}
