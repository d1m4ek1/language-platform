package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SigninPage(ctx *gin.Context) {
	_, existsUserId := ctx.Get("userId")

	if existsUserId {
		ctx.Redirect(http.StatusSeeOther, "/")
		ctx.AbortWithStatus(http.StatusOK)
	}

	ctx.HTML(http.StatusOK, "signin", nil)
}

func RegistrationPage(ctx *gin.Context) {
	_, existsUserId := ctx.Get("userId")

	if existsUserId {
		ctx.Redirect(http.StatusSeeOther, "/")
		ctx.AbortWithStatus(http.StatusOK)
	}

	ctx.HTML(http.StatusOK, "registration", nil)
}
