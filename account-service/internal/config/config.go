package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	GRPC GRPC `yaml:"grpc"`
}

type GRPC struct {
	Address string `yaml:"address" validate:"nonzero"`
}

func New() (*Config, error) {
	k := koanf.NewWithConf(koanf.Conf{
		Delim: ".",
	})

	file := file.Provider("files/config.yaml")
	if err := k.Load(file, yaml.Parser()); err != nil {
		log.Fatal().Err(err).Msg("failed to config load config")
	}

	var cfg *Config
	if err := k.UnmarshalWithConf("", &cfg, koanf.UnmarshalConf{}); err != nil {
		log.Fatal().Err(err).Msg("failed to config unmarshal config")
	}

	return cfg, nil
}
