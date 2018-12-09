package request

import (
	"errors"
	"strings"
)

type (
	UserCreateRequest struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
)

func (u UserCreateRequest) Validate() []error {
	var errs []error
	if strings.Trim(u.Email, " ") == "" {
		errs = append(errs, errors.New("email field is required"))
	}

	if strings.Trim(u.FirstName, " ") == "" {
		errs = append(errs, errors.New("first_name is required"))
	}

	if strings.Trim(u.LastName, " ") == "" {
		errs = append(errs, errors.New("last_name is required"))
	}

	return errs
}
