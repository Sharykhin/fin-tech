package company

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sharykhin/fin-tech/request"
	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var ccr request.UpdateCompanyRequest

	if err := json.NewDecoder(r.Body).Decode(&ccr); err != nil {
		logger.Error("could not parse request body into CreateCompanyRequest struct: %v", err)
		response.SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("oops")
	fmt.Println(ccr)

	w.Write([]byte("OK"))
}
