package Balance

import (
	"wash-bonus/internal/app/entity"
)

type BalanceSvc interface {
	GetBalance(id string) (*entity.Balance, error)
	AddBalance(userId string, balance float64) (*entity.Balance, error)
	EditBalance(Id string, balance float64) error
	DeleteBalance(id string, userId string) error
}

type Repository interface {
	GetBalance(id string) (*entity.Balance, error)
	AddBalance(userID string, balance float64) (*entity.Balance, error)
	EditBalance(id string, balance float64) error
	DeleteBalance(id string, userId string) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) BalanceSvc {
	return &Service{repo: repo}
}

func (a *Service) GetBalance(id string) (*entity.Balance, error) {
	return a.repo.GetBalance(id)
}

func (a *Service) AddBalance(userId string, balance float64) (*entity.Balance, error) {
	return a.repo.AddBalance(userId, balance)
}

func (a *Service) EditBalance(Id string, balance float64) error {
	return a.repo.EditBalance(Id, balance)
}

func (a *Service) DeleteBalance(id string, userId string) error {
	return a.repo.DeleteBalance(id, userId)
}
