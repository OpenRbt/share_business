package entity

import (
	uuid "github.com/satori/go.uuid"
)

type WashServer struct {
	ID          uuid.UUID
	Title       string
	Description string
	ServiceKey  string
	CreatedBy   string
}

type CreateWashServer struct {
	Title       string
	Description string
}

type UpdateWashServer struct {
	Title       *string
	Description *string
	Deleted     *bool
}
