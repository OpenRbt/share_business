package entity

import uuid "github.com/satori/go.uuid"

type User struct {
	ID              string
	Role            Role
	OrganizationIDs []uuid.UUID
	Deleted         bool
}

type UserUpdate struct {
	ID   string
	Role Role
}
