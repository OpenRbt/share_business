package dbmodels

import uuid "github.com/satori/go.uuid"

type UserMoneyReport struct {
	ID             int64         `db:"id"`
	StationID      int           `db:"station_id"`
	Banknotes      int           `db:"banknotes"`
	CarsTotal      int           `db:"cars_total"`
	Coins          int           `db:"coins"`
	Electronical   int           `db:"electronical"`
	Service        int           `db:"service"`
	Bonuses        int           `db:"bonuses"`
	SessionID      uuid.NullUUID `db:"session_id"`
	OrganizationID uuid.UUID     `db:"organization_id"`
	User           string        `db:"user"`
	Processed      bool          `db:"processed"`
	UUID           uuid.NullUUID `db:"uuid"`
}
