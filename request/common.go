package request

import (
	"encoding/json"
	"strings"
)

type (
	// ErrorBox is a customer error container that would be used for validating
	ErrorBox map[string]error
)

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
