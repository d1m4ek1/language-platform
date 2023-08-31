package student

import (
	auth2 "english/backend/internal/services/main/services/auth"
	"english/backend/internal/services/main/services/student/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine, authService *auth2.ServiceClient) {
	a := auth2.InitAuthMiddleware(authService)

	student := engine.Group("/student")
	student.Use(a.AuthRequired)
	student.Use(a.AuthValidateClient)
	{
		student.GET("", routes.StudentPageApplication)
		student.GET("/my-education", routes.StudentPageApplication)
	}
}
