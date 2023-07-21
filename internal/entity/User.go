package entity

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID             string
	Balance        decimal.Decimal
	PendingBalance decimal.Decimal
	Role           Role
	Deleted        bool
}

type UpdateUser struct {
	ID   string
	Role Role
}
