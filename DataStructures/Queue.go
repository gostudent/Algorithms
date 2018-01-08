package main

import "fmt"

type Queue struct {
	size int
	front *Node
}

type Node struct {
	value string
	next *Node
}

func (q *Queue) Length() int {
  return q.size
}

func (q *Queue) Enqueue(val string) {
	
	if(q.front == nil){
		q.front = &Node{val, nil}
	}else{
		rear := q.front
		for rear.next != nil {
		  rear = rear.next
		}
		rear.next = &Node{val, nil}
	}
	fmt.Printf("%v enqueued to queue\n",val)
	q.size++
}

func (q *Queue) Dequeue() (val string) {
  if q.size > 0 {
    val, q.front = q.front.value, q.front.next
    q.size--
    return 
  }
  return
}

func main() {
	queue := &Queue{}
	
	queue.Enqueue("10")
	queue.Enqueue("20")
	queue.Enqueue("30")
	
	fmt.Println()
	
	for queue.Length() > 0 {
		fmt.Printf("%s popped from queue\n", queue.Dequeue())
	}
}