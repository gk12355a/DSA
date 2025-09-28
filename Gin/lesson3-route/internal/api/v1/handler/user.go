package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User By Id V1",
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
