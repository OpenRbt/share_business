package dbmodels

import uuid "github.com/satori/go.uuid"

type MoneyReport struct {
	StationID    int           `db:"station_id"`
	Banknotes    int           `db:"banknotes"`
	CarsTotal    int           `db:"cars_total"`
	Coins        int           `db:"coins"`
	Electronical int           `db:"electronical"`
	Service      int           `db:"service"`
	Bonuses      int           `db:"bonuses"`
	SessionID    uuid.NullUUID `db:"session_id"`
	Processed    bool          `db:"processed"`
	UUID         uuid.NullUUID `db:"uuid"`
}
