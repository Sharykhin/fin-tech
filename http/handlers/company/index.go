package company

import (
	"net/http"

	"github.com/Sharykhin/fin-tech/controller/company"
	utils "github.com/Sharykhin/fin-tech/http/handlers/_utils"
	"github.com/Sharykhin/fin-tech/service/logger"
	"github.com/Sharykhin/fin-tech/service/response"
)

func Index(w http.ResponseWriter, r *http.Request) {
	ctrl := company.NewCompanyController()
	limit, err := utils.GetQueryInt64(r, "limit", 10)
	if err != nil {
		response.SendError(w, "limit has invalid format", http.StatusBadRequest)
		return
	}

	offset, err := utils.GetQueryInt64(r, "offset", 0)
	if err != nil {
		response.SendError(w, "offset has invalid format", http.StatusBadRequest)
		return
	}
	list, total, err := ctrl.Index(r.Context(), limit, offset)
	if err != nil {
		logger.Error("could not retreive companies list from controller: %v", err)
		response.SendError(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	response.SendSuccess(w, list, nil, map[string]int64{
		"limit":  limit,
		"offset": offset,
		"total":  total,
		"count":  10,
	}, http.StatusOK)
}
