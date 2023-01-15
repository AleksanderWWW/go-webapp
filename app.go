package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/AleksanderWWW/go-webapp/utils"
	"github.com/AleksanderWWW/go-webapp/backend"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, utils.IndexGreeting)
}

func retrieveAll(w http.ResponseWriter, r *http.Request) {
	products := backend.RetrieveAll()
	json.NewEncoder(w).Encode(products)
}

func handleRequests() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/all", retrieveAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
