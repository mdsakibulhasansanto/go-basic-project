package handlers

import (
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service *services.AuthService
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: service,
	}
}

// ✅ Register API
func (h *AuthHandler) Register(ctx *gin.Context) {
	var request RegisterRequest

	// JSON binding validation
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Validation error",
		})
		return
	}

	// Call service to register user
	err := h.service.RegisterUser(request.Username, request.Email, request.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Success response
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
	})
}
