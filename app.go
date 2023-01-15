package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

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
	url_path := strings.TrimPrefix(r.URL.Path, "/provisions/")
	splited_path := strings.Split(url_path, "/")
	id := splited_path[len(splited_path)-1]
	p := backend.RetrieveByID(id)

	json.NewEncoder(w).Encode(p)
}

func handleRequests() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/all", retrieveAll)
	http.HandleFunc("/product/", retrieveSingle)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
