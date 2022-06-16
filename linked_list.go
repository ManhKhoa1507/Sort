package main

type Node struct {
	Value int   // key
	Next  *Node // Next node
}

type List struct {
	Head *Node
	Tail *Node
}

// Insert 1 node to tail of linked list
func (l *List) Insert(value int) {
	node := createNewNode(value)
	if l.checkEmpty() == true {
		// If empty linked list add to head
		l.Head = node
		l.Tail = node

	} else {

		// Not empty linked list
		l.Tail.Next = node
		l.Tail = node
	}

}

// Check empty linked list
func (l *List) checkEmpty() bool {
	// Return true if empty list (head = tail = nil)
	// Return false if empty
	if l.Head == nil && l.Tail == nil {
		return true
	} else {
		return false
	}
}

// Create 1 node with value k
func createNewNode(value int) *Node {
	node := &Node{
		Value: value,
		Next:  nil,
	}
	return node
}
