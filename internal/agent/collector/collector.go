package collector

import "time"

type Metrics struct {
    CPU       CPU       `json:"cpu"`
    Memory    Memory    `json:"memory"`
    Disk      Disk      `json:"disk"`
    Timestamp time.Time `json:"timestamp"`
}

type CPU struct {
    Valid bool 	  `json:"-"`
    Usage float64 `json:"usage_percent"`
}

type Memory struct {
    Total uint64  `json:"total_bytes"`
    Used  uint64  `json:"used_bytes"`
    Usage float64 `json:"usage_percent"`
}

type Disk struct {
    Total uint64  `json:"total_bytes"`
    Used  uint64  `json:"used_bytes"`
    Usage float64 `json:"usage_percent"`
}

type Collector interface {
	Collect() (Metrics, error)
}

