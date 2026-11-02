package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pchchv/env"
	"github.com/pchchv/tmae/internal/config"
)

func init() {
	// Load values from .env into the system.
	if err := env.Load(); err != nil {
		log.Panic("No .env file found")
	}
}

func getEnvValue(v string) (string, error) {
	value, exist := os.LookupEnv(v)
	if !exist {
		return value, fmt.Errorf("value %v does not exist", v)
	}

	return value, nil
}

func main() {
	cfg := config.Config{
		ListenAddress: "0.0.0.0:3000",
	}

	if v, err := getEnvValue("LISTEN_ADDRESS"); err != nil || v == "" {
		log.Printf("error run: %v", err)
		os.Exit(1)
	} else {
		cfg.ListenAddress = v
	}

	if v, err := getEnvValue("TG_BOT_TOKEN"); err != nil || v == "" {
		log.Printf("error run: %v", err)
		os.Exit(1)
	} else {
		cfg.TgBotToken = v
	}
}
