package collector
import (
    "fmt"

    "github.com/shirou/gopsutil/v4/cpu"
)

func CollectCPU() (CPU, error) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		return CPU{}, err
	}
	if len(percentages) == 0{
		return CPU{}, fmt.Errorf("no cpu stats returned")
	}

	return CPU{
		Usage: percentages[0],
		Valid: true,
	}, nil
}