package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/OpenRbt/share_business/wash_rabbit/entity/admin"
	"github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/wagslane/go-rabbitmq"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/infrastructure/rabbit/models"
)

func (s *Service) ProcessMessage(d rabbitmq.Delivery) (action rabbitmq.Action) {
	// TODO: use context with timeout
	ctx := context.Background()

	switch vo.MessageType(d.Type) {
	case vo.AdminServerRegisteredMessageType:
		var msg admin.ServerRegistered
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
	case vo.AdminServerUpdatedMessageType:
		var msg admin.ServerUpdate
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
	case vo.SessionRequestMessageType:
		var msg session.RequestSessions
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
	case vo.SessionStateMessageType:
		var msg session.StateChange
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

		err = s.svcSessions.UpdateSessionState(ctx, sessionID, models.SessionState(msg.State))
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case vo.SessionBonusConfirmMessageType:
		var msg session.BonusChargeConfirm
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
	case vo.SessionBonusDiscardMessageType:
		var msg session.BonusChargeDiscard
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

func (s *Service) SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) (err error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return
	}

	switch service {
	case vo.WashBonusService:
		return s.washBonusPub.Publish(
			jsonMsg,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsExchange(string(service)),
		)
	default:
		return errors.New("unknown Service")
	}
}
