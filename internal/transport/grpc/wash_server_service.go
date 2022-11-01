package grpc

import (
	"context"
	"log"
	"sync"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/entity/vo"

	uuid "github.com/satori/go.uuid"
)

type WashServerRepository interface {
	GetWashServer(id string) (*entity.WashServer, error)
	ListWashServers(filter vo.ListFilter) ([]entity.WashServer, []string, error)
}

type WashServerService struct {
	WashServerRepo             WashServerRepository
	WashServerConnectionsMutex sync.Mutex
	WashServerConnections      map[string]*WashServerConnection
}

func NewWashServerService(washServerRepo WashServerRepository) (*WashServerService, error) {
	washList, _, err := washServerRepo.ListWashServers(vo.ListFilter{})
	if err != nil {
		return nil, err
	}

	connections := make(map[string]*WashServerConnection)
	for _, v := range washList {
		if v.ServiceKey != "" {
			connections[v.ServiceKey] = NewWashServerConnection(v)
		}
	}

	return &WashServerService{
		WashServerRepo:        washServerRepo,
		WashServerConnections: connections,
	}, nil
}

func (svc *WashServerService) InitConnection(ctx context.Context, msg *InitConnectionRequest) (*InitConnectionAnswer, error) {
	log.Println("InitConnection: ", msg.ServiceKey)

	svc.WashServerConnectionsMutex.Lock()
	defer svc.WashServerConnectionsMutex.Unlock()
	washServerConnection, ok := svc.WashServerConnections[msg.ServiceKey]

	if !ok || washServerConnection.Verify {
		return nil, ErrVerifyFailed
	}

	washServerConnection.Verify = true
	svc.WashServerConnections[msg.ServiceKey] = washServerConnection

	return &InitConnectionAnswer{Success: true}, nil
}

func (svc *WashServerService) StartSession(ctx context.Context, msg *StartSessionRequest) (*StartSessionAnswer, error) {
	log.Println("StartSession: ", msg.ServiceKey)

	svc.WashServerConnectionsMutex.Lock()
	defer svc.WashServerConnectionsMutex.Unlock()
	washServerConnection, ok := svc.WashServerConnections[msg.ServiceKey]

	if !ok || washServerConnection.Verify {
		return nil, ErrNotFound
	}

	sessionID, err := uuid.FromString(msg.SessionID)
	if err != nil {
		return nil, ErrBadID
	}

	washServerConnection.WashSessionsMutex.Lock()
	defer washServerConnection.WashSessionsMutex.Unlock()
	washServerConnection.WashSessions[msg.SessionID] = WashSession{
		ID: sessionID,
	}

	return &StartSessionAnswer{Success: true}, nil
}

func (svc *WashServerService) ConfirmSession(ctx context.Context, msg *ConfirmSessionRequest) (*ConfirmSessionAnswer, error) {
	log.Println("ConfirmSession: ", msg.ServiceKey)

	svc.WashServerConnectionsMutex.Lock()
	defer svc.WashServerConnectionsMutex.Unlock()
	washServerConnection, ok := svc.WashServerConnections[msg.ServiceKey]

	if !ok || washServerConnection.Verify {
		return nil, ErrNotFound
	}

	washServerConnection.WashSessionsMutex.Lock()
	defer washServerConnection.WashSessionsMutex.Unlock()
	washSeddion, ok := washServerConnection.WashSessions[msg.SessionID]
	if !ok {
		return nil, ErrNotFound
	}

	washSeddion.Confirm = true
	washServerConnection.WashSessions[msg.SessionID] = washSeddion

	return &ConfirmSessionAnswer{Success: true}, nil
}

func (svc *WashServerService) FinishSession(ctx context.Context, msg *FinishSessionRequest) (*FinishSessionAnswer, error) {
	log.Println("FinishSession: ", msg.ServiceKey)

	svc.WashServerConnectionsMutex.Lock()
	defer svc.WashServerConnectionsMutex.Unlock()
	washServerConnection, ok := svc.WashServerConnections[msg.ServiceKey]

	if !ok || washServerConnection.Verify {
		return nil, ErrNotFound
	}

	washServerConnection.WashSessionsMutex.Lock()
	defer washServerConnection.WashSessionsMutex.Unlock()
	delete(washServerConnection.WashSessions, msg.SessionID)

	return &FinishSessionAnswer{Success: true}, nil
}

func (svc *WashServerService) UpdateSession(stream WashServerService_UpdateSessionServer) error {
	return nil
}

func (svc *WashServerService) mustEmbedUnimplementedWashServerServiceServer() {}
