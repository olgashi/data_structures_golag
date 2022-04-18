package main

import(
	"fmt"
)

type node struct {
	data int
	next_node *node
}

type linkedList struct {
	head *node
}

func (list *linkedList) readByIndex(index int) (node *node) {
  current_node := list.head
	current_index := 0

	for current_index < index {
		current_node = current_node.next_node
		current_index++
		if current_node == nil {
			return nil
		}
	}
	return current_node
}

func (list *linkedList) getIndexOf(value int) (nodeIndex int) {
	current_node := list.head

	for current_node != nil {
		if current_node.data == value {
			return 
		}
		current_node = current_node.next_node
		nodeIndex++
	}
	return -1
}

func (list *linkedList) insertAtIndex(index int, newNode *node) bool {
	if index == 0 && list.head != nil {
		newNode.next_node = list.head
		list.head = newNode
		return true
	}
	
	current_index := 0
  
	current_node := list.head
	for current_node != nil {
		if current_index == index-1 {
			newNode.next_node = current_node.next_node
			current_node.next_node = newNode
			return true
		}
		current_node = current_node.next_node
		current_index++
	}
  return false
}

func(list *linkedList) delteByIndex(index int) bool {
	current_node := list.head
	if current_node == nil {
		return false
	}

	if index == 0 {
		list.head = current_node.next_node
		return true
	}

	current_index := 0

	for current_node != nil {
		if current_index == index - 1 {
			current_node.next_node = current_node.next_node.next_node
			return true
		}
		current_index++
		current_node = current_node.next_node
	}
	return false
}

func (list *linkedList) printListElements() bool {
  current_node := list.head
	if current_node == nil {
		return false
	}

	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next_node
	}
	return true
}

func (list *linkedList) getLastElement() (node *node) {
	current_node := list.head

	if current_node == nil {
		return nil
	}

	for current_node != nil {
		if current_node.next_node == nil {
			return current_node
		}
		
		current_node = current_node.next_node
	}
	return nil
}

func main() {
	node5 := &node{data:14}
	node4 := &node{data:13, next_node: node5}
	node3 := &node{data:12, next_node: node4}
	node2 := &node{data:11, next_node: node3}
	node1 := &node{data:10, next_node: node2}
	
	list1 := &linkedList{head: node1}
	node6 := &node{data:99}

	fmt.Println(list1.printListElements())
	fmt.Println("==============")

	fmt.Println(list1.insertAtIndex(0, node6))
	fmt.Println(list1.printListElements())
	fmt.Println("==============")

	fmt.Println(list1.getLastElement())
}