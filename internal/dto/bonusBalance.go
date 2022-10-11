package dto

import (
	"strconv"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/dal/dbmodel"
	"wash-bonus/transport/rest/restapi/models"
)

func ApiBonusBalance(a *entity.BonusBalance) *models.Balance {
	if a == nil {
		return nil
	}
	return &models.Balance{
		ID:      a.ID,
		UserID:  a.UserId,
		Balance: strconv.FormatFloat(a.Balance, 'f', 6, 64),
	}
}

func AppBonusBalance(m dbmodel.BonusBalance) *entity.BonusBalance {
	return &entity.BonusBalance{
		UserId:  m.UserID.String(),
		Balance: m.Balance.Float64,
	}
}
