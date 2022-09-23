package dal

import (
	"database/sql"
	"time"
)

const (
	sqlGetBonusBalance = `
	SELECT
		id,
		user_id,
		balance,
	FROM
		bonus_balance
	WHERE
		id=:id AND
		NOT deleted
	`

	sqlAddBonusBalance = `
	INSERT INTO bonus_balance(
		id,
		user_id,
		balance,
	) VALUES (
		:id,
		:user_id,
		:balance,
	)
	RETURNING
		id
	`

	sqlDeleteBonusBalance = `
	UPDATE
		bonus_balance
	SET
		deleted=true,
		deleted_at=:deleted_at,
		deleted_by=:deleted_by
	WHERE
		id=:id AND
		NOT deleted
	`

	sqlEditBonusBalance = `
	UPDATE
		bonus_balance
	SET
		balance=:balance,
	WHERE
		id=:id AND
		NOT deleted
	`
)

type (
	argGetBonusBalance struct {
		ID sql.NullString `db:"id"`
	}

	argAddBonusBalance struct {
		ID      string  `db:"id"`
		UserID  string  `db:"user_id"`
		balance float64 `db:"balance"`
	}

	argEditBonusBalance struct {
		ID      string  `db:"id"`
		balance float64 `db:"balance"`
	}
	argDeleteBonusBalance struct {
		ID        string     `db:"id"`
		DeletedAt *time.Time `db:"deleted_at"`
		DeletedBy string     `db:"deleted_by"`
	}
)
