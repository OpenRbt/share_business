package entity

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	Id          uuid.UUID
	Name        string
	Description string
}
