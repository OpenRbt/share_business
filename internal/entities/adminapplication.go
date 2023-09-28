package entities

import uuid "github.com/satori/go.uuid"

type ApplicationStatus string

const (
	Accepted ApplicationStatus = "accepted"
	Rejected ApplicationStatus = "rejected"
	Pending  ApplicationStatus = "pending"
)

type (
	FirebaseUser struct {
		ID    string
		Name  string
		Email string
	}

	AdminApplication struct {
		ID     uuid.UUID
		User   FirebaseUser
		Status ApplicationStatus
	}

	AdminApplicationCreation struct {
		User FirebaseUser
	}

	AdminApplicationReview struct {
		Status         ApplicationStatus
		OrganizationID *uuid.UUID
		Role           *Role
	}

	AdminApplicationFilter struct {
		Status *ApplicationStatus
		Pagination
	}
)
