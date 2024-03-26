package entities

import (
	uuid "github.com/satori/go.uuid"
)

type (
	SimleOrganization struct {
		ID      uuid.UUID
		Name    string
		Deleted bool
	}

	Organization struct {
		ID                            uuid.UUID
		Name                          string
		DisplayName                   string
		Description                   string
		UTCOffset                     int32
		IsDefault                     bool
		ReportsProcessingDelayMinutes int64
		BonusPercentage               int64
		Deleted                       bool
		Version                       int
	}

	OrganizationCreation struct {
		Name                          string
		DisplayName                   string
		Description                   string
		UTCOffset                     *int32
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
	}

	OrganizationUpdate struct {
		Name                          *string
		DisplayName                   *string
		Description                   *string
		UTCOffset                     *int32
		ReportsProcessingDelayMinutes *int64
		BonusPercentage               *int64
	}

	OrganizationFilter struct {
		Pagination
		OrganizationIDs []uuid.UUID
	}
)
