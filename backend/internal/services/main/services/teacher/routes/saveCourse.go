package routes

import (
	"context"
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/shared/errorlog"
	"english/backend/shared/uploads"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SaveCourseBodyRequest struct {
	DeleteFiles map[string]string              `json:"deleteFiles"`
	Main        *svccourse.CreateCourseRequest `json:"main"`
}

func SaveCourse(ctx *gin.Context, svc svccourse.CourseServiceClient) {
	userID, _ := ctx.Get("userId")

	body := SaveCourseBodyRequest{
		Main: &svccourse.CreateCourseRequest{
			UserId: userID.(int64),
		},
	}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Неизвестная ошибка",
		})
		return
	}

	if err := uploads.DeleteFile(body.DeleteFiles); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Неизвестная ошибка",
		})
		return
	}

	response, err := svc.SaveCourse(context.Background(), body.Main)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.Status), gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func SaveLessons(ctx *gin.Context, svc svccourse.CourseServiceClient) {
	userID, _ := ctx.Get("userId")
	body := &svccourse.AddLessonRequest{
		UserId: userID.(int64),
	}

	if err := ctx.BindJSON(&body); err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	if err := uploads.DeleteLessonFileItems(body.LessonItems); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadGateway, gin.H{
			"success": false,
			"error":   "Неизвестная ошибка",
		})
		return
	}

	if body.DeleteLessonFileItems != nil {
		for _, item := range body.DeleteLessonFileItems {
			if err := uploads.DeleteAllFile(item.Paths); err != nil {
				fmt.Println(err)
				ctx.JSON(http.StatusBadGateway, gin.H{
					"success": false,
					"error":   "Неизвестная ошибка",
				})
				return
			}
		}
		body.DeleteLessonFileItems = nil
	}

	response, err := svc.SaveLessons(context.Background(), body)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(int(response.Status), gin.H{
			"success": false,
			"error":   response.Error,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}
