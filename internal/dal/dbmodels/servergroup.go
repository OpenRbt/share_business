package dbmodels

import uuid "github.com/satori/go.uuid"

type ServerGroup struct {
	ID             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	OrganizationID uuid.UUID `db:"organization_id"`
	IsDefault      bool      `db:"is_default"`
	Deleted        bool      `db:"deleted"`
	Version        int       `db:"version"`
	CostPerDay     int64     `db:"cost_per_day"`
}

type ServerGroupCreation struct {
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	OrganizationID uuid.UUID `db:"organization_id"`
	IsDefault      bool      `db:"is_default"`
	CostPerDay     int64     `db:"cost_per_day"`
}

type ServerGroupUpdate struct {
	Name        *string `db:"name"`
	Description *string `db:"description"`
	CostPerDay  int64   `db:"cost_per_day"`
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID *uuid.UUID
}
