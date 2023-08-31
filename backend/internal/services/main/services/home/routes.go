package home

import (
	"english/backend/internal/services/main/services/auth"
	"english/backend/internal/services/main/services/home/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, authService *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authService)

	engine.Use(a.AuthCheckUser).GET("", routes.HomePage)
}
