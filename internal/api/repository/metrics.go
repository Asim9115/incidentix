package repository

import (
	"context"

	"github.com/asim9115/incidentix/internal/api/database"
)

type MetricsRepo struct {
	db *database.DataBase
}

func NewMetricsRepo(db *database.DataBase) *MetricsRepo {
	return &MetricsRepo{db: db}
}

func (r *MetricsRepo) Insert(ctx context.Context, cpu, mem, disk float64) error {
	return nil
}
