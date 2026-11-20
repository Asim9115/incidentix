package detector

import (
	"testing"
	"time"

	"github.com/asim9115/incidentix/internal/agent/collector"
)

var fake_metrics = []collector.Metrics{
	{
		CPU:       collector.CPU{Valid: true, Usage: 25.5},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 6_400_000_000, Usage: 40.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 200_000_000_000, Usage: 40.0},
		Timestamp: time.Now().Add(-19 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 32.1},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 7_200_000_000, Usage: 45.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 225_000_000_000, Usage: 45.0},
		Timestamp: time.Now().Add(-18 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 45.8},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 8_000_000_000, Usage: 50.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 250_000_000_000, Usage: 50.0},
		Timestamp: time.Now().Add(-17 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 38.4},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 7_680_000_000, Usage: 48.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 260_000_000_000, Usage: 52.0},
		Timestamp: time.Now().Add(-16 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 55.2},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 9_600_000_000, Usage: 60.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 275_000_000_000, Usage: 55.0},
		Timestamp: time.Now().Add(-15 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 62.3},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 9_920_000_000, Usage: 62.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 300_000_000_000, Usage: 60.0},
		Timestamp: time.Now().Add(-14 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 48.7},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 10_240_000_000, Usage: 64.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 310_000_000_000, Usage: 62.0},
		Timestamp: time.Now().Add(-13 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 70.5},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 10_880_000_000, Usage: 68.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 325_000_000_000, Usage: 65.0},
		Timestamp: time.Now().Add(-12 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 65.2},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 11_200_000_000, Usage: 70.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 340_000_000_000, Usage: 68.0},
		Timestamp: time.Now().Add(-11 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 72.8},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 11_520_000_000, Usage: 72.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 350_000_000_000, Usage: 70.0},
		Timestamp: time.Now().Add(-10 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 58.9},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 11_840_000_000, Usage: 74.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 360_000_000_000, Usage: 72.0},
		Timestamp: time.Now().Add(-9 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 75.4},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 12_160_000_000, Usage: 76.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 375_000_000_000, Usage: 75.0},
		Timestamp: time.Now().Add(-8 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 80.1},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 12_480_000_000, Usage: 78.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 385_000_000_000, Usage: 77.0},
		Timestamp: time.Now().Add(-7 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 68.3},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 12_800_000_000, Usage: 80.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 400_000_000_000, Usage: 80.0},
		Timestamp: time.Now().Add(-6 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 85.6},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 13_120_000_000, Usage: 82.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 410_000_000_000, Usage: 82.0},
		Timestamp: time.Now().Add(-5 * time.Minute),
	},

	// Last 5 records: High resource utilization
	{
		CPU:       collector.CPU{Valid: true, Usage: 91.5},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 14_720_000_000, Usage: 92.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 425_000_000_000, Usage: 85.0},
		Timestamp: time.Now().Add(-4 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 94.2},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 14_720_000_000, Usage: 92.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 425_000_000_000, Usage: 85.0},
		Timestamp: time.Now().Add(-3 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 96.8},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 14_720_000_000, Usage: 92.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 425_000_000_000, Usage: 85.0},
		Timestamp: time.Now().Add(-2 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 98.3},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 14_720_000_000, Usage: 92.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 425_000_000_000, Usage: 85.0},
		Timestamp: time.Now().Add(-1 * time.Minute),
	},
	{
		CPU:       collector.CPU{Valid: true, Usage: 99.5},
		Memory:    collector.Memory{Total: 16_000_000_000, Used: 14_720_000_000, Usage: 92.0},
		Disk:      collector.Disk{Total: 500_000_000_000, Used: 425_000_000_000, Usage: 85.0},
		Timestamp: time.Now(),
	},
}

func TestEngine(t *testing.T) {
	threshold := MetricsThreshold{
		CpuMax:    90,
		MemoryMax: 91,
		DiskMax:   86,
	}
	detector := New(threshold)
	incidents := detector.EvaluateMetrics(fake_metrics)
	for _, i := range incidents {
		t.Logf("incident: type=%s severity=%s metric=%.2f", i.Type, i.Severity, i.Metric)
	}
	
	if !hasType(incidents, CPUHigh) {
		t.Error("expected CPUHigh incident, got none")
	}
	
	if !hasType(incidents, MemoryPressure) {
		t.Error("expected MemoryPressure incident, got none")
	}
	
	if hasType(incidents, DiskFull) {
		t.Error("expected no DiskFull incident, but one was raised")
	}
}

func hasType(incidents []Incident, t Type) bool {
	for _, i := range incidents {
		if i.Type == t {
			return true
		}
	}
	return false
}
