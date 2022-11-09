package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	Active     bool
	CreatedAt  *time.Time
	ID         uuid.UUID
	IdentityID string
	ModifiedAt *time.Time
}
