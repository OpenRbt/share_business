package dal

import (
	"database/sql"
	"time"

	"wash-bonus/internal/app"

	"wash-bonus/internal/app/entity"
	"wash-bonus/internal/dal/dbmodel"
	"wash-bonus/internal/dto"
)

func (a *Repo) GetBalance(id string) (*entity.Balance, error) {
	var m dbmodel.Balance
	if err := a.db.NamedGet(&m, sqlGetBalance, argGetBalance{
		ID: newNullUUID(id),
	}); err != nil {
		if err == sql.ErrNoRows {
			return nil, app.ErrNotFound
		}
		return nil, err
	}
	return dto.BalanceFromDB(m), nil
}

func (a *Repo) AddBalance(userID string, balance float64) (*entity.Balance, error) {
	_, err := a.db.NamedExec(sqlAddBalance, argAddBalance{
		UserID:  userID,
		balance: balance,
	})
	if err != nil {
		return nil, err
	}
	return a.GetBalance(userID)
}

func (a *Repo) EditBalance(id string, balance float64) error {
	res, err := a.db.NamedExec(sqlEditBalance, argEditBalance{
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

func (a *Repo) DeleteBalance(id string, userId string) error {
	t := time.Now()
	res, err := a.db.NamedExec(sqlDeleteBalance, argDeleteBalance{
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
