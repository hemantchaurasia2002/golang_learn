package main

import (
	"fmt"
)

func secondMinElement(array []int) {
	var min int
	var secondMin int
	min = array[0]
	secondMin = 0
	for i := 0; i < len(array); i++ {
		if array[i] < min {
			Min = array[i]
			if secondMin > min {
				secondMin = array[i]
			}
		}
	}
	fmt.Println("Second min element is ", secondMin)
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
	secondMinElement(array)
}
