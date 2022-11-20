package entity

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	ID          uuid.UUID
	Name        string
	Description string
	APIKey      string
}
