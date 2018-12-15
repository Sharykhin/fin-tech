package request

type (
	CreateCompanyRequest struct {
		Symbol      string   `json:"symbol"`
		Name        string   `json:"name"`
		Exchange    string   `json:"exchange"`
		Website     string   `json:"website"`
		Description string   `json:"description"`
		Tags        []string `json:"tags"`
	}
)
