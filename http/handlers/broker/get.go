package broker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Sharykhin/fin-tech/entity"

	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		logger.LOG(logger.ERROR, "could not parse route param id: %v", err)
		response.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(ID)

	b := entity.Broker{
		User: entity.UserIdentity{
			ID:    1,
			Email: "chapal@inbox.ru",
		},
		Position: entity.NullInt{
			Valid: false,
		},
		CreatedAt: entity.Time(time.Now()),
		DeletedAt: entity.NullTime{
			Valid: false,
		},
	}

	err = json.NewEncoder(w).Encode(&b)
	if err != nil {
		fmt.Println(err)
	}
}
