package core

import "errors"

var (
	ErrTestNotFound            = errors.New("test not found")
	ErrUserTestSessionNotFound = errors.New("user test session not found")
)
