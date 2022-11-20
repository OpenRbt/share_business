package dal

import (
	"context"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"wash_bonus/internal/entity"
)

func (s *Storage) Get(ctx context.Context, userID uuid.UUID) (decimal.Decimal, error) {
	var d decimal.Decimal
	err := s.db.NewSession(nil).
		Select("balance").
		From("users").
		Where("id = ?", userID).
		LoadOneContext(ctx, &d)

	return d, err
}

func (s *Storage) Add(ctx context.Context, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Remove(ctx context.Context, userID uuid.UUID, amount uuid.UUID) (decimal.Decimal, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) LogAction(ctx context.Context, event entity.BalanceEvent) error {
	_, err := s.db.NewSession(nil).
		InsertInto("balance_events").
		Columns("user",
			"operation_kind",
			"old_amount",
			"new_amount",
			"wash_server",
			"session",
			"status",
			"error_msg",
			"date",
		).
		Record(event).
		ExecContext(ctx)

	return err
}
