package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const (
	ISO8601 = "2006-01-02T15:04:05Z"
)

type (
	// NullTime provides an ability to pass nil and return null valies
	NullTime struct {
		Time  Time
		Valid bool // Valid is true if Time is not NULL
	}

	NullInt struct {
		Valid bool
		Value int
	}

	// Time is a general time that is used across that project.
	// The main idea under custom time type is to use ISO 8601 as a returned value
	Time time.Time
)

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(ISO8601))
}

func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ni.Valid = false
		return nil
	}

	var tmpInt int
	if err := json.Unmarshal(data, &tmpInt); err != nil {
		return err
	}

	ni.Valid = true
	ni.Value = tmpInt

	return nil
}

func (ni NullInt) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Value)
	}

	return json.Marshal(nil)
}

func (nt *NullTime) Scan(value interface{}) (err error) {
	if value == nil {
		nt.Time, nt.Valid = Time(time.Time{}), false
		return
	}

	return nil

}

// Value implements the driver Valuer interface.
func (nt NullTime) Value() (driver.Value, error) {
	if !nt.Valid {
		return nil, nil
	}
	return nt.Time, nil
}

func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(nt.Time)
}
