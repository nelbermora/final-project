package config

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// este package es solo un emulador del paquete config de nuestro template oficial
// de esta forma el proyecto queda lo mas parecido posible a como se implementara en prod

type Configuration interface {
	GetString(key string, defaultValue string) string
}

type configuration struct {
	source AppConfigProperties
}

type AppConfigProperties map[string]string

func (c *configuration) GetString(key string, defaultValue string) string {
	if c.source[key] != "" {
		return c.source[key]
	}
	return defaultValue
}

func ResolveConfiguration() Configuration {
	config := AppConfigProperties{
		"port":     "8888",
		"password": "abc123",
		"ip":       "127.0.0.1",
	}
	file, err := os.Open("../application.yml")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}
	return &configuration{
		source: config,
	}
}
