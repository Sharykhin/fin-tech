package entity

type (
	Broker struct {
		User      UserIdentity `json:"user"`
		Position  NullInt      `json:"position"`
		Role      NullString   `json:"role"`
		CreatedAt Time         `json:"created_at"`
		DeletedAt NullTime     `json:"deleted_at"`
	}
)
