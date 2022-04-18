package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	nextNode := l.head
  for index := 0; index < l.length; index ++ {
		fmt.Println(nextNode.data)
		nextNode = nextNode.next
	}
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	} else {
		previousNode := l.head
		currentNode := l.head.next

		for currentNode != nil {
			if currentNode.data == value {
				previousNode.next = currentNode.next
				currentNode.next = nil
				l.length--
				return
			} else {
				previousNode = currentNode
				currentNode = currentNode.next
			}
		}

		return
	}
}

func main() {
	myList := linkedList{}
	myList.deleteWithValue(51)
	node1 := &node{data: 48}
	node2 := &node{data: 49}
	node3 := &node{data: 50}
	node4 := &node{data: 51}

	myList.prepend(node4)
	myList.prepend(node3)
	myList.prepend(node2)
	myList.prepend(node1)
	
	myList.printListData()
	fmt.Println("======")
	myList.printListData()
	myList.deleteWithValue(51)
	fmt.Println("======")
	myList.prepend(&node{data: 100})
	myList.printListData()
}