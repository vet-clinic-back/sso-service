package handlers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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

/*

- api
  - v1
	- vet
	  - auth
	- User
	  - auth

*/

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600,
	}))

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
		}
	}

	return router
}
