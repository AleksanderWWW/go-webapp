package main

import (
	"fmt"
	"log"
	"net/http"
)

const IndexGreeting string = "Hello from Go"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, IndexGreeting)
}

func handleRequests() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
