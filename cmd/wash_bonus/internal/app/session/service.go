package session

import (
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"wash_bonus/internal/entity"
)

type Service interface {
}

type WashRepo interface {
}

type UserRepo interface {
}

type SessionsRepo interface {
}

type Cache interface {
	GetSessions(sessionID uuid.UUID) *entity.Session
	SetSessions(sessionID uuid.UUID, session entity.Session)
	RefreshSessions(sessionID uuid.UUID, session entity.Session) error
}

type service struct {
	l            *zap.SugaredLogger
	washRepo     WashRepo
	userRepo     UserRepo
	sessionsRepo SessionsRepo
}

func New(l *zap.SugaredLogger, washRepo WashRepo, userRepo UserRepo, sessionsRepo SessionsRepo) *service {
	return &service{
		l:            l,
		washRepo:     washRepo,
		userRepo:     userRepo,
		sessionsRepo: sessionsRepo,
	}
}
