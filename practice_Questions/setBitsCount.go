//find setbits(1) of an integer
//Input - 15
//Output - 4

package main

import (
	"fmt"
)

//method 1
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

//Method 2
func setBitsCounter(n int) int {
	var count int = 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func main() {
	var number int
	fmt.Printf("Enter the integer which you want to count setbits : ")
	fmt.Scanf("%d", &number)
	fmt.Println("SetBits : ", setBitsCounter(number))
}
