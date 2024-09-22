package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/models"
)

// TODO регистрация и авторизация, возвращает JWT
// TODO обновление JWT

func (h *Handler) signUp(c *gin.Context) {
	var input models.User
	var handler *Handler

	if err := c.BindJSON(&input); err != nil {
		handler.newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	id, err := h.service.Auth.CreateUser(input)
	if err != nil {
		handler.newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string
	Password string
}

func (h *Handler) signIn(c *gin.Context) {
	var input models.User
	var handler *Handler

	if err := c.BindJSON(&input); err != nil {
		handler.newResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.service.Auth.CreateToken(input.Username, input.Password)
	if err != nil {
		handler.newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
