package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Connection ConnectionConfig `json:"connection"`
}

type ConnectionConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

func ReadConfig(configPath string) (Config, error) {
	var c Config

	data, err := os.ReadFile(configPath)

	if err != nil {
		return c, err
	}

	return c, json.Unmarshal(data, &c)
}

func MustReadConfig(configPath string) Config {
	c, err := ReadConfig(configPath)

	if err != nil {
		log.Fatal("Failed to read config:", err)
	}

	return c
}
