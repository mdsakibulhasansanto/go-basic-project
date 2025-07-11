package main

import (
	"final-project/cache"
	"final-project/db"
	"final-project/handler"
	"final-project/middlewares"
	"final-project/radis"
	"final-project/repositories"
	"final-project/routes"
	"final-project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	dbConn, err := db.NewMySqlDB()
	if err != nil {
		panic(err)
	}

	gormDb, err := db.NewMySqlGormDB()

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
	router.GET("/login", authHandler.Login)

	router.GET("/user", middlewares.JWTAuthMiddleware(), authHandler.GetUserByEmail)

	/*

		if err := db.MigrateProductTable(gormDb); err != nil {
			panic(err)
		}
	*/

	// Product route handler & dependency
	productRepo := repositories.NewMySQLProductRepository(gormDb)
	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)
	routes.RegisterProductRoutes(router, productHandler)
	router.Run(":8080")
}

// http://localhost:8080/register
