package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User V2",
	})
}
func (u *UserHandler) GetUserByIdV2(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get User By Id V2",
	})
}
func (u *UserHandler) PostUserV2(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post User V2",
	})
}
func (u *UserHandler) PutUserV2ById(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Put User By Id V2",
	})
}
func (u *UserHandler) DeleteUserV2ById(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"message": "Delete User By Id V2",
	})
}
