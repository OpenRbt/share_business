package dbmodels

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Wallet struct {
	ID             uuid.UUID       `db:"id"`
	UserID         string          `db:"user_id"`
	OrganizationID uuid.UUID       `db:"organization_id"`
	Balance        decimal.Decimal `db:"balance"`
	IsDefault      bool            `db:"is_default"`
}
