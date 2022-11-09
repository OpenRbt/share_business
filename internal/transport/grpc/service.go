package grpc

import (
	"sync"
	"wash-bonus/internal/app/entity"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Repository interface {
	GetWashServerByKey(key string) (*entity.WashServer, error)
	AddBonuses(id string, balance float64) error
}

type Service struct {
	repo             Repository
	connectionsMutex sync.RWMutex
	connections      map[string]*Connection
	UnimplementedServerServiceServer
	UnimplementedSessionServiceServer
}

type Connection struct {
	Valid  bool
	Server entity.WashServer

	sessionsMutex sync.RWMutex
	Sessions      map[string]*Session
}

type Session struct {
	ID            uuid.UUID
	PostID        int64
	User          *entity.User
	Balance       *entity.Balance
	Amount        decimal.Decimal
	ConsumeAmount decimal.Decimal
	Processed     bool
}

func NewSvc(repo Repository) *Service {
	return &Service{
		repo:        repo,
		connections: make(map[string]*Connection),
	}
}
