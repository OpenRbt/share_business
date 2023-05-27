package conversions

import (
	"github.com/OpenRbt/share_business/wash_rabbit/entity/session"
	uuid "github.com/satori/go.uuid"
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
)

func MoneyReportFromRabbit(r session.MoneyReport) (e entity.MoneyReport, err error) {
	sessionID, err := uuid.FromString(r.SessionID)

	reportUUID, err := uuid.FromString(r.UUID)
	if err != nil {
		return
	}

	e = entity.MoneyReport{
		StationID:    r.StationID,
		Banknotes:    r.Banknotes,
		CarsTotal:    r.CarsTotal,
		Coins:        r.Coins,
		Electronical: r.Electronical,
		Service:      r.Service,
		Bonuses:      r.Bonuses,
		UUID:         reportUUID,
	}
	if err == nil {
		e.SessionID = &sessionID
	}

	return
}

func MoneyReportFromDB(db dbmodels.MoneyReport) (e entity.MoneyReport) {
	e = entity.MoneyReport{
		StationID:    db.StationID,
		Banknotes:    db.Banknotes,
		CarsTotal:    db.CarsTotal,
		Coins:        db.Coins,
		Electronical: db.Electronical,
		Service:      db.Service,
		Bonuses:      db.Bonuses,
		Processed:    db.Processed,
		UUID:         db.UUID.UUID,
	}

	if db.SessionID.Valid {
		e.SessionID = &db.SessionID.UUID
	}
	return
}

func MoneyReportToDB(e entity.MoneyReport) (db dbmodels.MoneyReport) {
	db = dbmodels.MoneyReport{
		StationID:    e.StationID,
		Banknotes:    e.Banknotes,
		CarsTotal:    e.CarsTotal,
		Coins:        e.Coins,
		Electronical: e.Electronical,
		Service:      e.Service,
		Bonuses:      e.Bonuses,
		Processed:    e.Processed,
		UUID: uuid.NullUUID{
			UUID:  e.UUID,
			Valid: true,
		},
	}

	if e.SessionID != nil {
		db.SessionID = uuid.NullUUID{
			UUID:  *e.SessionID,
			Valid: true,
		}
	}

	return
}
