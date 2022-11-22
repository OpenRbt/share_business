package dbmodels

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	WashKey     string    `db:"wash_key"`
}
