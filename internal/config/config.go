package config

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	PrettyLogging  bool          `env:"PRETTY_LOGGING" envDefault:"true"`
	RequestTimeout time.Duration `env:"REQUEST_TIMEOUT" envDefault:"10s"`
	Address        string        `env:"ADDRESS" envDefault:"127.0.0.1:8080"`
	WriteTimeout   time.Duration `env:"WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout    time.Duration `env:"READ_TIMEOUT" envDefault:"15s"`
}

func New() Config {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
