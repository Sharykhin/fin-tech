package errs

import "errors"

var (
	InvalidJSON         = errors.New("json is invalid")
	EmailIsRequired     = errors.New("email is required")
	EmailIsInvalid      = errors.New("email is invalid")
	FirstNameIsRequired = errors.New("first_name is required")
	FirstNameIsTooLong  = errors.New("first_name can not be longer than %d characters")
	LastNameIsRequired  = errors.New("last_name is required")
	LastNameIsTooLong   = errors.New("last_name can not be longer than %d characters")
)
