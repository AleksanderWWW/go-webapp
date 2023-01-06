package main

import (
	"fmt"
	"net/http"
)

const IndexGreeting string = "Hello from Go"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, IndexGreeting)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
