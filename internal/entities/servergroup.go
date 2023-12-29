package entities

import (
	uuid "github.com/satori/go.uuid"
)

type ServerGroup struct {
	ID             uuid.UUID
	OrganizationID uuid.UUID
	Name           string
	Description    string
	UTCOffset      int32
	IsDefault      bool
	Deleted        bool
	Version        int
}

type ServerGroupCreation struct {
	OrganizationID uuid.UUID
	Name           string
	Description    string
	UTCOffset      *int32
}

type ServerGroupUpdate struct {
	Name        *string
	Description *string
	UTCOffset   *int32
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID *uuid.UUID
}
