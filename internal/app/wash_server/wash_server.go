package wash_server

import (
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"
)

type WashServerSvc interface {
	Get(prof entity.IdentityProfile, id string) (*entity.WashServer, error)
	Add(prof entity.IdentityProfile, s entity.WashServer) error
	Edit(prof entity.IdentityProfile, id string, update vo.WashServerUpdate) error
	Delete(prof entity.IdentityProfile, id string) error
	List(prof entity.IdentityProfile, filter vo.ListFilter) ([]entity.WashServer, []string, error)
}

type Repository interface {
	GetWashServer(id string) (*entity.WashServer, error)
	AddWashServer(s entity.WashServer) error
	EditWashServer(id string, update vo.WashServerUpdate, editedBy entity.User) error
	DeleteWashServer(id string, deletedBy entity.User) error
	ListWashServer(filter vo.ListFilter) ([]entity.WashServer, []string, error)
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) WashServerSvc {
	return &Service{repo: repo}
}

func (a *Service) Get(prof entity.IdentityProfile, id string) (*entity.WashServer, error) {
	return a.repo.GetWashServer(id)
}

func (a *Service) Add(prof entity.IdentityProfile, s entity.WashServer) error {
	return a.repo.AddWashServer(s)
}

func (a *Service) Edit(prof entity.IdentityProfile, id string, update vo.WashServerUpdate) error {
	return a.repo.EditWashServer(id, update, entity.User{})
}

func (a *Service) Delete(prof entity.IdentityProfile, id string) error {
	return a.repo.DeleteWashServer(id, entity.User{})
}

func (a *Service) List(prof entity.IdentityProfile, filter vo.ListFilter) ([]entity.WashServer, []string, error) {
	return a.repo.ListWashServer(filter)
}
