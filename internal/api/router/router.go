package router

import (
	"github.com/asim9115/incidentix/internal/api/database"
	"github.com/asim9115/incidentix/internal/api/handler"
	"github.com/asim9115/incidentix/internal/api/repository"
	"github.com/asim9115/incidentix/internal/api/service"
	"github.com/gin-gonic/gin"
)

func New(db *database.DataBase) *gin.Engine {
	r := gin.Default()
	repo := repository.NewMetricsRepo(db)
	metrics := service.NewMetricsService(repo)
	h := handler.New(db, metrics)

	r.POST("/v1/agents/metrics", h.HandleHeartBeats)
	return r
}