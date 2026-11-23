package service

import (
	"context"

	"github.com/asim9115/incidentix/internal/api/repository"
)

type MetricsService struct {
	repo *repository.MetricsRepo
}

func NewMetricsService(repo *repository.MetricsRepo) *MetricsService {
	return &MetricsService{repo: repo}
}

func (s *MetricsService) StoreMetrics(ctx context.Context, cpu, mem, disk float64) error {
	return s.repo.Insert(ctx, cpu, mem, disk)
}
