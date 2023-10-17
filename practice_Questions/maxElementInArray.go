package main

import (
	"fmt"
)

func maxElement(array []int) {
	var max int
	max = array[0]
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	fmt.Println("Max element is ", max)
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
	maxElement(array)
}
