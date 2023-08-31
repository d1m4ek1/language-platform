package services

import (
	"context"
	svcclient "english/backend/api/proto/svc-client"
	"english/backend/internal/services/svc-client/models"
	"english/backend/shared/errorlog"
	"github.com/jinzhu/copier"
	"net/http"
)

func (s *Server) SaveDataClient(ctx context.Context, req *svcclient.SaveDataClientRequest) (*svcclient.SaveDataClientResponse, error) {
	c := models.NewDataClient(req.DataClient)
	if err := copier.Copy(c, req.DataClient); err != nil {
		return &svcclient.SaveDataClientResponse{
			ResponseStatus: &svcclient.ResponseStatus{
				Status: http.StatusInternalServerError,
				Error:  "Неизвестная ошибка",
			},
		}, errorlog.NewError(err.Error())
	}

	response, err := c.UpdateDataClient(s.H.DB)
	if err != nil {
		response.ResponseStatus.Status = http.StatusInternalServerError
		response.ResponseStatus.Error = "Неизвестная ошибка"
		return response, err
	}

	return response, nil
}
