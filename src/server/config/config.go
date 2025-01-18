package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server ServerConfig `json:"server"`
	DB     DBConfig     `json:"database"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
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
