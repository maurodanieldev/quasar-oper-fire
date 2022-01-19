package config

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/gommon/log"
)

type Config struct {
	ServerHost string `required:"true" split_words:"true"`
	ServerPort int    `required:"true" split_words:"true"`
	Postfix    string `required:"false"`
}

var once sync.Once
var c Config

func Environments() Config {
	once.Do(func() {
		if err := envconfig.Process("", &c); err != nil {
			log.Panicf("Error parsing environment vars %#v", err)
		}
	})

	return c
}
