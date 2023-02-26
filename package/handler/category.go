package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCategory(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllCategories(c *gin.Context) {

}

func (h *Handler) getCategoryById(c *gin.Context) {

}

func (h *Handler) updateCategory(c *gin.Context) {

}

func (h *Handler) deleteCategory(c *gin.Context) {

}
