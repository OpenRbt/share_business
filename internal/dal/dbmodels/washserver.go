package dbmodels

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID             uuid.UUID `db:"id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	CreatedBy      string    `db:"created_by"`
	ServiceKey     string    `db:"service_key"`
	GroupID        uuid.UUID `db:"group_id"`
	OrganizationID uuid.UUID `db:"organization_id"`
}

type WashServerCreation struct {
	Title       string     `db:"title"`
	Description string     `db:"description"`
	ServiceKey  string     `db:"service_key"`
	CreatedBy   string     `db:"created_by"`
	GroupID     *uuid.UUID `db:"group_id"`
}

type WashServerUpdate struct {
	Name        *string `db:"title"`
	Description *string `db:"description"`
}

type WashServerFilter struct {
	Pagination
	OrganizationID *uuid.UUID
	GroupID        *uuid.UUID
}
