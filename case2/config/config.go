package config

import (
	"bytes"
	"io/ioutil"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB   DBCredential `yaml:"db"`
	Rest ServerConfig `yaml:"rest"`
	Grpc ServerConfig `yaml:"grpc"`
	Omdb OmdbConfig   `yaml:"omdb"`
}

type DBCredential struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type OmdbConfig struct {
	ApiKey string `yaml:"apikey"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return LoadConfigFromBytes(data)
}

func LoadConfigFromBytes(data []byte) (*Config, error) {
	v := viper.New()
	v.SetConfigType("yaml")

	if err := v.ReadConfig(bytes.NewBuffer(data)); err != nil {
		return nil, err
	}

	var conf Config
	err := v.Unmarshal(&conf)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	return &conf, nil
}
