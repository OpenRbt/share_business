package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type WashServer struct {
	CreatedAt   *time.Time
	ModifiedAt  *time.Time
	ID          uuid.UUID
	OwnerID     uuid.UUID
	ServiceKey  string
	Name        string
	Description string
}
