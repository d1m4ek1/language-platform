package services

import (
	"context"
	svcclient "english/backend/api/proto/svc-client"
	"english/backend/internal/services/svc-client/models"
	"net/http"
)

func (s *Server) GetDataClient(ctx context.Context, req *svcclient.GetDataClientRequest) (*svcclient.GetDataClientResponse, error) {
	c := models.NewDataClient(&svcclient.DataClient{UserId: req.UserId})

	response, err := c.GetDataClient(s.H.DB)
	if err != nil {
		response.ResponseStatus.Status = http.StatusInternalServerError
		response.ResponseStatus.Error = "Неизвестная ошибка"
		return response, err
	}

	return response, nil
}
