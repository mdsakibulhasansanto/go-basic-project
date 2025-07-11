package handler

import (
	"final-project/services"
	"final-project/utils"
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

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var request RegisterRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
		return
	}

	if err := h.service.RegisterUser(request.Username, request.Email, request.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(ctx *gin.Context) {

	var request LoginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Validation error"})
		return
	}

	user, error := h.service.Login(request.Email, request.Password)
	if error != nil {
		ctx.JSON(500, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	token, err := utils.GenerateJwt(user.Email)

	if err != nil {
		ctx.JSON(500, gin.H{

			"error": "Invalid email and password",
		})
	}

	ctx.JSON(201, gin.H{
		"message": "Login successfull",
		"user ": gin.H{
			"id":       user.Id,
			"username": user.Username,
			"email":    user.Email,
			"token":    token,
		},
	})
}

func (h *AuthHandler) GetUserByEmail(ctx *gin.Context) {

	//email := ctx.Query("email")
	email := ctx.GetString("email")
	user, err := h.service.GetUserByEmail(email)
	if err != nil || user == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":       user.Id,
		"username": user.Username,
		"email":    user.Email,
		"token":    user.VerificationToken,
	})
}
