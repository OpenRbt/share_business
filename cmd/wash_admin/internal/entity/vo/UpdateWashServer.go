package vo

import uuid "github.com/satori/go.uuid"

type UpdateWashServer struct {
	ID          uuid.UUID
	Name        *string
	Description *string
}
