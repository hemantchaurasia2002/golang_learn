//sum of N natural numbers (N is given by user)
package main

import (
	"fmt"
)
//fastest approach
func sum(n int) int {
	var sum int
	sum = n * (n + 1) / 2
	return sum
}

//first try
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
	var secondAns int
	fmt.Printf("Enter the value of N : ")
	fmt.Scanf("%d", &num)
	ans = sumOfN(num)
	secondAns = sumOfN(num)
	fmt.Printf("Sum of first %d natural numbers : %v\n", num, ans)
	fmt.Printf("Sum of first %d natural numbers : %v", num, secondAns)
}
