package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserPageApplication(ctx *gin.Context) {
	clientType := ctx.Param("clientType")
	ctx.HTML(http.StatusOK, clientType, nil)
}
