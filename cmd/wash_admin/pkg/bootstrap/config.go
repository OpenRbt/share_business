package bootstrap

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTPPort       string `env:"HTTP_PORT" envDefault:"8080"`
	BasePath       string `env:"HTTP_BASE_PATH" envDefault:""`
	AllowedOrigins string `env:"HTTP_ALLOWED_ORIGINS" envDefault:"*"`
	Host           string `env:"HTTP_HOST"  envDefault:""`
	LogLevel       string `env:"LOG_LEVEL" envDefault:""`
	DB             DBConfig
	FirebaseConfig FirebaseConfig
	RabbitMQConfig RabbitMQConfig
}

type DBConfig struct {
	Host     string `env:"DB_HOST" envDefault:"wash_db"`
	Port     string `env:"DB_PORT" envDefault:"5432"`
	Database string `env:"DB_DATABASE" envDefault:"wash_admin"`
	User     string `env:"DB_USER" envDefault:"wash_admin"`
	Password string `env:"DB_PASSWORD" envDefault:"wash_admin"`
}

type FirebaseConfig struct {
	FirebaseKeyFilePath string `env:"FB_KEYFILE_PATH" envDefault:"/app/firebase/fb_key.json"`
}

type RabbitMQConfig struct {
	Port      string `env:"RABBIT_SERVICE_PORT" envDefault:"5671"`
	Url       string `env:"RABBIT_SERVICE" envDefault:"wash_rabbit"`
	CertsPath string `env:"RABBIT_CERTS_PATH" envDefault:"/app/certs/"`
	User      string `env:"RABBIT_SERVICE_USER" envDefault:"wash_bonus_svc"`
	Password  string `env:"RABBIT_SERVICE_PASSWORD" envDefault:"wash_bonus_svc"`
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
