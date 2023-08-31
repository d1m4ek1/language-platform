package services

import (
	"context"
	"database/sql"
	svcauth "english/backend/api/proto/svc-auth"
	"english/backend/internal/services/svc-auth/models"
	"english/backend/shared/errorlog"
	sharedjwt "english/backend/shared/jwt"
	"errors"
	"net/http"
)

func (s *Server) Signin(ctx context.Context, req *svcauth.SigninRequest) (*svcauth.SigninResponse, error) {
	client := models.Client{
		Login: req.Login,
	}

	if status, err := client.GetClient(s.H.DB); errors.Is(err, sql.ErrNoRows) {
		return &svcauth.SigninResponse{
			Status: int64(status),
			Error:  "Не верный логин или пароль!",
		}, nil
	} else if err != nil {
		return &svcauth.SigninResponse{
			Status: int64(status),
			Error:  "Неизвестная ошибка!",
		}, errorlog.NewError(err.Error())
	}

	match := sharedjwt.CheckPasswordHash(req.Password, client.Password)
	if !match {
		return &svcauth.SigninResponse{
			Status: http.StatusOK,
			Error:  "Не верный логин или пароль!",
		}, nil
	}

	token, err := s.JWT.GenerateToken(client)
	if err != nil {
		return &svcauth.SigninResponse{
			Status: http.StatusNotFound,
			Error:  "Неизвестная ошибка!",
		}, err
	}

	return &svcauth.SigninResponse{
		Status:    http.StatusOK,
		Token:     token,
		IsTeacher: client.IsTeacher,
		UserId:    client.Id,
	}, nil
}
