package routes

import (
	"context"
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LoginRequestBody struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func Singin(ctx *gin.Context, c svcauth.AuthServiceClient) {
	body := LoginRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	response, err := c.Signin(context.Background(), &svcauth.SigninRequest{
		Login:    body.Login,
		Password: body.Password,
	})
	if err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	if response.Token != "" {
		ctx.SetCookie("isSignedIn", response.Token, int(time.Hour*12), "/", "localhost:8000", false, true)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":   true,
		"error":     response.Error,
		"isTeacher": response.IsTeacher,
	})
}
