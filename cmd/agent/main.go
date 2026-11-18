package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/asim9115/incidentix/internal/agent/application"
	"github.com/asim9115/incidentix/internal/agent/collector"
)


func main() {
	agent := application.New(collector.NewSystemCollector())

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("agent running...")
	if err := agent.Run(ctx); err != nil {
		log.Fatal(err)
	}
	log.Printf("agent stopped!")
}