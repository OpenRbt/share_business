package rabbit

import (
	"context"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"
	"washBonus/internal/infrastructure/rabbit/entity/session"
	rabbitVo "washBonus/internal/infrastructure/rabbit/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *rabbitService) CreatePool(ctx context.Context, serverID uuid.UUID, postId int64, amount int64) (session.PostSessions, error) {
	sessions := session.PostSessions{
		NewSessions: make([]string, amount),
		PostID:      postId,
	}

	for i := int64(0); i < amount; i++ {
		session, err2 := s.sessionSvc.Create(ctx, serverID, postId)
		if err2 != nil {
			err := err2
			return sessions, err
		}

		sessions.NewSessions[i] = session.ID.String()
	}

	return sessions, nil
}

func (s *rabbitService) UpdateState(ctx context.Context, sessionID uuid.UUID, state rabbitVo.SessionState) error {
	return s.sessionSvc.UpdateSessionState(ctx, sessionID, vo.SessionState(state))
}

func (s *rabbitService) ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) error {
	session, err := s.sessionSvc.Get(ctx, sessionID, nil)
	if err != nil {
		return err
	}

	if session.User == nil {
		return entity.ErrSessionNoUser
	}

	return s.sessionSvc.ConfirmBonuses(ctx, amount, sessionID)
}

func (s *rabbitService) DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) error {
	session, err := s.sessionSvc.Get(ctx, sessionID, nil)
	if err != nil {
		return err
	}

	if session.User == nil {
		return entity.ErrSessionNoUser
	}

	return s.sessionSvc.DiscardBonuses(ctx, amount, sessionID)
}

func (s *rabbitService) RewardBonuses(ctx context.Context, payload []byte, sessionID uuid.UUID, amount decimal.Decimal, messageUUID uuid.UUID) error {
	session, err := s.sessionSvc.Get(ctx, sessionID, nil)
	if err != nil {
		return err
	}

	if session.User == nil {
		return entity.ErrSessionNoUser
	}

	if err := s.sessionSvc.LogRewardBonuses(ctx, sessionID, payload, messageUUID); err != nil {
		return entity.ErrMessageDuplicate
	}

	return s.userSvc.AddBonuses(ctx, amount, session.User.ID)
}

func (s *rabbitService) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) error {
	return s.sessionSvc.SaveMoneyReport(ctx, report)
}
