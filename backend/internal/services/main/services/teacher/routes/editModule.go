package routes

import (
	"context"
	svcmodule "english/backend/api/proto/svc-module"
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetModuleByID(ctx *gin.Context, svc svcmodule.ModuleServiceClient) {
	moduleID, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	body := &svcmodule.ModuleDataRequest{
		ModuleId: int64(moduleID),
	}

	response, err := svc.GetModuleByID(context.Background(), body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.Status), gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
