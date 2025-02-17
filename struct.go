package main

import "fmt"

type Profile struct {
	FirstName, LasName, Address string
	Age                         int
}

func main() {
	farid := Profile{
		FirstName: "Farid",
		LasName:   "Fatkhurrozak",
		Address:   "abc",
		Age:       11,
	}

	fmt.Println("First Name: ", farid.FirstName)
	fmt.Println("Last Name: ", farid.LasName)
	fmt.Println("Address: ", farid.Address)
	fmt.Println("Age: ", farid.Age)
}