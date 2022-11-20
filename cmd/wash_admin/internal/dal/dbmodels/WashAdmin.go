package dbmodels

import uuid "github.com/satori/go.uuid"

type WashAdmin struct {
	ID uuid.NullUUID `db:"id"`
	Identity string `db:"identity"`
}