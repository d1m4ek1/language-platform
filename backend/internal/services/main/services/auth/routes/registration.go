package routes

import (
	"context"
	svcauth "english/backend/api/proto/svc-auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type RegistrationRequestBody struct {
	Login     string `json:"login"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	RegToken  string `json:"regToken"`
	Password  string `json:"password"`
}

func Registration(ctx *gin.Context, c svcauth.AuthServiceClient) {
	body := RegistrationRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	response, err := c.Registration(context.Background(), &svcauth.RegRequest{
		Login:             body.Login,
		Email:             body.Email,
		FirstName:         body.FirstName,
		LastName:          body.LastName,
		RegistrationToken: body.RegToken,
		Password:          body.Password,
	})
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.Status), gin.H{
			"success": false,
		})
		return
	}

	if response.Token != "" {
		ctx.SetCookie("isSignedIn", response.Token, int(time.Hour*12), "/", "localhost:8000", false, true)
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"error":   response.Error,
	})
}
