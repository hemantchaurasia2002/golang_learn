package main

import (
	"fmt"
)

func secondMaxElement(array []int) {
	var max int
	var secondMax int
	max = array[0]
	secondMax = 0
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			secondMax = array[i]
			if secondMax < max || secondMax > array[i-1] {
				secondMax = array[i-1]
			}
		}
	}
	fmt.Println("Second max element is ", secondMax)
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
	secondMaxElement(array)
}
