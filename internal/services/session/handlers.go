package session

import (
	"context"
	"errors"
	"washBonus/internal/conversions"
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"
	moneyreports "washBonus/pkg/moneyReports"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *sessionService) Create(ctx context.Context, serverID uuid.UUID, postID int64) (entity.Session, error) {
	sessionFromDB, err := s.sessionRepo.CreateSession(ctx, serverID, postID)
	if err != nil {
		return entity.Session{}, err
	}

	session := conversions.SessionFromDB(sessionFromDB)

	serverFromDB, err := s.washServerRepo.GetWashServerById(ctx, serverID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.Session{}, entity.ErrNotFound
		}
		return entity.Session{}, err
	}

	server := conversions.WashServerFromDB(serverFromDB)
	session.WashServer = server

	return session, nil
}

func (s *sessionService) Get(ctx context.Context, sessionID uuid.UUID, userID *string) (entity.Session, error) {
	sessionFromDB, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.Session{}, entity.ErrNotFound
		}
		return entity.Session{}, err
	}

	session := conversions.SessionFromDB(sessionFromDB)
	if session.User != nil && userID != nil && session.User.ID != *userID {
		return entity.Session{}, entity.ErrForbidden
	}

	washServerFromDB, err := s.washServerRepo.GetWashServerById(ctx, session.WashServer.ID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return entity.Session{}, entity.ErrNotFound
		}
		return entity.Session{}, err
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

func (s *sessionService) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error) {
	return s.sessionRepo.SaveMoneyReport(ctx, conversions.MoneyReportToDB(report))
}

func (s *sessionService) ProcessMoneyReports(ctx context.Context) (err error) {
	lastID := int64(0)

	for {
		reports, err := s.sessionRepo.GetUnprocessedMoneyReports(ctx, lastID, s.reportsProcessingDelayInMinutes)

		if err != nil {
			return err
		}

		if len(reports) == 0 {
			break
		}

		lastID = reports[len(reports)-1].ID

		for _, report := range reports {
			err = s.processMoneyReport(ctx, conversions.UserMoneyReportFromDB(report))
			if err != nil {
				s.logger.Warn("failed to process money report with id", report.ID, "error", err)
				break
			}
		}
	}

	return
}

func (s *sessionService) processMoneyReport(ctx context.Context, report entity.UserMoneyReport) (err error) {
	addAmount := moneyreports.ProcessBonusesReward(report, decimal.NewFromInt(int64(s.moneyReportsRewardPercentDefault)))

	err = s.walletRepo.ChargeBonusesByUserAndOrganization(ctx, addAmount, report.User, report.OrganizationID)
	if err != nil {
		return
	}

	err = s.sessionRepo.UpdateMoneyReport(ctx, report.ID, true)

	return
}

func (s *sessionService) GetUserOrganizationPendingBalance(ctx context.Context, userID string, organizationID uuid.UUID) (decimal.Decimal, error) {
	reports, err := s.sessionRepo.GetUnporcessedReportsByUserAndOrganization(ctx, userID, organizationID)
	if err != nil {
		if errors.Is(err, dbmodels.ErrNotFound) {
			return decimal.NewFromInt(0), nil
		}

		return decimal.Decimal{}, err
	}

	balance := decimal.Zero

	for _, report := range reports {
		addAmount := moneyreports.ProcessBonusesReward(conversions.UserMoneyReportFromDB(report), decimal.NewFromInt(int64(s.moneyReportsRewardPercentDefault)))
		balance = balance.Add(addAmount)
	}

	return balance, nil
}

func (s *sessionService) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error) {
	return s.sessionRepo.ChargeBonuses(ctx, amount, sessionID, userID)
}

func (s *sessionService) DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	return s.sessionRepo.DiscardBonuses(ctx, amount, sessionID)
}

func (s *sessionService) ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	return s.sessionRepo.ConfirmBonuses(ctx, amount, sessionID)
}

func (s *sessionService) LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error) {
	return s.sessionRepo.LogRewardBonuses(ctx, sessionID, payload, messageUuid)
}

func (s *sessionService) DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error) {
	return s.sessionRepo.DeleteUnusedSessions(ctx, SessionRetentionDays)
}
