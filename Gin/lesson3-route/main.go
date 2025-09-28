package main

import (
	"log"

	v1handler "ginroute/internal/api/v1/handler"
	v2handler "ginroute/internal/api/v2/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("server starting at :8080")

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		user := v1.Group("/users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUserV1)
			user.GET("/:id", userHandlerV1.GetUserByIdV1)
			user.POST("", userHandlerV1.PostUserV1)
			user.PUT("/:id", userHandlerV1.PutUserV1ById)
			user.DELETE("/:id", userHandlerV1.DeleteUserV1ById)
		}
		product := v1.Group("/products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductV1)
			product.GET("/:id", productHandlerV1.GetProductByIdV1)
			product.POST("", productHandlerV1.PostProductV1)
			product.PUT("/:id", productHandlerV1.PutProductV1ById)
			product.DELETE("/:id", productHandlerV1.DeleteProductV1ById)
		}

	}

	v2 := r.Group("/v2")
	{
		userv2 := v2.Group("/users")
		{
			userHandlerV2 := v2handler.NewUserHandler()
			userv2.GET("", userHandlerV2.GetUserV2)
			userv2.GET("/:id", userHandlerV2.GetUserByIdV2)
			userv2.POST("", userHandlerV2.PostUserV2)
			userv2.PUT("/:id", userHandlerV2.PutUserV2ById)
			userv2.DELETE("/:id", userHandlerV2.DeleteUserV2ById)
		}
	}

	r.Run(
		":8080",
	)
}
