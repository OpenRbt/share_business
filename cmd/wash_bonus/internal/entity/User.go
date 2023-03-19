package entity

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID      string
	Balance decimal.Decimal
	Deleted bool
}
