package entity

import uuid "github.com/satori/go.uuid"

type Balance struct {
	ID      uuid.UUID
	UserId  string
	Balance float64
}
