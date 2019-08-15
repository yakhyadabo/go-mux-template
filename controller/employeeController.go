package controller

import (
	"net/http"
	"github.com/yakhyadabo/go-mux-template/model"
	"encoding/json"
	"strconv"
	"math/rand"
	"log"
)

var employees = [] model.Employee{
	model.Employee{ID:"100", Firstname:"FN_100",Lastname:"LN_100", Company: &model.Company{ID:"100",Name:"N_100"}},
	model.Employee{ID:"200", Firstname:"FN_200",Lastname:"LN_200", Company: &model.Company{ID:"200",Name:"N_200"}},
	model.Employee{ID:"300", Firstname:"FN_300",Lastname:"LN_300", Company: &model.Company{ID:"300",Name:"N_300"}},
}

var GetEmployees = func(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(employees)
}

var GetEmployee = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	v := r.URL.Query()
	id := v.Get("id")

	log.Println("ID : ", id)
	for _, employee := range employees{
		if employee.ID == id {
			json.NewEncoder(w).Encode(employee)
			return
		}
	}

	json.NewEncoder(w).Encode(&model.Employee{})
}

var CreateEmployee = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	var employee model.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		// u.Respond(w, u.Message(false, "Error while decoding request body"))
		return
	}

	employee.ID = strconv.Itoa(rand.Intn(100000))
	employees = append(employees,employee)
	json.NewEncoder(w).Encode(employee)
}