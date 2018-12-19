package request

import (
	"errors"
	"strings"

	"github.com/Sharykhin/fin-tech/entity"
)

type (
	CreateCompanyRequest struct {
		Symbol      string      `json:"symbol"`
		Name        string      `json:"name"`
		Exchange    string      `json:"exchange"`
		Website     string      `json:"website"`
		Description string      `json:"description"`
		Tags        entity.Tags `json:"tags"`
	}
)

func (ccr CreateCompanyRequest) Validate() ErrorBox {
	errs := ErrorBox{}
	if err := validateSymbol(ccr.Symbol); err != nil {
		errs["symbol"] = err
	}

	if err := validateName(ccr.Name); err != nil {
		errs["name"] = err
	}

	if err := validateExchange(ccr.Exchange); err != nil {
		errs["exchange"] = err
	}

	if !_filledString(ccr.Website) {
		errs["website"] = errors.New("website is required")
	}

	if !_filledString(ccr.Description) {
		errs["description"] = errors.New("description is required")
	}

	return errs
}

func validateName(name string) error {
	trimmedName := strings.Trim(name, " ")
	if trimmedName == "" {
		return errors.New("name is required")
	}

	return nil
}

func validateSymbol(symbol string) error {
	trimmedSymbol := strings.Trim(symbol, " ")
	if trimmedSymbol == "" {
		return errors.New("symbol is required")
	}

	if len(trimmedSymbol) > 10 {
		return errors.New("symbol can not be longer than 10 characters")
	}
	return nil
}

func validateExchange(exchange string) error {
	trimmedExchange := strings.Trim(exchange, " ")
	if trimmedExchange == "" {
		return errors.New("exchange is required")
	}

	return nil
}
