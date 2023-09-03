package entity

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

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

	UserUpdateRole struct {
		ID   string
		Role Role
	}

	UserUpdate struct {
		ID    string
		Email string
		Name  string
	}

	UserPendingBalance struct {
		OrganizationID uuid.UUID
		PendingBalance decimal.Decimal
	}
)
