package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/wagslane/go-rabbitmq"
	"wash_bonus/internal/conversions"
	"wash_bonus/internal/infrastructure/rabbit/models"
	"wash_bonus/internal/infrastructure/rabbit/models/vo"
)

func (s *Service) ProcessMessage(d rabbitmq.Delivery) (action rabbitmq.Action) {
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

		_, err = s.svcWashServer.CreateWashServer(context.Background(), created)
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

		err = s.svcWashServer.UpdateWashServer(context.Background(), update)
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
		)
	default:
		return errors.New("unknown Service")
	}
}
