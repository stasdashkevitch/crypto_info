package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/stasdashkevitch/crypto_info/pkg/logger"
)

type Config struct {
	Env    string `yaml:"env" env-default:"local"`
	Server struct {
		BindIP       string `yaml:"bind_ip" env-default:"0.0.0.0"`
		Port         string `yaml:"port" env-default:"8080"`
		ReadTimeout  int    `yaml:"read_timeout" env-default:"5"`
		WriteTimeout int    `yaml:"write_timeout" env-default:"5"`
	} `yaml:"server"`
}

var instance *Config
var once sync.Once

func NewConfig() *Config {
	l := logger.GetLogger()
	once.Do(func() {
		l.Info("start parsing config")
		instance = &Config{}

		err := cleanenv.ReadConfig("../../config.yaml", instance)
		if err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			l.Error(help)
			l.Fatal(err.Error())
		}
	})

	return instance
}
