package dbmodels

import uuid "github.com/satori/go.uuid"

type Organization struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
	IsDefault   bool      `db:"is_default"`
	Deleted     bool      `db:"deleted"`
}

type OrganizationCreation struct {
	Name        string `db:"name"`
	Description string `db:"description"`
}

type OrganizationUpdate struct {
	Name        *string `db:"name"`
	Description *string `db:"description"`
}

type OrganizationFilter struct {
	Pagination
	OrganizationIDs []uuid.UUID
}
