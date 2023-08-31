package routes

import (
	"context"
	svcclient "english/backend/api/proto/svc-client"
	"english/backend/shared/errorlog"
	"english/backend/shared/uploads"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDataClient(ctx *gin.Context, svc svcclient.ClientServiceClient) {
	userId := ctx.Keys["userId"].(int64)

	response, err := svc.GetDataClient(context.Background(), &svcclient.GetDataClientRequest{
		UserId: userId,
	})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.ResponseStatus.Status), response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func SaveDataClient(ctx *gin.Context, svc svcclient.ClientServiceClient) {
	userId := ctx.Keys["userId"].(int64)

	body := &svcclient.SaveDataClientRequest{
		DataClient: &svcclient.DataClient{
			UserId: userId,
		},
	}

	if err := ctx.BindJSON(&body.DataClient); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, svcclient.ResponseStatus{
			Status: http.StatusBadRequest,
			Error:  "Неизвестная ошибка",
		})
	}

	if body.DataClient.DeleteAvatar != "" {
		if err := uploads.DeleteFile(map[string]string{"avatar": body.DataClient.DeleteAvatar}); err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusInternalServerError, &svcclient.ResponseStatus{
				Status: http.StatusInternalServerError,
				Error:  fmt.Sprintf("%s", uploads.ErrorFileIsNotSupported),
			})
			return
		}
	}

	response, err := svc.SaveDataClient(context.Background(), body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.ResponseStatus.Status), response.ResponseStatus)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
