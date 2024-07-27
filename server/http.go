package server

import (
	"fmt"
	"net/http"
)

type HttpServer struct {
	Hostname string
	Port     string

	server *http.Server

	handler *Handler
}

type Option func(*HttpServer)

func NewHttpServer(hostname string, port string, handler *Handler, opts ...Option) *HttpServer {

	me := &HttpServer{
		Hostname: hostname,
		Port:     port,
		handler:  handler,
	}

	for _, opt := range opts {
		opt(me)
	}

	return me
}

func (s *HttpServer) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.handler.GenerateReport)

	srv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.Hostname, s.Port),
		Handler: mux,
	}

	s.server = &srv

	return s.server.ListenAndServe()

}
