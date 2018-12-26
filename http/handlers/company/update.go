package company

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Sharykhin/fin-tech/controller/company"

	"github.com/Sharykhin/fin-tech/http/errs"
	"github.com/Sharykhin/fin-tech/request"
	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var ucr request.UpdateCompanyRequest
	vars := mux.Vars(r)

	ID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		logger.LOG(logger.ERROR, "could not parse route param id: %v", err)
		response.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&ucr); err != nil {
		logger.Error("could not parse request body into CreateCompanyRequest struct: %v", err)
		if err == io.EOF {
			response.SendError(w, errs.InvalidJSON.Error(), http.StatusBadRequest)
			return
		}
		response.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if valid, errsBox := ucr.Validate(); !valid {
		response.SendError(w, errsBox, http.StatusBadRequest)
		return
	}

	ctrl := company.NewCompanyController()
	err = ctrl.Update(r.Context(), ID, ucr)
	if err != nil {
		fmt.Println("e", err)
	}
	fmt.Println("oops")
	fmt.Println(ucr.Tags)

	w.Write([]byte("OK"))
}
