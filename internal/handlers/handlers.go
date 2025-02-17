package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/vet-clinic-back/sso-service/docs"
	"github.com/vet-clinic-back/sso-service/internal/logging"
	"github.com/vet-clinic-back/sso-service/internal/service"
)

type Handler struct {
	log     *logging.Logger
	service *service.Service
}

func NewHandler(log *logging.Logger, service *service.Service) *Handler {
	return &Handler{log: log, service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/auth")
	{
		v1 := api.Group("/v1")
		{
			signUp := v1.Group("/sign-up")
			{
				signUp.POST("/owner", h.signUpOwner)
				signUp.POST("/vet", h.signUpVet)
			}

			v1.POST("/sign-in", h.signIn)
			v1.GET("/owner", h.getOwners)
		}
	}

	return router
}
