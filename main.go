package main

import (
	"fmt"
)

func main() {
	customerSQLRepo := NewCustomerSQLRepository("fake-sql-session")
	customerProfile, errGetCustomerProfile := customerSQLRepo.GetCustomerProfileByID("1")
	if errGetCustomerProfile != nil {
		fmt.Print("can't get customer profile")
	}
	fmt.Print(customerProfile)

	customerMongoRepo := NewCustomerMongoRepository("fake-mgo-session")
	customerProfile, errGetCustomerProfile = customerMongoRepo.GetCustomerProfileByID("1")
	if errGetCustomerProfile != nil {
		fmt.Print("can't get customer profile")
	}

	fmt.Print(customerProfile)
}
