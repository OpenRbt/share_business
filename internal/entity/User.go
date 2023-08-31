package entity

import uuid "github.com/satori/go.uuid"

type (
	User struct {
		ID              string
		Email           *string
		Name            *string
		Role            Role
		OrganizationIDs []uuid.UUID
		Deleted         bool
	}

	UserCreation struct {
		ID    string
		Email string
		Name  string
	}

	UserUpdate struct {
		ID   string
		Role Role
	}
)
