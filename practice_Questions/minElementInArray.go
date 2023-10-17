package main

import (
	"fmt"
)

func minElement(array []int) {
	var min int
	min = array[0]
	for i := 0; i < len(array); i++ {
		if array[i] < min {
			min = array[i]
		}
	}
	fmt.Println("Min element is ", min)
}

func main() {
	fmt.Printf("Enter size of your array: ")
	var size int
	fmt.Scanln(&size)
	fmt.Printf("Enter array's elements : ")
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &array[i])
	}
	fmt.Println("Your array is: ", array)
	minElement(array)
}
