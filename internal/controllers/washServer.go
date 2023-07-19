package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/entity"
	"washBonus/internal/infrastructure/rabbit"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type washServerController struct {
	logger        *zap.SugaredLogger
	washServerSvc app.WashServerService
	userSvc       app.UserService
	rabbitSvc     rabbit.RabbitService
}

func NewWashServerController(l *zap.SugaredLogger, washServerSvc app.WashServerService, userSvc app.UserService, rabbitSvc rabbit.RabbitService) app.WashServerController {
	return &washServerController{
		logger:        l,
		washServerSvc: washServerSvc,
		userSvc:       userSvc,
		rabbitSvc:     rabbitSvc,
	}
}

func (ctrl *washServerController) CreateWashServer(ctx context.Context, userID string, newServer entity.CreateWashServer) (entity.WashServer, error) {
	user, err := ctrl.userSvc.GetOrCreate(ctx, userID)
	if err != nil {
		return entity.WashServer{}, err
	}

	switch user.Role {
	case entity.AdminRole:
		registered, err := ctrl.washServerSvc.CreateWashServer(ctx, user.ID, newServer)
		if err != nil {
			return entity.WashServer{}, err
		}

		err = ctrl.rabbitSvc.CreateRabbitUser(registered.ID.String(), registered.ServiceKey)
		if err != nil {
			return entity.WashServer{}, err
		}

		return registered, nil
	default:
		return entity.WashServer{}, entity.ErrAccessDenied
	}
}

func (ctrl *washServerController) GetWashServerById(ctx context.Context, userID string, id uuid.UUID) (entity.WashServer, error) {
	user, err := ctrl.userSvc.GetOrCreate(ctx, userID)
	if err != nil {
		return entity.WashServer{}, err
	}

	switch user.Role {
	case entity.AdminRole:
		return ctrl.washServerSvc.GetWashServerById(ctx, id)
	default:
		return entity.WashServer{}, entity.ErrAccessDenied
	}
}

func (ctrl *washServerController) UpdateWashServer(ctx context.Context, userID string, id uuid.UUID, updateWashServer entity.UpdateWashServer) (entity.WashServer, error) {
	user, err := ctrl.userSvc.GetOrCreate(ctx, userID)
	if err != nil {
		return entity.WashServer{}, err
	}

	switch user.Role {
	case entity.AdminRole:
		server, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
		if err != nil {
			return entity.WashServer{}, err
		}

		updatedServer, err := ctrl.washServerSvc.UpdateWashServer(ctx, server.ID, updateWashServer)
		if err != nil {
			return entity.WashServer{}, err
		}

		return updatedServer, nil

	default:
		return entity.WashServer{}, entity.ErrAccessDenied
	}
}

func (ctrl *washServerController) DeleteWashServer(ctx context.Context, userID string, id uuid.UUID) error {
	user, err := ctrl.userSvc.GetOrCreate(ctx, userID)
	if err != nil {
		return err
	}

	switch user.Role {
	case entity.AdminRole:
		_, err := ctrl.washServerSvc.GetWashServerById(ctx, id)
		if err != nil {
			return err
		}

		err = ctrl.washServerSvc.DeleteWashServer(ctx, id)
		if err != nil {
			return err
		}

		return nil
	default:
		return entity.ErrAccessDenied
	}
}

func (ctrl *washServerController) GetWashServers(ctx context.Context, userID string, pagination entity.Pagination) ([]entity.WashServer, error) {
	user, err := ctrl.userSvc.GetOrCreate(ctx, userID)
	if err != nil {
		return []entity.WashServer{}, err
	}

	switch user.Role {
	case entity.AdminRole:
		return ctrl.washServerSvc.GetWashServers(ctx, pagination)
	default:
		return []entity.WashServer{}, entity.ErrAccessDenied
	}
}
