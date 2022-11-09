package dal

import (
	"database/sql"
	"time"
)

const (
	sqlGetBalance = `
	SELECT
		id,
		user_id,
		balance
	FROM
		bonus_balance
	WHERE
		id=:id AND
		NOT deleted
	`

	sqlAddBalance = `
	INSERT INTO bonus_balance(
		user_id,
		balance
	) VALUES (
		:user_id,
		:balance
	)
	RETURNING
		id
	`

	sqlDeleteBalance = `
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

	sqlEditBalance = `
	UPDATE
		bonus_balance
	SET
		balance=:balance
	WHERE
		id=:id AND
		NOT deleted
	`
)

type (
	argGetBalance struct {
		ID sql.NullString `db:"id"`
	}

	argAddBalance struct {
		UserID  string  `db:"user_id"`
		balance float64 `db:"balance"`
	}

	argEditBalance struct {
		ID      string  `db:"id"`
		balance float64 `db:"balance"`
	}
	argDeleteBalance struct {
		ID        string     `db:"id"`
		DeletedAt *time.Time `db:"deleted_at"`
		DeletedBy string     `db:"deleted_by"`
	}
)
