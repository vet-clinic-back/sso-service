package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/models"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	// TODO Validate input

	if false {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}

	_, err := h.service.Auth.CreateUser(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, models.SuccessDTO{Token: "qwe"}) // TODO
}
