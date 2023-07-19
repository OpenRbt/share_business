package rabbit

import (
	"fmt"
	"washBonus/internal/app"
	"washBonus/internal/infrastructure/rabbit/entity/vo"
	"washBonus/rabbit-intapi/client"

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
}

type Service struct {
	l    *zap.SugaredLogger
	conn *amqp.Connection

	// wash bonus handlers
	washBonusPub    *rabbitmq.Publisher
	washBonusSvcSub *rabbitmq.Consumer

	// wash admin handler
	washServerSub *rabbitmq.Consumer
	rabbitSvc     app.RabbitService

	intApi     *client.RabbitIntAPI
	intApiAuth runtime.ClientAuthInfoWriter
}

func New(l *zap.SugaredLogger, url string, port string, user string, password string, rabbitSvc app.RabbitService) (svc *Service, err error) {
	svc = &Service{
		l:         l,
		rabbitSvc: rabbitSvc,
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

	intClient := client.New(httptransport.New(url+":15672", "", []string{"http"}), strfmt.Default)
	intAuth := httptransport.BasicAuth(user, password)

	svc.intApi = intClient
	svc.intApiAuth = intAuth

	return
}
