package app

import "errors"

var (
	ErrAccessDenied = errors.New("access denied")
	ErrNotFound     = errors.New("entity not found")
	ErrUserNotOwner = errors.New("user not owner")
)
