package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/AleksanderWWW/go-webapp/backend"
	"github.com/AleksanderWWW/go-webapp/resources"
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

func insertSingle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var p resources.Product
	json.Unmarshal(reqBody, &p)
	
	coll := backend.ConnectToMongo(utils.DatabaseName, utils.CollectionName)

	defer backend.CloseMongoConnection(coll.Database().Client())

	backend.InsertProduct(coll, p)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", indexHandler)
	myRouter.HandleFunc("/all", retrieveAll)
	myRouter.HandleFunc("/product/{id}", retrieveSingle)
	myRouter.HandleFunc("/product", insertSingle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	
	handleRequests()
}
