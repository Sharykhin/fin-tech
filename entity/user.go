package entity

import "time"

type (
	User struct {
		ID        int64
		Email     string
		FirstName string
		LastName  string
		CreatedAt time.Time
	}
)
