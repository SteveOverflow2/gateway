package rest

import (
	"Gateway/pkg/config"
	"Gateway/pkg/http/rest/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type server struct {
	environment string
	HTTPcfg     config.HTTPConfig
	URLcfg      config.URLConfig
	Rabbitcfg   config.RabbitMQ
	Server      *http.Server
	Router      *mux.Router
}

const serverLog string = "[Server]: "

func NewServer(cfg *config.Config, env string) *server {
	baseUrl := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	fmt.Printf("cfg.HTTP: %v\n", cfg.HTTP)
	fmt.Printf("cfg.URLS: %v\n", cfg.URLS)
	s := &server{
		environment: env,
		Server: &http.Server{
			Addr:         baseUrl,
			WriteTimeout: cfg.HTTP.WriteTimeOut,
			ReadTimeout:  cfg.HTTP.ReadTimeOut,
			IdleTimeout:  cfg.HTTP.IdleTimeOut,
		},
		Router: mux.NewRouter(),
	}

	log.Println(serverLog+"started api on base url: ", baseUrl)

	InitUrls(cfg.URLS)
	// Generic routes
	s.Router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	s.Server.Handler = s.Router

	return s
}

func (s *server) Init() {
	s.routes()
}

func (s *server) Run(name string) {
	var wait time.Duration

	s.Server.Handler = cors.Default().Handler(s.Router)

	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println(serverLog+name, "is running..")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	s.Server.Shutdown(ctx)

	log.Println(serverLog+name, "is shutting down..")

	os.Exit(0)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("404 - Endpoint was not found")
	handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
}
