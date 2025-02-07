package models

import "errors"

var (
	ErrInvalidStatus           = errors.New("Invalid status provided.")
	ErrNoNameOrCompany         = errors.New("Name and/or Company are required.")
	ErrInvalidCurrentStatus    = errors.New("Invalid status provided.")
	ErrInvalidStatusTransition = errors.New("Current status transition is invalid.")
)
