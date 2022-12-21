package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type Session struct {
	ConnectionID uuid.UUID
	UserID       *uuid.UUID
	Post         int64
	WashServer   WashServer

	PostBalance decimal.Decimal
}
