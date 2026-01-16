package main

import (
	"context"
	"fmt"
	"github.com/todo-app/internal/config"
	"github.com/todo-app/internal/db"
	"github.com/todo-app/internal/router"
	"github.com/todo-app/internal/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := config.Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db.Init()
	db.MigrateAll()

	serviceGroup := service.NewServiceGroup()
	routerSetup := router.Setup(serviceGroup)
	port := config.AppCfg.ServerConfig.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: routerSetup.Engine,
	}

	go func() {
		log.Printf("HTTP server started on %d\n", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shutdown signal received...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
	} else {
		log.Println("HTTP server shutdown completed")
	}

	log.Println("Application exited gracefully")
}
