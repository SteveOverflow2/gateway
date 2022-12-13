package rest

import (
	"net/http"
)

func (s *server) routes() {
	s.Router.PathPrefix("/api/import/").Handler(http.HandlerFunc(CreateReverseProxy())).Methods(http.MethodPost)
	s.Router.PathPrefix("/api/{service}").Handler(http.HandlerFunc(CreateReverseProxy())).Methods(http.MethodGet, http.MethodOptions)
}
