package rabbit

import (
	"fmt"
	"wash_bonus/internal/usecase/rabbit"

	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit interface {
	SendMessage(msg interface{}, service vo.Service, routingKey vo.RoutingKey, messageType vo.MessageType) (err error)
}

type UseCase interface {
	ProcessMessage(d rabbitmq.Delivery) (action rabbitmq.Action)
}

type Service struct {
	l    *zap.SugaredLogger
	conn *amqp.Connection

	// wash bonus handlers
	washBonusPub    *rabbitmq.Publisher
	washBonusSvcSub *rabbitmq.Consumer

	// wash admin handler
	washServerSub *rabbitmq.Consumer

	useCase rabbit.UseCase
}

func New(l *zap.SugaredLogger, url string, port string, user string, password string, useCase rabbit.UseCase) (svc *Service, err error) {
	svc = &Service{
		l:       l,
		useCase: useCase,
	}

	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/", user, password, url, port)
	rabbitConf := rabbitmq.Config{
		SASL: []amqp.Authentication{
			&amqp.PlainAuth{
				Username: user,
				Password: password,
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

	svc.washServerSub, err = rabbitmq.NewConsumer(conn,
		svc.ProcessMessage,
		string(vo.WashAdminServesEventsRoutingKey),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(vo.WashAdminService)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsRoutingKey(string(vo.WashAdminServesEventsRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)

	return
}
