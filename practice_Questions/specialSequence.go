// special subsequence count
//input - abgcagag
//output - 6

package main

import (
	"fmt"
)

func main() {
	var seq string
	fmt.Print("Enter the sequence : ")
	fmt.Scan(&seq)
	var a_count int = 0
	var total int = 0
	for i := 0; i < len(seq); i++ {
		if seq[i] == 'a' {
			a_count++
		}
		if seq[i] == 'g' {
			total += a_count
		}

	}
	fmt.Println("a Count : ", a_count)
	fmt.Println("Total : ", total)
}
