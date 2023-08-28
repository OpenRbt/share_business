package entity

import (
	uuid "github.com/satori/go.uuid"
)

type Organization struct {
	ID          uuid.UUID
	Name        string
	Description string
	IsDefault   bool
	Deleted     bool
}

type OrganizationCreation struct {
	Name        string
	Description string
}

type OrganizationUpdate struct {
	Name        *string
	Description *string
}

type OrganizationFilter struct {
	Pagination
	OrganizationIDs []uuid.UUID
}
