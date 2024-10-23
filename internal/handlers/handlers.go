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

	auth := router.Group("/auth")
	{
		v1 := auth.Group("/v1")
		{
			signUp := v1.Group("/sign-up")
			{
				signUp.POST("/owner", h.signUpOwner)
				signUp.POST("/vet", h.signUpVet)
			}
			v1.POST("/sign-in", h.signIn)
		}
	}

	info := router.Group("/info")
	{
		v1 := info.Group("/v1")
		{
			pet := v1.Group("/pet")
			{
				pet.POST("/create", h.createPet)
				pet.GET("/get", h.getAllPets)
				pet.GET("/get/:id", h.getPet)
				pet.PUT("/update/:id", h.updatePet)
				pet.DELETE("/delete/:id", h.deletePet)
			}
			owner := v1.Group("/owner")
			{
				owner.POST("/create", h.createOwner)
				owner.GET("/get/:id", h.getOwner)
				owner.GET("/get", h.getAllOwners)
				owner.PUT("/update/:id", h.updateOwner)
				owner.DELETE("/delete/:id", h.deleteOwner)
			}
		}
	}

	return router
}
