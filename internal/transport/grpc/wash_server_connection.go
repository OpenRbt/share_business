package grpc

import (
	"wash-bonus/internal/app/entity"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type WashSession struct {
	ID               uuid.UUID
	User             entity.User
	Amount           decimal.Decimal
	MaxConsumeAmount decimal.Decimal
}

type WashServerConnection struct {
	Verify                         bool
	WashServer                     entity.WashServer
	StreamSendMessage              WashServerService_SendMessageServer
	StreamSendMessageToOtherClient WashServerService_SendMessageToOtherClientServer
	WashSessions                   map[string]WashSession
}

func NewWashServerConnection(washServer entity.WashServer) WashServerConnection {
	return WashServerConnection{
		Verify:       false,
		WashServer:   washServer,
		WashSessions: make(map[string]WashSession),
	}
}

func NewWashSession(user entity.User) WashSession {
	return WashSession{
		ID:   uuid.NewV4(),
		User: user,
	}
}
