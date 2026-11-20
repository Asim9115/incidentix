package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/asim9115/incidentix/internal/agent/application"
	"github.com/asim9115/incidentix/internal/agent/collector"
	"github.com/asim9115/incidentix/internal/agent/detector"
)


func main() {
	c := collector.NewSystemCollector()

	threshold := detector.MetricsThreshold{
		CpuMax: 95,
		MemoryMax: 90,
		DiskMax: 90,
	}
	d := detector.New(threshold)
	agent := application.New(c, d)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	log.Printf("agent running...")
	if err := agent.Run(ctx); err != nil {
		log.Fatal(err)
	}
	log.Printf("agent stopped!")
}