package services

import (
	"context"
	svcmodule "english/backend/api/proto/svc-module"
	"english/backend/internal/services/svc-module/models"
	"net/http"
)

func (s *Server) CreateModule(ctx context.Context, req *svcmodule.CreateModuleRequest) (*svcmodule.CreateModuleResponse, error) {
	client := models.NewModule(req)

	moduleID, status, err := client.InsertModule(s.H.DB)
	if err != nil {
		return &svcmodule.CreateModuleResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svcmodule.CreateModuleResponse{
		Status:   http.StatusOK,
		ModuleId: int64(moduleID),
	}, nil
}

func (s *Server) AddLessonToCreatedModule(ctx context.Context, req *svcmodule.AddLessonRequest) (*svcmodule.AddLessonResponse, error) {
	client, err := models.NewAddLessons(req)
	if err != nil {
		return &svcmodule.AddLessonResponse{
			Status: http.StatusInternalServerError,
			Error:  "Неизвестная ошибка",
		}, err
	}

	status, err := client.AddLessonToCreatedModule(s.H.DB)
	if err != nil {
		return &svcmodule.AddLessonResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svcmodule.AddLessonResponse{
		Status: http.StatusOK,
	}, nil
}

func (s *Server) SetFileForModule(ctx context.Context, req *svcmodule.FileForModuleRequest) (*svcmodule.FileForModuleResponse, error) {
	client := models.NewFileForModule(req)

	if status, err := client.SetFileMain(s.H.DB); err != nil {
		return &svcmodule.FileForModuleResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка",
		}, err
	}

	return &svcmodule.FileForModuleResponse{
		Status: int64(http.StatusOK),
	}, nil
}
