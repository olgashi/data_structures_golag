package main

import "fmt"


type Queue struct {
	items []int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() int {
	itemToDequeue := q.items[0]
	q.items = q.items[1:]

	return itemToDequeue
}

func main() {

	myQueue := Queue{}

	myQueue.Enqueue(4)
	myQueue.Enqueue(6)
	myQueue.Enqueue(9)

	fmt.Println(myQueue)
	
	myQueue.Dequeue()
	fmt.Println(myQueue)
}