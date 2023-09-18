package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	AdminUser struct {
		ID             string     `db:"id"`
		Email          *string    `db:"email"`
		Name           *string    `db:"name"`
		Role           Role       `db:"role"`
		OrganizationID *uuid.UUID `db:"organization_id"`
		Deleted        bool       `db:"deleted"`
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
)

type Role string

const (
	SystemManagerRole Role = "system_manager"
	AdminRole         Role = "admin"
)
