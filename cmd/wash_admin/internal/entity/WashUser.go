package entity

import (
	"wash_admin/internal/app/role"

	uuid "github.com/satori/go.uuid"
)

type WashUser struct {
	ID       uuid.UUID
	Identity string
	Role     role.Role
}
