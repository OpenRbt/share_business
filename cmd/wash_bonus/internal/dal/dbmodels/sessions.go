package dbmodels

import uuid "github.com/satori/go.uuid"

type CreateSession struct {
	WashServer uuid.NullUUID `db:"wash_server"`
}

type Session struct {
	ID         uuid.NullUUID `db:"id"`
	WashServer uuid.NullUUID `db:"wash_server"`
	User       uuid.NullUUID `db:"user"`
	PostID     int64         `db:"post_id"`
	Started    bool          `db:"started"`
	Finished   bool          `db:"finished"`
}
