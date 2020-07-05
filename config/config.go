package config

import (
	"github.com/jinzhu/configor"
)

var Config = struct {
	App struct {
		HTTPAddr string
		HTTPPort string
	}

	DB struct {
		Driver   string
		Host     string
		Port     string `default:"3306"`
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
	}
}{}

func init() {
	configor.Load(&Config, "config.yaml")
}
