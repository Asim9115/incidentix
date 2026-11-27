package handler

import (
	"crypto/rand"
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
	server.Token = GenerateToken()

	h.RegisterServer(c, )
	
}

func GenerateToken()(string) {
	prefix := "idx-"
	id := rand.Text()
	return prefix+id[:8]
} 