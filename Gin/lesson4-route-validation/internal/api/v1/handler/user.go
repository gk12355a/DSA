package v1handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User V1",
	})
}
func (u *UserHandler) GetUserByIdV1(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID must be a positive integer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User By Id V1",
		"user_id": id,
	})
}

func (u *UserHandler) GetUserByUuIdV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")
	// go get github.com/google/uuid
	_, err := uuid.Parse(uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid UUID",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User By Id V1",
		"user_id": uuidStr,
	})
}
func (u *UserHandler) PostUserV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post User V1",
	})
}
func (u *UserHandler) PutUserV1ById(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Put User By Id V1",
	})
}
func (u *UserHandler) DeleteUserV1ById(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete User By Id V1",
	})
}
