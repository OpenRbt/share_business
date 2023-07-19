package session

import (
	"context"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *sessionService) Create(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error) {
	return s.sessionRepo.CreateSession(ctx, serverID, postID)
}

func (s *sessionService) Get(ctx context.Context, sessionID uuid.UUID, userID *string) (entity.Session, error) {
	session, err := s.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return session, err
	}

	if session.User != nil && userID != nil && session.User.ID != *userID {
		return session, entity.ErrForbidden
	}

	washServer, err := s.washServerRepo.GetWashServerById(ctx, session.WashServer.ID)
	if err != nil {
		return session, err
	}

	session.WashServer = washServer
	return s.sessionRepo.GetSession(ctx, sessionID)
}

func (s *sessionService) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error {
	return s.sessionRepo.UpdateSessionState(ctx, sessionID, state)
}

func (s *sessionService) SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
	return s.sessionRepo.SetSessionUser(ctx, sessionID, userID)
}

func (s *sessionService) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error) {
	return s.sessionRepo.SaveMoneyReport(ctx, report)
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
			err = s.processMoneyReport(ctx, report)
			if err != nil {
				s.logger.Warn("failed to process money report with id", report.ID, "error", err)
				break
			}
		}
	}

	return
}

func (s *sessionService) processMoneyReport(ctx context.Context, report entity.UserMoneyReport) (err error) {
	coins := decimal.NewFromInt(int64(report.Coins))
	banknotes := decimal.NewFromInt(int64(report.Banknotes))
	electonical := decimal.NewFromInt(int64(report.Electronical))

	percent := decimal.NewFromInt(s.moneyReportsRewardPercentDefault)

	divider := decimal.NewFromInt(100)

	addAmount := coins.Div(divider).Mul(percent)
	addAmount = addAmount.Add(banknotes.Div(divider).Mul(percent))
	addAmount = addAmount.Add(electonical.Div(divider).Mul(percent))

	err = s.userRepo.AddBonuses(ctx, addAmount, report.User)
	if err != nil {
		return
	}

	err = s.sessionRepo.UpdateMoneyReport(ctx, report.ID, true)

	return
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
