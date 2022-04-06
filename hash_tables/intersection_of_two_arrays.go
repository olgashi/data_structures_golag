package main

import (
	"fmt"
)

func arrayIntersection(slice1, slice2 []int) (intersectionSlice []int) {
	// technically I am using slice(s) here

	var largestSlice []int
	var smallestSlice []int

	if len(slice1) > len(slice2) {
		largestSlice = slice1
		smallestSlice = slice2
	} else {
		largestSlice = slice2
		smallestSlice = slice1
	}
  //O(N), N is length of the largest array/slice
	var largestSliceMap = make(map[int]bool)
	for i, _ := range largestSlice {
		largestSliceMap[largestSlice[i]] = true
	}

	for _, el := range smallestSlice {
		if _, ok := largestSliceMap[el]; ok {
			intersectionSlice = append(intersectionSlice, el)
		}
	}
	return
}

func main() {
	fmt.Println(arrayIntersection([]int{1, 2, 3, 4, 5}, []int{0, 2, 4, 6, 8})) // [2, 4]
	fmt.Println(arrayIntersection([]int{1, 2, 3, 4, 5}, []int{0, 0, 0, 0, 0})) // []
	fmt.Println(arrayIntersection([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5})) // [1, 2, 3, 4, 5]
}