package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product V1",
	})
}
func (p *ProductHandler) GetProductByIdV1(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Id V1",
	})
}
func (p *ProductHandler) PostProductV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post Product V1",
	})
}
func (p *ProductHandler) PutProductV1ById(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Put Product By Id V1",
	})
}
func (p *ProductHandler) DeleteProductV1ById(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete Product By Id V1",
	})
}
