package dbmodels

import uuid "github.com/satori/go.uuid"

type ServerGroup struct {
	ID                            uuid.UUID `db:"id"`
	OrganizationID                uuid.UUID `db:"organization_id"`
	Name                          string    `db:"name"`
	Description                   string    `db:"description"`
	ReportsProcessingDelayMinutes int64     `db:"processing_delay"`
	BonusPercentage               int64     `db:"bonus_percentage"`
	UTCOffset                     int32     `db:"utc_offset"`
	IsDefault                     bool      `db:"is_default"`
	Deleted                       bool      `db:"deleted"`
	Version                       int       `db:"version"`
}

type ServerGroupCreation struct {
	OrganizationID                uuid.UUID `db:"organization_id"`
	Name                          string    `db:"name"`
	Description                   string    `db:"description"`
	UTCOffset                     *int32    `db:"utc_offset"`
	IsDefault                     bool      `db:"is_default"`
	ReportsProcessingDelayMinutes *int64    `db:"processing_delay"`
	BonusPercentage               *int64    `db:"bonus_percentage"`
}

type ServerGroupUpdate struct {
	Name                          *string `db:"name"`
	Description                   *string `db:"description"`
	UTCOffset                     *int32  `db:"utc_offset"`
	ReportsProcessingDelayMinutes *int64  `db:"processing_delay"`
	BonusPercentage               *int64  `db:"bonus_percentage"`
}

type ServerGroupFilter struct {
	Pagination
	OrganizationID *uuid.UUID
}
