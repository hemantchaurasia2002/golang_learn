//reverse array
//INPUT ARRAY =. [1,2,3,4,5,6,7,8,9,10]
//OUTPUT ARRAY = [10,9,8,7,6,5,4,3,2,1]

package main

import (
	"fmt"
)
//USER INPUT SWAP ARRAY
// func main() {
// 	//declaring variables
// 	var arraySize,j int
// 	fmt.Println("Enter the size of array = ")
// 	fmt.Scanln(&arraySize)

// 	//declaring arrays
// 	array := make([]int, arraySize)
// 	reverseArray := make([]int, arraySize)

// 	fmt.Println("Enter element - ")
// 	//running a loop for storing values in array
// 	for i:=0;i<arraySize;i++ {
// 		fmt.Scan(&array[i])
// 	}
	
// 	j = 0
// 	//running a reverse loop of array and storing values of array into reverseArray
// 	for i:=arraySize-1; i>=0; i-- {
// 		reverseArray[j] = array[i]
// 		j++
// 	}
// 	fmt.Println("Reverse array - ", reverseArray)
// }


func main() {

	//declaring arrays
	array := [10]int{1,2,3,4,5,6,7,8,9,10}
	reverseArray := [10]int{}
	var arrayLen = len(array)

	for i:=0;i<10;i++ {
		fmt.Print(" ",array[i])
	}
	
	var j = 0
	for i:=; i>=0; i-- {
		reverseArray[j] = array[i]
		j++
	}
	fmt.Println("\n Reverse array - ", reverseArray)
}

func reverseArray(array []int) []int{
	for i, j :=0, len(array)-1; i<j; i,j = i+1, j-1 {
		array[i]=array[j]
		array[j]=array[i]
	}
	retrun array
}

func main(){
	fmt.Println(reverseArray([]int{1,2,3,4,5,6,7,8,9,10}))
}