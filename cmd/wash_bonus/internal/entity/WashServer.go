package entity

import uuid "github.com/satori/go.uuid"

type WashServer struct {
	Id          uuid.UUID
	Title       string
	Description string
}
