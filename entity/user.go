package entity

import "encoding/json"

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

	UserIdentity struct {
		ID    int64  `json:"id"`
		Email string `json:"email"`
	}
)

// TODO: just imagine that we have 50 properties and they can be changed from time to time
func (u User) Serialize(group string) ([]byte, error) {
	if group == "public" {
		return json.Marshal(&struct {
			ID        int64  `json:"id"`
			Email     string `json:"email"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			CreatedAt Time   `json:"created_at"`
		}{
			ID:        u.ID,
			Email:     u.Email,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			CreatedAt: u.CreatedAt,
		})
	}

	return json.Marshal(u)
}
