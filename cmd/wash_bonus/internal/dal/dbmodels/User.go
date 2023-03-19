package dbmodels

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID      string              `db:"id"`
	Balance decimal.NullDecimal `db:"balance"`
	Deleted bool                `db:"deleted"`
	// TODO: Add user activity management
}
