package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, message)
}
