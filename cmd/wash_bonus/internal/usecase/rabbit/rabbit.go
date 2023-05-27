package rabbit

import (
	"context"
	rabbitSession "github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	rabbitVo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
)

func (u *useCase) CreatePool(ctx context.Context, serverID uuid.UUID, postId int64, amount int64) (sessions rabbitSession.PostSessions, err error) {
	sessions = rabbitSession.PostSessions{
		NewSessions: make([]string, amount),
		PostID:      postId,
	}

	for i := int64(0); i < amount; i++ {
		session, err2 := u.SessionSvc.Create(ctx, serverID, postId)
		if err2 != nil {
			err = err2
			return
		}

		sessions.NewSessions[i] = session.ID.String()
	}

	return
}

func (u *useCase) UpdateState(ctx context.Context, sessionID uuid.UUID, state rabbitVo.SessionState) error {
	return u.SessionSvc.UpdateSessionState(ctx, sessionID, vo.SessionState(state))
}

func (u *useCase) ConfirmBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	session, err := u.SessionSvc.Get(ctx, sessionID)
	if err != nil {
		return
	}

	if session.User == nil {
		return entity.ErrSessionNoUser
	}

	return u.SessionSvc.ConfirmBonuses(ctx, amount, sessionID)
}

func (u *useCase) DiscardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	session, err := u.SessionSvc.Get(ctx, sessionID)
	if err != nil {
		return
	}

	if session.User == nil {
		return entity.ErrSessionNoUser
	}

	return u.SessionSvc.DiscardBonuses(ctx, amount, sessionID)
}

func (u *useCase) RewardBonuses(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	session, err := u.SessionSvc.Get(ctx, sessionID)
	if err != nil {
		return
	}

	if session.User == nil {
		return entity.ErrSessionNoUser
	}

	return u.UserSvc.AddBonuses(ctx, amount, session.User.ID)
}

func (u *useCase) CreateWashServer(ctx context.Context, server entity.WashServer) (entity.WashServer, error) {
	return u.WashSvc.CreateWashServer(ctx, server)
}

func (u *useCase) UpdateWashServer(ctx context.Context, update vo.WashServerUpdate) error {
	return u.WashSvc.UpdateWashServer(ctx, update)
}

func (u *useCase) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error) {
	return u.SessionSvc.SaveMoneyReport(ctx, report)
}
