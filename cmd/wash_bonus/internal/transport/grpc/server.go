package grpc

import (
	"context"
	"errors"
	"time"
	"wash_bonus/intapi"
	"wash_bonus/internal/entity"
)

func (s *Service) HealthCheck(request *intapi.HealthCheckRequest, server intapi.ServerService_HealthCheckServer) error {
	svcKey := request.GetServiceKey()
	if svcKey == "" {
		server.Send(&intapi.HealthCheckResponse{
			Response: false,
			Msg:      entity.ErrEmptyWashServerKey.Error(),
		})
		return nil
	}

	// setting up recover functions if wash server stopped responding
	defer func() {
		if r := recover(); r != nil {
			s.l.Named("transport").Named("grpc").Error(r)
			s.connectionsMutex.Lock()
			delete(s.connections, svcKey)
			s.connectionsMutex.Unlock()
		}
	}()

	s.connectionsMutex.RLock()
	_, exists := s.connections[svcKey]
	s.connectionsMutex.RUnlock()

	if exists {
		server.Send(&intapi.HealthCheckResponse{
			Response: false,
			Msg:      entity.ErrEmptyWashServerKey.Error(),
		})
		return nil
	}

	ctx := context.WithValue(context.TODO(), "washKey", svcKey)

	washServer, err := s.washServerSvc.GetWashServerByKey(ctx, svcKey)
	switch {
	case errors.Is(err, entity.ErrNotFound):
		server.Send(&intapi.HealthCheckResponse{
			Response: false,
			Msg:      entity.ErrWashServerNotFound.Error(),
		})
		return nil
	case err != nil:
		server.Send(&intapi.HealthCheckResponse{
			Response: false,
			Msg:      entity.ErrWashServerConnectionInit.Error(),
		})
		return nil
	}

	connection := entity.WashServerConnection{
		WashServer: washServer,
		Sessions:   make(map[string]*entity.Session),
	}

	s.connectionsMutex.Lock()
	s.connections[svcKey] = &connection
	s.connectionsMutex.Unlock()

	s.healthCheckConnection(server, svcKey)
	return nil
}

func (s *Service) healthCheckConnection(server intapi.ServerService_HealthCheckServer, svcKey string) {
	var err error
	for {
		err = server.Send(&intapi.HealthCheckResponse{
			Response: true,
			Msg:      "OK",
		})
		if err != nil {
			s.connectionsMutex.Lock()
			delete(s.connections, svcKey)
			s.connectionsMutex.Unlock()
			break
		}
		time.Sleep(time.Second * 10)
	}
}
