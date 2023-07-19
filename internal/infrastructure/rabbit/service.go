package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"washBonus/internal/conversions"
	"washBonus/internal/infrastructure/rabbit/entity/session"
	"washBonus/internal/infrastructure/rabbit/entity/vo"
	"washBonus/rabbit-intapi/client/operations"
	"washBonus/rabbit-intapi/models"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/wagslane/go-rabbitmq"
)

func (svc *Service) ProcessMessage(d rabbitmq.Delivery) (action rabbitmq.Action) {
	// TODO: use context with timeout
	ctx := context.Background()

	switch vo.MessageType(d.Type) {
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

		pool, err := svc.rabbitSvc.CreatePool(ctx, serverID, msg.PostID, msg.NewSessionsAmount)

		eventErr := svc.SendMessage(pool, vo.WashBonusService, vo.RoutingKey(serverID.String()), vo.SessionCreatedMessageType)
		if eventErr != nil {
			svc.l.Errorw("failed to send server event", "session pool creation", "target server", serverID.String(), "error", eventErr)
		}

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

		err = svc.rabbitSvc.UpdateState(ctx, sessionID, msg.State)
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

		amount := decimal.NewFromInt(msg.Amount)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = svc.rabbitSvc.ConfirmBonuses(ctx, sessionID, amount)
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

		amount := decimal.NewFromInt(msg.Amount)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = svc.rabbitSvc.DiscardBonuses(ctx, sessionID, amount)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case vo.SessionBonusRewardMessageType:
		var msg session.BonusReward
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

		amount := decimal.NewFromInt(int64(msg.Amount))
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		messageUuid, err := uuid.FromString(msg.UUID)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		err = svc.rabbitSvc.RewardBonuses(ctx, d.Body, sessionID, amount, messageUuid)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case vo.SessionMoneyReportMessageType:
		var msg session.MoneyReport
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

		report, _ := conversions.MoneyReportFromRabbit(msg)

		err = svc.rabbitSvc.SaveMoneyReport(ctx, report)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	default:
		action = rabbitmq.NackDiscard
	}

	return
}

func (svc *Service) SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) (err error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return
	}

	switch service {
	case vo.WashBonusService:
		return svc.washBonusPub.Publish(
			jsonMsg,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsExchange(string(service)),
		)
	default:
		return errors.New("unknown Service")
	}
}

func (s *Service) CreateRabbitUser(userID, userKey string) error {
	ctx := context.TODO()

	tags := ""
	vhost := "/"

	_, _, err := s.intApi.Operations.CreateUser(&operations.CreateUserParams{
		Body: &models.CreateUser{
			Password: &userKey,
			Tags:     &tags,
		},
		UserID:     userID,
		Context:    ctx,
		HTTPClient: nil,
	}, s.intApiAuth)
	if err != nil {
		return err
	}

	_, _, err = s.intApi.Operations.SetUserPerms(&operations.SetUserPermsParams{
		Body: &models.ManagePermissions{
			Configure: fmt.Sprintf("%s.*", userID),
			Read:      fmt.Sprintf("(wash_bonus_service)|(%s).*", userID),
			Write:     fmt.Sprintf("(wash_bonus_service)|(%s).*", userID),
		},
		UserID:     userID,
		Vhost:      vhost,
		Context:    ctx,
		HTTPClient: nil,
	}, s.intApiAuth)
	if err != nil {
		return err
	}

	return nil
}
