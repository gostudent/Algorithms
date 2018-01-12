//No need of custom input
//Values already added
package main

import "fmt"

// value is the value of node; next is the pointer to next node 
type node struct {
	value int
	next *node
}

// first node, called head. It points from first node to last node 
var head *node = nil

func (l *node) pushFront(val int) *node {
	// if there's no nodes, head points to l (first node) 
	if head == nil {
		l.value = val
		l.next = nil
		head = l
		return l
	} else {
		// creating a new node equals to head 
		nnode := new(node)
		nnode = head
		// creating a second node with new value and `next -> nnode`
		//  is this way, nnode2 is before nnode
		nnode2 := &node {
			value: val,
			next: nnode,
		}
		// now head is equals nnode2 
		head = nnode2
		return head
	}
}

func (l *node) pushBack(val int) *node {
	// if there's no nodes, head points to l (first node) 
	if head == nil {
		l.value = val
		l.next = nil
		head = l
		return l
	} else {
		// read all list to last node 
		for l.next != nil {
			l = l.next
		}
		// allocate a new portion of memory 
		l.next = new(node)
		l.next.value = val
		l.next.next = nil
		return l
	}
}

func (l *node) popFront() *node {
	if head == nil {
		return head
	}
	// creating a new node equals to first node pointed by head 
	cpnode := new(node)
	cpnode = head.next
	
	// now head is equals cpnode (second node) 
	head = cpnode

	return head
}

func (l *node) popBack() *node {
	if head == nil {
		return head
	}
	// creating a new node equals to head 
	cpnode := new(node)
	cpnode = head
	
	// read list to the penultimate node 
	for cpnode.next.next != nil {
		cpnode = cpnode.next
	}
	// the penultimate node points to null. In this way the last node is deleted 
	cpnode.next = nil
	return head
}

func main() {
	lista := new(node)
	lista.pushBack(99).pushBack(65).pushBack(5) /* lista: 99 65 5 */
	lista.pushBack(44) /* lista: 99 65 5 44 */
	lista.pushFront(873) /* lista: 873 99 65 5 44  */
	lista.popFront() /* lista: 99 65 5 44 */
	lista.popBack() /* lista: 99 65 5 */
	
	// read the list until head is not nil 
	for head != nil {
		fmt.Printf("%d ",head.value)
		head = head.next /*head points to next node */
	}
}
