package main

import "fmt"

type List struct {
	head  *ListNode
	count int
}

type ListNode struct {
	value   int
	nxtNode *ListNode
}

func (list *List) Size() int {
	return list.count
}
func (list *List) IsEmpty() bool {
	return list.count == 0
}
func (list *List) AddHead(val int) {

	list.head = &ListNode{value: val, nxtNode: list.head}
	list.count++
}

func (list *List) AddTail(val int) {

	newNode := &ListNode{value: val, nxtNode: nil}
	list.count++

	if list.head == nil {
		list.head = newNode
		return
	}

	curr := list.head

	for curr.nxtNode != nil {

		curr = curr.nxtNode
	}
	curr.nxtNode = newNode
}

func (list *List) Find(data int) bool {
	temp := list.head
	for temp != nil {
		if temp.value == data {
			return true
		}
		temp = temp.nxtNode
	}
	return false

}

func (list *List) Print() {

	temp := list.head

	if temp == nil {
		fmt.Println("empty linkedList")
	}

	for temp != nil {
		fmt.Println(temp.value)
		temp = temp.nxtNode
	}
	fmt.Println("")

}

func (list *List) SortedInsert(val int) {
	newNode := &ListNode{value: val}

	if list.head == nil || list.head.value > val {
		newNode.nxtNode = list.head
		list.head = newNode
		return
	}
	curr := list.head

	for curr.nxtNode != nil && curr.nxtNode.value < val {
		curr = curr.nxtNode
	}

	newNode.nxtNode = curr.nxtNode
	curr.nxtNode = newNode
}

func (list *List) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("LinkedList Empty")
		return 0, false
	}
	value := list.head.value
	list.head = list.head.nxtNode
	list.count--

	return value, true
}
func main1() {
	ll := new(List)
	ll.AddHead(1)
	ll.AddHead(2)
	ll.AddHead(3)
	ll.Print()
	fmt.Println("Size :", ll.Size())
	fmt.Println("IsEmpty :", ll.IsEmpty())
	ll.RemoveHead()

	ll.Print()
}

func main2() {
	ll := new(List)
	ll.AddTail(1)
	ll.AddTail(2)
	ll.AddTail(3)
	ll.Print()
	fmt.Println("Size :", ll.Size())
	fmt.Println("IsEmpty :", ll.IsEmpty())
}

func main3() {

	ll := new(List)
	ll.SortedInsert(4)
	ll.SortedInsert(6)
	ll.SortedInsert(3)
	ll.SortedInsert(1)
	ll.SortedInsert(8)
	ll.SortedInsert(5)
	ll.Print()

}

func main() {
	main1()

}
