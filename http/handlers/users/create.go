package users

import (
	"encoding/json"
	"net/http"

	"github.com/Sharykhin/fin-tech/service/logger"

	"github.com/Sharykhin/fin-tech/database"

	"github.com/Sharykhin/fin-tech/controller/user"
	"github.com/Sharykhin/fin-tech/http/errs"
	"github.com/Sharykhin/fin-tech/request"
	"github.com/Sharykhin/fin-tech/service/response"
)

// TODO: we should use logging when user could not be created

// Create handler creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var ur request.UserCreateRequest
	err := decoder.Decode(&ur)
	if err != nil {
		logger.Error("could not validate input data: %v", err)
		response.SendError(w, errs.InvalidJSON.Error(), http.StatusBadRequest)
		return
	}

	if err := ur.Validate(); err != nil {
		logger.Error("some data was invalid: %v", err)
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	if u, err := database.UserStorage.FindByEmail(r.Context(), ur.Email); err != nil || u != nil {
		if u != nil {
			response.SendError(w, map[string]string{"email": errs.EmailAlreadyExists.Error()}, http.StatusBadRequest)
			return
		} else if err != nil && err != errs.UserWasNotFound {
			response.SendError(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	uc := user.NewController()
	u, err := uc.Create(r.Context(), ur)
	if err != nil {
		logger.Error("could not create a new user: %v", err)
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	s, err := u.Serialize("public")

	if err != nil {
		logger.Emergency("could not serialize just created user, something really weird went wrong: %v", err)
	}

	//response.Send(w, s, http.StatusCreated)
	response.SendSuccess(w, s, nil, nil, http.StatusCreated)
}
