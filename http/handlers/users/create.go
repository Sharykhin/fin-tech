package users

import (
	"encoding/json"
	"net/http"

	"github.com/Sharykhin/fin-tech/controller/user"
	"github.com/Sharykhin/fin-tech/http/errs"
	"github.com/Sharykhin/fin-tech/request"
	"github.com/Sharykhin/fin-tech/service/response"
)

// TODO: we should use logging when user could not be created
// TODO: think about serialization group

// Create handler creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ur request.UserCreateRequest
	err := decoder.Decode(&ur)
	if err != nil {
		response.SendError(w, errs.InvalidJSON, http.StatusBadRequest)
		return
	}

	if err := ur.Validate(); err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	uc := user.NewUserController()
	u, err := uc.Create(r.Context(), ur)
	if err != nil {
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	response.SendSuccess(w, u, nil, nil, http.StatusCreated)
}
