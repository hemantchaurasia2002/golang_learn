
/*Write a program to take number of input as N then take N values as input. 
Negative value is restricted. 
You have to create a map with key as input number and value should me array of factors of that numbers.
N = 4 
arr = [ 3,4,6,2]

res = {
	3: [1, 3],
	4: [1,2,4],
	6: [1,2,3,6],
	2: [1,2]
}
*/

package main 

import (
	"fmt"
	// "math"
)

func findFactor(num int) []int {
	//input int num
	//returns array of factors
	var factorArray = make([]int, 0)
	factorArray = append(factorArray, 1)
	for i:=2; i<=num; i++ {
		if num%i==0 {
			// fmt.Println(i)
			factorArray = append(factorArray,i)
		}
	} 
	return factorArray
	
	// fmt.Println("FACTORS ARRAY - ", factorArray)

}

func main() {
	fmt.Printf("Enter the size of the array - ")
	var size int
	fmt.Scanln(&size)
	var array = make([]int, size)
	resultmap := map[int] []int{}
	for i:=0; i<size; i++ {
		fmt.Printf("%dth Element - ", i)
		fmt.Scanf("%d\n", &array[i])
		resultmap[array[i]]= findFactor(array[i])
	}
	fmt.Println("Result - ", resultmap)
	// fmt.Println("ARRAY - ", array)
}



