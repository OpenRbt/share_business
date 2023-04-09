package vo

import uuid "github.com/satori/go.uuid"

type WashServerUpdate struct {
	ID          uuid.UUID
	Title       *string
	Description *string
	Deleted     *bool
}
