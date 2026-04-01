package http

import (
	"context"
	"fmt"
	"gateway/internal/usecase"
	"net/http"
)

type Server struct {
	server *http.Server
}

func NewServer(port int, service *usecase.GatewayService) Server {
	mux := http.NewServeMux()
	handler := NewHandler(service)

	mux.HandleFunc("GET /repository/get", handler.GetRepository)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	return Server{server: server}
}

func (s Server) Run(ctx context.Context) {
	go func() {
		s.server.ListenAndServe()
	}()
}

func (s Server) Shutdown() {
	s.server.Shutdown(context.Background())
}
