package handler

import (
	"net/http"

	"github.com/asim9115/incidentix/internal/api/repository"
	"github.com/gin-gonic/gin"
)

type ServerRegistration struct {
	Name string `json:"name"`
}

func (h *Handler)RegisterServer(c *gin.Context) {
	var s ServerRegistration
	if err := c.ShouldBindBodyWithJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	var server repository.Server
	server.Name = s.Name
	if err := h.servers.CreateServer(c, &server); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err})
		return
	}

	c.JSON(http.StatusCreated, "created successfully")
}

func (h *Handler)GetAllServers(c *gin.Context) {
	servers, err := h.servers.GetAll(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, servers)
}
