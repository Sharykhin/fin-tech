package company

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Sharykhin/fin-tech/entity"
	"github.com/Sharykhin/fin-tech/service/serialize"
	companySerializer "github.com/Sharykhin/fin-tech/service/serialize/company"

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

	err = ctrl.Update(r.Context(), ID, ucr)
	if err != nil {
		if err == errs.NothingToUpdate {
			response.SendError(w, err.Error(), http.StatusBadRequest)
			return
		}
		logger.Logger.Error("controller could not update a company: %v", err)
		response.SendError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c = updateCompanyEntity(c, ucr)
	s, err := companySerializer.Item(*c, serialize.PUBLIC)
	if err != nil {
		log.Panicf("could not serialize comany: %v", err)
	}
	response.SendSuccess(w, s, nil, nil, http.StatusOK)
}

func updateCompanyEntity(c *entity.Company, ucr request.UpdateCompanyRequest) *entity.Company {
	if ucr.Symbol.Valid {
		c.Symbol = ucr.Symbol.Value
	}

	if ucr.Name.Valid {
		c.Name = ucr.Name.Value
	}

	if ucr.Exchange.Valid {
		c.Exchange = ucr.Exchange.Value
	}

	if ucr.Website.Valid {
		c.Website = ucr.Website.Value
	}

	if ucr.Description.Valid {
		c.Description = ucr.Description.Value
	}

	if ucr.Tags.Valid {
		c.Tags = ucr.Tags.Value
	}

	return c
}
