package company

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Sharykhin/fin-tech/http/errs"

	"github.com/Sharykhin/fin-tech/service/logger"

	"github.com/Sharykhin/fin-tech/request"
	"github.com/Sharykhin/fin-tech/service/response"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var ucr request.UpdateCompanyRequest

	if err := json.NewDecoder(r.Body).Decode(&ucr); err != nil {
		logger.Error("could not parse request body into CreateCompanyRequest struct: %v", err)
		if err == io.EOF {
			response.SendError(w, errs.InvalidJSON.Error(), http.StatusBadRequest)
			return
		}
		response.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if ok, err := ucr.Validate(); !ok {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	fmt.Println("oops")
	fmt.Println(ucr.Description.Valid, ucr.Description.Value)

	w.Write([]byte("OK"))
}
