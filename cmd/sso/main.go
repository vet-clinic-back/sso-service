package main

import (
	"context"
	"flag"

	"github.com/vet-clinic-back/sso-service/internal/handlers"
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/server"
	"github.com/vet-clinic-back/sso-service/internal/service"
	"github.com/vet-clinic-back/sso-service/internal/storage"
)

func main() {
	isLocal := flag.Bool("local", false, "is it local? can make logs pretty")
	idDebug := flag.Bool("debug", false, "is it local? can make logs pretty")
	flag.Parse()

	log := logging.NewLogger(isLocal, idDebug)
	log.Info("logger initialized")

	// TODO init config
	log.Info("initializing config")

	log.Info("initializing storage")
	storage := storage.New(log)

	log.Info("initializing service")
	service := service.New(log, storage)

	log.Info("initializing handler")
	hander := handlers.NewHandler(log, service)

	log.Info("initializing server")
	server := server.NewServer()

	log.Info("starting server")
	server.Run("8080", hander.InitRoutes()) // TODO use port from config

	log.Info("graceful shutdown")
	server.Shutdown(context.Background())
}
