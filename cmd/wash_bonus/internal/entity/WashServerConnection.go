package entity

import uuid "github.com/satori/go.uuid"

type WashServerConnection struct {
	WashServer   WashServer
	ConnectionID uuid.UUID
}
