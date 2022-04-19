package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	leftChild *TreeNode
	righChild *TreeNode
}

func (currentNode *TreeNode) search(value int) bool {

	if currentNode == nil {
		return false
	}

	if currentNode.value > value {
		return currentNode.leftChild.search(value)
	}
	
	if currentNode.value < value {
		return currentNode.righChild.search(value)
	}
	return true
}

func main() {
	node1 := &TreeNode{value: 10}
	node2 := &TreeNode{value: 20}
	root := &TreeNode{value: 15, leftChild: node1, righChild: node2}

	fmt.Println(root.search(10)) // true
	fmt.Println(root.search(11)) // false
	fmt.Println(root.search(15)) // true
}