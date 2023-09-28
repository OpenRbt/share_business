package dbmodels

import (
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type (
	User struct {
		ID      string  `db:"id"`
		Email   *string `db:"email"`
		Name    *string `db:"name"`
		Deleted bool    `db:"deleted"`
	}

	UserCreation struct {
		ID    string `db:"id"`
		Email string `db:"email"`
		Name  string `db:"name"`
	}

	UserUpdate struct {
		ID    string `db:"id"`
		Email string `db:"email"`
		Name  string `db:"name"`
	}

	UserPendingBalance struct {
		OrganizationID uuid.UUID       `db:"organization_id"`
		PendingBalance decimal.Decimal `db:"pending_balance"`
	}
)