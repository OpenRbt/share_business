package entity

import "errors"

var (
	ErrNotFound     = errors.New("entity not found")
	ErrUserNotOwner = errors.New("user not owner")
)
