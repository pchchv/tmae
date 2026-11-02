package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"

	"github.com/pchchv/env"
	"github.com/pchchv/tmae/internal/application"
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
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

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

	if err := run(ctx, cancel, cfg); err != nil {
		log.Printf("error run: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, cancel context.CancelFunc, cfg config.Config) error {
	ln, errLn := net.Listen("tcp", cfg.ListenAddress)
	if errLn != nil {
		return fmt.Errorf("failed to listen: %w", errLn)
	}
	defer ln.Close()

	app, errApp := application.New(cfg.TgBotToken)
	if errApp != nil {
		return fmt.Errorf("failed to create application: %w", errApp)
	}

	log.Printf("start")

	var wg sync.WaitGroup
	wg.Add(1)
	go app.Run(ctx, cancel, &wg, ln)

	<-ctx.Done()

	wg.Wait()

	log.Printf("done")

	return nil
}
