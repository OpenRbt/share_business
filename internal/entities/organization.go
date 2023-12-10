package entities

import (
	uuid "github.com/satori/go.uuid"
)

type (
	Organization struct {
		ID                            uuid.UUID
		Name                          string
		DisplayName                   string
		Description                   string
		IsDefault                     bool
		ReportsProcessingDelayMinutes int64
		BonusPercentage               int64
		Deleted                       bool
		Version                       int
		CostPerDay                    int64
	}

	OrganizationCreation struct {
		Name                          string
		DisplayName                   string
		Description                   string
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
		CostPerDay                    int64
	}

	OrganizationUpdate struct {
		Name                          *string
		DisplayName                   *string
		Description                   *string
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
		CostPerDay                    int64
	}

	OrganizationFilter struct {
		Pagination
		OrganizationIDs []uuid.UUID
	}
)
