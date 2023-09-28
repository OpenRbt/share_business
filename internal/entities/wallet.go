package entities

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type OrganizationForWallet struct {
	ID   uuid.UUID
	Name string
}

type Wallet struct {
	ID             uuid.UUID
	UserID         string
	Organization   OrganizationForWallet
	Balance        decimal.Decimal
	PendingBalance decimal.Decimal
}
