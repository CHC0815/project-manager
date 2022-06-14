package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	ImportDirs []string
}

type ProjectConfig struct {
	Title string
	Desc []string
	Languages []string
	Version string
	Author string
}

func ReadConfig() *Config{
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error reading config:", err)
	}
	return &config
}

func ReadProjectConfig(path string) *ProjectConfig {
	file, err := os.Open(path + "/project.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		// some other error occurred
		// maybe log it
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := ProjectConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error reading config:", err)
	}
	return &config
}