package dbmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type BonusReport struct {
	ID            uuid.UUID            `db:"id"`
	Amount        decimal.Decimal      `db:"amount"`
	Date          time.Time            `db:"date"`
	UserID        string               `db:"user_id"`
	OperationType BalanceOperationType `db:"operation_type"`
	Organization  SimleOrganization
}

type BalanceOperationType string

const (
	DepositOperationType    BalanceOperationType = "deposit"
	WithdrawalOperationType BalanceOperationType = "withdrawal"
)
