package services

import (
	"context"
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/internal/services/svc-auth/models"
	"net/http"
)

func (s *Server) Validate(ctx context.Context, req *svcauth.ValidateRequest) (*svcauth.ValidateResponse, error) {
	claims, err := s.JWT.ValidateToken(req.Token)
	if err != nil {
		return &svcauth.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  "Неизвестная ошибка!",
		}, err
	}

	client := models.Client{Login: claims.Login}
	if status, err := client.GetClient(s.H.DB); err != nil {
		return &svcauth.ValidateResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка!",
		}, err
	}

	return &svcauth.ValidateResponse{
		Status:    http.StatusOK,
		UserId:    client.Id,
		IsTeacher: client.IsTeacher,
	}, nil
}
