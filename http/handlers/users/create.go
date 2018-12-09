package users

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Sharykhin/fin-tech/controller/user"
	"github.com/Sharykhin/fin-tech/request"
)

// Create handler creates a new user
func Create(w http.ResponseWriter, r *http.Request) {
	uc := user.NewUserController()
	decoder := json.NewDecoder(r.Body)
	var ur request.UserCreateRequest
	err := decoder.Decode(&ur)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), 400)
		return
	}
	if errs := ur.Validate(); errs != nil {
		//http.Error(w)
	}
	u, err := uc.Create(r.Context(), ur)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&u)
}
