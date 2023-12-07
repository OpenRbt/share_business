package controllers

import (
	"context"
	"washbonus/internal/app"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	"washbonus/internal/infrastructure/rabbit"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"

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

func (ctrl *sessionController) GetSession(ctx context.Context, auth app.Auth, sessionID uuid.UUID) (entities.Session, error) {
	session, err := ctrl.sessionSvc.Get(ctx, sessionID, &auth.User.ID)
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

func (ctrl *sessionController) ChargeBonuses(ctx context.Context, auth app.Auth, amount decimal.Decimal, sessionID uuid.UUID) error {
	session, err := ctrl.sessionSvc.Get(ctx, sessionID, &auth.User.ID)
	if err != nil {
		return err
	}

	if session.User == nil || session.Finished {
		return entities.ErrForbidden
	}

	err = ctrl.sessionSvc.ChargeBonuses(ctx, amount, sessionID, auth.User.ID)
	if err != nil {
		return err
	}

	washServerID := session.WashServer.ID.String()
	eventErr := ctrl.rabbitSvc.SendMessage(conversions.SessionBonusCharge(sessionID, amount, session.Post), rabbitEntities.WashBonusService, rabbitEntities.RoutingKey(washServerID), rabbitEntities.SessionBonusChargeMessageType)
	if eventErr != nil {
		ctrl.logger.Errorw("failed to send charge bonuses event", "session", sessionID.String(), "amount", amount.String(), "error", eventErr)
	}

	return nil
}

func (ctrl *sessionController) AssignUserToSession(ctx context.Context, auth app.Auth, sessionID uuid.UUID) error {
	session, err := ctrl.sessionSvc.Get(ctx, sessionID, nil)
	if err != nil {
		return err
	}

	if (session.User != nil && session.User.ID != auth.User.ID) || session.Finished {
		return entities.ErrForbidden
	}

	err = ctrl.sessionSvc.SetSessionUser(ctx, sessionID, auth.User.ID)
	if err != nil {
		return err
	}

	washServerID := session.WashServer.ID.String()
	eventErr := ctrl.rabbitSvc.SendMessage(conversions.SessionUserAssign(sessionID, auth.User.ID, session.Post), rabbitEntities.WashBonusService, rabbitEntities.RoutingKey(washServerID), rabbitEntities.SessionUserMessageType)
	if eventErr != nil {
		ctrl.logger.Errorw("failed to send server event", "session pool creation", "target server", washServerID, "error", eventErr)
	}

	return nil
}
