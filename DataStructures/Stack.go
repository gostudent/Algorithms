package main

import "fmt"

type Stack struct {
	size int
	top *Node
}

type Node struct {
	value string
	next *Node
}

func (s *Stack) Length() int {
  return s.size
}

func (s *Stack) Push(val string) {
	s.top = &Node{val, s.top}
	fmt.Printf("%v pushed to stack\n",val)
	s.size++
}

func (s *Stack) Pop() (val string) {
  if s.size > 0 {
    val, s.top = s.top.value, s.top.next
    s.size--
    return 
  }
  return
}

func main() {
	stack := &Stack{}
	
	stack.Push("10")
	stack.Push("20")
	stack.Push("30")
	
	fmt.Println()
	
	for stack.Length() > 0 {
		fmt.Printf("%s popped from stack\n", stack.Pop())
	}
}