package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"

	uuid "github.com/satori/go.uuid"
)

func UserMoneyReportsFromDB(db []dbmodels.UserMoneyReport) []entity.UserMoneyReport {
	res := make([]entity.UserMoneyReport, len(db))

	for i, report := range db {
		res[i] = UserMoneyReportFromDB(report)
	}

	return res
}

func UserMoneyReportFromDB(db dbmodels.UserMoneyReport) entity.UserMoneyReport {
	return entity.UserMoneyReport{
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

func UserMoneyReportToDB(e entity.UserMoneyReport) dbmodels.UserMoneyReport {
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
