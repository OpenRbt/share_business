package entities

import (
	"time"
	"washbonus/internal/entities/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type BalanceEvent struct {
	ID            uuid.UUID
	User          uuid.UUID
	WalletID      uuid.UUID
	OperationKind vo.OperationKind
	OldAmount     decimal.Decimal
	NewAmount     decimal.Decimal
	WashServer    uuid.UUID
	Session       string
	Status        bool
	ErrorMsg      string
	Date          time.Time
}
