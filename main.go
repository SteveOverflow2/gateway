package main

import (
	"Gateway/pkg/config"
	"Gateway/pkg/http/rest"
	"Gateway/pkg/rabbitmq"
	"errors"
	"fmt"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	cfg := config.NewConfig()

	err := cfg.LoadConfig()
	if err != nil {
		return errors.New(err.Error())
	}

	server := rest.NewServer(cfg, cfg.Environment)
	fmt.Printf("%v\n", "Oh wow the service is online")
	rabbitmq.StartServer(cfg.Rabbit)
	server.Init()

	server.Run(cfg.Name)
	return nil
}
