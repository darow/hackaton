package server

import (
	"fmt"
	"hackaton/internal/model"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *server) createUser(c *gin.Context) {
	u := &model.User{
		Username: c.PostForm("name"),
		Password: c.PostForm("password"),
	}

	err := s.store.User().Create(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}
	s.createSession(c)
}

func (s *server) createSession(c *gin.Context) {
	u, err := s.store.User().Login(c.PostForm("name"), c.PostForm("password"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
		return
	}

	session, err := s.store.Session().Create(u)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	c.SetCookie("session", session.Token, int(session.ExpirationTime.Sub(time.Now()).Seconds()), "", "", false, false)
	c.JSON(http.StatusCreated, u)
}

func (s *server) whoAmI(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Logout
// @Tags auth
// @Description delete session cookie
// @ID logout
// @Success 200
// @Success 400 {object} any
// @Router /auth/logout [get]
func (s *server) logout(c *gin.Context) {
	u, ok := c.Get(ctxUserKey)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объекта с ключом не существует", ErrNotFoundInContext))
		return
	}
	user, ok := u.(*model.User)
	if !ok {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("%w объект имеет некорректный тип", ErrNotFoundInContext))
		return
	}

	user.IsOnline = false
	c.SetCookie("session", "", -1, "", "", false, false)
	c.Status(http.StatusOK)
}
