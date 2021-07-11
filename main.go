package main

import "fmt"


func main() {
	nico := map[string]string{"name":"nico", "age":"12"}
	for key, value := range nico {
		fmt.Println(key , value)
	}
	// name nico
	// age 12
	for _, value := range nico {
		fmt.Println(value)
	}
	// nico
	// 12
	for key, _ := range nico {
		fmt.Println(key)
	}
	// name
	// age
}
