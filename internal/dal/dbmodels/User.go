package dbmodels

import (
	"github.com/shopspring/decimal"
)

type (
	User struct {
		ID                 string              `db:"id"`
		Email              *string             `db:"email"`
		Name               *string             `db:"name"`
		Balance            decimal.NullDecimal `db:"balance"`
		Role               string              `db:"role"`
		OrganizationIDsRaw string              `db:"organization_ids"`
		Deleted            bool                `db:"deleted"`
	}

	UserCreation struct {
		ID    string `db:"id"`
		Email string `db:"email"`
		Name  string `db:"name"`
	}

	UserUpdate struct {
		ID   string `db:"id"`
		Role string `db:"role"`
	}
)

const (
	AdminRole    string = "admin"
	UserRole     string = "user"
	EngineerRole string = "engineer"
)
