package routes

import (
	"final-project/handler"
	"final-project/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(router *gin.Engine, handler *handler.ProductHandler) {

	authGroup := router.Group("/product")
	authGroup.Use(middlewares.JWTAuthMiddleware())

	authGroup.POST("/create", handler.CreateProduct)
	authGroup.GET("/all", handler.GetAllProducts)
}
