package businesses

import "errors"

var (
	ErrInternalServer = errors.New("something's gone wrong, contact administrator")

	ErrNotFound = errors.New("data not found")

	ErrIDNotFound = errors.New("id not found")

	ErrThreadsIDResource = errors.New("(ThreadsID) not found or empty")

	ErrThreadsTitleResource = errors.New("(ThreadsTitle) not found or empty")

	ErrCategoryNotFound = errors.New("category not found")

	ErrDuplicateData = errors.New("duplicate data")

	ErrUsernamePasswordNotFound = errors.New("(Username) or (Password) empty")
)