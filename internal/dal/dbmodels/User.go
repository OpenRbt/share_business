package dbmodels

import (
	"github.com/shopspring/decimal"
)

type User struct {
	ID      string              `db:"id"`
	Balance decimal.NullDecimal `db:"balance"`
	Role    string              `db:"role"`
	Deleted bool                `db:"deleted"`
}

const (
	AdminRole    string = "admin"
	UserRole     string = "user"
	EngineerRole string = "engineer"
)

type UpdateUser struct {
	ID   string `db:"id"`
	Role string `db:"role"`
}
