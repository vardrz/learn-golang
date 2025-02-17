package main

import "fmt"

type Customer struct {
	Name, Address string
	Phone         int
}

func setCustomerPhone(customer *Customer, phone int) {
	customer.Phone = phone
}

func main() {
	customer1 := (Customer{"Dono", "BDG", 0})

	setCustomerPhone(&customer1, 628150021000)

	fmt.Println(customer1)
}