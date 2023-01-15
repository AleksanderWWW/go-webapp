package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/AleksanderWWW/go-webapp/utils"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, utils.IndexGreeting)
}

func handleRequests() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
