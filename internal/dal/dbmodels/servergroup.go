package dbmodels

import uuid "github.com/satori/go.uuid"

type ServerGroup struct {
	ID             uuid.UUID `db:"id"`
	OrganizationID uuid.UUID `db:"organization_id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	UTCOffset      int32     `db:"utc_offset"`
	IsDefault      bool      `db:"is_default"`
	Deleted        bool      `db:"deleted"`
	Version        int       `db:"version"`
}

type ServerGroupCreation struct {
	OrganizationID uuid.UUID `db:"organization_id"`
	Name           string    `db:"name"`
	Description    string    `db:"description"`
	UTCOffset      *int32    `db:"utc_offset"`
	IsDefault      bool      `db:"is_default"`
}

type ServerGroupUpdate struct {
	Name        *string `db:"name"`
	Description *string `db:"description"`
	UTCOffset   *int32  `db:"utc_offset"`
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID *uuid.UUID
}
