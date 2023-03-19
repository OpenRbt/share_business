package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/infrastructure/rabbit/models"
	"wash_bonus/internal/infrastructure/rabbit/models/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/wagslane/go-rabbitmq"
)

func (s *Service) ProcessMessage(d rabbitmq.Delivery) (action rabbitmq.Action) {
	// TODO: use context with timeout
	ctx := context.Background()

	messageType := vo.MessageTypeFromString(d.Type)
	switch messageType {
	case vo.WashAdminServerRegistered:
		var msg models.ServerRegistered
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		created, err := conversions.WashServerCreationFromRabbit(msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		_, err = s.svcWashServer.CreateWashServer(ctx, created)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case vo.WashAdminServerUpdated:
		var msg models.ServerUpdate
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		update, err := conversions.WashServerUpdateFromRabbit(msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = s.svcWashServer.UpdateWashServer(ctx, update)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case vo.BonusSessionRequest:
		var msg models.SessionRequest
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
		serverID, err := uuid.FromString(d.UserId)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		_, err = s.svcSessions.CreateSessionPool(ctx, serverID, msg.PostID, msg.NewSessionsAmount)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case vo.BonusSessionStateChange:
		var msg models.SessionStateChange
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		sessionID, err := uuid.FromString(msg.SessionID)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = s.svcSessions.UpdateSessionState(ctx, sessionID, msg.State)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

	case vo.BonusSessionBonusConfirm:
		var msg models.SessionBonusChargeConfirm
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		sessionID, err := uuid.FromString(msg.SessionID)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = s.svcSessions.ConfirmBonuses(ctx, sessionID, decimal.NewFromInt(msg.Amount))
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

	case vo.BonusSessionBonusDiscard:
		var msg models.SessionBonusChargeDiscard
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		sessionID, err := uuid.FromString(msg.SessionID)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = s.svcSessions.DiscardBonuses(ctx, sessionID, decimal.NewFromInt(msg.Amount))
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	default:
		action = rabbitmq.NackDiscard
	}

	return
}

func (s *Service) SendMessage(msg interface{}, service string, target string, messageType int) (err error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return
	}

	switch service {
	case vo.WashBonusService:
		return s.washBonusPub.Publish(
			jsonMsg,
			[]string{target},
			rabbitmq.WithPublishOptionsType(vo.MessageType(messageType).String()),
			rabbitmq.WithPublishOptionsExchange(vo.WashBonusService),
		)
	default:
		return errors.New("unknown Service")
	}
}
