package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/app"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/infrastructure/rabbit/models/vo"
)

func (s *service) AssignRabbit(handler func(msg interface{}, service string, target string, messageType int) error) {
	s.rabbitPublisherFunc = handler
}

func (s *service) CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error) {
	session, err = s.sessionRepo.CreateSession(ctx, serverID)
	if err != nil {
		return
	}
	eventErr := s.rabbitPublisherFunc(conversions.SessionToRabbit(session), vo.WashBonusService, serverID.String(), int(vo.BonusSessionCreated))
	if eventErr != nil {
		s.l.Errorw("failed to send server event", "created session", session, "target server", serverID.String(), "error", eventErr)
	}
	return
}

func (s *service) GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error) {
	return s.sessionRepo.GetSession(ctx, sessionID)
}

func (s *service) GetUserSession(ctx context.Context, auth *app.Auth, sessionID uuid.UUID) (session entity.Session, err error) {
	user, err := s.userRepo.Get(ctx, auth.UID)
	if err != nil {
		return
	}

	session, err = s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}

	switch {
	case session.User != nil && session.User.Identity != user.Identity:
		fallthrough
	case session.Finished:
		err = entity.ErrForbidden
		return
	}

	assignErr := s.AssignSessionUser(ctx, session.WashServer.Id, sessionID, user)
	if assignErr != nil {
		s.l.Errorw("failed to assign user to session event", "session", session, "target user", user, "error", assignErr)
	}

	return
}

func (s *service) AssignSessionUser(ctx context.Context, serverID uuid.UUID, sessionID uuid.UUID, user entity.User) (err error) {
	err = s.sessionRepo.SetSessionUser(ctx, sessionID, user.ID)
	if err != nil {
		return
	}

	eventErr := s.rabbitPublisherFunc(conversions.SessionUserAssign(sessionID, user), vo.WashBonusService, serverID.String(), int(vo.BonusSessionUserAssign))
	if eventErr != nil {
		s.l.Errorw("failed to send server event", "assign session user for session", sessionID.String(), "target user", user, "error", eventErr)
	}

	return
}

func (s *service) ChargeBonuses(ctx context.Context, auth *app.Auth, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	user, err := s.userRepo.Get(ctx, auth.UID)
	if err != nil {
		return
	}

	session, err := s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}

	if session.User.ID != user.ID {
		err = entity.ErrForbidden
		return
	}

	subtractAmount := amount.Neg()

	err = s.userRepo.UpdateBalance(ctx, user.ID, subtractAmount)
	if err != nil {
		return
	}

	err = s.sessionRepo.UpdateSessionBalance(ctx, sessionID, amount)

	return
}

func (s *service) ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	subtractAmount := amount.Neg()

	err = s.sessionRepo.UpdateSessionBalance(ctx, sessionID, subtractAmount)

	return
}

func (s *service) DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	subtractAmount := amount.Neg()

	session, err := s.GetSession(ctx, sessionID)
	if err != nil {
		return
	}

	err = s.sessionRepo.UpdateSessionBalance(ctx, sessionID, subtractAmount)
	if err != nil {
		return
	}

	err = s.userRepo.UpdateBalance(ctx, session.User.ID, subtractAmount)

	return
}
