package handler

import (
	"github.com/asim9115/incidentix/internal/api/service"
)

type Handler struct {
	metrics *service.MetricsService
	servers *service.ServersService
}


func New( metrics *service.MetricsService, servers *service.ServersService) *Handler{
	return &Handler{

		metrics: metrics,
		servers: servers,
	}
}