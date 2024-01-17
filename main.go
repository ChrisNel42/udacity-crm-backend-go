package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	Id 			int		`json:"id"`
	Name 		string  `json:"name"`
	Role 		string	`json:"role"`
	Email 		string  `json:"email"`
	Phone 		string  `json:"phone"`
	Contacted 	bool    `json:"contacted"`
}

type CustomerDB interface {
	addCustomerToDB()
	deleteCustomerFromDB()
}

// This map represents the Database
var CustomerMap map[int]Customer

// Functions for Customer objects

func (c Customer) addCustomerToDB () (error) {

	_, exists := CustomerMap[c.Id]
	if exists {
		return errors.New("Customer with this id already exists")
	}
	CustomerMap[c.Id] = c
	return nil
}

func (c Customer) deleteCustomerFromDB () (error) {
	_, exists := CustomerMap[c.Id]
	if !exists {
		return errors.New("Customer not found in DB")
	}
	delete(CustomerMap, c.Id)
	return nil
}

// Handlers for API

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(CustomerMap); err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	customerId, err := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customer, exists := CustomerMap[customerId]
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(customer); err == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {

	var customer Customer
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&customer); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	
	if err := customer.addCustomerToDB(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(customer)

	w.WriteHeader(http.StatusCreated)
}

func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId, err := strconv.Atoi(vars["id"])

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	customer := CustomerMap[customerId]

	if err := customer.deleteCustomerFromDB(); err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}
	
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(CustomerMap); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main () {
	customer1 := Customer{1, "John Doe", "Customer", "email@email.com", "+49 12452 1234632", false}
	customer2 := Customer{2, "Bob", "Customer", "cats@email.com", "+49 4542 123684932", false}
	customer3 := Customer{3, "Amanda Smith", "Customer", "amanda@email.com", "+49 4542 123684932", false}

	CustomerMap = make(map[int]Customer)
	customer1.addCustomerToDB()
	customer2.addCustomerToDB()
	customer3.addCustomerToDB()

	r := mux.NewRouter()

	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", addCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	fmt.Println("Starting Server on port :3000 ...")
	http.ListenAndServe(":3000", r)
}