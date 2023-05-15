package app

import (
	uuid "github.com/satori/go.uuid"
)

type WashUser struct {
	ID       uuid.UUID
	Identity string
	Role     Role
}
