package service

import (
	"context"
	"time"

	"github.com/asim9115/incidentix/internal/api/repository"
)

type MetricsService struct {
	repo *repository.MetricsRepo
}

func NewMetricsService(repo *repository.MetricsRepo) *MetricsService {
	return &MetricsService{repo: repo}
}

func (s *MetricsService) StoreMetrics(ctx context.Context, id, serverID string, cpu, ram, space float64, ts time.Time) error {
	return s.repo.Insert(ctx, id, serverID , cpu, ram, space , ts)
}

func (s *MetricsService) DisplayAll(ctx context.Context) (*[]repository.Metrics, error) {
	return s.repo.DisplayAll(ctx)
}
