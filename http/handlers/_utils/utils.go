package _utils

import (
	"net/http"
	"strconv"
)

func GetQueryInt64(r *http.Request, name string, def int64) (int64, error) {
	val := r.FormValue(name)
	if val == "" {
		return def, nil
	}

	return strconv.ParseInt(val, 10, 64)
}
