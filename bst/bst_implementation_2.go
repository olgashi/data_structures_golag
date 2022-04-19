package main

import (
	"fmt"
)

type TreeNode struct {
	value int
	leftChild *TreeNode
	rightChild *TreeNode
}

func (currentNode *TreeNode) search(value int) bool {

	if currentNode == nil {
		return false
	}

	if currentNode.value > value {
		return currentNode.leftChild.search(value)
	}
	
	if currentNode.value < value {
		return currentNode.rightChild.search(value)
	}
	return true
}

func (startingNode *TreeNode) insert(value int) {
  if value < startingNode.value {

		if startingNode.leftChild == nil {
			startingNode.leftChild = &TreeNode{value: value}
		} else {
			startingNode.leftChild.insert(value)
		}
	} else {
		if startingNode.rightChild == nil {
			startingNode.rightChild = &TreeNode{value: value}
		} else {
			startingNode.rightChild.insert(value)
		}
	}
}

func(startingNode *TreeNode) delete(value int) *TreeNode {
	if startingNode != nil {
		return startingNode
	
	} else if value < startingNode.value {
		startingNode.leftChild = startingNode.leftChild.delete(value)
	  return startingNode
		
	} else if value > startingNode.value {
	  startingNode.rightChild = startingNode.rightChild.delete(value)
	  return startingNode
	
	} else if value == startingNode.value {

		if startingNode.leftChild == nil {
			return startingNode.rightChild
		
		} else if startingNode.rightChild == nil {
			return startingNode.leftChild
		
		} else {
			startingNode.rightChild = startingNode.rightChild.lift(startingNode)
			return startingNode
		}
	}
	return nil
}

func(node *TreeNode) lift(nodeToDelete *TreeNode) *TreeNode {
	if node.leftChild != nil {
		node.leftChild = node.leftChild.lift(nodeToDelete)
		return node
	
	} else {
		nodeToDelete.value = node.value
		return node.rightChild
	}

}

func (startingNode *TreeNode) findeGreatest() int {
  if startingNode.rightChild != nil {
		return startingNode.rightChild.findeGreatest()
	}
	return startingNode.value
}

func (startingNode *TreeNode) inOrderTraversalAndPrint() {
	if startingNode == nil {
		return
	}
	startingNode.leftChild.inOrderTraversalAndPrint()
	fmt.Println(startingNode.value)
	startingNode.rightChild.inOrderTraversalAndPrint()
}

func main() {
	// node1 := &TreeNode{value: 10}
	// node2 := &TreeNode{value: 20}
	root := &TreeNode{value: 50, leftChild: nil, rightChild: nil}
	root.insert(25)
	root.insert(75)
	root.insert(10)
	root.insert(33)
	root.insert(56)
	root.insert(89)
	root.insert(4)
	root.insert(11)
	root.insert(30)
	root.insert(40)
	root.insert(52)
	root.insert(61)
	root.insert(82)
	root.insert(95)

	fmt.Println(root.search(95)) // true
	fmt.Println(root.search(10)) // true
	fmt.Println(root.search(11)) // true
	fmt.Println(root.search(15)) // false
	fmt.Println(root.search(11)) // true
	fmt.Println(root.search(16)) // false

	fmt.Println("Greatest", root.findeGreatest())// 95
	root.inOrderTraversalAndPrint() // 4, 10, 11, 25, 30, 33, 40, 50, 52, 56, 61, 75, 82, 89, 95
}