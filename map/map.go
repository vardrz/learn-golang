package main

import "fmt"

func main() {
	abc := map[string]any{
		"firstname": "Farid",
		"lastname":  "FR",
		"number":  123,
	}

	delete(abc, "number")

	fmt.Println(abc)
}