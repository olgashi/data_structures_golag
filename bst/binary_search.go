package main

import (
	"fmt"
)


func binarySearch(inputList []int, val int) int {
	lowerBound := 0
	upperBound := len(inputList) - 1
	
	for lowerBound <= upperBound {
		midpoint := (upperBound + lowerBound) / 2
		midpointValue := inputList[midpoint] 
		if midpointValue == val {
			return midpoint
		
		} else if midpointValue > val {
			upperBound = midpoint - 1

		} else if midpointValue < val {
			lowerBound = midpoint + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarySearch([]int{1, 4, 5, 8, 10, 11, 15, 21, 22}, 10))
	fmt.Println(binarySearch([]int{1, 4, 5, 8, 10, 21, 25, 31, 70}, 1))
	fmt.Println(binarySearch([]int{11, 101, 110, 150, 200, 223}, 223))
	fmt.Println(binarySearch([]int{1, 4, 5, 8, 10, 11, 15, 21, 22}, 99))
}