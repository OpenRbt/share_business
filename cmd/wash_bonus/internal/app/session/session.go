package session

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/entity"
	"wash_bonus/internal/entity/vo"
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

func (s *service) UpdateSessionBalance(ctx context.Context, sessionID uuid.UUID, amount decimal.Decimal) (err error) {
	return s.sessionRepo.UpdateSessionBalance(ctx, sessionID, amount)
}

func (s *service) SaveMoneyReport(ctx context.Context, report entity.MoneyReport) (err error) {
	return s.sessionRepo.SaveMoneyReport(ctx, report)
}

func (s *service) ProcessMoneyReports(ctx context.Context) (err error) {
	reports, err := s.sessionRepo.GetUnprocessedMoneyReports(ctx)

	for _, report := range reports {
		err = s.processMoneyReport(ctx, report)
		if err != nil {
			break
		}
	}

	return
}

func (s *service) processMoneyReport(ctx context.Context, report entity.UserMoneyReport) (err error) {
	coins := decimal.NewFromInt(int64(report.Coins))
	banknotes := decimal.NewFromInt(int64(report.Banknotes))
	electonical := decimal.NewFromInt(int64(report.Electronical))

	percent := decimal.NewFromInt(5) //TODO: GET VALUE FROM CONFIGURATION

	divider := decimal.NewFromInt(100)

	addAmount := coins.Div(divider).Mul(percent)
	addAmount = addAmount.Add(banknotes.Div(divider).Mul(percent))
	addAmount = addAmount.Add(electonical.Div(divider).Mul(percent))

	err = s.userRepo.UpdateBalance(ctx, report.User, addAmount)
	if err != nil {
		return
	}

	err = s.sessionRepo.UpdateMoneyReport(ctx, report.ID, true)

	return
}
