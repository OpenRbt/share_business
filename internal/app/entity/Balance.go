package entity

import "github.com/google/uuid"

type Balance struct {
	ID      uuid.UUID
	UserId  string
	Balance float64
}
