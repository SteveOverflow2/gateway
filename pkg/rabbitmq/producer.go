package rabbitmq

import (
	"Gateway/pkg/config"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ch *amqp.Channel
	q  amqp.Queue
)

func StartServer(cfg config.RabbitMQ) {
	fmt.Println("Starting rabbitmq")
	fmt.Println(cfg.Host + ":" + cfg.Port)
	conn, err := amqp.Dial("amqp://guest:guest@" + cfg.Host + ":" + cfg.Port)
	fmt.Println("Done dialing")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	q, err = ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
}
func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func SendMessage() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		ctx = context.WithValue(ctx, "Subject", "123")
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("the err was: %v\n", err)
		}

		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        body,
				Headers:     amqp.Table{"Subject": "123"},
			})
		failOnError(err, "Failed to publish a message")

		log.Printf(" [x] Sent %s\n", body)

	}
}
