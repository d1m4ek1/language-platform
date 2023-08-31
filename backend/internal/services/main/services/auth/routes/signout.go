package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signout(ctx *gin.Context) {
	token, _ := ctx.Cookie("isSignedIn")

	if token != "" {
		ctx.SetCookie("isSignedIn", "", -1, "/", "localhost:8000", false, true)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
