package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   Server
	Postgres Postgres
}

type Server struct {
	Port int
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func New(configPath, configName string) (*Config, error) {
	v := viper.New()
	v.AddConfigPath(configPath)
	v.SetConfigName(configName)
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}

	if err := v.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c Config) Print() {
	log.Println(c)
}
