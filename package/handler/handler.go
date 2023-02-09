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

	api := router.Group("/api/")
	{
		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getAllCategories)
			categories.GET("/:id", h.getCategoryById)
			categories.PUT("/:id", h.updateCategory)
			categories.DELETE("/:id", h.deleteCategory)

			products := categories.Group(":id/products")
			{
				products.POST("/", h.createProduct)
				products.GET("/", h.getAllProducts)
				products.GET("/:product_id", h.getProductById)
				products.PUT("/:product_id", h.updateProduct)
				products.DELETE("/:product_id", h.deleteProduct)
			}
		}
	}

	return router
}
