package dbmodels

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type User struct {
	ID       uuid.NullUUID       `db:"id"`
	Identity string              `db:"identity"`
	Balance  decimal.NullDecimal `db:"balance"`
	Active   bool                `db:"active"`
}
