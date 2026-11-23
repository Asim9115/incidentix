package handler

import (
	"github.com/asim9115/incidentix/internal/api/database"
	"github.com/asim9115/incidentix/internal/api/service"
)

type Handler struct {
	db *database.DataBase
	metrics *service.MetricsService
}


func New(db *database.DataBase, metrics *service.MetricsService) *Handler{
	return &Handler{
		db:db,
		metrics: metrics,
	}
}