package detector

import "time"

type Incident struct {
	Type Type 
	Severity string
	DetectedAt time.Time
	Metric float64
}

type Type string

const (
	MemoryPressure Type = "MemoryPressure"
	CPUHigh Type = "CPUHigh"
	DiskFull Type = "DiskFull"
)

type MetricsThreshold struct {
	CpuMax float64
	MemoryMax float64
	DiskMax float64
}
