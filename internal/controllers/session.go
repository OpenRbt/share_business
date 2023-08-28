package controllers

import (
	"context"
	"washBonus/internal/app"
	"washBonus/internal/conversions"
	"washBonus/internal/entity"
	"washBonus/internal/infrastructure/rabbit"
	rabbitVo "washBonus/internal/infrastructure/rabbit/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type sessionController struct {
	logger     *zap.SugaredLogger
	sessionSvc app.SessionService
	userSvc    app.UserService
	washSvc    app.WashServerService
	rabbitSvc  rabbit.RabbitService
}

func NewSessionController(l *zap.SugaredLogger, sessionSvc app.SessionService, userSvc app.UserService, washSvc app.WashServerService, rabbitSvc rabbit.RabbitService) app.SessionController {
	return &sessionController{
		logger:     l,
		sessionSvc: sessionSvc,
		userSvc:    userSvc,
		washSvc:    washSvc,
		rabbitSvc:  rabbitSvc,
	}
}

func (ctrl *sessionController) GetSession(ctx context.Context, sessionID uuid.UUID, userID string) (entity.Session, error) {
	session, err := ctrl.sessionSvc.Get(ctx, sessionID, &userID)
	if err != nil {
		return session, err
	}

	washServer, err := ctrl.washSvc.GetWashServerById(ctx, session.WashServer.ID)
	if err != nil {
		return session, err
	}

	session.WashServer = washServer
	return session, nil
}

func (ctrl *sessionController) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, authUser entity.User) error {
	session, err := ctrl.sessionSvc.Get(ctx, sessionID, &authUser.ID)
	if err != nil {
		return err
	}

	if session.User == nil || session.Finished {
		return entity.ErrForbidden
	}

	err = ctrl.sessionSvc.ChargeBonuses(ctx, amount, sessionID, authUser.ID)
	if err != nil {
		return err
	}

	washServerID := session.WashServer.ID.String()
	eventErr := ctrl.rabbitSvc.SendMessage(conversions.SessionBonusCharge(sessionID, amount, session.Post), rabbitVo.WashBonusService, rabbitVo.RoutingKey(washServerID), rabbitVo.SessionBonusChargeMessageType)
	if eventErr != nil {
		ctrl.logger.Errorw("failed to send charge bonuses event", "session", sessionID.String(), "amount", amount.String(), "error", eventErr)
	}

	return nil
}

func (ctrl *sessionController) AssignUserToSession(ctx context.Context, sessionID uuid.UUID, authUser entity.User) error {
	session, err := ctrl.sessionSvc.Get(ctx, sessionID, nil)
	if err != nil {
		return err
	}

	if (session.User != nil && session.User.ID != authUser.ID) || session.Finished {
		return entity.ErrForbidden
	}

	err = ctrl.sessionSvc.SetSessionUser(ctx, sessionID, authUser.ID)
	if err != nil {
		return err
	}

	washServerID := session.WashServer.ID.String()
	eventErr := ctrl.rabbitSvc.SendMessage(conversions.SessionUserAssign(sessionID, authUser.ID, session.Post), rabbitVo.WashBonusService, rabbitVo.RoutingKey(washServerID), rabbitVo.SessionUserMessageType)
	if eventErr != nil {
		ctrl.logger.Errorw("failed to send server event", "session pool creation", "target server", washServerID, "error", eventErr)
	}

	return nil
}
