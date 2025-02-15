package main

import "fmt"

func main() {
	abc := [...]int{1,2,3,4,5,6,7,8}

	slice1 := abc[5:]
	slice2 := append(slice1, 9)

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(abc)
}