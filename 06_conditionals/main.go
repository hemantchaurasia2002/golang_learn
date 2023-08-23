package main

import "fmt"

func main() {
	// x := 15
	// y := 10

	//if x < y {
	//	fmt.Printf("%d is less than %d\n", x,y)}

	//else if
	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue" {
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue and red")
	}

	//switch
	switch color {
	case "red":
		fmt.Println("color s red")
	case "blue":
		fmt.Println("color is blue")
	}
}
