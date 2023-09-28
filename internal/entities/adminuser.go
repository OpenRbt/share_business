package entities

import uuid "github.com/satori/go.uuid"

type (
	AdminUser struct {
		ID           string
		Email        *string
		Name         *string
		Role         Role
		Organization *AdminOrganization
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

	AdminOrganization struct {
		ID          uuid.UUID
		Name        string
		DisplayName string
		Description string
		Deleted     bool
	}
)

type Role string

const (
	SystemManagerRole Role = "systemManager"
	AdminRole         Role = "admin"
	NoAccessRole      Role = "noAccess"
)
