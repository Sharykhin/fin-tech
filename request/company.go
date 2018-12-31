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
		Symbol      NullString `json:"symbol"`
		Name        NullString `json:"name"`
		Exchange    NullString `json:"exchange"`
		Website     NullString `json:"website"`
		Description NullString `json:"description"`
		Tags        NullTags   `json:"tags"`
	}

	NullTags struct {
		Valid bool
		Value entity.Tags
	}
)

func (nt *NullTags) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nt.Valid = false
		return nil
	}

	nt.Valid = true
	var tags entity.Tags
	if err := json.Unmarshal(data, &tags); err != nil {
		return err
	}

	nt.Valid = true
	nt.Value = tags
	return nil
}

func (ucr UpdateCompanyRequest) Validate() (bool, ErrorBox) {
	box := ErrorBox{}

	if ucr.Symbol.Valid {
		if err := validateCompanySymbol(ucr.Symbol.Value); err != nil {
			box["symbol"] = err
		}
	}

	if ucr.Name.Valid {
		if err := validateCompanyName(ucr.Name.Value); err != nil {
			box["name"] = err
		}
	}

	if ucr.Exchange.Valid {
		if err := validateCompanyExchange(ucr.Exchange.Value); err != nil {
			box["exchange"] = err
		}
	}

	if ucr.Website.Valid {
		if err := validateCompanyWebsite(ucr.Website.Value); err != nil {
			box["website"] = err
		}
	}

	if ucr.Description.Valid {
		if err := validateCompanyDescription(ucr.Description.Value); err != nil {
			box["description"] = err
		}
	}

	if len(box) > 0 {
		return false, box
	}

	return true, box
}

func (ccr CreateCompanyRequest) Validate() (bool, ErrorBox) {
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
		errs["description"] = err
	}

	if len(errs) > 0 {
		return false, errs
	}
	return true, errs
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
