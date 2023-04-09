package conversions

import (
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/openapi/models"
)

func UserFromDb(dbUser dbmodels.User) entity.User {
	return entity.User{
		Balance: dbUser.Balance.Decimal,
		ID:      dbUser.ID,
	}
}

func UserToRest(user entity.User) models.Profile {
	return models.Profile{
		Balance: user.Balance.IntPart(),
		ID:      user.ID,
	}
}
