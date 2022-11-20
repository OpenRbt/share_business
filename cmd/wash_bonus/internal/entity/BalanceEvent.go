package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	"time"
	"wash_bonus/internal/entity/vo"
)

type BalanceEvent struct {
	ID            uuid.UUID
	User          uuid.UUID
	OperationKind vo.OperationKind
	OldAmount     decimal.Decimal
	NewAmount     decimal.Decimal
	WashServer    uuid.UUID
	Session       string
	Status        bool
	ErrorMsg      string
	Date          time.Time
}
