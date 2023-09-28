package teacher

import (
	"english/backend/config"
	"english/backend/internal/services/main/services/auth"
	"english/backend/internal/services/main/services/teacher/routes"
	"github.com/gin-gonic/gin"
)

func (svc *ServiceClient) SaveLessons(ctx *gin.Context) {
	routes.SaveLessons(ctx, svc.Client)
}

func (svc *ServiceClient) SaveCourse(ctx *gin.Context) {
	routes.SaveCourse(ctx, svc.Client)
}

func (svc *ServiceClient) GetCourseByID(ctx *gin.Context) {
	routes.GetCourseByID(ctx, svc.Client)
}

func (svc *ServiceClient) GetListOfCourses(ctx *gin.Context) {
	routes.GetListOfCourses(ctx, svc.Client)
}

func (svc *ServiceClient) AddLessonToCreatedCourse(ctx *gin.Context) {
	routes.AddLessonToCreatedCourse(ctx, svc.Client)
}

func (svc *ServiceClient) CreateCourse(ctx *gin.Context) {
	routes.CreateCourse(ctx, svc.Client)
}

func (svc *ServiceClient) UploadFileCourse(ctx *gin.Context) {
	routes.UploadFileCourse(ctx, svc.Client)
}

func RegisterRoutes(engine *gin.Engine, authService *auth.ServiceClient, cfg *config.MainConfig) {
	a := auth.InitAuthMiddleware(authService)

	svc := &ServiceClient{
		Client: InitServiceClient(cfg),
	}

	teacher := engine.Group("/teacher")
	teacher.Use(a.AuthRequired)
	teacher.Use(a.AuthValidateClient)
	{
		teacher.GET("", routes.TeacherPageApplication)
		teacher.GET("/create-course", routes.TeacherPageApplication)
		teacher.GET("/list-courses", routes.TeacherPageApplication)
		teacher.GET("/check-lessons", routes.TeacherPageApplication)
		teacher.GET("/students", routes.TeacherPageApplication)
		teacher.GET("/profile", routes.TeacherPageApplication)
		teacher.GET("/profile/settings", routes.TeacherPageApplication)

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
