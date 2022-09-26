package vo

import uuid "github.com/satori/go.uuid"

type WashServerUpdate struct {
	ServiceKey  string
	Name        string
	Description string
	OwnerID     uuid.UUID
}
