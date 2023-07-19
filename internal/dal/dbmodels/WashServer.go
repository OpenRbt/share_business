package dbmodels

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.NullUUID `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	CreatedBy   string        `db:"created_by"`
	ServiceKey  string        `db:"service_key"`
}

type RegisterWashServer struct {
	Title       string `db:"title"`
	Description string `db:"description"`
	ServiceKey  string `db:"service_key"`
	CreatedBy   string `db:"created_by"`
}

type UpdateWashServer struct {
	Name        *string `db:"name"`
	Description *string `db:"description"`
}
