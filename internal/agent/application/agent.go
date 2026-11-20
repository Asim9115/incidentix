package application

import (
	"context"
	"log"
	"time"

	"github.com/asim9115/incidentix/internal/agent/collector"
	"github.com/asim9115/incidentix/internal/agent/detector"
	"github.com/asim9115/incidentix/internal/agent/transport"
)

// Metrics is a collection of the records in memory that contains cpu, ram and disk
// maxSize contains the maximum size of the window
type MetricsWindow struct {
	Metrics []collector.Metrics
	maxSize int
}

func NewMetricWindow(maxSize int) *MetricsWindow {
	return &MetricsWindow{
		Metrics: make([]collector.Metrics, 0, maxSize),
		maxSize: maxSize,
	}
}

func (m *MetricsWindow) Push(metric collector.Metrics) {
	m.Metrics = append(m.Metrics, metric)
	if len(m.Metrics) > m.maxSize {
		m.Metrics = m.Metrics[1:]
	}
}

//Agent calls collector.collect
type Agent struct {
	collector collector.Collector
	detector  detector.Detector
	window    *MetricsWindow
	transport *transport.Client
}

func New(c collector.Collector, d detector.Detector, t *transport.Client) *Agent {
	return &Agent{
		collector: c,
		detector: d,
		window:    NewMetricWindow(24),
		transport: t,
	}
}

func (a *Agent) Run(ctx context.Context) error {
	ticker := time.NewTicker(2 *time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			metrics, err := a.collector.Collect()
			if err != nil {
				continue

			}
			a.window.Push(metrics)
			if err := a.transport.SendMetrics(ctx, metrics); err != nil {
				log.Printf("failed to send metrics : %v", err)
			}
			incidents := a.detector.EvaluateMetrics(a.window.Metrics)
			for _, inc := range incidents {
   				log.Printf("[INCIDENT] type=%s severity=%s metric=%.2f", inc.Type, inc.Severity, inc.Metric)

			}
		}
	}
}