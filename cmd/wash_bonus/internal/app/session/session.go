package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/app"
	"wash_bonus/internal/entity"
)

func (s *service) CreateSession(ctx context.Context, connectionID uuid.UUID, postID int64, washKey string) (session entity.Session, err error) {
	washServer, err := s.washRepo.GetWashServerByKey(ctx, washKey)
	if err != nil {
		return
	}

	session = entity.Session{
		ID:           uuid.NewV4(),
		ConnectionID: connectionID,
		User:         nil,
		Post:         postID,
		WashServer:   washServer,
		PostBalance:  decimal.Decimal{},
	}

	s.cache.SetSession(session)

	return
}

func (s *service) GetUserSession(ctx context.Context, auth app.Auth, sessionID uuid.UUID) (session entity.Session, err error) {
	cacheSession := s.cache.GetSession(sessionID)
	if cacheSession == nil {
		err = entity.ErrNotFound
		return
	}

	if cacheSession.User != nil {
		if cacheSession.User.Identity != auth.UID {
			err = entity.ErrForbidden
			return
		}
	}

	err = s.AssignUser(ctx, auth, sessionID)
	if err != nil {
		return
	}

	session = *cacheSession

	return
}

func (s *service) GetSession(ctx context.Context, sessionID uuid.UUID) (session entity.Session, err error) {
	cacheSession := s.cache.GetSession(sessionID)
	if cacheSession == nil {
		err = entity.ErrNotFound
		return
	}

	return *cacheSession, nil
}

func (s *service) AssignUser(ctx context.Context, auth app.Auth, sessionID uuid.UUID) (err error) {
	user, err := s.userRepo.Get(ctx, auth.UID)
	if err != nil {
		return
	}

	session, err := s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}

	if session.Closed {
		err = entity.ErrForbidden
		return
	}

	if session.User != nil {
		if *session.User != user {
			err = entity.ErrForbidden
			return
		}
	}

	err = s.cache.RefreshSession(session)
	if err != nil {
		return
	}

	session.User = &user

	return
}

func (s *service) RefreshSession(ctx context.Context, sessionID uuid.UUID, PostBalance decimal.Decimal) (session entity.Session, err error) {
	session, err = s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}
	if session.Closed {
		err = entity.ErrForbidden
		return
	}

	session.PostBalance = PostBalance

	err = s.cache.RefreshSession(session)

	return
}

func (s *service) EndSession(ctx context.Context, sessionID uuid.UUID) (err error) {
	session, err := s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}

	if session.Closed {
		err = entity.ErrForbidden
		return
	}

	session.Closed = true

	err = s.cache.RefreshSession(session)

	return
}

func (s *service) ConsumeMoney(ctx context.Context, sessionID uuid.UUID) (err error) {
	session, err := s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}

	if session.Closed {
		err = entity.ErrForbidden
		return
	}

	session.AddAmount = decimal.Zero

	err = s.cache.RefreshSession(session)

	return
}
