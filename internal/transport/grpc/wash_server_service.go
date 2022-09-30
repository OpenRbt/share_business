package grpc

import (
	context "context"

	uuid "github.com/satori/go.uuid"
)

type app interface {
	VerifyClient(*Verify) (bool, error)
	SendMessage(WashServerService_SendMessageServer) error
	SendMessageToOtherClient(WashServerService_SendMessageToOtherClientServer) error
}

type WashServer struct {
	ID                             uuid.UUID
	OwnerID                        uuid.UUID
	StreamSendMessage              WashServerService_SendMessageServer
	StreamSendMessageToOtherClient WashServerService_SendMessageToOtherClientServer
}

type washServerService struct {
	app app
}

func (svc *washServerService) VerifyClient(ctx context.Context, msg *Verify) (*VerifyAnswer, error) {
	success, err := svc.app.VerifyClient(msg)
	return &VerifyAnswer{Success: success}, err
}

func (svc *washServerService) SendMessage(stream WashServerService_SendMessageServer) error {
	return svc.app.SendMessage(stream)
}

func (svc *washServerService) SendMessageToOtherClient(stream WashServerService_SendMessageToOtherClientServer) error {
	return svc.app.SendMessageToOtherClient(stream)
}

func (svc *washServerService) mustEmbedUnimplementedWashServerServiceServer() {}
