package routes

import (
	"context"
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/shared/errorlog"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetCourseByID(ctx *gin.Context, svc svccourse.CourseServiceClient) {
	courseID, err := strconv.Atoi(ctx.Query("id"))
	if err != nil {
		fmt.Println(errorlog.NewError(err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
		})
		return
	}

	body := &svccourse.CourseDataRequest{
		CourseId: int64(courseID),
	}

	response, err := svc.GetCourseByID(context.Background(), body)
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
