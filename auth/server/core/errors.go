package core

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrInvalidSession     = errors.New("invalid session")
	ErrAccountNotVerified = errors.New("account not verified")
)
