package entity

import (
	"database/sql/driver"
	"encoding/json"
)

type (
	// Company represents general company info
	Company struct {
		ID          int64  `json:"id"`
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Exchange    string `json:"exchange"`
		Website     string `json:"website"`
		Description string `json:"description"`
		Tags        Tags   `json:"tags"`
	}
	// Tags is a slice of tags. It used as custom param to be able to override Scan and Value properties
	Tags []Tag
	// Tag is special type to identify tag value
	Tag string
)

func (t *Tags) Scan(value interface{}) (err error) {
	return json.Unmarshal(value.([]byte), t)
}

// Value implements the driver Valuer interface.
func (t Tags) Value() (driver.Value, error) {
	//log.Fatal("value", t)
	return json.Marshal(t)
}
