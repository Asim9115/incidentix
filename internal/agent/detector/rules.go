package detector

import (
	"time"

	"github.com/asim9115/incidentix/internal/agent/collector"
)

func checkCPU(window []collector.Metrics, maxCPU float64) *Incident{
	if len(window) < 5 {
		return nil
	}

	last5 := window[len(window)-5:]
	for _, m := range last5{
		if m.CPU.Usage < maxCPU{
			return nil
		}
	}

	return &Incident{
		Type: CPUHigh,
		Severity: "High",
		DetectedAt: time.Now(),
		Metric: last5[0].CPU.Usage,
	}
}

func checkMemory(window []collector.Metrics, maxMemory float64) *Incident {
	if len(window) < 5 {
		return nil
	}

	last5 := window[len(window)-5:]
	for _, m := range last5{
		if m.Memory.Usage < maxMemory{
			return nil
		}
	}

	return &Incident{
		Type: MemoryPressure,
		Severity: "High",
		DetectedAt: time.Now(),
		Metric: last5[0].Memory.Usage,
	}
}

func checkDisk(window []collector.Metrics, maxDisk float64) *Incident {
	if len(window) < 5 {
		return nil
	}

	last5 := window[len(window)-5:]
	for _, m := range last5{
		if m.Disk.Usage < maxDisk{
			return nil
		}
	}

	return &Incident{
		Type: DiskFull,
		Severity: "High",
		DetectedAt: time.Now(),
		Metric: last5[0].Disk.Usage,
	}
}