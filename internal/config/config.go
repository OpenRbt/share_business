package config

import (
	"errors"
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		HTTPPort        string `env:"HTTP_PORT" envDefault:"8080"`
		BasePath        string `env:"HTTP_BASE_PATH" envDefault:""`
		AllowedOrigins  string `env:"HTTP_ALLOWED_ORIGINS" envDefault:"*"`
		Host            string `env:"HTTP_HOST"  envDefault:""`
		LogLevel        string `env:"LOG_LEVEL" envDefault:""`
		DB              DBConfig
		FirebaseConfig  FirebaseConfig
		RabbitMQConfig  RabbitMQConfig
		SchedulerConfig SchedulerConfig
		SessionsConfig  SessionsConfig
	}

	DBConfig struct {
		Host     string `env:"DB_HOST" envDefault:"wash_db"`
		Port     string `env:"DB_PORT" envDefault:"5432"`
		Database string `env:"DB_DATABASE" envDefault:"wash_bonus"`
		User     string `env:"DB_USER" envDefault:"wash_bonus"`
		Password string `env:"DB_PASSWORD" envDefault:"wash_bonus"`
	}

	FirebaseConfig struct {
		FirebaseKeyFilePath string `env:"FB_KEYFILE_PATH" envDefault:"/app/firebase/fb_key.json"`
	}

	SessionsConfig struct {
		ReportsProcessingDelayInMinutes int64 `env:"MONEY_REPORTS_PROCESSING_DELAY_MINUTES" envDefault:"14400"`
		MoneyReportRewardPercentDefault int64 `env:"MONEY_REPORT_REWARD_PERCENT_DEFAULT" envDefault:"5"`
	}

	RabbitMQConfig struct {
		Port     string `env:"RABBIT_SERVICE_PORT" envDefault:"5672"`
		Url      string `env:"RABBIT_SERVICE" envDefault:"wash_rabbit"`
		User     string `env:"RABBIT_SERVICE_USER" envDefault:"wash_bonus_svc"`
		Password string `env:"RABBIT_SERVICE_PASSWORD" envDefault:"wash_bonus_svc"`
	}

	SchedulerConfig struct {
		DelayMinutes int `env:"SCHEDULER_DELAY_MINUTES" envDefault:"1"`
	}
)

func NewConfig(configFiles ...string) (*Config, error) {
	var c Config
	err := godotenv.Load(configFiles...)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}
	}

	err = env.Parse(&c, env.Options{RequiredIfNoDef: true})
	if err != nil {
		return nil, err
	}

	return &c, checkConfig(&c)
}

func checkConfig(c *Config) (err error) {
	if c.SessionsConfig.MoneyReportRewardPercentDefault <= 0 {
		err = errors.New("bad money reports reward value")
	}

	if c.SessionsConfig.ReportsProcessingDelayInMinutes < 0 {
		err = errors.New("bad money report processing delay value")
	}

	return
}
