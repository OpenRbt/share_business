package dal

import (
	"database/sql"
	"strings"
	"time"

	"wash-bonus/internal/app"

	"github.com/google/uuid"
)

type BonusBalance struct {
	ID      uuid.UUID       `db:"id"`
	UserID  uuid.UUID       `db:"user_id`
	balance sql.NullFloat64 `db:"balance"`
}

func (a *Repo) GetBonusBalance(id string) (*app.BonusBalance, error) {
	var m BonusBalance
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

func (a *Repo) AddBonusBalance(userID string, balance float64) (*app.BonusBalance, error) {
	ID := uuid.New().String()
	if err := a.db.NamedGet(&ID, sqlAddBonusBalance, argAddBonusBalance{
		ID:      ID,
		UserID:  userID,
		balance: balance,
	}); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return nil, app.ErrDuplicateID
		}
		return nil, err
	}
	return a.GetBonusBalance(ID)
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

func appBonusBalance(m BonusBalance) *app.BonusBalance {
	return &app.BonusBalance{
		UserId:  m.UserID.String(),
		Balance: m.balance.Float64,
	}
}
