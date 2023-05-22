package dbmodels

import (
	uuid "github.com/satori/go.uuid"
)

type WashUser struct {
	ID       uuid.NullUUID `db:"id"`
	Identity string        `db:"identity"`
	Role     string        `db:"role"`
}

const (
	AdminRole    string = "admin"
	UserRole     string = "user"
	EngineerRole string = "engineer"
)

type UpdateUser struct {
	ID   uuid.NullUUID `db:"id"`
	Role string        `db:"role"`
}
