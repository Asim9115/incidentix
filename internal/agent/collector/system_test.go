package collector

import (
    "testing"
    "time"

)

func TestSystemCollector(t *testing.T) {
    c := NewSystemCollector()

    // First collection establishes CPU baseline.
    _, err := c.Collect()
    if err != nil {
        t.Fatalf("first collection failed: %v", err)
    }

    time.Sleep(1 * time.Second)

    // Second collection calculates CPU usage.
    second, err := c.Collect()
    if err != nil {
        t.Fatalf("second collection failed: %v", err)
    }

    if !second.CPU.Valid {
        t.Error("second CPU reading should be valid")
    }

    if second.Memory.Total == 0 {
        t.Error("memory total is zero")
    }

    if second.Disk.Total == 0 {
        t.Error("disk total is zero")
    }

    if second.Timestamp.IsZero() {
        t.Error("timestamp should not be zero")
    }

    t.Logf("CPU: %.2f%%", second.CPU.Usage)
    t.Logf("Memory: %.2f%%", second.Memory.Usage)
    t.Logf("Disk: %.2f%%", second.Disk.Usage)
}