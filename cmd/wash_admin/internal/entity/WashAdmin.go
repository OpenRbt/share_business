package entity

import uuid "github.com/satori/go.uuid"

type WashAdmin struct {
	ID       uuid.UUID
	Identity string
	Role     string
}
