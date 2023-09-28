package routes

import (
	"context"
	svccourse "english/backend/api/proto/svc-module"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetListOfCourses(ctx *gin.Context, svc svccourse.CourseServiceClient) {
	userID, _ := ctx.Get("userId")
	body := &svccourse.ListOfCoursesRequest{UserID: userID.(int64)}

	response, err := svc.GetListOfCourses(context.Background(), body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.Status), gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"list":    response.CourseData,
		"success": true,
	})
}
