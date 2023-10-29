package main

import (
	"fmt"
)

func sumOfN(n int) int {
	var sum int
	sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

func main() {
	var num int
	var ans int
	fmt.Printf("Enter the value of N : ")
	fmt.Scanf("%d", &num)
	ans = sumOfN(num)
	fmt.Printf("Sum of first %d natural numbers : %v", num, ans)
}
