package services

import (
	"context"
	svccourse "english/backend/api/proto/svc-module"
	"english/backend/internal/services/svc-module/models"
	"net/http"
)

func (s *Server) GetListOfCourses(ctx context.Context, req *svccourse.ListOfCoursesRequest) (*svccourse.ListOfCoursesResponse, error) {
	client := &models.ListOfCourses{
		UserID: req.UserID,
	}

	if status, err := client.GetListOfCourses(s.H.DB); err != nil {
		return &svccourse.ListOfCoursesResponse{
			Status: int64(status),
			Error:  "Курсы не найдены, неизвестная ошибка!",
		}, err
	}

	return &svccourse.ListOfCoursesResponse{
		CourseData: client.ProtoCourseData,
		Status:     http.StatusOK,
	}, nil
}

func (s *Server) GetCourseByID(ctx context.Context, req *svccourse.CourseDataRequest) (*svccourse.CourseDataResponse, error) {
	response, err := models.SelectCourseByID(s.H.DB, req.CourseId)
	if err != nil {
		return response, err
	}

	return response, nil
}
