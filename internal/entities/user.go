package entities

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	User struct {
		ID      string
		Email   *string
		Name    *string
		Deleted bool
	}

	UserCreation struct {
		ID    string
		Email string
		Name  string
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
