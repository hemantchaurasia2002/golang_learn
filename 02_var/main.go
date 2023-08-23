package main

import "fmt"

func main() {
	//shorthant method
	//name := "Hemant"
	//size := 1.3
	// name, email := "Hemant", "chaurasiahemant2002@gmail.com"
	var name string = "Hemant Chaurasia"
	var age int = 21
	var isCool = true
	fmt.Println(name, age, isCool)
	fmt.Printf("%T\n%T\n%T\n", name,age,isCool)
}
