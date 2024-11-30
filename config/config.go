package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)


type Config struct {
    Collectors []CollectorConfig `yaml:"collectors"`
    Senders []SendersConfig `yaml:"senders"`
}

type CollectorConfig struct {
    Type string `yaml:"type"`
    Title string `yaml:"title"`
    Category string `yaml:"category"`
    ScrapeDuration time.Duration `yaml:"scrapeDuration"`
    Link string `yaml:"link"`
}

type SendersConfig struct {
    Type string `yaml:"type"`
    Token string `yaml:"token"`
    Proxy bool `yaml:"proxy"`
}

func ReadConfig(configPath string) (*Config, error) {
    f, err := os.Open(configPath)
    if err != nil {
        return nil, err
    }
    
    var c Config
    decoder := yaml.NewDecoder(f)
    err = decoder.Decode(&c)
    if err != nil {
        return nil, err
    }

    return &c, nil
}
