package entities

import uuid "github.com/satori/go.uuid"

type (
	AdminUser struct {
		ID             string
		Email          *string
		Name           *string
		Role           Role
		OrganizationID *uuid.UUID
	}

	AdminUserCreation struct {
		ID             string
		Email          string
		Name           string
		OrganizationID *uuid.UUID
	}

	AdminUserUpdate struct {
		ID    string
		Email *string
		Name  *string
	}

	AdminUserRoleUpdate struct {
		ID   string
		Role Role
	}

	AdminUserFilter struct {
		Pagination
		Role      *Role
		IsBlocked *bool
	}
)

type Role string

const (
	SystemManagerRole Role = "systemManager"
	AdminRole         Role = "admin"
	NoAccessRole      Role = "noAccess"
)