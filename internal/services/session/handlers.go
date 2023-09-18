package session

import (
	"context"
	"errors"
	"washbonus/internal/conversions"
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	"washbonus/internal/entities/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *sessionService) Create(ctx context.Context, serverID uuid.UUID, postID int64) (entities.Session, error) {
	sessionFromDB, err := s.sessionRepo.CreateSession(ctx, serverID, postID)
	if err != nil {
		return entities.Session{}, err
	}

	session := conversions.SessionFromDB(sessionFromDB)

	serverFromDB, err := s.washServerRepo.GetWashServerById(ctx, serverID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.Session{}, entities.ErrNotFound
		}
		return entities.Session{}, err
	}

	server := conversions.WashServerFromDB(serverFromDB)
	session.WashServer = server

	return session, nil
}

func (s *sessionService) Get(ctx context.Context, sessionID uuid.UUID, userID *string) (entities.Session, error) {
	sessionFromDB, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.Session{}, entities.ErrNotFound
		}
		return entities.Session{}, err
	}

	session := conversions.SessionFromDB(sessionFromDB)
	if session.User != nil && userID != nil && session.User.ID != *userID {
		return entities.Session{}, entities.ErrForbidden
	}

	washServerFromDB, err := s.washServerRepo.GetWashServerById(ctx, session.WashServer.ID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entities.Session{}, entities.ErrNotFound
		}
		return entities.Session{}, err
	}

	washServer := conversions.WashServerFromDB(washServerFromDB)

	session.WashServer = washServer
	return session, nil
}

func (s *sessionService) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error {
	return s.sessionRepo.UpdateSessionState(ctx, sessionID, dbmodels.SessionState(state))
}

func (s *sessionService) SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
	return s.sessionRepo.SetSessionUser(ctx, sessionID, userID)
}

func (s *sessionService) SaveMoneyReport(ctx context.Context, report entities.MoneyReport) (err error) {
	return s.sessionRepo.SaveMoneyReport(ctx, conversions.MoneyReportToDB(report))
}

func (s *sessionService) ProcessMoneyReports(ctx context.Context) (err error) {
	lastID := int64(0)

	for {
		reports, err := s.sessionRepo.ProcessAndChargeMoneyReports(ctx, lastID)
		if err != nil {
			s.logger.Warn("failed to process money reports: ", err)
			break
		}

		if len(reports) == 0 {
			break
		}

		lastID = reports[len(reports)-1].ID
	}

	return
}

func (s *sessionService) GetUserPendingBalanceByOrganization(ctx context.Context, userID string, organizationID uuid.UUID) (decimal.Decimal, error) {
	pendingBalance, err := s.sessionRepo.GetUserPendingBalanceByOrganization(ctx, userID, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return decimal.NewFromInt(0), nil
		}

		return decimal.Decimal{}, err
	}

	return pendingBalance, nil
}

func (s *sessionService) GetUserPendingBalances(ctx context.Context, userID string) ([]entities.UserPendingBalance, error) {
	pendingBalances, err := s.sessionRepo.GetUserPendingBalances(ctx, userID)
	if err != nil {
		return nil, err
	}

	return conversions.UserPendingBalancesFromDB(pendingBalances), nil
}

func (s *sessionService) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error) {
	err = s.sessionRepo.ChargeBonuses(ctx, amount, sessionID, userID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.ErrForbidden
		}

		return err
	}

	return err
}

func (s *sessionService) DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	err = s.sessionRepo.DiscardBonuses(ctx, amount, sessionID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.ErrForbidden
		}

		return err
	}

	return err
}

func (s *sessionService) ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	err = s.sessionRepo.ConfirmBonuses(ctx, amount, sessionID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrBadRequest) {
			return entities.ErrForbidden
		}

		return err
	}

	return err
}

func (s *sessionService) LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error) {
	return s.sessionRepo.LogRewardBonuses(ctx, sessionID, payload, messageUuid)
}

func (s *sessionService) DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error) {
	return s.sessionRepo.DeleteUnusedSessions(ctx, SessionRetentionDays)
}
