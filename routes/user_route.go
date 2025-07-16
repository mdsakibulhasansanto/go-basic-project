package routes

import (
	"final-project/handler"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, handler *handler.UserHandler) {
	userGroup := router.Group("/user")
	userGroup.PUT("/update/:id", handler.UpdateUser)
	userGroup.DELETE("/delete", handler.DeleteUser)
}
