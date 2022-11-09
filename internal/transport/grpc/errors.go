package grpc

import "errors"

var (
	ErrVerifyFailed         = errors.New("verify failed")
	ErrNotFound             = errors.New("not found")
	ErrSessionAlreadyExists = errors.New("session already exist")
	ErrBadFormat            = errors.New("bad id format")
	ErrBadAmount            = errors.New("bad amount value")
)
