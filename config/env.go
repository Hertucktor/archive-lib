package config

import (
	"io/ioutil"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v3"
)

type Config struct {
	DBName string `yaml:"dbName"`
	DBPort string `yaml:"dbPort"`
	DBCollectionAllcards string `yaml:"dbCollectionAllcards"`
	DBCollectionMycards string `yaml:"dbCollectionMycards"`
	DBCollectionSetimages string `yaml:"dbCollectionSetimages"`
	DBCollectionSetNames string `yaml:"dbCollectionSetNames"`
}

func GetConfig(configFile string) (Config, error) {
	var c Config

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error().Err(err)
		return c, err
	}

	if err = yaml.Unmarshal(buf, &c); err != nil {
		log.Error().Err(err)
		return c, err
	}
	return c, err
}