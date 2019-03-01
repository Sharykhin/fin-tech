package response

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type (
	response struct {
		Data     interface{} `json:"data"`
		Errors   interface{} `json:"errors"`
		Warnings interface{} `json:"warnings"`
		Meta     interface{} `json:"meta"`
	}
)

//TODO: this about using middleware to send a proper headers
func setHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

// SendSuccess send a general success response
func SendSuccess(w http.ResponseWriter, data interface{}, warnings interface{}, meta interface{}, statusCode int) {
	setHeaders(w)
	w.WriteHeader(statusCode)
	if statusCode == http.StatusNoContent {
		return
	}
	var err error
	if raw, ok := data.([]byte); ok {
		err = json.NewEncoder(w).Encode(&struct {
			Data     json.RawMessage `json:"data"`
			Errors   interface{}     `json:"errors"`
			Warnings interface{}     `json:"warnings"`
			Meta     interface{}     `json:"meta"`
		}{
			Data:     raw,
			Errors:   nil,
			Warnings: warnings,
			Meta:     meta,
		})
	} else {
		err = json.NewEncoder(w).Encode(&response{
			Data:     data,
			Errors:   nil,
			Warnings: warnings,
			Meta:     meta,
		})
	}

	if err != nil {
		log.Panic(fmt.Errorf("could not encode http response to json: %v", err))
	}
}

// SendError method sends an error response back to a client
func SendError(w http.ResponseWriter, errors interface{}, statusCode int) {
	setHeaders(w)
	w.WriteHeader(statusCode)

	res, err := json.Marshal(response{
		Data:     nil,
		Errors:   errors,
		Warnings: nil,
		Meta:     nil,
	})

	if err != nil {
		// TODO: use log.Panic or just panic but use one style
		panic(fmt.Errorf("could not marshal json response: %v", err))
	}

	_, err = w.Write(res)
	if err != nil {
		panic(fmt.Errorf("could not return a http response: %v", err))
	}
}
