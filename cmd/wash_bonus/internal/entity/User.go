package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	Active   bool
	Balance  decimal.Decimal
	ID       uuid.UUID
	Identity string
}
