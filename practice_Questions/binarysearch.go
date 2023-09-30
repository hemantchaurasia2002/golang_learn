//Given a sorted array of size N and an integer K, find the position(0-based indexing) at which K is present in the array using binary search.
/*Input:
N = 5
arr[] = {1 2 3 4 5} 
K = 4
Output: 3
Explanation: 4 appears at index 3.*/
//----------------------------------------------------------------------------------------

package main

import (
	"fmt"
)

func binarySearch(arr []int, n int) {
	for i:=0;i<len(arr);i++ {
		if arr[i] == n {
			fmt.Printf("%v ", i)
			return
		}
	}
	fmt.Println("-1")
}

func main() {
	var K int
	fmt.Println("Enter the value which you want to check present in array - ")
	fmt.Scan(&K)
	arr := []int{1,2,3,4,5,6,7,8,9,10}
	binarySearch(arr, K)
}