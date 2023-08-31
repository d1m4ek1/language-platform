package services

import (
	"context"
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/internal/services/svc-auth/models"
	sharedjwt "english/backend/shared/jwt"
	"net/http"
)

func (s *Server) Registration(ctx context.Context, req *svcauth.RegRequest) (*svcauth.RegResponse, error) {
	client := models.Client{
		Login:     req.Login,
		Email:     req.Email,
		RegToken:  req.RegistrationToken,
		Password:  sharedjwt.HashPassword(req.Password),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	if messageError, status, err := client.CheckHaveClientByLoginAndEmail(s.H.DB); err != nil {
		return &svcauth.RegResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка!",
		}, err
	} else if err == nil && messageError != "" {
		return &svcauth.RegResponse{
			Status: int64(status),
			Error:  messageError,
		}, nil
	}

	if isRegTokenValid, status, err := client.CheckRegToken(s.H.DB); err != nil {
		return &svcauth.RegResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка!",
		}, err
	} else if err == nil && !isRegTokenValid {
		return &svcauth.RegResponse{
			Status: int64(status),
			Error:  "Токен не существует или истек срок активности",
		}, nil
	}

	if status, err := client.InsertClient(s.H.DB); err != nil {
		return &svcauth.RegResponse{
			Status: int64(status),
			Error:  err.Error(),
		}, err
	}

	token, err := s.JWT.GenerateToken(client)
	if err != nil {
		return &svcauth.RegResponse{
			Status: http.StatusNotFound,
			Error:  err.Error(),
		}, err
	}

	return &svcauth.RegResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
