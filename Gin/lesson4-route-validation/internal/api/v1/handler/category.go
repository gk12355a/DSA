package v1handler

import (
	"ginroute/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

var validCategory = map[string]bool{
	"php":    true,
	"js":     true,
	"golang": true,
}

// private var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get Category V1",
	})
}
func (c *CategoryHandler) GetCategoryByIdV1(ctx *gin.Context) {
	category := ctx.Param("category")
	// if !validCategory[category] {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid category",
	// 	})
	// 	return
	// }
	if err := utils.ValidationInList("category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Get Category By Id V1",
		"category": category,
	})
}
func (c *CategoryHandler) PostCategoryV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post Category V1",
	})
}
func (c *CategoryHandler) PutCategoryV1ById(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Put Category By Id V1",
	})
}
func (c *CategoryHandler) DeleteCategoryV1ById(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete Category By Id V1",
	})
}
