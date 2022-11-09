package grpc

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"sync"
	"wash-bonus/internal/app/entity"
)

type Repository interface {
	GetWashServerByKey(key string) (*entity.WashServer, error)
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
