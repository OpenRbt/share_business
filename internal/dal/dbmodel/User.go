package dbmodel

import (
	"database/sql"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Active     bool           `db:"active"`
	CreatedAt  *time.Time     `db:"created_at"`
	ID         uuid.UUID      `db:"id"`
	IdentityID sql.NullString `db:"identity_id"`
	ModifiedAt *time.Time     `db:"modified_at"`
}
