package handlers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/utils"
	"net/http"
)

// @Summary Get all owners
// @Description Get all owners with pagination
// @Security ApiKeyAuth
// @Tags owners
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Produce json
// @Success 200 {object} []models.Owner "Successfully found owners"
// @Failure 404 {object} models.ErrorDTO "Not found in db"
// @Failure 500 {object} models.ErrorDTO "Internal server error"
// @Router  /auth/v1/owner [get]
func (h *Handler) getOwners(c *gin.Context) {
	log := h.log.WithField("op", "Handler.getOwners")

	filters, err := utils.ParseOwnerFilters(c)
	if err != nil {
		log.Error("failed to parse filters: ", err.Error())
		h.newErrorResponse(c, http.StatusBadRequest, "failed to parse filters")
		return
	}

	log.WithField("filters", filters).Info("filters updated")

	log.Debug("retrieving all owners")
	owners, err := h.service.Auth.GetOwners(filters)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("owners not found: ", err.Error())
			h.newErrorResponse(c, http.StatusNotFound, "owners not found")
			return
		}
		log.Error("failed to get owners ", err.Error())
		h.newErrorResponse(c, http.StatusInternalServerError, "failed to get owners")
		return
	}

	log.Info("successfully retrieved all owners")
	c.JSON(http.StatusOK, owners)
}
