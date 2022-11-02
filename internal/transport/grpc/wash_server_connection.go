package grpc

import (
	sync "sync"
	"wash-bonus/internal/app/entity"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type WashSession struct {
	ID               uuid.UUID
	PostID           uuid.UUID
	User             entity.User
	Amount           decimal.Decimal
	MaxConsumeAmount decimal.Decimal
	Confirm          bool
}

type WashServerConnection struct {
	Verify              bool
	WashServer          entity.WashServer
	StreamUpdateSession WashServerService_UpdateSessionServer
	WashSessionsMutex   sync.Mutex
	WashSessions        map[string]WashSession
}

func NewWashServerConnection(washServer entity.WashServer) *WashServerConnection {
	return &WashServerConnection{
		WashServer:   washServer,
		WashSessions: make(map[string]WashSession),
	}
}
