package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) addValue(value int) {
	newNode := &Node{data: value}

	if l.head == nil {
		l.head = newNode
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func printList(l *List) {
	current := l.head
	for current != nil {
		fmt.Printf("%d->", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := &List{}
	list.addValue(10)
	list.addValue(20)
	list.addValue(30)
	list.addValue(40)

	fmt.Println("Linkedlist - ")
	printList(list)
}
