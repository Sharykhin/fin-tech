package http

import (
	"net/http"
	"time"

	"github.com/Sharykhin/fin-tech/http/handlers/company"

	"github.com/Sharykhin/fin-tech/http/handlers/users"

	"github.com/Sharykhin/fin-tech/http/handlers"
	"github.com/gorilla/mux"
)

func router() http.Handler {
	r := mux.NewRouter()

	s := r.StrictSlash(true).PathPrefix("/api").Subrouter()
	s.HandleFunc("/users", users.Index).Methods("GET")
	s.HandleFunc("/users", users.Create).Methods("POST")
	s.HandleFunc("/_healthcheck", handlers.HealthCheck).Methods("GET")

	s.HandleFunc("/companies", company.Index).Methods("GET")
	s.HandleFunc("/companies", company.Create).Methods("POST")
	s.HandleFunc("/companies/{id:[0-9]+}", company.Update).Methods("PUT")

	return s
}

// ListenAndServe starts listening http service on a provided addr
func ListenAndServe(addr string) error {
	s := &http.Server{
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  time.Second * 60,
		Handler:      router(),
		Addr:         addr,
	}

	return s.ListenAndServe()
}
