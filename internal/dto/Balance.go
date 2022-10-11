package dto

import (
	"strconv"
	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/dal/dbmodel"
	"wash-bonus/transport/rest/restapi/models"
)

func BalanceToRest(a *entity.Balance) *models.Balance {
	if a == nil {
		return nil
	}
	return &models.Balance{
		ID:      a.ID.String(),
		UserID:  a.UserId,
		Balance: strconv.FormatFloat(a.Balance, 'f', 6, 64),
	}
}

func BalanceFromDB(m dbmodel.Balance) *entity.Balance {
	return &entity.Balance{
		UserId:  m.UserID.String(),
		Balance: m.Balance.Float64,
	}
}
