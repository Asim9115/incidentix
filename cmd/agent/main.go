package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/asim9115/incidentix/internal/agent/application"
	"github.com/asim9115/incidentix/internal/agent/collector"
	"github.com/asim9115/incidentix/internal/agent/detector"
	"github.com/asim9115/incidentix/internal/agent/transport"
)


func main() {
	c := collector.NewSystemCollector()

	threshold := detector.MetricsThreshold{
		CpuMax: 50,
		MemoryMax: 90,
		DiskMax: 90,
	}
	t := transport.New("http://localhost:8000")
	d := detector.New(threshold)
	agent := application.New(c, d, t)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("agent running...")
	if err := agent.Run(ctx); err != nil {
		log.Fatal(err)
	}
	log.Printf("agent stopped!")
}