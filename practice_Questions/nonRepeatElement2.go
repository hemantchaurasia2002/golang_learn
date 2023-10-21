// find non repeated element in unsorted array

package main

import (
	"fmt"
)

func nonRepeatedElement(array []int) {
	var xor int
	xor = 0000
	for i := 0; i < len(array); i++ {
		xor = xor ^ array[i]
		// if xor == array[i] {
		// 	fmt.Println("Non Repeated Element:", xor)
		// }
	}
	fmt.Println("Non Repeated Element:", xor)
	return
}

func main() {
	var size int
	fmt.Printf("Enter the size of the array : ")
	fmt.Scanf("%d", &size)
	fmt.Printf("Enter array's elements : ")
	var array = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanf("%d", &array[i])
	}
	nonRepeatedElement(array)
}
