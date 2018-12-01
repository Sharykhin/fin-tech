package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		panic(fmt.Errorf("could not write into the http output: %v", err))
	})
	log.Fatal(http.ListenAndServe(os.Getenv("WEB_API_ADDRESS"), nil))
}
