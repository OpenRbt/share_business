package grpc

import (
	context "context"
	"log"
	"sync"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/app/wash_server"
)

type WashServerConnection struct {
	Verify                         bool
	WashServer                     entity.WashServer
	StreamSendMessage              WashServerService_SendMessageServer
	StreamSendMessageToOtherClient WashServerService_SendMessageToOtherClientServer
}

type WashServerService struct {
	WashServerRepo        wash_server.Repository
	Mutex                 sync.Mutex
	WashServerConnections map[string]WashServerConnection
}

func NewWashServerService(washServerRepo wash_server.Repository, washServerConnections map[string]WashServerConnection) *WashServerService {
	return &WashServerService{
		WashServerRepo:        washServerRepo,
		WashServerConnections: washServerConnections,
	}
}

func (svc *WashServerService) VerifyClient(ctx context.Context, msg *Verify) (*VerifyAnswer, error) {
	svc.Mutex.Lock()
	washServer, ok := svc.WashServerConnections[msg.ServiceKey]

	var err error
	if ok {
		washServer.Verify = true
		svc.WashServerConnections[msg.ServiceKey] = washServer
	} else {
		log.Println("Verify failed for wash server ", msg.ServiceKey)
		err = ErrVerifyFailed
	}

	svc.Mutex.Unlock()
	return &VerifyAnswer{Success: ok}, err
}

func (svc *WashServerService) SendMessage(stream WashServerService_SendMessageServer) error {
	msg, err := stream.Recv()
	if err != nil {
		log.Println("Failed to recv: ", err)
		return err
	}

	svc.Mutex.Lock()
	washServer, ok := svc.WashServerConnections[msg.ServiceKey]
	if !ok || !washServer.Verify {
		log.Println("Verify failed for wash server")
		svc.Mutex.Unlock()
		return ErrVerifyFailed
	}
	washServer.StreamSendMessage = stream
	svc.WashServerConnections[msg.ServiceKey] = washServer
	svc.Mutex.Unlock()

	for {
		msg, err = stream.Recv()
		if err != nil {
			log.Println("Failed to recv: ", err)
			return err
		}

		log.Println("Message received: ", msg.Msg)
		err = stream.Send(&MessageAnswer{Msg: "Message received: " + msg.Msg})
		if err != nil {
			log.Println("Failed to send: ", err)
			return err
		}
	}
}

func (svc *WashServerService) SendMessageToOtherClient(stream WashServerService_SendMessageToOtherClientServer) error {
	msg, err := stream.Recv()
	if err != nil {
		log.Println("Failed to recv: ", err)
		return err
	}

	svc.Mutex.Lock()
	washServer, ok := svc.WashServerConnections[msg.ServiceKey]
	if !ok || !washServer.Verify {
		log.Println("Verify failed for wash server")
		svc.Mutex.Unlock()
		return ErrVerifyFailed
	}
	washServer.StreamSendMessageToOtherClient = stream
	svc.WashServerConnections[msg.ServiceKey] = washServer
	svc.Mutex.Unlock()

	for {
		msg, err = stream.Recv()
		if err != nil {
			log.Println("Failed to recv: ", err)
			return err
		}

		log.Println("Message received: ", msg.Msg, " to user ", msg.TargetUuid)

		washServerFromBD, err := svc.WashServerRepo.GetWashServer(msg.TargetUuid)
		if err != nil {
			stream.Send(&MessageToOtherAnswer{Success: false})
		}

		if washServer, ok := svc.WashServerConnections[washServerFromBD.ServiceKey]; ok && washServer.Verify {
			washServer.StreamSendMessage.Send(&MessageAnswer{Msg: msg.Msg})
			stream.Send(&MessageToOtherAnswer{Success: true})
		} else {
			stream.Send(&MessageToOtherAnswer{Success: false})
		}
	}
}

func (svc *WashServerService) mustEmbedUnimplementedWashServerServiceServer() {}
