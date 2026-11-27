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

	mrepo := repository.NewMetricsRepo(db)
	metrics := service.NewMetricsService(mrepo)
	srepo := repository.NewServersService(db)
	servers := service.NewServersService(srepo)

	h := handler.New(db, metrics, servers)

	r.POST("/v1/agents/metrics", h.HandleHeartBeats)
	r.GET("/v1/agents/metrics/all", h.DisplayAll)

	r.POST("/v1/agents/server", h.RegisterServer)
	r.GET("/v1/agents/server/all", h.GetAllServers)
	return r
}