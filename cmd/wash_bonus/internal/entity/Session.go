package entity

import (
	uuid "github.com/satori/go.uuid"
)

type Session struct {
	ID         uuid.UUID
	User       *User
	Post       int64
	WashServer WashServer
	Finished   bool
}
