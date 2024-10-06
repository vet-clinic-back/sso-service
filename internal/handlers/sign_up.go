package handlers

import (
	"database/sql"
	"errors"
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

	log.Debug("creating owner")
	_, err = h.service.Auth.GetOwner(models.Owner{User: models.User{
		Email: input.Email,
		Phone: input.Phone,
	}})
	if err == nil { // if found owner with same email
		log.Error("failed to create new owner. Owner with same email or phone already exists")
		h.newErrorResponse(c, http.StatusConflict, "owner with same email or phone already exists")
		return
	}

	if !errors.Is(err, sql.ErrNoRows) {
		log.Error(err)
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to find owner in db")
		return
	}

	ownerID, err := h.service.Auth.CreateOwner(input)
	if err != nil {
		log.Error("failed to create new owner: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create new owner")
		return
	}

	token, err := h.service.Auth.CreateToken(ownerID, input.FullName, false)
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

	log.Debug("creating vet")
	_, err = h.service.Auth.GetVet(models.Vet{User: models.User{Email: input.Email}})
	if err == nil { // if found vet with same email
		log.Error("failed to create new vet. Vet with same email already exists")
		h.newErrorResponse(c, http.StatusConflict, "vet with same email already exists")
		return
	}

	if !errors.Is(err, sql.ErrNoRows) {
		log.Error(err)
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to find vet in db")
		return
	}

	vetID, err := h.service.Auth.CreateVet(input)
	if err != nil {
		log.Error("failed to create new vet: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.Auth.CreateToken(vetID, input.FullName, false)
	if err != nil {
		log.Error("failed to create token: ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to create token")
		return
	}

	log.Info("successfully signed up")
	c.JSON(http.StatusOK, models.SuccessDTO{Token: token}) // TODO
}
