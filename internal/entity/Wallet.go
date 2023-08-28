package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID             uuid.UUID
	UserID         string
	OrganizationID uuid.UUID
	Balance        decimal.Decimal
	PendingBalance decimal.Decimal
}
