package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting server...")
	r := gin.Default()
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World"})
	})
	r.GET("/users/:user_id", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "List of users"})
	})
	r.GET("/user/:user_id", func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")
		ctx.JSON(http.StatusOK, gin.H{"data": "User ID",
			"user_id": user_id,
		})
	})
	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "List of products"})
	})
	r.GET("/product/:product_id", func(ctx *gin.Context) {
		product_id := ctx.Param("product_id")
		ctx.JSON(http.StatusOK, gin.H{"data": "product_id", "Product_id": product_id})

		prices := ctx.Query("prices")
		colors := ctx.Query("colors")
		ctx.JSON(http.StatusOK, gin.H{"prices": prices, "colors": colors})
	})

	r.Run(":8080")
}
