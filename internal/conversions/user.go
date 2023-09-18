package conversions

import (
	"washbonus/internal/dal/dbmodels"
	"washbonus/internal/entities"
)

func UserFromDb(dbUser dbmodels.User) entities.User {
	return entities.User{
		ID:      dbUser.ID,
		Name:    dbUser.Name,
		Email:   dbUser.Email,
		Deleted: dbUser.Deleted,
	}
}

func UserUpdateToDB(user entities.UserUpdate) dbmodels.UserUpdate {
	return dbmodels.UserUpdate{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}
}

func UserCreationToDB(e entities.UserCreation) dbmodels.UserCreation {
	return dbmodels.UserCreation{
		ID:    e.ID,
		Email: e.Email,
		Name:  e.Name,
	}
}

func UserPendingBalancesFromDB(e []dbmodels.UserPendingBalance) []entities.UserPendingBalance {
	balances := make([]entities.UserPendingBalance, len(e))

	for idx, balance := range e {
		balances[idx].OrganizationID = balance.OrganizationID
		balances[idx].PendingBalance = balance.PendingBalance
	}

	return balances
}
