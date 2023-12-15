package rabbit

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"washbonus/internal/conversions"
	"washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"
	"washbonus/rabbit-intapi/client/operations"
	"washbonus/rabbit-intapi/models"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"github.com/wagslane/go-rabbitmq"
)

const MessageProcessingTimeout = time.Second * 15

func (svc *Service) ProcessMessage(d rabbitmq.Delivery) (action rabbitmq.Action) {
	ctx, cancel := context.WithTimeout(context.Background(), MessageProcessingTimeout)
	defer cancel()

	switch rabbitEntities.MessageType(d.Type) {
	case rabbitEntities.SessionRequestMessageType:
		var msg rabbitEntities.RequestSessions
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

		eventErr := svc.SendMessage(pool, rabbitEntities.WashBonusService, rabbitEntities.RoutingKey(serverID.String()), rabbitEntities.SessionCreatedMessageType)
		if eventErr != nil {
			svc.l.Errorw("failed to send server event", "session pool creation", "target server", serverID.String(), "error", eventErr)
		}

		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
	case rabbitEntities.SessionStateMessageType:
		var msg rabbitEntities.ChangeSessionState
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
	case rabbitEntities.SessionBonusConfirmMessageType:
		var msg rabbitEntities.BonusChargeConfirm
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
	case rabbitEntities.SessionBonusDiscardMessageType:
		var msg rabbitEntities.BonusChargeDiscard
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
	case rabbitEntities.SessionBonusRewardMessageType:
		var msg rabbitEntities.BonusReward
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
	case rabbitEntities.SessionMoneyReportMessageType:
		var msg rabbitEntities.MoneyReport
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
	case rabbitEntities.CreateUserType:
		var msg rabbitEntities.CreateUser
		err := json.Unmarshal(d.Body, &msg)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}
		err = svc.createRabbitUser(msg.ID, msg.ServiceKey, msg.Exchange)
		if err != nil {
			action = rabbitmq.NackDiscard
			return
		}

	case rabbitEntities.RequestAdminDataMessageType:
		replyQueue := d.ReplyTo

		err := svc.sendOrganizations(ctx, replyQueue)
		if err != nil {
			return rabbitmq.NackDiscard
		}

		err = svc.sendGroups(ctx, replyQueue)
		if err != nil {
			return rabbitmq.NackDiscard
		}

		err = svc.sendAdminUsers(ctx, replyQueue)
		if err != nil {
			return rabbitmq.NackDiscard
		}

	default:
		action = rabbitmq.NackDiscard
	}

	return
}

func (svc *Service) sendOrganizations(ctx context.Context, key string) error {
	offset := int64(0)

	for {
		orgs, err := svc.orgSvc.GetAll(ctx, entities.Pagination{
			Offset: offset,
			Limit:  25,
		})
		if err != nil {
			return err
		}

		if len(orgs) == 0 {
			break
		}

		rabbitOrgs := conversions.OrganizationsToRabbit(orgs)
		for _, org := range rabbitOrgs {
			err := svc.SendMessage(org, rabbitEntities.AdminsExchange, rabbitEntities.RoutingKey(key), rabbitEntities.OrganizationMessageType)
			if err != nil {
				svc.l.Errorf("unable to send organization with id - %s to external services: %w", org.ID, err)
			}
		}

		offset += 25
	}

	return nil
}

func (svc *Service) sendGroups(ctx context.Context, key string) error {
	offset := int64(0)

	for {
		groups, err := svc.groupSvc.GetAll(ctx, entities.Pagination{
			Offset: offset,
			Limit:  25,
		})
		if err != nil {
			return err
		}

		if len(groups) == 0 {
			break
		}

		rabbitGroups := conversions.ServerGroupsToRabbit(groups)
		for _, group := range rabbitGroups {
			err := svc.SendMessage(group, rabbitEntities.AdminsExchange, rabbitEntities.RoutingKey(key), rabbitEntities.ServerGroupMessageType)
			if err != nil {
				svc.l.Errorf("unable to send group with id - %s to external services: %w", group.ID, err)
			}
		}

		offset += 25
	}

	return nil
}

func (svc *Service) sendAdminUsers(ctx context.Context, key string) error {
	offset := int64(0)

	for {
		admins, err := svc.adminSvc.GetAll(ctx, entities.Pagination{
			Offset: offset,
			Limit:  25,
		})
		if err != nil {
			return err
		}

		if len(admins) == 0 {
			break
		}

		rabbitAdmins := conversions.AdminUsersToRabbit(admins)
		for _, admin := range rabbitAdmins {
			err := svc.SendMessage(admin, rabbitEntities.AdminsExchange, rabbitEntities.RoutingKey(key), rabbitEntities.AdminUserMessageType)
			if err != nil {
				svc.l.Errorf("unable to send admin user with id - %s to external services: %w", admin.ID, err)
			}
		}

		offset += 25
	}

	return nil
}

func (svc *Service) SendMessage(msg interface{}, service rabbitEntities.Service, routingKey rabbitEntities.RoutingKey, messageType rabbitEntities.MessageType) (err error) {
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		return
	}

	var publisher *rabbitmq.Publisher

	switch service {
	case rabbitEntities.WashBonusService:
		publisher = svc.washBonusPub
	case rabbitEntities.AdminsExchange:
		publisher = svc.adminsPub
	default:
		return errors.New("unknown Service")
	}

	return publisher.Publish(
		jsonMsg,
		[]string{string(routingKey)},
		rabbitmq.WithPublishOptionsType(string(messageType)),
		rabbitmq.WithPublishOptionsExchange(string(service)),
	)
}

func (s *Service) CreateRabbitUser(userID, userKey string) error {
	return s.createRabbitUser(userID, userKey, "wash_bonus_service")
}

func (s *Service) createRabbitUser(userID, userKey, exchange string) error {
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
			Read:      fmt.Sprintf("(%s)|(%s).*", exchange, userID),
			Write:     fmt.Sprintf("(%s)|(%s).*", exchange, userID),
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

func (s *Service) DeleteRabbitUser(ctx context.Context, userID string) error {
	_, err := s.intApi.Operations.DeleteUser(&operations.DeleteUserParams{
		UserID:  userID,
		Context: ctx,
	}, s.intApiAuth)
	if err != nil {
		return err
	}

	return nil
}
