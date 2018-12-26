package request

import (
	"encoding/json"
	"strings"
)

type (
	// ErrorBox is a customer error container that would be used for validating
	ErrorBox map[string]error
	// NullString allows to accept string or null values
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

// MarshalJSON implements Marshaller interface to parse error and return string
func (eb ErrorBox) MarshalJSON() ([]byte, error) {
	box := make(map[string]string)

	for key, err := range eb {
		box[key] = err.Error()
	}

	return json.Marshal(&box)
}

func _filledString(value string) bool {
	trimmedValue := strings.Trim(value, " ")
	if trimmedValue == "" {
		return false
	}

	return true
}
