package domain

import "errors"

var (
	ErrRepositoryNotFound = errors.New("repository not found")
	ErrInternal           = errors.New("internal server error")
	ErrInvalidArgument    = errors.New("invalid argument")
	ErrUnknown            = errors.New("unknown error")
)
