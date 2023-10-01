package main

import (
	"fmt"
)

func main() {
	array := make([]string, 0, 3)
	array = append(array, "Phoenix", "Skye", "Brimstone")

	fmt.Println("Duelist is - ", array[0])
	fmt.Println("Initiator is - ", array[1])
	fmt.Println("Controller is - ", array[2])
}