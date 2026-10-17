package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rideflow/internal/config"
	"rideflow/internal/health"
	"rideflow/internal/platform/mongodb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	cfg := config.Load()

	log.Printf("[INFO] Starting RideFlow API in %s environment on port %s", cfg.AppEnv, cfg.ServerPort)

	initCtx, cancelInit := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelInit()

	mongoCfg := mongodb.DefaultConfig(cfg.MongoURI, cfg.MongoDBName)
	mongoClient, err := mongodb.NewClient(initCtx, mongoCfg)
	if err != nil {
		log.Fatalf("[FATAL] Could not connect to MongoDB: %v", err)
	}
	defer func() {
		closeCtx, cancelClose := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelClose()
		if err := mongoClient.Close(closeCtx); err != nil {
			log.Printf("[ERROR] Error closing MongoDB connection: %v", err)
		}
	}()
	log.Printf("[INFO] Connected to MongoDB database '%s' at %s", cfg.MongoDBName, cfg.MongoURI)

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(60 * time.Second))

	healthHandler := health.NewHandler(mongoClient)
	router.Mount("/", healthHandler.Routes())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.ServerPort),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	serverErrors := make(chan error, 1)
	go func() {
		log.Printf("[INFO] HTTP server listening on http://localhost:%s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		log.Fatalf("[FATAL] Server error: %v", err)
	case sig := <-shutdownSignal:
		log.Printf("[INFO] Received signal '%s', initiating graceful shutdown...", sig)

		shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancelShutdown()

		if err := server.Shutdown(shutdownCtx); err != nil {
			log.Printf("[ERROR] Graceful HTTP server shutdown failed: %v", err)
			_ = server.Close()
		}

		log.Println("[INFO] RideFlow API shutdown completed cleanly")
	}
}
