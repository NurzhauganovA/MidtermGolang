package handler

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createCategory(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input endpoint.Category
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Category.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

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
