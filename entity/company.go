package entity

type (
	Company struct {
		ID          int64  `json:"id"`
		Symbol      string `json:"symbol"`
		Name        string `json:"name"`
		Exchange    string `json:"exchange"`
		Website     string `json:"website"`
		Description string `json:"description"`
		Tags        []Tag  `json:"tags"`
	}

	Tag string
)
