package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pchchv/env"
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
}
