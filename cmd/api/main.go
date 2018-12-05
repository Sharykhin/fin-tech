package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/Sharykhin/fin-tech/database/mysql"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("OK"))
		if err != nil {
			panic(fmt.Errorf("could not write into the http output: %v", err))
		}
	})

	addr := os.Getenv("WEB_API_ADDRESS")
	fmt.Println("Start listening on: ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))

}
