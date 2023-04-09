package dbmodels

import uuid "github.com/satori/go.uuid"

type RegisterWashServer struct {
	Title       string        `db:"title"`
	Description string        `db:"description"`
	Owner       uuid.NullUUID `db:"owner"`
	ServiceKey  string        `db:"service_key"`
}
