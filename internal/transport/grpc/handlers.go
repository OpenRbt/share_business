package grpc

import (
	"context"
)

// InitConnection TODO: Improve logic with case when connection already exists
func (s *Service) InitConnection(ctx context.Context, request *InitRequest) (*InitAnswer, error) {
	err := s.createConnectionIfMotExist(request.ServiceKey)
	if err != nil {
		return &InitAnswer{
			Success: false,
			Error:   "connection already taken",
		}, err
	}

	return &InitAnswer{
		Success: true,
	}, nil
}

func (s *Service) Begin(ctx context.Context, request *BeginRequest) (*BeginAnswer, error) {
	con, err := s.getConnection(request.GetServiceKey())
	if err != nil {
		return &BeginAnswer{Success: false}, nil
	}

	err = con.createSessionIfNotExist(request.GetSessionID())
	if err != nil {
		return &BeginAnswer{Success: false}, nil
	}

	return &BeginAnswer{
		Success: true,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, request *RefreshRequest) (*RefreshAnswer, error) {
	con, err := s.getConnection(request.GetServiceKey())
	if err != nil {
		return &RefreshAnswer{Success: false}, nil
	}

	session, err := con.getSession(request.GetSessionID())
	if err != nil {
		return &RefreshAnswer{Success: false}, nil
	}

	con.updateSession(request.GetSessionID(), request.GetEnteredAmount())

	var userID string
	if session.User != nil {
		userID = session.User.ID.String()
	}

	return &RefreshAnswer{
		Success:       true,
		UserID:        userID,
		ConsumeAmount: session.ConsumeAmount.String(),
	}, nil
}

func (s *Service) Confirm(ctx context.Context, request *ConfirmRequest) (*ConfirmAnswer, error) {
	con, err := s.getConnection(request.GetServiceKey())
	if err != nil {
		return &ConfirmAnswer{Success: false}, nil
	}

	err = con.processSession(request.GetSessionID())
	if err != nil {
		return &ConfirmAnswer{Success: false}, nil
	}

	// TODO: Subtract bonus from user

	return &ConfirmAnswer{Success: true}, nil
}

func (s *Service) End(ctx context.Context, request *FinishRequest) (*FinishAnswer, error) {
	con, err := s.getConnection(request.GetServiceKey())
	if err != nil {
		return &FinishAnswer{Success: false}, nil
	}

	// TODO: Calculate bonus amount
	_, err = con.getSession(request.GetSessionID())
	if err != nil {
		return &FinishAnswer{Success: false}, nil
	}

	err = con.finishSession(request.GetSessionID())
	if err != nil {
		return &FinishAnswer{Success: false}, nil
	}

	// TODO: Add bonus to user

	return &FinishAnswer{Success: true}, nil
}
