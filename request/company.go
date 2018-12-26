package request

import (
	"encoding/json"
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

	UpdateCompanyRequest struct {
		Description NullString  `json:"description"`
		Tags        entity.Tags `json:"tags"`
	}

	NullString struct {
		Valid bool
		Value string
	}
)

func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		// The key was set to null
		ns.Valid = false
		return nil
	}

	var tmp string
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	ns.Valid = true
	ns.Value = tmp

	return nil
}

func (ccr CreateCompanyRequest) Validate() ErrorBox {
	errs := ErrorBox{}
	if err := validateCompanySymbol(ccr.Symbol); err != nil {
		errs["symbol"] = err
	}

	if err := validateCompanyName(ccr.Name); err != nil {
		errs["name"] = err
	}

	if err := validateCompanyExchange(ccr.Exchange); err != nil {
		errs["exchange"] = err
	}

	if err := validateCompanyWebsite(ccr.Website); err != nil {
		errs["website"] = err
	}

	if err := validateCompanyDescription(ccr.Description); err != nil {
		errs["description"] = errors.New("description is required")
	}

	return errs
}

func validateCompanySymbol(symbol string) error {
	trimmedSymbol := strings.Trim(symbol, " ")
	if trimmedSymbol == "" {
		return errors.New("symbol is required")
	}

	if len(trimmedSymbol) > 10 {
		return errors.New("symbol can not be longer than 10 characters")
	}
	return nil
}

func validateCompanyName(name string) error {
	trimmedName := strings.Trim(name, " ")
	if trimmedName == "" {
		return errors.New("name is required")
	}

	return nil
}

func validateCompanyExchange(exchange string) error {
	trimmedExchange := strings.Trim(exchange, " ")
	if trimmedExchange == "" {
		return errors.New("exchange is required")
	}

	return nil
}

func validateCompanyWebsite(website string) error {
	if !_filledString(website) {
		return errors.New("website is required")
	}
	return nil
}

func validateCompanyDescription(description string) error {
	if !_filledString(description) {
		return errors.New("description is required")
	}

	return nil
}
