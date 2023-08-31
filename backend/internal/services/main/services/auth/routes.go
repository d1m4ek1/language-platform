package auth

import (
	"english/backend/config"
	routes2 "english/backend/internal/services/main/services/auth/routes"
	"github.com/gin-gonic/gin"
)

func (svc *ServiceClient) Singin(ctx *gin.Context) {
	routes2.Singin(ctx, svc.Client)
}

func (svc *ServiceClient) Registration(ctx *gin.Context) {
	routes2.Registration(ctx, svc.Client)
}

func RegisterRoutes(engine *gin.Engine, cfg *config.MainConfig) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}

	a := InitAuthMiddleware(svc)

	auth := engine.Group("/auth")
	{
		auth.Use(a.AuthCheckUser).GET("/signin", routes2.SigninPage)
		auth.Use(a.AuthCheckUser).GET("/registration", routes2.RegistrationPage)
		auth.GET("/signout", routes2.Signout)

		auth.POST("/signin", svc.Singin)
		auth.POST("/registration", svc.Registration)
	}

	return svc
}
