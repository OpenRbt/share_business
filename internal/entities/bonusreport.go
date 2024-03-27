package entities

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type BonusReport struct {
	ID            uuid.UUID
	Amount        decimal.Decimal
	Date          time.Time
	OperationType BalanceOperationType
	UserID        string
	Organization  SimleOrganization
}

type BalanceOperationType string

const (
	DepositOperationType    BalanceOperationType = "deposit"
	WithdrawalOperationType BalanceOperationType = "withdrawal"
)

type BonusReportFilter struct {
	Filter
	OrganizationID       *uuid.UUID
	BalanceOperationType *BalanceOperationType
}
