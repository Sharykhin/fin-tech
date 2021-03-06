package entity

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
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
		Int   int
	}

	NullString struct {
		Valid  bool
		String string
	}

	// Time is a general time that is used across that project.
	// The main idea under custom time type is to use ISO 8601 as a returned value
	Time time.Time
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
	ns.String = tmp

	return nil
}

func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.Valid, ns.String = false, ""
		return nil
	}

	if v, ok := value.([]uint8); ok {
		ns.Valid, ns.String = true, string(v)
	}

	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(ISO8601))
}

/** NullInt **/

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
	ni.Int = tmpInt

	return nil
}

func (ni NullInt) MarshalJSON() ([]byte, error) {
	if ni.Valid {
		return json.Marshal(ni.Value)
	}

	return json.Marshal(nil)
}

func (ni *NullInt) Scan(value interface{}) error {
	if value == nil {
		ni.Valid, ni.Int = false, 0
		return nil
	}

	if v, ok := value.(int64); ok {
		ni.Valid, ni.Int = true, int(v)
	}

	return nil
}

func (ni NullInt) Value() (driver.Value, error) {
	fmt.Println("HA HA", ni)
	if !ni.Valid {
		return nil, nil
	}
	return ni.Int, nil
}

/** NullTime **/

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
