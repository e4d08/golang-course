package domain

import "errors"

var (
	ErrRepositoryNotFound = errors.New("repository not found")
	ErrInternal           = errors.New("internal server error")
)
