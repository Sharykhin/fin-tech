package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sharykhin/fin-tech/database"

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
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	s, err := u.Serialize("index")
	fmt.Println("serialized user", s)
	if err != nil {
		// TODO: if we could create user but not serialize we should not terminate app
		log.Fatal(err)
	}

	response.SendSuccess(w, u, nil, nil, http.StatusCreated)
}
