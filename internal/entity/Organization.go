package entity

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Organization struct {
		ID          uuid.UUID
		Name        string
		Description string
		IsDefault   bool
		Deleted     bool
	}

	OrganizationCreation struct {
		Name        string
		Description string
	}

	OrganizationUpdate struct {
		Name        *string
		Description *string
	}

	OrganizationFilter struct {
		Pagination
		OrganizationIDs []uuid.UUID
		IsManagedByMe   bool
	}

	OrganizationSettings struct {
		ID                            uuid.UUID
		OrganizationID                uuid.UUID
		ReportsProcessingDelayMinutes int64
		BonusPercentage               int64
	}

	OrganizationSettingsUpdate struct {
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
	}
)
