package conversions

import (
	"washBonus/internal/dal/dbmodels"
	"washBonus/internal/entity"
	"washBonus/internal/entity/vo"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

func BalanceEventFromDB(dbEvent dbmodels.BalanceEvent) entity.BalanceEvent {
	return entity.BalanceEvent{
		ID:            dbEvent.ID.UUID,
		User:          dbEvent.User.UUID,
		WalletID:      dbEvent.WalletID,
		OperationKind: vo.OperationKind(dbEvent.OperationKind),
		OldAmount:     dbEvent.OldAmount.Decimal,
		NewAmount:     dbEvent.NewAmount.Decimal,
		WashServer:    dbEvent.WashServer.UUID,
		Session:       dbEvent.Session,
		Status:        dbEvent.Status,
		ErrorMsg:      dbEvent.ErrorMsg,
		Date:          dbEvent.Date,
	}
}

func BalanceEventToDB(event entity.BalanceEvent) dbmodels.BalanceEvent {
	return dbmodels.BalanceEvent{
		ID: uuid.NullUUID{
			UUID:  event.ID,
			Valid: true,
		},
		User: uuid.NullUUID{
			UUID:  event.User,
			Valid: true,
		},
		WalletID:      event.WalletID,
		OperationKind: 0,
		OldAmount: decimal.NullDecimal{
			Decimal: event.OldAmount,
			Valid:   true,
		},
		NewAmount: decimal.NullDecimal{
			Decimal: event.NewAmount,
			Valid:   true,
		},
		WashServer: uuid.NullUUID{
			UUID:  event.WashServer,
			Valid: true,
		},
		Session:  event.Session,
		Status:   event.Status,
		ErrorMsg: event.ErrorMsg,
		Date:     event.Date,
	}
}
