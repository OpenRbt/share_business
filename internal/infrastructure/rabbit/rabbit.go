package rabbit

import (
	"context"
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/config"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"
	"washbonus/rabbit-intapi/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitService interface {
	SendMessage(msg interface{}, service rabbitEntities.Service, routingKey rabbitEntities.RoutingKey, messageType rabbitEntities.MessageType) (err error)
	CreateRabbitUser(userID, userKey string) error
	DeleteRabbitUser(ctx context.Context, userID string) error
}

type Service struct {
	l         *zap.SugaredLogger
	conn      *amqp.Connection
	rabbitSvc app.RabbitService

	orgSvc   app.OrganizationService
	groupSvc app.ServerGroupService
	adminSvc app.AdminService

	washBonusPub *rabbitmq.Publisher
	washBonusSub *rabbitmq.Consumer

	adminsPub *rabbitmq.Publisher

	intApi     *client.RabbitIntAPI
	intApiAuth runtime.ClientAuthInfoWriter
}

func New(l *zap.SugaredLogger, cfg config.RabbitMQConfig, rabbitSvc app.RabbitService, services app.Services) (*Service, error) {
	svc := &Service{
		l:         l,
		rabbitSvc: rabbitSvc,

		orgSvc:   services.Org,
		groupSvc: services.Group,
		adminSvc: services.Admin,
	}

	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.User, cfg.Password, cfg.Url, cfg.Port)
	rabbitConf := rabbitmq.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: cfg.User,
				Password: cfg.Password,
			},
		},
		Vhost:      "/",
		ChannelMax: 0,
		FrameSize:  0,
		Heartbeat:  0,
		Properties: nil,
		Locale:     "",
		Dial:       nil,
	}

	l.Info("conf: %s", rabbitConf)
	conn, err := rabbitmq.NewConn(
		connString,
		rabbitmq.WithConnectionOptionsConfig(rabbitmq.Config{}),
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitConf),
	)
	if err != nil {
		return svc, err
	}

	svc.washBonusPub, err = setupPublisher(conn, string(rabbitEntities.WashBonusService), "direct")
	if err != nil {
		return svc, err
	}

	svc.washBonusSub, err = rabbitmq.NewConsumer(conn,
		svc.ProcessMessage,
		string(rabbitEntities.WashBonusRoutingKey),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(rabbitEntities.WashBonusService)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsRoutingKey(string(rabbitEntities.WashBonusRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return svc, err
	}

	svc.adminsPub, err = setupPublisher(conn, string(rabbitEntities.AdminsExchange), "fanout")
	if err != nil {
		return svc, err
	}

	intClient := client.New(httptransport.New(cfg.Url+":15672", "", []string{"http"}), strfmt.Default)
	intAuth := httptransport.BasicAuth(cfg.User, cfg.Password)

	svc.intApi = intClient
	svc.intApiAuth = intAuth

	return svc, nil
}

func setupPublisher(conn *rabbitmq.Conn, exchangeName, exchangeKind string) (*rabbitmq.Publisher, error) {
	return rabbitmq.NewPublisher(conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeName(exchangeName),
		rabbitmq.WithPublisherOptionsExchangeKind(exchangeKind),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
}
