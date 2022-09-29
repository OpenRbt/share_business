package grpc

import context "context"

type app interface {
	VerifyClient()
	SendMessage()
	SendMessageToOtherClient()
}

type washServerService struct {
	app app
}

func (svc *washServerService) VerifyClient(ctx context.Context, msg *Verify) (*VerifyAnswer, error) {
	svc.app.VerifyClient()
	return nil, nil
}

func (svc *washServerService) SendMessage(stream WashServerService_SendMessageServer) error {
	svc.app.SendMessage()
	return nil
}

func (svc *washServerService) SendMessageToOtherClient(stream WashServerService_SendMessageToOtherClientServer) error {
	svc.app.SendMessageToOtherClient()
	return nil
}

func (svc *washServerService) mustEmbedUnimplementedWashServerServiceServer() {}
