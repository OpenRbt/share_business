package rabbit

import (
	"context"
	"fmt"
	"washbonus/internal/app"
	"washbonus/internal/config"
	"washbonus/internal/infrastructure/rabbit/entities/vo"
	"washbonus/rabbit-intapi/client"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitService interface {
	SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) (err error)
	CreateRabbitUser(userID, userKey string) error
	DeleteRabbitUser(ctx context.Context, userID string) error
}

type Service struct {
	l    *zap.SugaredLogger
	conn *amqp.Connection

	washBonusPub    *rabbitmq.Publisher
	washBonusSvcSub *rabbitmq.Consumer
	rabbitSvc       app.RabbitService

	intApi     *client.RabbitIntAPI
	intApiAuth runtime.ClientAuthInfoWriter
}

func New(l *zap.SugaredLogger, cfg config.RabbitMQConfig, rabbitSvc app.RabbitService) (svc *Service, err error) {
	svc = &Service{
		l:         l,
		rabbitSvc: rabbitSvc,
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
		return
	}

	svc.washBonusPub, err = rabbitmq.NewPublisher(conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeName(string(vo.WashBonusService)),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return
	}

	svc.washBonusSvcSub, err = rabbitmq.NewConsumer(conn,
		svc.ProcessMessage,
		string(vo.WashBonusRoutingKey),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(vo.WashBonusService)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsRoutingKey(string(vo.WashBonusRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)
	if err != nil {
		return
	}

	intClient := client.New(httptransport.New(cfg.Url+":15672", "", []string{"http"}), strfmt.Default)
	intAuth := httptransport.BasicAuth(cfg.User, cfg.Password)

	svc.intApi = intClient
	svc.intApiAuth = intAuth

	return
}
