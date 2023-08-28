package entity

import (
	uuid "github.com/satori/go.uuid"
)

type WashServer struct {
	ID             uuid.UUID
	Title          string
	Description    string
	ServiceKey     string
	CreatedBy      string
	GroupID        uuid.UUID
	OrganizationID uuid.UUID
}

type WashServerCreation struct {
	Title       string
	Description string
	GroupID     uuid.NullUUID
}

type WashServerUpdate struct {
	Title       *string
	Description *string
}

type WashServerFilter struct {
	Pagination
	OrganizationID uuid.UUID
	GroupID        uuid.UUID
}
