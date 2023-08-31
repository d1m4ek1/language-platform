package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StudentPageApplication(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "student", nil)
}
