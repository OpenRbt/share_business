package grpc

type washServer interface {
	VerifyClient(*Verify) (bool, error)
	SendMessage(WashServerService_SendMessageServer) error
	SendMessageToOtherClient(WashServerService_SendMessageToOtherClientServer) error
}
