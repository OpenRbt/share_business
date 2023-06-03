package dbmodels

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.NullUUID `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	CreatedBy   string        `db:"created_by"`
	ServiceKey  string        `db:"service_key"`
}
