package dbmodels

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type BalanceEvent struct {
	ID            uuid.NullUUID       `db:"id"`
	User          uuid.NullUUID       `db:"user"`
	WalletID      uuid.UUID           `db:"wallet_id"`
	OperationKind int64               `db:"operation_kind"`
	OldAmount     decimal.NullDecimal `db:"old_amount"`
	NewAmount     decimal.NullDecimal `db:"new_amount"`
	WashServer    uuid.NullUUID       `db:"wash_server"`
	Session       string              `db:"session"`
	Status        bool                `db:"status"`
	ErrorMsg      string              `db:"error_msg"`
	Date          time.Time           `db:"date"`
}

type UpdateBalanceEvent struct {
	UserID    string
	WalletID  uuid.UUID
	OldAmount decimal.Decimal
	NewAmount decimal.Decimal
}
