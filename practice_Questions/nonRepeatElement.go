package main

import (
	"fmt"
)

func nonRepeatElement(array []int) {
	var count int
	var nre int
	for i := 0; i < len(array)-1; i++ {
		count = 0
		for j := i + 1; j < len(array); j++ {
			if array[i] == array[j] {
				count++
			}
		}
		if count == 0 {
			nre = array[i]
		}
	}
	fmt.Println("NRE : ", nre)
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
	nonRepeatElement(array)

}
