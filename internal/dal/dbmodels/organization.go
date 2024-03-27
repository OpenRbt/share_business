package dbmodels

import (
	uuid "github.com/satori/go.uuid"
)

type (
	SimleOrganization struct {
		ID      uuid.UUID `db:"org_id"`
		Name    string    `db:"org_name"`
		Deleted bool      `db:"org_deleted"`
	}

	Organization struct {
		ID                            uuid.UUID `db:"id"`
		Name                          string    `db:"name"`
		DisplayName                   string    `db:"display_name"`
		Description                   string    `db:"description"`
		UTCOffset                     int32     `db:"utc_offset"`
		IsDefault                     bool      `db:"is_default"`
		ReportsProcessingDelayMinutes int64     `db:"processing_delay"`
		BonusPercentage               int64     `db:"bonus_percentage"`
		Deleted                       bool      `db:"deleted"`
		Version                       int       `db:"version"`
	}

	OrganizationCreation struct {
		Name                          string `db:"name"`
		DisplayName                   string `db:"display_name"`
		Description                   string `db:"description"`
		UTCOffset                     *int32 `db:"utc_offset"`
		ReportsProcessingDelayMinutes *int64 `db:"processing_delay"`
		BonusPercentage               *int64 `db:"bonus_percentage"`
	}

	OrganizationUpdate struct {
		Name                          *string `db:"name"`
		DisplayName                   *string `db:"display_name"`
		Description                   *string `db:"description"`
		UTCOffset                     *int32  `db:"utc_offset"`
		ReportsProcessingDelayMinutes *int64  `db:"processing_delay"`
		BonusPercentage               *int64  `db:"bonus_percentage"`
	}

	OrganizationFilter struct {
		Pagination
		OrganizationIDs []uuid.UUID
	}
)
