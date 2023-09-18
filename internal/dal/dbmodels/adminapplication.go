package dbmodels

import uuid "github.com/satori/go.uuid"

type ApplicationStatus string

const (
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
	Pending  ApplicationStatus = "pending"
)

type (
	AdminApplication struct {
		ID          uuid.UUID         `db:"id"`
		AdminUserID string            `db:"admin_user_id"`
		Name        string            `db:"name"`
		Email       string            `db:"email"`
		Status      ApplicationStatus `db:"status"`
	}

	AdminApplicationCreation struct {
		AdminUserID string `db:"admin_user_id"`
		Name        string `db:"name"`
		Email       string `db:"email"`
	}

	AdminApplicationReview struct {
		Status         ApplicationStatus `db:"status"`
		OrganizationID *uuid.UUID        `db:"organization_id"`
	}

	AdminApplicationFilter struct {
		Status *ApplicationStatus `db:"status"`
		Pagination
	}
)
