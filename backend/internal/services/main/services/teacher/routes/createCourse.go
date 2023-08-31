package routes

import (
	"context"
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(ctx *gin.Context, c svccourse.CourseServiceClient) {
	body := svccourse.CreateCourseRequest{}
	userID, _ := ctx.Get("userId")
	body.UserId = userID.(int64)

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	response, err := c.CreateCourse(context.Background(), &body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success":  true,
		"error":    "dfdf",
		"courseId": 1,
	})
}

func AddLessonToCreatedCourse(ctx *gin.Context, c svccourse.CourseServiceClient) {
	userID, _ := ctx.Get("userId")

	body := svccourse.AddLessonRequest{
		UserId: userID.(int64),
	}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	response, err := c.AddLessonToCreatedCourse(context.Background(), &body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"error":   response.Error,
	})
}
