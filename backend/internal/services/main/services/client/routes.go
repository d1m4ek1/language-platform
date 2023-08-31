package client

import (
	"english/backend/config"
	auth2 "english/backend/internal/services/main/services/auth"
	routes2 "english/backend/internal/services/main/services/client/routes"
	"github.com/gin-gonic/gin"
)

func (svc *ServiceClient) GetDataClient(ctx *gin.Context) {
	routes2.GetDataClient(ctx, svc.Client)
}

func (svc *ServiceClient) SaveDataClient(ctx *gin.Context) {
	routes2.SaveDataClient(ctx, svc.Client)
}

func RegisterRoutes(engine *gin.Engine, authService *auth2.ServiceClient, cfg *config.MainConfig) {
	a := auth2.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: IniServiceClient(cfg),
	}

	user := engine.Group("/:clientType")
	user.Use(a.AuthRequired)
	{
		user.GET("profile", routes2.UserPageApplication)
		user.Use(CheckUserOnType).GET("profile/settings", routes2.UserPageApplication)
		user.Use(CheckUserOnType).GET("profile/data-client", svc.GetDataClient)

		user.POST("profile/data-client/save", svc.SaveDataClient)
	}
}
