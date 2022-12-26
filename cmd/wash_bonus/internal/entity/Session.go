package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Session struct {
	ID           uuid.UUID
	ConnectionID uuid.UUID
	User         *User
	Post         int64
	WashServer   WashServer

	PostBalance decimal.Decimal
	AddAmount   decimal.Decimal

	Closed bool
}
