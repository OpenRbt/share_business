package rabbit

import (
	"context"
	"errors"
	"fmt"
	"wash_admin/internal/infrastructure/rabbit-intapi/client/operations"
	"wash_admin/internal/infrastructure/rabbit-intapi/models"

	rabbit_vo "github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"github.com/wagslane/go-rabbitmq"
)

func (s *Service) SendMessage(msg interface{}, service rabbit_vo.Service, routingKey rabbit_vo.RoutingKey, messageType rabbit_vo.MessageType) (err error) {
	bytes, ok := msg.([]byte)
	if !ok {
		return errors.New("wrong message body")
	}

	switch service {
	case rabbit_vo.WashAdminService:
		return s.eventsPublisher.Publish(
			bytes,
			[]string{string(routingKey)},
			rabbitmq.WithPublishOptionsType(string(messageType)),
			rabbitmq.WithPublishOptionsExchange(string(service)),
		)
	default:
		return errors.New("unknown Service")
	}
}

func (s *Service) CreateRabbitUser(userID, userKey string) (err error) {
	ctx := context.TODO()

	tags := ""
	vhost := "/"

	_, _, err = s.intApi.Operations.CreateUser(&operations.CreateUserParams{
		Body: &models.CreateUser{
			Password: &userKey,
			Tags:     &tags,
		},
		UserID:     userID,
		Context:    ctx,
		HTTPClient: nil,
	}, s.intApiAuth)

	if err != nil {
		return
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

	return
}
