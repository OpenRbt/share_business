package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
	rabbitEntities "washbonus/internal/infrastructure/rabbit/entities"

	uuid "github.com/satori/go.uuid"
)

func MoneyReportFromRabbit(r rabbitEntities.MoneyReport) (e entities.MoneyReport, err error) {
	sessionID, err := uuid.FromString(r.SessionID)
	if err != nil {
		return
	}

	reportUUID, err := uuid.FromString(r.UUID)
	if err != nil {
		return
	}

	e = entities.MoneyReport{
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

func MoneyReportFromDB(db dbmodels.MoneyReport) (e entities.MoneyReport) {
	e = entities.MoneyReport{
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

func MoneyReportToDB(e entities.MoneyReport) (db dbmodels.MoneyReport) {
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

func UserMoneyReportsFromDB(db []dbmodels.UserMoneyReport) []entities.UserMoneyReport {
	res := make([]entities.UserMoneyReport, len(db))

	for i, report := range db {
		res[i] = UserMoneyReportFromDB(report)
	}

	return res
}

func UserMoneyReportFromDB(db dbmodels.UserMoneyReport) entities.UserMoneyReport {
	return entities.UserMoneyReport{
		ID:             db.ID,
		StationID:      db.StationID,
		Banknotes:      db.Banknotes,
		CarsTotal:      db.CarsTotal,
		Coins:          db.Coins,
		Electronical:   db.Electronical,
		Service:        db.Service,
		Bonuses:        db.Bonuses,
		SessionID:      db.SessionID.UUID,
		OrganizationID: db.OrganizationID,
		User:           db.User,
		Processed:      db.Processed,
		UUID:           db.UUID.UUID,
	}
}

func UserMoneyReportToDB(e entities.UserMoneyReport) dbmodels.UserMoneyReport {
	return dbmodels.UserMoneyReport{
		ID:             e.ID,
		StationID:      e.StationID,
		Banknotes:      e.Banknotes,
		CarsTotal:      e.CarsTotal,
		Coins:          e.Coins,
		Electronical:   e.Electronical,
		Service:        e.Service,
		Bonuses:        e.Bonuses,
		SessionID:      uuid.NullUUID{UUID: e.SessionID, Valid: true},
		OrganizationID: e.OrganizationID,
		User:           e.User,
		Processed:      e.Processed,
		UUID:           uuid.NullUUID{UUID: e.UUID, Valid: true},
	}
}
