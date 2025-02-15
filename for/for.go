package main

import "fmt"

func main() {
	abc := map[any]any {
		"abc": "abc",
		"bcd": "bcd",
		12: 12,
	}

	// fmt.Println(abc["abc"])

	for in, val := range abc {
		fmt.Println(in)
		fmt.Println(val)
	}
}