package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/utils"
)

func (h *Handler) signUp(c *gin.Context) {
	op := "Handler.signUp"
	log := h.log.WithField("op", op)

	var input models.User

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. Failed to parse")
		return
	}

	log.Debug("validating input")
	err := utils.ValidateSignUpDTO(input)
	if err != nil {

		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body values")
		return
	}

	log.Debug("creating user")
	_, err = h.service.Auth.CreateUser(input)
	if err != nil {
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Info("successfully signed up")
	c.JSON(http.StatusOK, models.SuccessDTO{Token: "qwe"}) // TODO
}
