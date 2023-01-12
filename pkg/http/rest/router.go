package rest

import (
	"Gateway/pkg/rabbitmq"
	"net/http"
)

func (s *server) routes() {
	// s.Router.PathPrefix("/import/").Handler(http.HandlerFunc(CreateReverseProxy())).Methods(http.MethodPost)
	s.Router.PathPrefix("/{service}").Handler(VerifyJWT(http.HandlerFunc(CreateReverseProxy()))).Methods(http.MethodGet)
	s.Router.PathPrefix("/{service}").Handler(VerifyJWT(http.HandlerFunc(rabbitmq.SendMessage())))
	s.Router.HandleFunc("/", PrintRequest())
}
