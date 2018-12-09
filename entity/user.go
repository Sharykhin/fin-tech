package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

//TODO @QUESTIONS:
// 1. Don't you think that general custom types can be moved to a separate file or even package?

const (
	ISO8601 = "2006-01-02T15:04:05Z"
)

type (
	// User struct represents general User model
	User struct {
		ID        int64    `json:"id"`
		Email     string   `json:"email"`
		FirstName string   `json:"first_name"`
		LastName  string   `json:"last_name"`
		CreatedAt Time     `json:"created_at"`
		DeletedAt NullTime `json:"deleted_at"`
	}
	// NullTime provides an ability to pass nil and return null valies
	NullTime struct {
		Time  Time
		Valid bool // Valid is true if Time is not NULL
	}

	// Time is a general time that is used across that project.
	// The main idea under custom time type is to use ISO 8601 as a returned value
	Time time.Time
)

func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(ISO8601))
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
