package main

import "fmt"

const ArraySize = 7
// HashTable structure

type HashTable struct {
	array[ArraySize] *bucket
}

// bucket structure
type bucket struct {
	head *bucketNode
}

// bucketNode structure

type bucketNode struct {
	key string
	next *bucketNode
}

// Insert

func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
	
}
// Search
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)

	
}
// Delete
func (h *HashTable) Delete(key string) int {
	index := hash(key)
	return index

}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println("Already exists", k)
	}
}

func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func (b *bucket) delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}
	previousNode := b.head

	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
	}
	
}

// hash
func hash(key string) int {
	sum := 0

	for _, char := range key {
		sum += int(char)
	}
	return sum % ArraySize
}

func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

func main() {
	testHashTable := Init()
	list := []string{
		"ABC",
		"ABCD",
		"ABCDEF",
		"ABCDEFG",
	}
	for _, v := range list {
		testHashTable.Insert(v)
	}
	fmt.Println(testHashTable.Delete("ABC"))
	fmt.Println("ABC", testHashTable.Search("ABC"))
}