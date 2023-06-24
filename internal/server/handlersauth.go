package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *server) createSession(c *gin.Context) {
	var role string
	c.Bind(role)

	session, err := s.store.Session().Create(role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.SetCookie("session", session.Token, 9999, "", "", false, false)
	c.JSON(http.StatusCreated, session)
}

func (s *server) logout(c *gin.Context) {
	c.SetCookie("session", "", -1, "", "", false, false)
	c.Status(http.StatusOK)
}
