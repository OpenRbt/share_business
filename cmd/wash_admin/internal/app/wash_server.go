package app

import (
	"context"
	"wash_admin/internal/conversions"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"
	vo2 "wash_admin/internal/infrastructure/rabbit/models/vo"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type WashServerService interface {
	GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (entity.WashServer, error)
	RegisterWashServer(ctx context.Context, auth *Auth, newServer vo.RegisterWashServer) (entity.WashServer, error)
	UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer vo.UpdateWashServer) error
	DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error
	GetWashServerList(ctx context.Context, auth *Auth, getWashServerList vo.Pagination) ([]entity.WashServer, error)
}

type Repository interface {
	GetOrCreateAdminIfNotExists(ctx context.Context, identity string) (entity.WashAdmin, error)
	GetWashAdmin(ctx context.Context, identity string) (entity.WashAdmin, error)
	GetWashServer(ctx context.Context, ownerId uuid.UUID, id uuid.UUID) (entity.WashServer, error)

	RegisterWashServer(ctx context.Context, owner uuid.UUID, newServer vo.RegisterWashServer) (entity.WashServer, error)
	UpdateWashServer(ctx context.Context, updateWashServer vo.UpdateWashServer) error
	DeleteWashServer(ctx context.Context, id uuid.UUID) error
	GetWashServerList(ctx context.Context, ownerId uuid.UUID, pagination vo.Pagination) ([]entity.WashServer, error)
}

type WashServerSvc struct {
	l    *zap.SugaredLogger
	repo Repository

	r RabbitSvc
}

type RabbitSvc interface {
	CreateRabbitUser(userID, userKey string) (err error)
	SendMessage(msg interface{}, service string, target string, messageType int) error
}

func NewWashServerService(logger *zap.SugaredLogger, repo Repository, rabbit RabbitSvc) WashServerService {
	return &WashServerSvc{
		l:    logger,
		repo: repo,
		r:    rabbit,
	}
}

func (svc *WashServerSvc) RegisterWashServer(ctx context.Context, auth *Auth, newServer vo.RegisterWashServer) (entity.WashServer, error) {
	owner, err := svc.repo.GetOrCreateAdminIfNotExists(ctx, auth.UID)

	if err != nil {
		return entity.WashServer{}, err
	}

	registered, err := svc.repo.RegisterWashServer(ctx, owner.ID, newServer)
	if err != nil {
		return entity.WashServer{}, err
	}

	err = svc.r.CreateRabbitUser(registered.ID.String(), registered.ServiceKey)
	if err != nil {
		return entity.WashServer{}, err
	}

	eventErr := svc.r.SendMessage(conversions.WashServerToRabbit(registered), vo2.WashAdminService, vo2.WashAdminServers, int(vo2.WashAdminServerRegistered))
	if eventErr != nil {
		svc.l.Errorw("failed to send server event", "registered server", registered, "error", eventErr)
	}

	return registered, nil
}

func (svc *WashServerSvc) GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (entity.WashServer, error) {
	owner, err := svc.repo.GetOrCreateAdminIfNotExists(ctx, auth.UID)

	if err != nil {
		return entity.WashServer{}, err
	}

	return svc.repo.GetWashServer(ctx, owner.ID, id)
}

func (svc *WashServerSvc) UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer vo.UpdateWashServer) error {
	owner, err := svc.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return err
	}

	washServer, err := svc.repo.GetWashServer(ctx, owner.ID, updateWashServer.ID)

	if err != nil {
		return err
	}

	if washServer.Owner != owner.ID {
		return entity.ErrUserNotOwner
	}

	err = svc.repo.UpdateWashServer(ctx, updateWashServer)
	if err != nil {
		return err
	}

	eventErr := svc.r.SendMessage(conversions.WashServerUpdateToRabbit(updateWashServer, false), vo2.WashAdminService, vo2.WashAdminServers, int(vo2.WashAdminServerUpdated))
	if eventErr != nil {
		svc.l.Errorw("failed to send server event", "update server", updateWashServer, "error", eventErr)
	}

	return nil
}

func (svc *WashServerSvc) DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error {
	owner, err := svc.repo.GetWashAdmin(ctx, auth.UID)
	if err != nil {
		return err
	}

	washServer, err := svc.repo.GetWashServer(ctx, owner.ID, id)
	if err != nil {
		return err
	}

	if washServer.Owner != owner.ID {
		return entity.ErrUserNotOwner
	}
	err = svc.repo.DeleteWashServer(ctx, id)
	if err != nil {
		return err
	}

	eventErr := svc.r.SendMessage(conversions.WashServerUpdateToRabbit(vo.UpdateWashServer{ID: id}, true), vo2.WashAdminService, vo2.WashAdminServers, int(vo2.WashAdminServerUpdated))
	if eventErr != nil {
		svc.l.Errorw("failed to send server event", "deleted server", id.String(), "error", eventErr)
	}

	return nil
}

func (svc *WashServerSvc) GetWashServerList(ctx context.Context, auth *Auth, pagination vo.Pagination) ([]entity.WashServer, error) {
	owner, err := svc.repo.GetOrCreateAdminIfNotExists(ctx, auth.UID)

	if err != nil {
		return []entity.WashServer{}, err
	}

	return svc.repo.GetWashServerList(ctx, owner.ID, pagination)
}
