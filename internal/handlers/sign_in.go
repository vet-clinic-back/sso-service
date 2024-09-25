package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/utils"
)

func (h *Handler) signIn(c *gin.Context) {
	op := "Handler.signIn"
	log := h.log.WithField("op", op)

	var input models.User

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		log.Error("failed to bind json: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	log.Debug("validating input")
	if err := utils.ValidateSignInDTO(input); err != nil {
		log.Error("failed to validate input: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. email & password required")
		return
	}

	log.Debug("creating token")
	token, err := h.service.Auth.CreateToken(input.Email, input.Password)
	if err != nil {
		log.Error("failed to create token: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create token")
		return
	}

	log.Info("successfully signed in")
	c.JSON(http.StatusOK, models.SuccessDTO{Token: token})
}
