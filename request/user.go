package request

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/Sharykhin/fin-tech/http/errs"
)

type (
	// UserCreateRequest represents a struct of income request body
	// for creating a new user
	UserCreateRequest struct {
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		isValid   bool
	}
)

var reEmail = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// Validate method validates itself and returns a map of errors
func (u *UserCreateRequest) Validate() *ErrorBox {
	errBox := ErrorBox{}
	u.isValid = true

	if err := validateEmail(u.Email); err != nil {
		errBox["email"] = err
		u.isValid = false
	}

	if err := validateFirstName(u.FirstName); err != nil {
		errBox["first_name"] = err
		u.isValid = false
	}

	if err := validateLastName(u.LastName); err != nil {
		errBox["last_name"] = err
		u.isValid = false
	}

	if !u.isValid {
		return &errBox
	}

	return nil
}

func validateEmail(email string) error {
	trimmedEmail := strings.Trim(email, " ")
	if trimmedEmail == "" {
		return errs.EmailIsRequired
	}

	if !reEmail.MatchString(trimmedEmail) {
		return errs.EmailIsInvalid
	}

	return nil
}

func validateFirstName(firstName string) error {
	trimmedFirstName := strings.Trim(firstName, " ")
	if trimmedFirstName == "" {
		return errs.FirstNameIsRequired
	}

	if len(trimmedFirstName) > 20 {
		return errors.New(fmt.Sprintf(errs.FirstNameIsTooLong.Error(), 20))
	}

	return nil
}

func validateLastName(lastName string) error {
	trimmedLastName := strings.Trim(lastName, " ")
	if trimmedLastName == "" {
		return errs.LastNameIsRequired
	}

	if len(trimmedLastName) > 20 {
		return errors.New(fmt.Sprintf(errs.LastNameIsTooLong.Error(), 20))
	}

	return nil
}
