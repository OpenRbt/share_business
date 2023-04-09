package session

import (
	"context"
	"errors"
	rabbitSession "github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	rabbitVo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/entity"
)

func (u *useCase) Get(ctx context.Context, sessionID uuid.UUID, user string) (session entity.Session, err error) {
	session, err = u.SessionSvc.Get(ctx, sessionID)
	if err != nil {
		return
	}

	if session.User != nil && session.User.ID != user {
		err = entity.ErrForbidden
		return
	}

	washServer, err := u.WashSvc.GetWashServer(ctx, session.WashServer.Id)
	if err != nil {
		return
	}

	session.WashServer = washServer

	return

}

func (u *useCase) AssignUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
	_, err = u.UserSvc.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			_, err = u.UserSvc.Create(ctx, userID)
			if err != nil {
				return err
			}
		} else {
			return
		}
	}

	session, err := u.SessionSvc.Get(ctx, sessionID)
	if err != nil {
		return
	}

	if (session.User != nil && session.User.ID != userID) || session.Finished {
		return entity.ErrForbidden
	}

	err = u.SessionSvc.SetSessionUser(ctx, sessionID, userID)
	if err != nil {
		return err
	}

	msg := rabbitSession.UserAssign{
		SessionID: sessionID.String(),
		UserID:    userID,
	}
	eventErr := u.RabbitSvc.SendMessage(msg, rabbitVo.WashBonusService, rabbitVo.RoutingKey(session.WashServer.Id.String()), rabbitVo.SessionUserMessageType)
	if eventErr != nil {
		u.l.Errorw("failed to send server event", "session pool creation", "target server", session.WashServer.Id.String(), "error", eventErr)
	}

	return
}

func (u *useCase) ChargeBonuses(ctx context.Context, sessionID uuid.UUID, userID string, amount decimal.Decimal) (err error) {
	user, err := u.UserSvc.Get(ctx, userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			_, err = u.UserSvc.Create(ctx, userID)
			if err != nil {
				return err
			}
		} else {
			return
		}
	}

	session, err := u.SessionSvc.Get(ctx, sessionID)
	if err != nil {
		return
	}

	if session.User == nil || session.User.ID != user.ID {
		return entity.ErrForbidden
	}

	subtractAmount := amount.Neg()

	_, err = u.UserSvc.UpdateBalance(ctx, userID, subtractAmount)
	if err != nil {
		return
	}

	err = u.SessionSvc.UpdateSessionBalance(ctx, sessionID, amount)
	if err != nil {
		return
	}

	eventErr := u.RabbitSvc.SendMessage(conversions.SessionBonusCharge(session, amount), rabbitVo.WashBonusService, rabbitVo.RoutingKey(session.WashServer.Id.String()), rabbitVo.SessionBonusChargeMessageType)
	if eventErr != nil {
		u.l.Errorw("failed to send charge bonuses event", "session", session.ID.String(), "amount", amount.String(), "error", eventErr)
	}

	return
}
