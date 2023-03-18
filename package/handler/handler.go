package handler

import (
	"github.com/NurzhauganovA/online-store/package/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/", h.userIdentity)
	{
		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
			categories.GET("/:id", h.getCategoryById)

			products := categories.Group(":id/products")
			{
				products.POST("/", h.createProduct)
				products.GET("/", h.getAllProducts)
			}
		}

		products := api.Group("products")
		{
			products.GET("/:id", h.getProductById)
		}
	}

	return router
}
