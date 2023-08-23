package main

import (
	"fmt"
	"strconv"
)

// Define person struct 
type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// greeting message (value reciver)
func (p Person) greet() string {
	return "Hello, my name is "  p.firstName + " " p.lastName + "and I am " + strconv.Itoa(p.age)

}

func main() {
	//Init person using struct
	person1 := Person{firstName: "Hemant", lastName: "Chaurasia", city: "Lucknow", gender: "Male", age: 25}
	//person2 := Person{"Satyam","Chaurasia","Lucknow","Male",21}

	fmt.Println(person1.greet())
	
}

