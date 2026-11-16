package collector

import "github.com/shirou/gopsutil/v4/disk"

func CollectDisk(path string) (Disk, error) {
    stats, err := disk.Usage(path)
    if err != nil {
        return Disk{}, err
    }

    return Disk{
        Total: stats.Total,
        Used:  stats.Used,
        Usage: stats.UsedPercent,
    }, nil
}