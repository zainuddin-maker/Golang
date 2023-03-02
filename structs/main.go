package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	// this not the right way to implement to struct
	// alex := person{"Alex","Anderson"}

	// this the corret way
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println((alex))
}