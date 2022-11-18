package conversions

import (
	"wash_bonus/internal/dal/dbmodels"
	"wash_bonus/internal/entity"
	"wash_bonus/openapi/models"
)

func UserFromDb(dbUser dbmodels.User) entity.User {
	return entity.User{
		Active:   dbUser.Active,
		Balance:  dbUser.Balance.Decimal,
		ID:       dbUser.ID.UUID,
		Identity: dbUser.Identity,
	}
}

func UserToRest(user entity.User) models.Profile {
	return models.Profile{
		Active:  user.Active,
		Balance: user.Balance.String(),
		ID:      user.ID.String(),
	}
}
