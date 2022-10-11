package dbmodel

import (
	"database/sql"

	uuid "github.com/satori/go.uuid"
)

type Balance struct {
	ID      uuid.UUID       `db:"id"`
	UserID  uuid.UUID       `db:"user_id"`
	Balance sql.NullFloat64 `db:"balance"`
}
