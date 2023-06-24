package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) getOnlineUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}
