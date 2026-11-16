package collector

import "github.com/shirou/gopsutil/v4/mem"

func CollectMemory() (Memory, error) {
    stats, err := mem.VirtualMemory()
    if err != nil {
        return Memory{}, err
    }

    return Memory{
        Total: stats.Total,
        Used:  stats.Used,
        Usage: stats.UsedPercent,
    }, nil
}