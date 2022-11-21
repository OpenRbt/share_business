package dbmodels

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.NullUUID `db:"id"`
	Name        string        `db:"name"`
	Description string        `db:"description"`
	APIKey      string        `db:"wash_key"`
	Owner       uuid.NullUUID `db:"owner"`
}
