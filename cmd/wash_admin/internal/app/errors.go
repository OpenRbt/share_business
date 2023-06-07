package app

import "errors"

var (
	ErrAccessDenied   = errors.New("access denied")
	ErrNotFound       = errors.New("entity not found")
	ErrPrepareMessage = errors.New("failed to prepare rabbit message")
)
