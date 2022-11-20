package app

import (
	"context"
	"go.uber.org/zap"
	"wash_bonus/internal/entity"
)

type UserService interface {
	GetProfile(ctx context.Context, auth *Auth) (entity.User, error)
}

type Repository interface {
	GetProfileOrCreateIfNotExists(ctx context.Context, identity string) (entity.User, error)
}

type UserSvc struct {
	l    *zap.SugaredLogger
	repo Repository
}

func NewUserService(logger *zap.SugaredLogger, repo Repository) *UserSvc {
	return &UserSvc{
		l:    logger,
		repo: repo,
	}
}

func (u *UserSvc) GetProfile(ctx context.Context, auth *Auth) (entity.User, error) {
	res, err := u.repo.GetProfileOrCreateIfNotExists(nil, auth.UID)
	if err != nil {
		u.l.Named("user").Errorw("failed to get user", "firebase_identity", auth.UID)
		return entity.User{}, err
	}

	return res, nil
}
