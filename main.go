package main

import (
	"Gateway/pkg/config"
	"Gateway/pkg/http/rest"
	"Gateway/pkg/kafka"
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
	server.Init()

	// go StartKafka()
	kafka.Init(cfg.Kafka)

	server.Run(cfg.Name)
	return nil
}

func StartKafka() {
	// conf := kafka.ReaderConfig{
	// 	Brokers: []string{"145.93.76.209:9092"},
	// 	Topic:   "fromBrand",
	// 	GroupID: "g1",
	// }

	// reader := kafka.NewReader(conf)

	// fmt.Println("Started the reader")
	// for {
	// 	m, err := reader.ReadMessage(context.Background())
	// 	if err != nil {
	// 		fmt.Println("Error", err)
	// 		continue
	// 	}
	// 	fmt.Println("Message:", string(m.Value))
	// }
}
