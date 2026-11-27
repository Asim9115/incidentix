package handler

import (
	"crypto/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TempMetrics struct {
	CPU       CPU       `json:"cpu"`
	Memory    Memory    `json:"memory"`
	Disk      Disk      `json:"disk"`
	Timestamp time.Time `json:"timestamp"`
}

type CPU struct {
	Valid bool    `json:"-"`
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

func (h *Handler) HandleHeartBeats(c *gin.Context) {
	var m TempMetrics
	if err := c.ShouldBindBodyWithJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	if err := h.metrics.StoreMetrics(c, rand.Text(), "HYMBRX6P", m.CPU.Usage, m.Memory.Usage, m.Disk.Usage, m.Timestamp); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "stored")
}

func (h *Handler) DisplayAll(c *gin.Context) {
	metrics, err := h.metrics.DisplayAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, metrics)
}
