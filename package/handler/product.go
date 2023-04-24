package handler

import (
	"github.com/NurzhauganovA/online-store/endpoint"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
	"strconv"
)

func (h *Handler) getQueryParam(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productTitle, err := strconv.Atoi(c.Param("title"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}

	product, err := h.services.Product.GetById(userId, productTitle)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)

}

func (h *Handler) createProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}

	var input endpoint.Product
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.Create(userId, categoryId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// create rating
	rating := endpoint.Rating{
		ProductId: id,
		Rating:    input.Rating,
	}

	_, err = h.services.Product.CreateRating(rating)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// create comment
	comment := endpoint.Comment{
		ProductId: id,
		Comment:   input.Comment,
	}

	_, err = h.services.Product.CreateComment(comment)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllProducts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	categoryId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}

	_, err = strconv.ParseFloat(c.Query("min_price"), 64)
	if err != nil {
		_ = 0
	}

	_, err = strconv.ParseFloat(c.Query("max_price"), 64)
	if err != nil {
		_ = math.MaxFloat64
	}

	_, err = strconv.Atoi(c.Query("min_rating"))
	if err != nil {
		_ = 0
	}

	products, err := h.services.Product.GetAll(userId, categoryId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *Handler) getProductById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid category id param")
		return
	}

	product, err := h.services.Product.GetById(userId, productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	title := c.Query("title")

	if title != "" {
		product, err := h.services.GetQueryParam(string(rune(userId)), title)

		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, product)
	}

	c.JSON(http.StatusOK, product)
}

type ratingInput struct {
	Rating int `json:"rating" binding:"required"`
}

func (h *Handler) rateProduct(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id")
		return
	}

	var input ratingInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	rating := input.Rating
	if rating < 1 {
		newErrorResponse(c, http.StatusBadRequest, "Ensure this value is greater than or equal to 1.")
		return
	}

	if rating > 10 {
		newErrorResponse(c, http.StatusBadRequest, "Ensure this value is less than or equal to 10.")
		return
	}

	_, err = h.services.Product.Rate(productID, rating)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "rating submitted successfully"})
}

func (h *Handler) getFilteredProducts(c *gin.Context) {
	var filter struct {
		MinPrice  float64 `form:"min_price"`
		MaxPrice  float64 `form:"max_price"`
		MinRating int     `form:"min_rating"`
	}

	if err := c.BindQuery(&filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	products, _ := h.services.Product.GetFilteredProducts(filter.MinPrice, filter.MaxPrice, filter.MinRating)
	if err := c.BindQuery(&filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}
