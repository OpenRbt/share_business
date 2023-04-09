package conversions

import (
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
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
		ID:           db.ID,
		StationID:    db.StationID,
		Banknotes:    db.Banknotes,
		CarsTotal:    db.CarsTotal,
		Coins:        db.Coins,
		Electronical: db.Electronical,
		Service:      db.Service,
		Bonuses:      db.Bonuses,
		SessionID:    db.SessionID.UUID,
		User:         db.User,
		Processed:    db.Processed,
	}
}
