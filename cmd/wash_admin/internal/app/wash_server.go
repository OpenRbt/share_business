package app

import (
	"context"

	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type WashServerService interface {
	GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (WashServer, error)
	RegisterWashServer(ctx context.Context, auth *Auth, newServer RegisterWashServer) (WashServer, error)
	UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer UpdateWashServer) error
	DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error
	GetWashServerList(ctx context.Context, auth *Auth, getWashServerList Pagination) ([]WashServer, error)

	UpdateUser(ctx context.Context, auth *Auth, user UpdateUser) error
}

type Repository interface {
	GetOrCreateUserIfNotExists(ctx context.Context, identity string) (WashUser, error)
	GetWashUser(ctx context.Context, identity string) (WashUser, error)
	GetWashServer(ctx context.Context, id uuid.UUID) (WashServer, error)

	UpdateUserRole(ctx context.Context, updateUser UpdateUser) error

	RegisterWashServer(ctx context.Context, user uuid.UUID, newServer RegisterWashServer) (WashServer, error)
	UpdateWashServer(ctx context.Context, updateWashServer UpdateWashServer) error
	DeleteWashServer(ctx context.Context, id uuid.UUID) error
	GetWashServerList(ctx context.Context, pagination Pagination) ([]WashServer, error)
}

type WashServerSvc struct {
	l    *zap.SugaredLogger
	repo Repository

	r RabbitSvc
}

type RabbitSvc interface {
	CreateRabbitUser(userID, userKey string) (err error)
	SendMessage(msg interface{}, service rabbit_vo.Service, routingKey rabbit_vo.RoutingKey, messageType rabbit_vo.MessageType) error
}

func NewWashServerService(logger *zap.SugaredLogger, repo Repository, rabbit RabbitSvc) WashServerService {
	return &WashServerSvc{
		l:    logger,
		repo: repo,
		r:    rabbit,
	}
}

func (svc *WashServerSvc) RegisterWashServer(ctx context.Context, auth *Auth, newServer RegisterWashServer) (WashServer, error) {
	user, err := svc.repo.GetOrCreateUserIfNotExists(ctx, auth.UID)

	if err != nil {
		return WashServer{}, err
	}

	switch user.Role {
	case AdminRole:
		registered, err := svc.repo.RegisterWashServer(ctx, user.ID, newServer)
		if err != nil {
			return WashServer{}, err
		}

		err = svc.r.CreateRabbitUser(registered.ID.String(), registered.ServiceKey)
		if err != nil {
			return WashServer{}, err
		}

		eventErr := svc.r.SendMessage(WashServerToRabbit(registered), rabbit_vo.WashAdminService, rabbit_vo.WashAdminServesEventsRoutingKey, rabbit_vo.AdminServerRegisteredMessageType)
		if eventErr != nil {
			svc.l.Errorw("failed to send server event", "registered server", registered, "error", eventErr)
			return WashServer{}, eventErr
		}

		return registered, nil
	default:
		return WashServer{}, ErrAccessDenied
	}
}

func (svc *WashServerSvc) GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (WashServer, error) {
	user, err := svc.repo.GetOrCreateUserIfNotExists(ctx, auth.UID)

	if err != nil {
		return WashServer{}, err
	}

	switch user.Role {
	case AdminRole:
		return svc.repo.GetWashServer(ctx, id)
	default:
		return WashServer{}, ErrAccessDenied
	}
}

func (svc *WashServerSvc) UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer UpdateWashServer) error {
	user, err := svc.repo.GetWashUser(ctx, auth.UID)

	if err != nil {
		return err
	}

	switch user.Role {
	case AdminRole:

		_, err := svc.repo.GetWashServer(ctx, updateWashServer.ID)

		if err != nil {
			return err
		}

		err = svc.repo.UpdateWashServer(ctx, updateWashServer)
		if err != nil {
			return err
		}

		eventErr := svc.r.SendMessage(WashServerUpdateToRabbit(updateWashServer, false), rabbit_vo.WashAdminService, rabbit_vo.WashAdminServesEventsRoutingKey, rabbit_vo.AdminServerUpdatedMessageType)
		if eventErr != nil {
			svc.l.Errorw("failed to send server event", "update server", updateWashServer, "error", eventErr)
			return eventErr
		}

		return nil

	default:
		return ErrAccessDenied
	}
}

func (svc *WashServerSvc) DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error {
	user, err := svc.repo.GetWashUser(ctx, auth.UID)
	if err != nil {
		return err
	}

	switch user.Role {
	case AdminRole:
		_, err := svc.repo.GetWashServer(ctx, id)
		if err != nil {
			return err
		}

		err = svc.repo.DeleteWashServer(ctx, id)
		if err != nil {
			return err
		}

		eventErr := svc.r.SendMessage(WashServerUpdateToRabbit(UpdateWashServer{ID: id}, true), rabbit_vo.WashAdminService, rabbit_vo.WashAdminServesEventsRoutingKey, rabbit_vo.AdminServerUpdatedMessageType)
		if eventErr != nil {
			svc.l.Errorw("failed to send server event", "deleted server", id.String(), "error", eventErr)
			return eventErr
		}

		return nil

	default:
		return ErrAccessDenied
	}
}

func (svc *WashServerSvc) GetWashServerList(ctx context.Context, auth *Auth, pagination Pagination) ([]WashServer, error) {
	user, err := svc.repo.GetOrCreateUserIfNotExists(ctx, auth.UID)

	switch user.Role {
	case AdminRole:
		if err != nil {
			return []WashServer{}, err
		}

		return svc.repo.GetWashServerList(ctx, pagination)
	default:
		return []WashServer{}, ErrAccessDenied
	}
}
