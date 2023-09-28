package auth

import (
	"context"
	svcauth "english/backend/api/proto/svc-auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (a *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authCookie, _ := ctx.Cookie("isSignedIn")

	if authCookie == "" {
		ctx.Redirect(http.StatusSeeOther, "/auth/signin")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	response, err := a.svc.Client.Validate(context.Background(), &svcauth.ValidateRequest{
		Token: authCookie,
	})
	if err != nil || response.Status != http.StatusOK {
		ctx.Redirect(http.StatusSeeOther, "/auth/signin")
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("userId", response.UserId)
	ctx.Set("isTeacher", response.IsTeacher)
	ctx.Next()
}

func (a *AuthMiddlewareConfig) AuthCheckUser(ctx *gin.Context) {
	authCookie, _ := ctx.Cookie("isSignedIn")

	if authCookie == "" {
		ctx.Next()
		return
	}

	response, err := a.svc.Client.Validate(context.Background(), &svcauth.ValidateRequest{
		Token: authCookie,
	})
	if err != nil || response.Status != http.StatusOK {
		ctx.Next()
		return
	}

	ctx.Set("userId", response.UserId)
	ctx.Set("isTeacher", response.IsTeacher)
	ctx.Next()
}

func (a *AuthMiddlewareConfig) AuthValidateClient(ctx *gin.Context) {
	isTeacher, _ := ctx.Get("isTeacher")

	if isTeacher == false {
		validateWordInPath := strings.Split(ctx.Request.URL.Path[1:], "/")[0]
		if validateWordInPath == "teacher" {
			ctx.Redirect(http.StatusSeeOther, "/student")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	} else {
		validateWordInPath := strings.Split(ctx.Request.URL.Path[1:], "/")[0]
		if validateWordInPath == "student" {
			ctx.Redirect(http.StatusSeeOther, "/teacher")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}

	ctx.Next()
}
