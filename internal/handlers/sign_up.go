package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"github.com/vet-clinic-back/sso-service/internal/utils"
)

func (h *Handler) signUpOwner(c *gin.Context) {
	op := "Handler.signUp"
	log := h.log.WithField("op", op)

	var input models.Owner

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. Failed to parse")
		return
	}

	log.Debug("validating input")
	err := utils.ValidateSignUpOwner(input)
	if err != nil {
		log.Error("failed to validate input: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body values")
		return
	}

	log.Debug("creating user")
	_, err = h.service.Auth.CreateOwner(input)
	if err != nil {
		log.Error("failed to create new owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.Auth.CreateToken(input.Email, input.Password, false)
	if err != nil {
		log.Error("failed to create token: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create token")
		return
	}

	log.Info("successfully signed up")
	c.JSON(http.StatusOK, models.SuccessDTO{Token: token}) // TODO
}

func (h *Handler) signUpVet(c *gin.Context) {
	op := "Handler.signUp"
	log := h.log.WithField("op", op)

	var input models.Vet

	log.Debug("binding json")
	if err := c.BindJSON(&input); err != nil {
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body. Failed to parse")
		return
	}

	log.Debug("validating input")
	err := utils.ValidateSignUpVet(input)
	if err != nil {
		log.Error("failed to validate input: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "invalid input body values")
		return
	}

	log.Debug("creating user")
	_, err = h.service.Auth.CreateVet(input)
	if err != nil {
		log.Error("failed to create new owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.Auth.CreateToken(input.Email, input.Password, false)
	if err != nil {
		log.Error("failed to create token: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create token")
		return
	}

	log.Info("successfully signed up")
	c.JSON(http.StatusOK, models.SuccessDTO{Token: token}) // TODO
}
