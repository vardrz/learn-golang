package main

import "fmt"

type Customer struct {
	Name, Address string
	Phone         int
}

func main() {
	customer1 := Customer{
		Name:    "Dono",
		Address: "JKT",
		Phone:   6281502100,
	}

	customer2 := &customer1
	customer2.Address = "BDG"

	customer3 := new(Customer)
	customer3.Name = "Kasino"
	customer3.Address = "SBY"
	customer3.Phone = 628155221

	fmt.Println(customer1)
	fmt.Println(customer2)
	fmt.Println(customer3)
}