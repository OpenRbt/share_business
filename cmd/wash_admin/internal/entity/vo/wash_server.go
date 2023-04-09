package vo

import uuid "github.com/satori/go.uuid"

type RegisterWashServer struct {
	Title       string
	Description string
}

type UpdateWashServer struct {
	ID          uuid.UUID
	Title       *string
	Description *string
}
