package handlers

import (
	"database/sql"
	"errors"
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

	var token string

	log.Debug("creating token")
	owner, err := h.service.Auth.GetOwner(models.Owner{User: models.User{Email: input.Email, Password: input.Password}})
	if err == nil { // if found owner with same email
		token, err = h.service.Auth.CreateToken(owner.ID, owner.FullName, false)
		if err != nil {
			log.Error(err)
			h.newErrorResponse(c, http.StatusUnauthorized, "failed to create token")
			return
		}
	}

	var vet models.Vet
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		log.Warn("failed to create token as owner: ", err.Error())
		vet, err = h.service.Auth.GetVet(models.Vet{User: models.User{Email: input.Email, Password: input.Password}})
		if err == nil { // if found vet with same email
			token, err = h.service.Auth.CreateToken(vet.ID, vet.FullName, true)
			if err != nil {
				log.Error(err)
				h.newErrorResponse(c, http.StatusUnauthorized, "failed to create token")
				return
			}
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			log.Error("User not found ", err)
			h.newErrorResponse(c, http.StatusUnauthorized, "user not found")
			return
		}
		log.Error("failed to create token ", err)
		h.newErrorResponse(c, http.StatusUnauthorized, "failed to creato token")
	}

	log.Info("successfully signed in")
	c.JSON(http.StatusOK, models.SuccessDTO{Token: token})
}
