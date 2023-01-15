package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AleksanderWWW/go-webapp/backend"
	"github.com/AleksanderWWW/go-webapp/utils"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, utils.IndexGreeting)
}

func retrieveAll(w http.ResponseWriter, r *http.Request) {
	products := backend.RetrieveAll()
	json.NewEncoder(w).Encode(products)
}

func retrieveSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	p := backend.RetrieveByID(id)

	json.NewEncoder(w).Encode(p)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexHandler)
	myRouter.HandleFunc("/all", retrieveAll)
	myRouter.HandleFunc("/product/{id}", retrieveSingle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	handleRequests()
}
