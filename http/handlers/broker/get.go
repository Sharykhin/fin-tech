package broker

import (
	"net/http"
	"strconv"

	"github.com/Sharykhin/fin-tech/http/errs"

	"github.com/Sharykhin/fin-tech/controller/broker"
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

	ctrl := broker.NewBrokerController()
	b, err := ctrl.Get(r.Context(), ID)
	if err != nil {
		if err == errs.ResourceNotFound {
			response.SendError(w, err.Error(), http.StatusNotFound)
			return
		}
		logger.Logger.Error("failed to get broker: %v", err)
		response.SendError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	response.SendSuccess(w, b, nil, nil, http.StatusOK)
}
