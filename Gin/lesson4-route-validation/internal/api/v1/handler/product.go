package v1handler

import (
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
	if search == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Search Product required",
		})
		return
	}
	if len(search) < 3 || len(search) > 50 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid search query length",
		})
		return
	}
	if !searchRegex.MatchString(search) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Search query contains invalid characters",
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
	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid slug format",
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
