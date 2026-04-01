package main

import (
	"collector/internal/adapter/http"
	"collector/internal/handler/grpc"
	"collector/internal/usecase"
	"context"
	"log"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env")
	}

	adapter := http.NewHttpAdapter(os.Getenv("GITHUB_API"))
	collector := usecase.NewCollectorService(adapter)

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Failed to read PORT from .env")
	}

	grpcServer := grpc.NewServer(port, collector)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		log.Printf("Starting grpc server on port %d...", port)
		grpcServer.Run(ctx)
	}()

	<-ctx.Done()
	log.Print("Shutting down...")
}
