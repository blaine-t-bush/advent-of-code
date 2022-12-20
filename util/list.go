package util

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}

	if L.head != nil {
		L.head.prev = list
	}

	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}

	L.tail = l
}

func (L *List) InsertAt(index int, key interface{}) {
	newNode := &Node{
		next: L.head,
		key:  key,
	}

	nodeToBreak := L.head
	// do "next" index times
	for i := 0; i < index-1; i++ {
		nodeToBreak = nodeToBreak.next
	}

	newNode.prev = nodeToBreak
	newNode.next = nodeToBreak.next
	nodeToBreak.next = newNode

	// change head or tail if needed
}

func (L *List) Delete(index int) {
	nodeToDelete := L.head
	for i := 0; i < index; i++ {
		nodeToDelete = nodeToDelete.next
	}

	nodeToDelete.prev.next = nodeToDelete.next
	nodeToDelete.next.prev = nodeToDelete.prev

	// change head or tail if needed
}

func (L *List) Display() {
	list := L.head
	for list != nil {
		fmt.Printf("%+v -> ", list.key)
		list = list.next
	}
	fmt.Println()
}
