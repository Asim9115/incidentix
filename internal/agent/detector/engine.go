package detector

import (
	"github.com/asim9115/incidentix/internal/agent/collector"
)

func EvaluateMetrics(window []collector.Metrics, threshold MetricsThreshold)[]Incident {
	var incidents []Incident

	if incident := checkCPU(window, threshold.CpuMax); incident != nil {
		incidents = append(incidents, *incident)
	}

	if incident := checkMemory(window, threshold.MemoryMax); incident != nil {
		incidents = append(incidents, *incident)
	}

	if incident := checkDisk(window, threshold.DiskMax); incident != nil {
		incidents = append(incidents, *incident)
	}
	return incidents
}