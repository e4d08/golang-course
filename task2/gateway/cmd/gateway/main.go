package main

import (
	"context"
	adapter "gateway/internal/adapter/grpc"
	"gateway/internal/handler/http"
	"gateway/internal/usecase"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
)

//	@title			Gateway API
//	@version		1.0
//	@description	Gateway for the Collector service which can fetch information about GitHub repositories.

//	@license.url	https://unlicense.org

//	@host		localhost:8080
//	@BasePath	/

// @externalDocs.description	OpenAPI Specification
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Failed to read PORT from .env")
	}

	adapter, err := adapter.NewGrpcAdapter(os.Getenv("GRPC_ADDRESS"))
	if err != nil {
		log.Fatalf("Failed to init grpc adapter: %v", err)
	}
	defer adapter.Close()

	gateway := usecase.NewGatewayService(adapter)
	handler := http.NewServer(port, gateway)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("Starting server on port %d...", port)
	handler.Run(ctx)

	<-ctx.Done()
	log.Print("Shutting down...")
	handler.Shutdown()
}
