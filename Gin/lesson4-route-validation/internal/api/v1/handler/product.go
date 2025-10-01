package v1handler

import (
	"ginroute/utils"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)
var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductV1(ctx *gin.Context) {
	search := ctx.Query("search")
	if err := utils.ValidationRequired("search", search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := utils.ValidationStringLength("search", search, 3, 50); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if err := utils.ValidationRegex("search", search, searchRegex, "contains invalid characters"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	limitStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid limit",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product V1",
		"search":  search,
		"limit":   limit,
	})
}
func (p *ProductHandler) GetProductBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")
	if err := utils.ValidationRegex("slug", slug, slugRegex, "contains invalid characters"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Product By Slug V1",
		"slug":    slug,
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
