package main

import (
	"fmt"
	"sort"
)

func findMissingNumber(list []int) int {
// with sorting
// total time complexity is O(N LogN)
sort.Ints(list)
fmt.Println(list)

for index, el := range list {
	if el != index {
		return el - 1
	}
}
return -1
}

func findMissingNumber_2(list []int) int {
	// total time complexity is O(N), space is O(N)
	emptyNumsList := make([]int, len(list) + 1)

	for _, el := range list {
		emptyNumsList[el] = 1
	}

	for index, el := range emptyNumsList {
		if el != 1 {
			return index
		}
	}
  return -1
}


func main() {
  fmt.Println(findMissingNumber([]int{9,3,2,5,6,7,1,0,4}))
  fmt.Println(findMissingNumber_2([]int{9,3,2,5,6,7,1,0,4}))
}