package app

import (
	"context"
	"wash_admin/internal/entity"
	"wash_admin/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
)

type WashServerService interface {
	GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (entity.WashServer, error)
	AddWashServer(ctx context.Context, auth *Auth, addWashServer vo.AddWashServer) error
	UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer vo.UpdateWashServer) error
	DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error
	GetWashServerList(ctx context.Context, auth *Auth, getWashServerList vo.Pagination) ([]entity.WashServer, error)
}

type Repository interface {
	GetWashAdmin(ctx context.Context, identity string) (entity.WashAdmin, error)
	GetWashServer(ctx context.Context, ownerId uuid.UUID, id uuid.UUID) (entity.WashServer, error)
	AddWashServer(ctx context.Context, addWashServer vo.AddWashServer, ownerId uuid.UUID) error
	UpdateWashServer(ctx context.Context, updateWashServer vo.UpdateWashServer) error
	DeleteWashServer(ctx context.Context, id uuid.UUID) error
	GetWashServerList(ctx context.Context, ownerId uuid.UUID, pagination vo.Pagination) ([]entity.WashServer, error)
}

type WashServerSvc struct {
	l    *zap.SugaredLogger
	repo Repository
}

func NewWashServerService(logger *zap.SugaredLogger, repo Repository) WashServerService {
	return &WashServerSvc{
		l:    logger,
		repo: repo,
	}
}

func (wa *WashServerSvc) GetWashServer(ctx context.Context, auth *Auth, id uuid.UUID) (entity.WashServer, error) {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return entity.WashServer{}, err
	}

	return wa.repo.GetWashServer(ctx, owner.ID, id)
}

func (wa *WashServerSvc) AddWashServer(ctx context.Context, auth *Auth, addWashServer vo.AddWashServer) error {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return err
	}

	return wa.repo.AddWashServer(ctx, addWashServer, owner.ID)
}

func (wa *WashServerSvc) UpdateWashServer(ctx context.Context, auth *Auth, updateWashServer vo.UpdateWashServer) error {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return err
	}

	washServer, err := wa.repo.GetWashServer(ctx, owner.ID, updateWashServer.ID)

	if err != nil {
		return err
	}

	if washServer.Owner != owner.ID {
		return entity.ErrUserNotOwner
	}

	return wa.repo.UpdateWashServer(ctx, updateWashServer)
}

func (wa *WashServerSvc) DeleteWashServer(ctx context.Context, auth *Auth, id uuid.UUID) error {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)
	if err != nil {
		return err
	}

	washServer, err := wa.repo.GetWashServer(ctx, owner.ID, id)
	if err != nil {
		return err
	}

	if washServer.Owner != owner.ID {
		return entity.ErrUserNotOwner
	}

	return wa.repo.DeleteWashServer(ctx, id)
}

func (wa *WashServerSvc) GetWashServerList(ctx context.Context, auth *Auth, pagination vo.Pagination) ([]entity.WashServer, error) {
	owner, err := wa.repo.GetWashAdmin(ctx, auth.UID)

	if err != nil {
		return []entity.WashServer{}, err
	}

	return wa.repo.GetWashServerList(ctx, owner.ID, pagination)
}