package teacher

import (
	"english/backend/config"
	auth2 "english/backend/internal/services/main/services/auth"
	routes2 "english/backend/internal/services/main/services/teacher/routes"
	"github.com/gin-gonic/gin"
)

func (svc *ServiceClient) SaveLessons(ctx *gin.Context) {
	routes2.SaveLessons(ctx, svc.Client)
}

func (svc *ServiceClient) SaveCourse(ctx *gin.Context) {
	routes2.SaveCourse(ctx, svc.Client)
}

func (svc *ServiceClient) GetCourseByID(ctx *gin.Context) {
	routes2.GetCourseByID(ctx, svc.Client)
}

func (svc *ServiceClient) GetListOfCourses(ctx *gin.Context) {
	routes2.GetListOfCourses(ctx, svc.Client)
}

func (svc *ServiceClient) AddLessonToCreatedCourse(ctx *gin.Context) {
	routes2.AddLessonToCreatedCourse(ctx, svc.Client)
}

func (svc *ServiceClient) CreateCourse(ctx *gin.Context) {
	routes2.CreateCourse(ctx, svc.Client)
}

func (svc *ServiceClient) UploadFileCourse(ctx *gin.Context) {
	routes2.UploadFileCourse(ctx, svc.Client)
}

func RegisterRoutes(engine *gin.Engine, authService *auth2.ServiceClient, cfg *config.MainConfig) {
	a := auth2.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}

	teacher := engine.Group("/teacher")
	teacher.Use(a.AuthRequired)
	teacher.Use(a.AuthValidateClient)
	{
		teacher.GET("", routes2.TeacherPageApplication)
		teacher.GET("/create-course", routes2.TeacherPageApplication)
		teacher.GET("/list-courses", routes2.TeacherPageApplication)
		teacher.GET("/check-lessons", routes2.TeacherPageApplication)
		teacher.GET("/profile", routes2.TeacherPageApplication)
		teacher.GET("/profile/settings", routes2.TeacherPageApplication)

		//create course
		teacher.POST("create-course", svc.CreateCourse)
		teacher.POST("add-lessons", svc.AddLessonToCreatedCourse)

		//list of courses
		teacher.GET("get-list-courses", svc.GetListOfCourses)
		teacher.GET("get-course", svc.GetCourseByID)
		teacher.POST("save-course", svc.SaveCourse)
		teacher.POST("save-lessons", svc.SaveLessons)

		teacher.POST("upload-file", svc.UploadFileCourse)
	}
}
