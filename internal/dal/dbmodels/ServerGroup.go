package dbmodels

import uuid "github.com/satori/go.uuid"

type ServerGroup struct {
	ID             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	OrganizationID uuid.UUID `db:"organization_id"`
	IsDefault      bool      `db:"is_default"`
	Deleted        bool      `db:"deleted"`
}

type ServerGroupCreation struct {
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	OrganizationID uuid.UUID `db:"organization_id"`
	IsDefault      bool      `db:"is_default"`
}

type ServerGroupUpdate struct {
	Name        *string `db:"name"`
	Description *string `db:"description"`
	IsDefault   *bool   `db:"is_default"`
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID uuid.UUID
	IsManagedByMe  bool
}
