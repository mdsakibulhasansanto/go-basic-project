package main

import (
	"final-project/db"
	handlers "final-project/handler"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	database, err := db.NewMySqlDB()
	if err != nil {
		panic(err)
	}

	userRepo := repositories.NewUserRepository(database)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/register", authHandler.Register)

	router.Run(":8080")
}
