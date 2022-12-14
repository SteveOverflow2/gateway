package rest

import (
	"net/http"
)

func (s *server) routes() {
	// s.Router.PathPrefix("/import/").Handler(http.HandlerFunc(CreateReverseProxy())).Methods(http.MethodPost)
	s.Router.PathPrefix("/{service}").Handler(http.HandlerFunc(CreateReverseProxy())).Methods(http.MethodGet, http.MethodOptions)
	s.Router.HandleFunc("/", PrintRequest())
}
