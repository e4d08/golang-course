package http

import (
	"context"
	"fmt"
	_ "gateway/cmd/gateway/docs"
	"gateway/internal/usecase"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

type Server struct {
	server *http.Server
}

func NewServer(port int, service *usecase.GatewayService) Server {
	mux := http.NewServeMux()
	handler := NewHandler(service)

	mux.HandleFunc("GET /repository/get", handler.GetRepository)
	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		httpSwagger.DocExpansion("full"),
		httpSwagger.DeepLinking(true),
	))

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
