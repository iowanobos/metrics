package config

import (
	"os"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Configuration struct {
	Currencies []ConfigurationCurrency `yaml:"currencies"`
}

type ConfigurationCurrency struct {
	Name  string  `yaml:"name"`
	Value float64 `yaml:"value"`
}

var cfg Configuration

func init() {
	log.Debug("Initialize configuration")
	file, err := os.Open("config/metrics.yml")
	if err != nil {
		log.WithError(err).Fatal("Read configuration failed")
	}

	if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
		log.WithError(err).Fatal("Decode configuration failed")
	}
}

func Get() Configuration {
	return cfg
}
