package main

import (
	"fmt"
	"math"
	"sort"
)

func findGreatest(list []int) int {
	// O(N^2)
  max := -1
	for i, el := range list {
		max = el
		for j, el2 := range list {
			if i != j {
				if max < el2 {
					max = el2
				}
			}
		}
	}
	return max
}

func findGreatest_2(list []int) int {
	// O(N LogN)
	sort.Ints(list)
	return list[len(list) - 1]
}

func findGreatest_3(list []int) int {
	// O(N)
  max := int(math.Inf(-1))
	for _, num := range list {
		if (int(max) < num) {
			max = num
		}
	}
	return max
}

func main() {
	fmt.Println(findGreatest([]int{1, 77, 100, 2, 5, 6, 12, 99, 109, 2, 1004, 1, 2, 5, 8, 11, 109, 99, 1003}))
	fmt.Println(findGreatest_2([]int{1, 77, 100, 2, 5, 6, 12, 99, 109, 2, 1004, 1, 2, 5, 8, 11, 109, 99, 1003}))
	fmt.Println(findGreatest_3([]int{1, 77, 100, 2, 5, 6, 12, 99, 109, 2, 1004, 1, 2, 5, 8, 11, 109, 99, 1003}))
}
