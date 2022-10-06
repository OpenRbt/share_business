package dal

import (
	"database/sql"
	"time"

	"wash-bonus/internal/app"

	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/dal/dbmodel"
)

func (a *Repo) GetBonusBalance(id string) (*entity.BonusBalance, error) {
	var m dbmodel.BonusBalance
	if err := a.db.NamedGet(&m, sqlGetBonusBalance, argGetBonusBalance{
		ID: newNullUUID(id),
	}); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, err
	}
	return appBonusBalance(m), nil
}

func (a *Repo) AddBonusBalance(userID string, balance float64) error {
	_, err := a.db.NamedExec(sqlAddBonusBalance, argAddBonusBalance{
		UserID:  userID,
		balance: balance,
	})
	if err != nil {
		return err
	}
	return nil
}

func (a *Repo) EditBonusBalance(id string, balance float64) error {
	res, err := a.db.NamedExec(sqlEditBonusBalance, argEditBonusBalance{
		ID:      id,
		balance: balance,
	})
	if err != nil {
		return err
	}

	if count, _ := res.RowsAffected(); count == 0 {
		return app.ErrNotFound
	}

	return nil
}

func (a *Repo) DeleteBonusBalance(id string, userId string) error {
	t := time.Now()
	res, err := a.db.NamedExec(sqlDeleteBonusBalance, argDeleteBonusBalance{
		ID:        id,
		DeletedAt: &t,
		DeletedBy: userId,
	})
	if err != nil {
		return err
	}
	if count, _ := res.RowsAffected(); count == 0 {
		return app.ErrNotFound
	}

	return nil
}

func appBonusBalance(m dbmodel.BonusBalance) *entity.BonusBalance {
	return &entity.BonusBalance{
		UserId:  m.UserID.String(),
		Balance: m.Balance.Float64,
	}
}
