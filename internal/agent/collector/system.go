package collector

import "time"

type SystemCollector struct {
    mountPath string
}

func NewSystemCollector() *SystemCollector {
    return &SystemCollector{
        mountPath: "/",
    }
}

func (s *SystemCollector) Collect() (Metrics, error) {

    // Collect CPU
    cpuStats, err := CollectCPU()
    if err != nil {
        return Metrics{}, err
    }

    // Collect Memory
    memoryStats, err := CollectMemory()
    if err != nil {
        return Metrics{}, err
    }

    // Collect Disk
    diskStats, err := CollectDisk(s.mountPath)
    if err != nil {
        return Metrics{}, err
    }

    return Metrics{
        CPU:       cpuStats,
        Memory:    memoryStats,
        Disk:      diskStats,
        Timestamp: time.Now().UTC(),
    }, nil
}