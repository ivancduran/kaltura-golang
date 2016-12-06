package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Url    string
	Token  string
	Email  string
	Id     string
	Expire string
	Format string
}

func New() Config {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}
	return config
}
