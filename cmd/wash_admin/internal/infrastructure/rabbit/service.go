package rabbit

import (
	"encoding/json"
	"errors"
	"github.com/wagslane/go-rabbitmq"
	"wash_admin/internal/infrastructure/rabbit/models/vo"
)

func (s *Service) SendMessage(msg interface{}, service string, target string, messageType int) (err error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return
	}

	switch service {
	case vo.WashAdminService:
		return s.eventsPublisher.Publish(
			jsonMsg,
			[]string{target},
			rabbitmq.WithPublishOptionsType(vo.MessageType(messageType).String()),
		)
	default:
		return errors.New("unknown Service")
	}
}
