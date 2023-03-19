package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"wash_admin/internal/infrastructure/rabbit-intapi/client/operations"
	"wash_admin/internal/infrastructure/rabbit-intapi/models"
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
