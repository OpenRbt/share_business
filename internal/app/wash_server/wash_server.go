package wash_server

import (
	"sync"
	"wash-bonus/internal/app"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"
	"wash-bonus/internal/app/user"
	"wash-bonus/internal/transport/grpc"

	"github.com/golang-jwt/jwt/v4"

	"crypto/rand"
	"crypto/rsa"

	"os"
)

type WashServerSvc interface {
	Get(prof entity.IdentityProfile, id string) (*entity.WashServer, error)
	Add(prof entity.IdentityProfile, s entity.WashServer) error
	Edit(prof entity.IdentityProfile, id string, update vo.WashServerUpdate) error
	Delete(prof entity.IdentityProfile, id string) error
	List(prof entity.IdentityProfile, filter vo.ListFilter) ([]entity.WashServer, []string, error)
	GenerateServiceKey(prof entity.IdentityProfile, wash_server_id string) (*string, error)
}

type Repository interface {
	GetWashServer(id string) (*entity.WashServer, error)
	AddWashServer(s entity.WashServer) error
	EditWashServer(id string, update vo.WashServerUpdate, editedBy entity.User) error
	DeleteWashServer(id string, deletedBy entity.User) error
	ListWashServers(filter vo.ListFilter) ([]entity.WashServer, []string, error)
}

type Service struct {
	repo                           Repository
	userSvc                        user.UserSvc
	rsaKey                         *rsa.PrivateKey
	washServerGRPCConnectionsMutex sync.Mutex
	washServerGRPCConnections      map[string]grpc.WashServerConnection
}

func NewService(repo Repository, userSvc user.UserSvc, kefilePath string, connections map[string]grpc.WashServerConnection) (WashServerSvc, error) {
	file, err := os.Open(kefilePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return nil, err
	}
	return &Service{
		userSvc:                   userSvc,
		repo:                      repo,
		rsaKey:                    key,
		washServerGRPCConnections: connections,
	}, nil
}

func (a *Service) Get(prof entity.IdentityProfile, id string) (*entity.WashServer, error) {
	return a.repo.GetWashServer(id)
}

func (a *Service) Add(prof entity.IdentityProfile, s entity.WashServer) error {
	return a.repo.AddWashServer(s)
}

func (a *Service) Edit(prof entity.IdentityProfile, id string, update vo.WashServerUpdate) error {
	editor, err := a.userSvc.GetByIdentityID(prof)
	if err != nil {
		return err
	}
	return a.repo.EditWashServer(id, update, *editor)
}

func (a *Service) Delete(prof entity.IdentityProfile, id string) error {
	editor, err := a.userSvc.GetByIdentityID(prof)
	if err != nil {
		return err
	}
	return a.repo.DeleteWashServer(id, *editor)
}

func (a *Service) List(prof entity.IdentityProfile, filter vo.ListFilter) ([]entity.WashServer, []string, error) {
	return a.repo.ListWashServers(filter)
}

func (a *Service) GenerateServiceKey(prof entity.IdentityProfile, wash_server_id string) (*string, error) {
	editor, err := a.userSvc.GetByIdentityID(prof)
	if err != nil {
		return nil, err
	}

	wash, err := a.repo.GetWashServer(wash_server_id)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"owner_id":       wash.Owner.ID,
		"wash_server_id": wash.ID,
	})

	tokenString, err := token.SignedString(a.rsaKey)
	if err != nil {
		return nil, app.ErrGenerateJWT
	}

	err = a.repo.EditWashServer(wash_server_id, vo.WashServerUpdate{
		Name:        wash.Name,
		Description: wash.Description,
		ServiceKey:  tokenString,
		OwnerID:     wash.Owner.ID,
	}, *editor)
	if err != nil {
		return nil, err
	}

	wash, err = a.repo.GetWashServer(wash_server_id)
	if err != nil {
		return nil, err
	}

	a.washServerGRPCConnectionsMutex.Lock()
	a.washServerGRPCConnections[tokenString] = grpc.WashServerConnection{
		WashServer: *wash,
		Verify:     false,
	}
	a.washServerGRPCConnectionsMutex.Unlock()

	return &tokenString, nil
}
