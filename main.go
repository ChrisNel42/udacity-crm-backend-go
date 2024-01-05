package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
}

// This map represents the Database
var CustomerMap map[int]Customer

func (c Customer) addCustomerToDB () (error) {

	_, exists := CustomerMap[c.Id]
	if exists {
		return errors.New("Customer already exists")
	}
	CustomerMap[c.Id] = c
	return nil
}

func getCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(CustomerMap); err == nil {
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

	http.HandleFunc("/customers", getCustomers)

	fmt.Println("Starting Server on port :3000 ...")
	http.ListenAndServe(":3000", nil)
}