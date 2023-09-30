package main

import (
	"fmt"
)

func findPair(arr []int, n, sum int) {
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == sum {
				fmt.Printf("Pair found at index %d and %d", i, j)
				return
			}
		}
	}
	fmt.Println("Pair not found")
}

func main() {
	arr := []int{8, 7, 5, 3, 1}
	findPair(arr, len(arr), 10)
}
