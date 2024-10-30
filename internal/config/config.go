package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stasdashkevitch/crypto_info/pkg/logger"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Server struct {
		IdleTimeout  int    `yaml:"idle_timeout" env-default:"100"`
		Port         string `yaml:"port" env-default:":8080"`
		ReadTimeout  int    `yaml:"read_timeout" env-default:"5"`
		WriteTimeout int    `yaml:"write_timeout" env-default:"5"`
	} `yaml:"server"`
	DB struct {
		Host     string `yaml:"host" env-default:"localhost"`
		Port     string `yaml:"port" env-default:"5432"`
		User     string `yaml:"user" env-default:"postgres"`
		Password string `yaml:"password" env-defaul:"123456"`
		DBName   string `yaml:"dbname" env-default:"crypto_info"`
		SSLMode  string `yaml:"sslmode" env-default:"disable"`
		Timezone string `yaml:"timezone" env-default:"Europe/Minsk"`
	} `yaml:"db"`
	Cache struct {
		Host     string `yaml:"host" env-default:"localhost"`
		Port     string `yaml:"port" env-default:"6379"`
		Password string `yaml:"password" env-default:"123456"`
		DB       int    `yaml:"db" env-default:"0"`
	} `yaml:"cache"`
	SMTP struct {
		SMTPHost string `yaml:"smtp_host"`
		SMTPPort string `yaml:"smtp_port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"smtp"`
}

var instance *Config
var once sync.Once

func NewConfig() *Config {
	l := logger.GetLogger()
	once.Do(func() {
		l.Info("start parsing config")
		instance = &Config{}

		err := cleanenv.ReadConfig("config.yaml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			l.Error(help)
			l.Fatal(err.Error())
		}
	})

	return instance
}
