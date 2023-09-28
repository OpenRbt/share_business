package dbmodels

import "errors"

var (
	ErrNotFound       = errors.New("entity not found")
	ErrBadRequest     = errors.New("bad request")
	ErrNotEnoughMoney = errors.New("not enough money")
	ErrAlreadyExists  = errors.New("record already exists")
)
