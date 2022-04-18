package main

import (
	"fmt"
)

type node struct {
	next_node *node
	prev_node *node
	data int
}

type linkedList struct {
	head *node
	tail *node
}

func (list *linkedList) instertAtEnd(value int) {
	newNode := &node{data: value}
	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	}

	list.tail.next_node = newNode
	newNode.prev_node = list.tail

	list.tail = newNode
}

func(list *linkedList) printListElements() bool {
  if list.head == nil {
		return false
	}

	current_node := list.head
  
	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next_node
	}
	return true
}

func(list *linkedList) printListElementsInReversedOrder() bool {
  if list.tail == nil {
		return false
	}

	current_node := list.tail

	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.prev_node
	}
	return true
}

func main() {
	headNode := &node{data:900}
	list1 := &linkedList{head: headNode, tail: headNode}
	list1.instertAtEnd(700)
	list1.instertAtEnd(400)
	list1.instertAtEnd(200)
	list1.instertAtEnd(500)
	list1.printListElements()
	list1.printListElementsInReversedOrder()
}