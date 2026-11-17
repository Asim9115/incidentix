package application

import (
    "context"
    "testing"
    "time"

    "github.com/asim9115/incidentix/internal/agent/collector"
)

func TestAgentIntegration(t *testing.T) {
    agent := New(collector.NewSystemCollector())

    ctx, cancel := context.WithTimeout(context.Background(), 17*time.Second)
    defer cancel()

    agent.Run(ctx)

    for _, m := range agent.window.Metrics {
        t.Logf("ts: %v | cpu: %.2f%% | mem: %.2f%% | disk: %.2f%%",
            m.Timestamp,
            m.CPU.Usage,
            m.Memory.Usage,
            m.Disk.Usage,
        )
    }

    if len(agent.window.Metrics) == 0 {
        t.Error("expected at least one real metric sample")
    }
}