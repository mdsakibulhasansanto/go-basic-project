package handler

import (
	"final-project/models"
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) Create(ctx *gin.Context) {
	var product models.Product

	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	userEmail := ctx.GetString("email")
	if userEmail == "" {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	product.UserEmail = userEmail

	if err := h.service.Create(&product); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
	})
}
