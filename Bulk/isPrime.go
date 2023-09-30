package main

import (
	"fmt"
	"math"
)

func isPrime(num int) {
	if num < 2 {
		fmt.Println("Enter the digit greater than 2")
		return
	}
	squareRoot := int(math.Sqrt(float64(num)))
	for i:=2; i<=squareRoot; i++ {
		if num%i==0 {
			fmt.Println("It's not a prime number")
			return
		}
	}
	fmt.Println("It's a prime number")
}

func main(){

	isPrime(1000000007)
	
}





