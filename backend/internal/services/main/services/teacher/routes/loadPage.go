package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func TeacherPageApplication(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "teacher", nil)
}
