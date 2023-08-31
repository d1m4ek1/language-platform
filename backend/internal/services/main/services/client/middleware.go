package client

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckUserOnType(ctx *gin.Context) {
	clientType := ctx.Param("clientType")
	isTeacher := ctx.Keys["isTeacher"].(bool)

	if isTeacher && clientType == "student" {
		ctx.Redirect(http.StatusSeeOther, "/profile/teacher")
		ctx.AbortWithStatus(http.StatusOK)
		return
	}

	if !isTeacher && clientType == "teacher" {
		ctx.Redirect(http.StatusSeeOther, "/profile/student")
		ctx.AbortWithStatus(http.StatusOK)
		return
	}

	ctx.Next()
}
