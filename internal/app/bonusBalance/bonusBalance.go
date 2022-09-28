package bonusBalance

import (
	"wash-bonus/internal/app/entity"
)

type BonusBalanceSvc interface {
	GetBonusBalance(id string) (*entity.BonusBalance, error)
	AddBonusBalance(userId string, balance float64) (*entity.BonusBalance, error)
	EditBonusBalance(Id string, balance float64) error
	DeleteBonusBalance(id string, userId string) error
}

type Repository interface {
	GetBonusBalance(id string) (*entity.BonusBalance, error)
	AddBonusBalance(userID string, balance float64) (*entity.BonusBalance, error)
	EditBonusBalance(id string, balance float64) error
	DeleteBonusBalance(id string, userId string) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) BonusBalanceSvc {
	return &Service{repo: repo}
}

func (a *Service) GetBonusBalance(id string) (*entity.BonusBalance, error) {
	return a.repo.GetBonusBalance(id)
}

func (a *Service) AddBonusBalance(userId string, balance float64) (*entity.BonusBalance, error) {
	return a.repo.AddBonusBalance(userId, balance)
}

func (a *Service) EditBonusBalance(Id string, balance float64) error {
	return a.repo.EditBonusBalance(Id, balance)
}

func (a *Service) DeleteBonusBalance(id string, userId string) error {
	return a.repo.DeleteBonusBalance(id, userId)
}
