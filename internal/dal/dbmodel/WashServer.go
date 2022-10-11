package dbmodel

import (
	"database/sql"
	"time"

	uuid "github.com/satori/go.uuid"
)

type WashServer struct {
	CreatedAt   *time.Time     `db:"created_at"`
	ModifiedAt  *time.Time     `db:"modified_at"`
	ID          uuid.UUID      `db:"id"`
	OwnerID     uuid.UUID      `db:"owner_id"`
	ServiceKey  sql.NullString `db:"service_key"`
	Name        string         `db:"name"`
	Description string         `db:"description"`
}
