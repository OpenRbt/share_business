package rabbit

import (
	"fmt"
	"wash_admin/internal/infrastructure/rabbit-intapi/client"

	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	l               *zap.SugaredLogger
	conn            *amqp.Connection
	eventsConsumer  *rabbitmq.Consumer
	eventsPublisher *rabbitmq.Publisher

	intApi     *client.RabbitIntAPI
	intApiAuth runtime.ClientAuthInfoWriter
}

//go:generate rm -rf ../rabbit-intapi/model ../rabbit-intapi/client
//go:generate swagger generate client -t ../rabbit-intapi -f ../rabbit-intapi/swagger.yaml --strict-responders --strict-additional-properties
func New(l *zap.SugaredLogger, url, port, certsPath, user, password string) (svc *Service, err error) {
	svc = &Service{
		l: l,
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
		rabbitmq.WithConnectionOptionsLogging,
		rabbitmq.WithConnectionOptionsConfig(rabbitConf),
	)

	if err != nil {
		return
	}

	svc.eventsPublisher, err = rabbitmq.NewPublisher(conn,
		rabbitmq.WithPublisherOptionsLogging,
		rabbitmq.WithPublisherOptionsExchangeDeclare,
		rabbitmq.WithPublisherOptionsExchangeName(string(vo.WashAdminService)),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return
	}
	svc.eventsConsumer, err = rabbitmq.NewConsumer(conn,
		func(d rabbitmq.Delivery) (action rabbitmq.Action) {
			l.Error("received unexpected message with type: ", vo.MessageType(d.Type))

			return rabbitmq.NackDiscard
		},
		string(vo.WashAdminRoutingKey),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(string(vo.WashAdminService)),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsRoutingKey(string(vo.WashAdminRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsConsumerExclusive,
	)

	intClient := client.New(httptransport.New(url+":15672", "", []string{"http"}), strfmt.Default)
	intAuth := httptransport.BasicAuth(user, password)

	svc.intApi = intClient
	svc.intApiAuth = intAuth

	return
}
