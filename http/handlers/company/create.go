package company

import (
	"encoding/json"
	"net/http"

	"github.com/Sharykhin/fin-tech/controller/company"
	"github.com/Sharykhin/fin-tech/request"
	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
)

// Create handler creates a new company
func Create(w http.ResponseWriter, r *http.Request) {
	var ccr request.CreateCompanyRequest

	err := json.NewDecoder(r.Body).Decode(&ccr)
	if err != nil {
		logger.Error("could not parse request body into CreateCompanyRequest struct: %v", err)
		response.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctrl := company.NewCompanyController()

	c, err := ctrl.Create(r.Context(), ccr)
	if err != nil {
		logger.Error("failed to create a new company: %v", err)
		response.SendError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccess(w, c, nil, nil, http.StatusCreated)

}
