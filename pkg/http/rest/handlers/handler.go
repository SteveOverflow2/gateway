package handlers

import (
	"Gateway/pkg/kafka"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func KafkaRequest() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := kafka.CreateKafkaCall(w, r)
		if err != nil {
			RenderErrorResponse(w, err.Error(), "CreateKafkaCall()", err)
		} else {
			RenderResponse(w, http.StatusOK, fmt.Sprintf("You posted successfuly on: %v", mux.Vars(r)["topic"]))
		}
		return
	}
}
