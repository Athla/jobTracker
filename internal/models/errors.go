package models

import "errors"

var (
	ErrInvalidStatus           = errors.New("invalid status provided")
	ErrNoNameOrCompany         = errors.New("name and/or company are required")
	ErrInvalidCurrentStatus    = errors.New("invalid current status")
	ErrInvalidStatusTransition = errors.New("invalid status transition")
	ErrInvalidCredentials      = errors.New("invalid credentials")
	ErrUserNotFound            = errors.New("user not found")
	ErrInvalidToken            = errors.New("invalid token")
)
