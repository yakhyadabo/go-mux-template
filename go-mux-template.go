package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yakhyadabo/go-mux-template/controller"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/employees", controller.GetEmployees).Methods("GET")
	router.HandleFunc("/api/employee", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/employee", controller.GetEmployee).Queries("id", "{id}").Methods("GET")

	err := http.ListenAndServe(":"+"8000", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		log.Print(err)
	}
}
