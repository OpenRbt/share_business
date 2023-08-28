package dbmodels

import "errors"

var (
	ErrNotFound       = errors.New("entity not found")
	ErrBadValue       = errors.New("bad value")
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrAlreadyExists  = errors.New("record already exists")
)
