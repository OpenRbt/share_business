package session

import (
	"context"
	"wash_bonus/internal/app"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/infrastructure/rabbit/models"

	rabbit_session "github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *service) AssignRabbit(handler func(msg interface{}, service rabbit_vo.Service, target rabbit_vo.RoutingKey, messageType rabbit_vo.MessageType) error) {
	s.rabbitPublisherFunc = handler
}

func (s *service) CreateSession(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error) {
	session, err = s.sessionRepo.CreateSession(ctx, serverID, postID)
	if err != nil {
		return
	}

	eventErr := s.rabbitPublisherFunc(conversions.SessionToRabbit(session), rabbit_vo.WashBonusService, rabbit_vo.RoutingKey(serverID.String()), rabbit_vo.SessionCreatedMessageType)
	if eventErr != nil {
		s.l.Errorw("failed to send server event", "created session", session, "target server", serverID.String(), "error", eventErr)
	}

	return
}

func (s *service) CreateSessionPool(ctx context.Context, serverID uuid.UUID, postID int64, sessionsAmount int64) (postSessions models.SessionCreation, err error) {
	postSessions = models.SessionCreation{
		NewSessions: make([]string, sessionsAmount),
		PostID:      postID,
	}

	for i := int64(0); i < sessionsAmount; i++ {
		session, err := s.sessionRepo.CreateSession(ctx, serverID, postID)
		if err != nil {
			s.l.Errorw("failed to create session", "server", serverID, "post", postID, "session#", i, "total sessions requested", sessionsAmount)
			break
		}

		postSessions.NewSessions[i] = session.ID.String()
	}

	eventErr := s.rabbitPublisherFunc(postSessions, rabbit_vo.WashBonusService, rabbit_vo.RoutingKey(serverID.String()), rabbit_vo.SessionCreatedMessageType)
	if eventErr != nil {
		s.l.Errorw("failed to send server event", "session pool creation", "target server", serverID.String(), "error", eventErr)
	}

	return
}

func (s *service) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state models.SessionState) error {
	return s.sessionRepo.UpdateSessionState(ctx, sessionID, state)
}

func (s *service) GetSession(ctx context.Context, sessionID uuid.UUID) (entity.Session, error) {
	session, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return entity.Session{}, err
	}

	washServer, err := s.washRepo.GetWashServer(ctx, session.WashServer.Id)
	if err != nil {
		return entity.Session{}, err
	}

	session.WashServer = washServer

	return session, nil
}

func (s *service) GetUserSession(ctx context.Context, auth *app.Auth, sessionID uuid.UUID) (session entity.Session, err error) {
	_, err = s.userRepo.GetByID(ctx, auth.UID)
	if err != nil {
		return
	}

	session, err = s.GetSession(ctx, sessionID)
	return
}

func (s *service) AssignSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
	user, err := s.userRepo.GetByID(ctx, userID)
	if err != nil {
		user, err = s.userRepo.Create(ctx, userID)
		if err != nil {
			return err
		}
	}

	session, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	if session.User != nil || session.Finished {
		return entity.ErrForbidden
	}

	assignUser := rabbit_session.UserAssign{
		SessionID: sessionID.String(),
		UserID:    userID,
	}

	eventErr := s.rabbitPublisherFunc(assignUser, vo.WashBonusService, rabbit_vo.RoutingKey(session.WashServer.Id.String()), rabbit_vo.SessionUserMessageType)
	if eventErr != nil {
		s.l.Errorw("failed to send server event", "session pool creation", "target server", session.WashServer.Id.String(), "error", eventErr)
	}

	return s.sessionRepo.SetSessionUser(ctx, sessionID, user.ID)
}

func (s *service) ChargeBonuses(ctx context.Context, sessionID uuid.UUID, userID string, amount decimal.Decimal) (err error) {
	user, err := s.userRepo.GetByID(ctx, userID)
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

	eventErr := s.rabbitPublisherFunc(conversions.SessionBonusCharge(session, amount), rabbit_vo.WashBonusService, rabbit_vo.RoutingKey(session.WashServer.Id.String()), rabbit_vo.SessionBonusChargeMessageType)
	if eventErr != nil {
		s.l.Errorw("failed to send charge bonuses event", "session", session.ID.String(), "amount", amount.String(), "error", eventErr)
	}

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

	err = s.userRepo.UpdateBalance(ctx, session.User.ID, amount)

	return
}

func (s *service) RewardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	session, err := s.GetSession(ctx, sessionID)
	if err != nil {
		return err
	}

	if session.User != nil {
		err = s.sessionRepo.UpdateSessionBalance(ctx, sessionID, amount.Neg())
		if err != nil {
			return
		}
		err = s.userRepo.UpdateBalance(ctx, session.User.ID, amount)
	}
	return
}
