//linkedlist which stores odd numbers from 1to20 


package main

import (
	"fmt"
)

type node struct {
	Val int
	Next *Node
}


func createLinkedList(head **Node, x int) *Node{
	temp := Node{Next: nil, Val: x}
	current := head
	if head == nil {
		head = &temp
	} else{
		for i := 0;i<=n;i++ {
			if i%2!=0 {
				current.Next = &temp
			}
		}
		return head
	}
}

func displayList(head *Node) {
	fmt.Printf("The list is : ")
	for; head !=nil; head.Next {
		fmt.Print(head.Data, " ")
	}
	fmt.Println()
}



func main(){
	var p *Node
	p = nil
	displayList(p)
	
}