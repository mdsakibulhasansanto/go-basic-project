package main

import (
	"final-project/cache"
	"final-project/db"
	"final-project/handler"
	"final-project/radis"
	"final-project/repositories"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	dbConn, err := db.NewMySqlDB()
	if err != nil {
		panic(err)
	}

	redisClient, err := radis.NewRedisClient()
	if err != nil {
		panic(err)
	}

	redisCache := cache.NewRedisCache(redisClient)
	baseRepo := repositories.NewMySQLUserRepository(dbConn)
	userRepo := repositories.NewRedisMySQLUserRepository(baseRepo, redisCache)

	authService := services.NewAuthService(userRepo)
	authHandler := handler.NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/register", authHandler.Register)
	router.GET("/user", authHandler.GetUserByEmail)

	router.Run(":8080")
}
