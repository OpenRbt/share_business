package entity

import "errors"

var (
	ErrNotFound        = errors.New("entity not found")
	ErrAccessDenied    = errors.New("access denied")
	ErrProfileInactive = errors.New("profile inactive")

	ErrBadValue = errors.New("bad value")

	ErrForbidden      = errors.New("forbidden operation")
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrSessionNoUser  = errors.New("session without user")

	ErrWashServerConnectionInit = errors.New("failed to init wash server connection")
	ErrWashServerNotFound       = errors.New("wash server with specified key not exists")
	ErrEmptyWashServerKey       = errors.New("wash server key must be not empty")
	ErrWashServerKeyAlreadyUsed = errors.New("wash server key already in use")
)
