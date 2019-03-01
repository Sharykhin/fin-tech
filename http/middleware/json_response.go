package middleware

import (
	"fmt"
	"net/http"
)

func EncodeJSON(h func(http.ResponseWriter, *http.Request) (int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, err := h(w, r)

		fmt.Println("hahah", res, err)
	}
}

func Foo(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("fooo")
		h.ServeHTTP(w, r)
	})
}

