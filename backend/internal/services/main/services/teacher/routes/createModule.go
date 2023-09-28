package routes

import (
	"context"
	svcmodule "english/backend/api/proto/svc-module"
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateModule(ctx *gin.Context, c svcmodule.ModuleServiceClient) {
	body := svcmodule.CreateModuleRequest{}
	userID, _ := ctx.Get("userId")
	body.UserId = userID.(int64)

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	response, err := c.CreateModule(context.Background(), &body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"moduleId": response.ModuleId,
	})
}

func AddLessonToCreatedModule(ctx *gin.Context, c svcmodule.ModuleServiceClient) {
	userID, _ := ctx.Get("userId")

	body := svcmodule.AddLessonRequest{
		UserId: userID.(int64),
	}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	response, err := c.AddLessonToCreatedModule(context.Background(), &body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   response.Error,
	})
}
