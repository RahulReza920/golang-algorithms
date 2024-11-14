package main

import "fmt"

type List struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func (list *List) Insert(v int) *List {

	if list == nil {
		fmt.Println("List is nil")
		return nil

	}

	temp := new(Node)
	temp.value = v
	temp.next = list.head
	list.head = temp
	return list

}

func (list *List) Print() {
	if list == nil {
		fmt.Println("List not created")
	}
	temp := list.head

	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.next
	}
}

func main() {
	lst := List{}
	lst.Insert(1)
	lst.Insert(2)
	lst.Insert(3)
	lst.Insert(4)
	lst.Insert(5)
	lst.Insert(6)
	lst.Print()

}
