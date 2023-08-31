package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(ctx *gin.Context) {
	isTeacher, _ := ctx.Get("isTeacher")
	_, existsUserId := ctx.Get("userId")

	ctx.HTML(http.StatusOK, "home", gin.H{
		"isTeacher": isTeacher,
		"isAuth":    existsUserId,
	})
}
