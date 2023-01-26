package rabbit

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
	"io/ioutil"
	"wash_admin/internal/infrastructure/rabbit/models/vo"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Service struct {
	l               *zap.SugaredLogger
	conn            *amqp.Connection
	eventsConsumer  *rabbitmq.Consumer
	eventsPublisher *rabbitmq.Publisher
}

func New(l *zap.SugaredLogger, url, port, certsPath, user, password string) (svc *Service, err error) {
	svc = &Service{
		l: l,
	}

	caCert, err := ioutil.ReadFile(certsPath + "root_ca.pem")
	if err != nil {
		return nil, err
	}

	cert, err := tls.LoadX509KeyPair(certsPath+"client.pem", certsPath+"client_key.pem")
	if err != nil {
		return nil, err
	}
	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)

	tlsConf := &tls.Config{
		RootCAs:            rootCAs,
		Certificates:       []tls.Certificate{cert},
		ServerName:         "localhost", // Optional
		InsecureSkipVerify: true,
	}

	connString := fmt.Sprintf("amqps://%s:%s@%s:%s/", user, password, url, port)
	rabbitConf := rabbitmq.Config{
		SASL:            nil,
		Vhost:           "",
		ChannelMax:      0,
		FrameSize:       0,
		Heartbeat:       0,
		TLSClientConfig: tlsConf,
		Properties:      nil,
		Locale:          "",
		Dial:            nil,
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
		rabbitmq.WithPublisherOptionsExchangeName(vo.WashAdminService),
		rabbitmq.WithPublisherOptionsExchangeKind("direct"),
		rabbitmq.WithPublisherOptionsExchangeDurable,
	)
	if err != nil {
		return
	}
	svc.eventsConsumer, err = rabbitmq.NewConsumer(conn,
		func(d rabbitmq.Delivery) (action rabbitmq.Action) {
			l.Error("received unexpected message with type: ", vo.MessageTypeFromString(d.Type))

			return rabbitmq.NackDiscard
		},
		vo.WashAdminSvc,
		rabbitmq.WithConsumerOptionsExchangeDeclare,
		rabbitmq.WithConsumerOptionsExchangeName(vo.WashAdminService),
		rabbitmq.WithConsumerOptionsExchangeKind("direct"),
		rabbitmq.WithConsumerOptionsExchangeDurable,
		rabbitmq.WithConsumerOptionsConsumerExclusive,
	)

	return
}
