package dbmodels

import uuid "github.com/satori/go.uuid"

type DeleteWashServer struct {
	ID uuid.NullUUID `db:"id"`
}
