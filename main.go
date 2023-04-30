package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-pagination/controller"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load(".env")
	router := gin.Default()
	router.Use(cors.Default())
	port := os.Getenv("PORT")
	router.POST("/api/products/populate", controller.Populate)
	router.GET("/api/products/frontend", controller.GetProducts)
	router.GET("/api/products/backend", controller.GetSingleProduct)
	router.Run("0.0.0.0:"+ port)
}
