package kafka

import (
	"Gateway/pkg/config"
	"context"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
)

var (
	writer kafka.Writer
)

func Init(cfg config.KafkaConfig) {
	writer = kafka.Writer{
		Addr:                   kafka.TCP(cfg.Host + ":" + cfg.Port),
		AllowAutoTopicCreation: false,
	}
}

func CreateKafkaCall(w http.ResponseWriter, r *http.Request) error {
	writer.Topic = mux.Vars(r)["topic"]
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return writer.WriteMessages(context.Background(), kafka.Message{Value: b, Key: []byte(mux.Vars(r)["method"])})
}
