package main

import "fmt"

type Filter func(string) string

func main() {
	getFilter("a", filter)
}

func filter(name string) string{
	if name == "a"{
		return "...."
	}else{
		return name
	}
}

func getFilter(name string, filter Filter) {
	abc := filter(name)
	fmt.Println(abc)
}