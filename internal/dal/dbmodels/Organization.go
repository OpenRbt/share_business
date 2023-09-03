package dbmodels

import uuid "github.com/satori/go.uuid"

type (
	Organization struct {
		ID          uuid.UUID `db:"id"`
		Name        string    `db:"name"`
		Description string    `db:"description"`
		IsDefault   bool      `db:"is_default"`
		Deleted     bool      `db:"deleted"`
	}

	OrganizationCreation struct {
		Name        string `db:"name"`
		Description string `db:"description"`
	}

	OrganizationUpdate struct {
		Name        *string `db:"name"`
		Description *string `db:"description"`
	}

	OrganizationFilter struct {
		Pagination
		OrganizationIDs []uuid.UUID
		IsManagedByMe   bool
	}

	OrganizationSettings struct {
		ID                            uuid.UUID `db:"id"`
		OrganizationID                uuid.UUID `db:"organization_id"`
		ReportsProcessingDelayMinutes int64     `db:"processing_delay"`
		BonusPercentage               int64     `db:"bonus_percentage"`
	}

	OrganizationSettingsUpdate struct {
		ReportsProcessingDelayMinutes *int64 `db:"processing_delay"`
		BonusPercentage               *int64 `db:"bonus_percentage"`
	}
)
