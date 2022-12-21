package grpc

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"wash_bonus/intapi"
	"wash_bonus/internal/entity"
)

func (s *Service) InitConn(ctx context.Context, request *intapi.InitConnRequest) (response *intapi.InitConnResponse, err error) {
	var code int64
	var msg string

	defer func() {
		if err != nil {
			response = &intapi.InitConnResponse{
				Error: &intapi.ErrorMsg{
					Code: code,
					Msg:  msg,
				},
			}
			err = nil
		}
	}()

	server, err := s.washServerSvc.GetWashServerByKey(ctx, request.ServiceKey)
	if err != nil {
		switch {
		case errors.Is(err, entity.ErrNotFound):
			code = http.StatusNotFound
			msg = "server not found"
		default:
			code = http.StatusInternalServerError
			msg = "internal error"
		}
	}

	conn := &entity.WashServerConnection{
		WashServer:   server,
		ConnectionID: uuid.NewV4(),
	}

	s.connectionsCache.Set(request.ServiceKey, *conn)
	response = &intapi.InitConnResponse{
		ConnectionID: conn.ConnectionID.String(),
		Error:        nil,
	}
	return
}

func (s *Service) HealthCheck(ctx context.Context, request *intapi.HealthCheckRequest) (response *intapi.HealthCheckResponse, err error) {
	var code int64
	var msg string

	defer func() {
		if err != nil {
			response = &intapi.HealthCheckResponse{
				Error: &intapi.ErrorMsg{
					Code: code,
					Msg:  msg,
				},
			}
			err = nil
		}
	}()

	if !s.isValidConnection(request.ServiceKey, request.ConnectionID) {
		code = http.StatusForbidden
		msg = "bad connection"
		err = fmt.Errorf(msg)
	}

	return
}

func (s *Service) isValidConnection(serviceKey string, connectionID *string) bool {
	conn := s.connectionsCache.Get(serviceKey)
	if conn == nil {
		return false
	}

	switch {
	case connectionID == nil:
		fallthrough
	case conn.ConnectionID.String() != *connectionID:
		return false
	}

	err := s.connectionsCache.Refresh(serviceKey, *conn)
	if err != nil {
		return false
	}

	return true
}
