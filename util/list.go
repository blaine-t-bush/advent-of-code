package util

import (
	"fmt"
	"log"
)

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

// insert a new node at the tail end of the list
func (L *List) Insert(key interface{}) *Node {
	// create new node pointer
	node := &Node{
		prev: L.tail,
		key:  key,
	}

	// if list already has a tail (i.e. this is not the first entry),
	// then it becomes our new node's parent. otherwise, this is the first entry
	// so our new node is both the head and tail
	if L.tail != nil {
		L.tail.next = node
	} else {
		L.head = node
	}

	// our new node becomes the tail
	L.tail = node

	return node
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

func (L *List) MoveForward(N *Node) {
	// 4 cases:
	// 1. node is head
	// 2. node is tail
	// 3. node's child is tail
	// 4. other
	if N == L.head {
		child := N.next
		grandchild := N.next.next
		// node's child becomes new parent (and head)
		child.prev = nil
		child.next = N
		N.prev = child
		L.head = child
		// node's grandchild becomes new child
		N.next = grandchild
		grandchild.prev = N
	} else if N == L.tail {
		parent := N.prev
		head := L.head
		headchild := head.next
		// node's parent becomes tail
		parent.next = nil
		L.tail = parent
		// head becomes new parent
		N.prev = head
		head.next = N
		// head's child becomes new child
		headchild.prev = N
		N.next = headchild
	} else if N.next == L.tail {
		parent := N.prev
		child := N.next
		// node becomes tail
		N.next = nil
		L.tail = N
		// node's child becomes new parent
		child.next = N
		N.prev = child
		// node's parent becomes parent of new parent
		parent.next = child
		child.prev = parent
	} else {
		parent := N.prev
		child := N.next
		grandchild := child.next
		// node's parent becomes parent of child
		parent.next = child
		child.prev = parent
		// node becomes child of child
		child.next = N
		N.prev = child
		// node becomes parent of grandchild
		N.next = grandchild
		grandchild.prev = N
	}
}

func (L *List) MoveBackward(N *Node) {
	// 4 cases:
	// 1. node is tail
	// 2. node is head
	// 3. node's parent is head
	// 4. other
	if N == L.tail {
		parent := N.prev
		grandparent := parent.prev
		// node's parent becomes new child (and tail)
		N.next = parent
		parent.prev = N
		parent.next = nil
		L.tail = parent
		// node's grandparent becomes new parent
		grandparent.next = N
		N.prev = grandparent
	} else if N == L.head {
		child := N.next
		tail := L.tail
		tailparent := L.tail.prev
		// node's child becomes new head
		child.prev = nil
		L.head = child
		// node becomes child of tail's parent
		tailparent.next = N
		N.prev = tailparent
		// node becomes parent of tail
		N.next = tail
		tail.prev = N
	} else if N.prev == L.head {
		parent := N.prev
		child := N.next
		// node's parent becomes new child
		N.next = parent
		parent.prev = N
		// node's child becomes child of parent
		child.prev = parent
		parent.next = child
		// node becomes head
		N.prev = nil
		L.head = N
	} else {
		grandparent := N.prev.prev
		parent := N.prev
		child := N.next
		// node's parent becomes new child
		N.next = parent
		parent.prev = N
		// node's grandparent becomes new parent
		grandparent.next = N
		N.prev = grandparent
		// node's child becomes child of new child
		parent.next = child
		child.prev = parent
	}
}

func (L *List) MoveX(N *Node, steps int) {
	if steps > 0 {
		for i := 0; i < steps; i++ {
			L.MoveForward(N)
		}
	} else if steps < 0 {
		for i := 0; i < -steps; i++ {
			L.MoveBackward(N)
		}
	}
}

func (L *List) Length() int {
	if L.head == nil {
		return 0
	}

	head := L.head
	length := 1
	for {
		if head.next == nil {
			break
		}
		head = head.next
		length++
	}

	return length
}

func (L *List) Index(N *Node) int {
	index := 0
	head := L.head
	for {
		if head == N {
			return index
		}

		if head.next == nil {
			break
		}

		head = head.next
		index++
	}

	log.Fatal("L.Index: could not find given node")
	return index
}

func (L *List) Display() {
	current := L.head
	for current != nil {
		fmt.Printf("%+v -> ", current.key)
		current = current.next
	}
	fmt.Println()
}

func (L *List) NextCyclic(N *Node) *Node {
	if N.next != nil {
		return N.next
	}

	return L.head
}

func (N *Node) Key() interface{} {
	return N.key
}
