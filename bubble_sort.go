package main

import (
	"fmt"
)

func bubbleSort(numList []int) []int {

	isSorted := false
	endIndex := len(numList) - 1
	index := 0
	for !isSorted {
		isSorted = true
		for index < endIndex {
			if numList[index] > numList[index + 1] {
				isSorted = false
				numList[index], numList[index + 1] = numList[index + 1], numList[index]
			}
			index++
		}
		endIndex--
		index = 0
	}
	return numList
}

func main() {
  fmt.Println(bubbleSort([]int{11, 7, 8, 32, 8, 90}))
}