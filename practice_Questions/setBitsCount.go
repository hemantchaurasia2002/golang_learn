package main

import (
	"fmt"
)

func setBitsCount(num int) {
	var count int
	count = 0
	for i := 0; i < 32; i++ {
		if num&(1<<i) != 0 {
			count++
		}
	}
	fmt.Println("Number of setbits : ", count)
	return
}

func main() {
	var number int
	fmt.Printf("Enter the integer which you want to count setbits : ")
	fmt.Scanf("%d", &number)
	setBitsCount(number)
}
