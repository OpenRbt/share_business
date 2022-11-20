package entity

import "errors"

var (
	ErrNotFound        = errors.New("entity not found")
	ErrAccessDenied    = errors.New("access denied")
	ErrProfileInactive = errors.New("profile inactive")
)
