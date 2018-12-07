package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Sharykhin/fin-tech/http"
)

func main() {
	addr := os.Getenv("WEB_API_ADDRESS")
	fmt.Println("Start listening on: ", addr)
	log.Fatal(http.ListenAndServe(addr))
}
