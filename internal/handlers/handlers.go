package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/vet-clinic-back/sso-service/internal/services"
	"github.com/vet-clinic-back/sso-service/logger"
)

type Handler struct {
	service *services.Service
	log     *logger.Logger
}

func NewHandler(services *services.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) newResponse(c *gin.Context, statusCode int, message string) {
	h.log.Error(message)
	c.AbortWithStatusJSON(statusCode, message)
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up")
		auth.POST("/sign-in")
	}
	return router
}
