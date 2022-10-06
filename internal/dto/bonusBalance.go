package dto

import (
	"strconv"
	"wash-bonus/internal/api/restapi/models"
	"wash-bonus/internal/app/entity"
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
