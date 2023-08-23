package main

import "fmt"

func main(){
	//arrays
	//var fruitArr [2]string

	//assign values
	// fruitArr [0] = "apple"
	// fruitArr [1] = "orange"

	//Declare and assign
	//fruitArr := [2]string{"apple", "orange"}

	fruitSlice := []string{"apple", "orange"}
	fmt.Println(fruitSlice)
	fmt.Println("Length of slice - ", len(fruitSlice))
}