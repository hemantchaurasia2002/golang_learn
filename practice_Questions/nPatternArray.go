package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Printf("Enter the value of n : ")
	fmt.Scanf("%d", &n)
	array := make([][]int, n)
	var value int = 1
	for i := 0; i < n; i++ {
		array[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if j >= n-i-1 {
				array[i][j] = value
				value++
			}
		}
		value = 1
	}
	fmt.Println("2D array : ")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(array[i][j], " ")
		}
		fmt.Println()
	}
}
