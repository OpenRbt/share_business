package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	AdminUser struct {
		ID                      string     `db:"id"`
		Email                   string     `db:"email"`
		Name                    string     `db:"name"`
		Role                    Role       `db:"role"`
		Version                 int        `db:"version"`
		OrganizationID          *uuid.UUID `db:"organization_id"`
		OrganizationName        *string    `db:"organization_name"`
		OrganizationDisplayName *string    `db:"organization_display_name"`
		OrganizationDescription *string    `db:"organization_description"`
		OrganizationDeleted     *bool      `db:"organization_deleted"`
	}

	AdminUserCreation struct {
		ID             string     `db:"id"`
		Email          string     `db:"email"`
		Name           string     `db:"name"`
		OrganizationId *uuid.UUID `db:"organization_id"`
	}

	AdminUserUpdate struct {
		ID    string  `db:"id"`
		Email *string `db:"email"`
		Name  *string `db:"name"`
	}

	AdminUserRoleUpdate struct {
		ID   string `db:"id"`
		Role Role   `db:"role"`
	}

	AdminUserFilter struct {
		Pagination
		Role      *Role
		IsBlocked *bool
	}
)

type Role string

const (
	SystemManagerRole Role = "system_manager"
	AdminRole         Role = "admin"
	NoAccessRole      Role = "no_access"
)
