package broker

import (
	"github.com/Sharykhin/fin-tech/http/errs"
	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (h HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// TODO: use wrapper to handle response and return resonse as type
	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		logger.LOG(logger.ERROR, "could not parse route param id: %v", err)
		response.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	b, err := h.controller.Get(r.Context(), ID)
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
