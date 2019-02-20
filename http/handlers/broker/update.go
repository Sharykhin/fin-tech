package broker

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/Sharykhin/fin-tech/controller/broker"
	"github.com/Sharykhin/fin-tech/http/errs"
	"github.com/Sharykhin/fin-tech/request"

	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
	"github.com/gorilla/mux"
)

// Update updates a broker
func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		logger.LOG(logger.ERROR, "could not parse route param id: %v", err)
		response.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var ubr request.UpdateBrokerRequest

	if err := json.NewDecoder(r.Body).Decode(&ubr); err != nil {
		logger.Logger.Error("could not parse request body into CreateCompanyRequest struct: %v", err)
		if err == io.EOF {
			response.SendError(w, errs.InvalidJSON.Error(), http.StatusBadRequest)
			return
		}
		response.SendError(w, err.Error(), http.StatusBadRequest)
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

	err = ctrl.Update(r.Context(), ID, ubr)
	if err != nil {
		if err == errs.NothingToUpdate {
			response.SendError(w, err.Error(), http.StatusBadRequest)
			return
		}
		logger.Logger.Error("controller could not update a broker: %v", err)
		response.SendError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if ubr.Position.Valid {
		b.Position = ubr.Position
	}

	if ubr.Role.Valid {
		b.Role = ubr.Role
	}

	response.SendSuccess(w, b, nil, nil, http.StatusOK)

}
