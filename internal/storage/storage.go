package storage

import "errors"

var (
	ErrURLExists   = errors.New("Todo already exists")
	ErrURLNotFound = errors.New("Todo not found")
)
