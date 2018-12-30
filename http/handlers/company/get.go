package company

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Sharykhin/fin-tech/service/serialize"
	companySerializer "github.com/Sharykhin/fin-tech/service/serialize/company"

	"github.com/Sharykhin/fin-tech/controller/company"
	"github.com/Sharykhin/fin-tech/http/errs"

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

	ctrl := company.NewCompanyController()

	c, err := ctrl.Find(r.Context(), ID)
	if err != nil {
		if err == errs.ResourceNotFound {
			response.SendError(w, "company could not be found", http.StatusNotFound)
			return
		}
		logger.Logger.Error("could not get a company by its ID: %v", err)
		response.SendError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s, err := companySerializer.Item(*c, serialize.PUBLIC)
	if err != nil {
		log.Panicf("could not serialize comany: %v", err)
	}

	response.SendSuccess(w, s, nil, nil, http.StatusOK)
}
