package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/models"
	"strconv"
)

func ParseOwnerFilters(c *gin.Context) (models.PaginationFilter, error) {
	var filters models.PaginationFilter

	offset, err := getUint64Param("offset", c)
	if err != nil {
		return filters, err
	}
	filters.Offset = offset

	limit, err := getUint64Param("limit", c)
	if err != nil {
		return filters, err
	}
	filters.Limit = limit

	return filters, nil
}

// getUint64Param returns *uint param. On error returns error and nil if param not exists
func getUint64Param(param string, c *gin.Context) (*uint, error) {
	stringParam, ok := c.GetQuery(param)
	if ok {
		paramUint64, err := strconv.ParseUint(stringParam, 10, 32)
		if err != nil {
			return nil, err
		}
		result := uint(paramUint64)
		return &result, nil
	} else {
		return nil, nil
	}
}
