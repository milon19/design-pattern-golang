package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

type Config struct {
	settings map[string]string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			settings: make(map[string]string),
		}
		instance.loadConfigFromJson()
	})
	return instance
}

func (c *Config) loadConfigFromJson() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return
	}
	dir := filepath.Dir(filename)
	configFilePath := filepath.Join(dir, "config.json")
	configData, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	err = json.Unmarshal(configData, &c.settings)
	if err != nil {
		fmt.Println("Error unmarshalling config data:", err)
		return
	}
}

func (c *Config) loadConfigFromEnv() {
	var keys []string

	for _, key := range keys {
		if value, exists := os.LookupEnv(key); exists {
			c.settings[key] = value
		}
	}
}

func (c *Config) Get(key string) (string, bool) {
	value, exists := c.settings[key]
	return value, exists
}
