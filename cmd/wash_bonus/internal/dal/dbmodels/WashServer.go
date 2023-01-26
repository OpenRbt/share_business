package dbmodels

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.NullUUID `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
}
