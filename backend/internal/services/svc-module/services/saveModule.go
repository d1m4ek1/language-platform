package services

import (
	"context"
	svccourse "english/backend/api/proto/svc-module"
	"english/backend/internal/services/svc-module/models"
	"net/http"
)

func (s *Server) SaveCourse(ctx context.Context, req *svccourse.CreateCourseRequest) (*svccourse.CreateCourseResponse, error) {
	client := models.NewCourse(req)

	courseID, status, err := client.SetCourse(s.H.DB)
	if err != nil {
		return &svccourse.CreateCourseResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svccourse.CreateCourseResponse{
		Status:   http.StatusOK,
		CourseId: courseID,
	}, nil
}

func (s *Server) SaveLessons(ctx context.Context, req *svccourse.AddLessonRequest) (*svccourse.AddLessonResponse, error) {
	client, err := models.NewAddLessons(req)
	if err != nil {
		return &svccourse.AddLessonResponse{
			Status: http.StatusInternalServerError,
			Error:  "Неизвестная ошибка",
		}, err
	}

	status, err := client.SetLessons(s.H.DB)
	if err != nil {
		return &svccourse.AddLessonResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svccourse.AddLessonResponse{
		Status: http.StatusOK,
	}, nil
}
