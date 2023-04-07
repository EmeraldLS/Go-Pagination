package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-pagination/controller"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	router.POST("/api/products/populate", controller.Populate)
	router.GET("/api/products/frontend", controller.GetProducts)
	router.GET("/api/products/backend", controller.GetSingleProduct)
	router.Run(":8080")
}
