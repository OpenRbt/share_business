package grpc

import (
	"context"
	"wash_bonus/intapi"
)

func (s *Service) Begin(ctx context.Context, request *intapi.BeginRequest) (*intapi.BeginAnswer, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Refresh(ctx context.Context, request *intapi.RefreshRequest) (*intapi.RefreshAnswer, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) Confirm(ctx context.Context, request *intapi.ConfirmRequest) (*intapi.ConfirmAnswer, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) End(ctx context.Context, request *intapi.FinishRequest) (*intapi.FinishAnswer, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Service) mustEmbedUnimplementedSessionServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (s *Service) mustEmbedUnimplementedServerServiceServer() {
	//TODO implement me
	panic("implement me")
}
