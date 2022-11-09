package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config struct {
	serverPort int
	logLevel   string
}

func NewConfig() *Config {
	config := Config{}

	config.serverPort = 8080
	config.logLevel = "info"
	config.initFromEnvVars()

	return &config
}

func (c *Config) initFromEnvVars() {
	if v, ok := os.LookupEnv("LOG_LEVEL"); ok {
		c.logLevel = v
	}
	if v, ok := os.LookupEnv("SERVER_PORT"); ok {
		var err error
		c.serverPort, err = strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Cannot parse SERVER_PORT variable %s", v)
		}
	}
}

func (c *Config) String() string {
	return fmt.Sprintf("Server port: %d, Log level: %s", c.serverPort, c.logLevel)
}

func (c *Config) LogLevel() string {
	return c.logLevel
}
func (c *Config) ServerPort() int {
	return c.serverPort
}
