package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Session struct {
	ID     uuid.UUID
	PostID int64
	User   *User

	EnteredAmount decimal.Decimal
	ConsumeAmount decimal.Decimal

	Processed bool
}
