package entities

import (
	uuid "github.com/satori/go.uuid"
)

type ServerGroup struct {
	ID             uuid.UUID
	Name           string
	Description    string
	OrganizationID uuid.UUID
	IsDefault      bool
	Deleted        bool
	Version        int
	CostPerDay     int64
}

type ServerGroupCreation struct {
	Name           string
	Description    string
	OrganizationID uuid.UUID
	CostPerDay     int64
}

type ServerGroupUpdate struct {
	Name        *string
	Description *string
	CostPerDay  int64
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID *uuid.UUID
}
