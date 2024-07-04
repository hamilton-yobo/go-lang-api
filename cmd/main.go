package main

import (
	"github.com/gin-gonic/gin"

	"go-api/controller"
	"go-api/db"
	"go-api/usecase"
)

func main() {

	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	ProductUseCase := usecase.NewProductUseCase()
	// Camada de controllers
	ProductController := controller.NewProductContoller(ProductUseCase)

	server.GET("/PING", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)

	server.Run(":8000")
}
