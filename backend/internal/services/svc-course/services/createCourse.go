package services

import (
	"context"
	svccourse "english/backend/api/proto/svc-course"
	"english/backend/internal/services/svc-course/models"
	"net/http"
)

func (s *Server) CreateCourse(ctx context.Context, req *svccourse.CreateCourseRequest) (*svccourse.CreateCourseResponse, error) {
	client := models.NewCourse(req)

	courseID, status, err := client.InsertCourse(s.H.DB)
	if err != nil {
		return &svccourse.CreateCourseResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svccourse.CreateCourseResponse{
		Status:   http.StatusOK,
		CourseId: int64(courseID),
	}, nil
}

func (s *Server) AddLessonToCreatedCourse(ctx context.Context, req *svccourse.AddLessonRequest) (*svccourse.AddLessonResponse, error) {
	client := models.NewAddLessons(req)

	status, err := client.AddLessonToCreatedCourse(s.H.DB)
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

func (s *Server) SetFileForCourse(ctx context.Context, req *svccourse.FileForCourseRequest) (*svccourse.FileForCourseResponse, error) {
	client := models.NewFileForCourse(req)

	if status, err := client.SetFileMain(s.H.DB); err != nil {
		return &svccourse.FileForCourseResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svccourse.FileForCourseResponse{
		Status: int64(http.StatusOK),
	}, nil
}
