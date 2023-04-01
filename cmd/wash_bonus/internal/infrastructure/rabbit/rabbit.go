package rabbit

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/OpenRbt/share_business/wash_rabbit/entity/vo"
	"github.com/wagslane/go-rabbitmq"
	"go.uber.org/zap"
	"io/ioutil"
	"wash_bonus/internal/app/session"
	"wash_bonus/internal/app/wash_server"

	amqp "github.com/rabbitmq/amqp091-go"
)

type WorkerService interface {
	ProcessBonusMessage(msg any, target string) (err error)
	SendMessage(msg any, target string) (err error)
}

type Service struct {
	l    *zap.SugaredLogger
	conn *amqp.Connection

	// wash bonus handlers
	washBonusPub    *rabbitmq.Publisher
	washBonusSvcSub *rabbitmq.Consumer
	customConsumers map[string]*rabbitmq.Consumer

	// wash admin handler
	washServerSub *rabbitmq.Consumer

	//worker services
	svcWashServer wash_server.Service
	svcSessions   session.Service
}

func New(l *zap.SugaredLogger, url string, port string, certsPath string, user string, password string, washServerSvc wash_server.Service, sessionsSvc session.Service) (svc *Service, err error) {
	svc = &Service{
		l:             l,
		svcWashServer: washServerSvc,
		svcSessions:   sessionsSvc,
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
		InsecureSkipVerify: true,
	}

	connString := fmt.Sprintf("amqps://%s:%s@%s:%s/", user, password, url, port)
	rabbitConf := rabbitmq.Config{
		SASL:            nil,
		Vhost:           "/",
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
		rabbitmq.WithConsumerOptionsRoutingKey(string(vo.WashAdminRoutingKey)),
		rabbitmq.WithConsumerOptionsExchangeDurable,
	)

	return
}
