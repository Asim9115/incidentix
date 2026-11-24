package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/asim9115/incidentix/internal/api/config"
	"github.com/asim9115/incidentix/internal/api/database"
	"github.com/asim9115/incidentix/internal/api/router"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	//Load env
	config, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load env %v", err)
	}
	log.Print(config.POSTGRES_URL)
	//Database
	db := database.New(config.POSTGRES_URL)
	if err := db.Connect(ctx); err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	router := router.New(db)

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	//start the server
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
		log.Print("server running on port :8080")
	}()

	<-ctx.Done()
	stop()
	
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("server forced to shutdown: %v", err)
	}
}