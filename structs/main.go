package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact contactInfo
}

// we can write like 
// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo
// }
//  same like :
// type person struct {
// 	firstName string
// 	lastName  string
// 	contactInfo contactInfo
// }


func main() {
	// this not the right way to implement to struct
	// alex := person{"Alex","Anderson"}

	// this the corret way
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println((alex))

	// this is another way

	// var alex person 
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// //%+v sill print out all the different field names and their values from Alex.
	// fmt.Printf("%+v",alex)

	//embedded struct

	jim:= person{

		firstName:  "Jim",
		lastName: "Party",
		contact: contactInfo{
			email: "jim@gmail.com",
			zipCode: 94000,
		},
		// must have coma after value declaration.
	}

	jim.updateName(("jimmy"))
	jim.print()

	// fmt.Printf("%+v",jim)
}

// with this below function, jim will not update , so we use pointer
func (p person) updateName(newFirstName string)  {
	p.firstName = newFirstName
}

func (p person) print()  {
	fmt.Printf("%+v",p)
}