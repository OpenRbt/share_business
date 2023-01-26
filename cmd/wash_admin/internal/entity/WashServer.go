package entity

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.UUID
	Title       string
	Description string
	ServiceKey  string
	Owner       uuid.UUID
}
