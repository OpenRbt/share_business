package session

import (
	"context"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func (s *service) Create(ctx context.Context, serverID uuid.UUID, postID int64) (session entity.Session, err error) {
	session, err = s.sessionRepo.CreateSession(ctx, serverID, postID)

	return
}

func (s *service) Get(ctx context.Context, sessionID uuid.UUID) (session entity.Session, err error) {
	session, err = s.sessionRepo.GetSession(ctx, sessionID)

	return
}

func (s *service) UpdateSessionState(ctx context.Context, sessionID uuid.UUID, state vo.SessionState) error {
	return s.sessionRepo.UpdateSessionState(ctx, sessionID, state)
}

func (s *service) SetSessionUser(ctx context.Context, sessionID uuid.UUID, userID string) (err error) {
	return s.sessionRepo.SetSessionUser(ctx, sessionID, userID)
}

func (s *service) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error) {
	return s.sessionRepo.SaveMoneyReport(ctx, report)
}

func (s *service) ProcessMoneyReports(ctx context.Context) (err error) {
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
				s.l.Warn("failed to process money report with id", report.ID, "error", err)
				break
			}
		}
	}

	return
}

func (s *service) processMoneyReport(ctx context.Context, report entity.UserMoneyReport) (err error) {
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

func (s *service) ChargeBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID, userID string) (err error) {
	return s.sessionRepo.ChargeBonuses(ctx, amount, sessionID, userID)
}

func (s *service) DiscardBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	return s.sessionRepo.DiscardBonuses(ctx, amount, sessionID)
}

func (s *service) ConfirmBonuses(ctx context.Context, amount decimal.Decimal, sessionID uuid.UUID) (err error) {
	return s.sessionRepo.ConfirmBonuses(ctx, amount, sessionID)
}

func (s *service) LogRewardBonuses(ctx context.Context, sessionID uuid.UUID, payload []byte, messageUuid uuid.UUID) (err error) {
	return s.sessionRepo.LogRewardBonuses(ctx, sessionID, payload, messageUuid)
}

func (s *service) DeleteUnusedSessions(ctx context.Context, SessionRetentionDays int64) (int64, error) {
	return s.sessionRepo.DeleteUnusedSessions(ctx, SessionRetentionDays)
}
