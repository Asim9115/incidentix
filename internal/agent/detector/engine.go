package detector

import (
	"github.com/asim9115/incidentix/internal/agent/collector"
)

type Detector interface {
	EvaluateMetrics(window []collector.Metrics) []Incident
}

type SystemDetector struct {
	threshold MetricsThreshold
}

func New(threshold MetricsThreshold) *SystemDetector{
	return &SystemDetector{threshold: threshold}
}

func (d *SystemDetector)EvaluateMetrics(window []collector.Metrics)[]Incident {
	var incidents []Incident

	if incident := checkCPU(window, d.threshold.CpuMax); incident != nil {
		incidents = append(incidents, *incident)
	}

	if incident := checkMemory(window, d.threshold.MemoryMax); incident != nil {
		incidents = append(incidents, *incident)
	}

	if incident := checkDisk(window, d.threshold.DiskMax); incident != nil {
		incidents = append(incidents, *incident)
	}
	return incidents
}