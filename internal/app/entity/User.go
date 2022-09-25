package entity

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Active     bool
	CreatedAt  *time.Time
	ID         uuid.UUID
	IdentityID string
	ModifiedAt *time.Time
}
