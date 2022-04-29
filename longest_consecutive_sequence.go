package main

import (
	"fmt"
)

func longestConsequtiveSequence(nums []int) int {
  mapCount := make(map[int]bool) 

	for _, num := range nums {
		mapCount[num] = true
	}

	counter := 0
	for _, el := range nums {
		if _, ok := mapCount[el - 1]; !ok {
			subCounter := 1
			nextEl := el + 1;
			for true {
				if _, ok := mapCount[nextEl]; ok {
					subCounter++
					} else {
						break
					}
					nextEl++
			}
			if subCounter > counter {
				counter = subCounter
			}
		}
	}
	fmt.Println(mapCount)
	return counter
}

func main() {
	fmt.Println(longestConsequtiveSequence([]int{10, 5, 12, 3, 55, 30, 4, 11, 2}))
	fmt.Println(longestConsequtiveSequence([]int{19, 13, 15, 12, 18, 14, 17, 11}))
	fmt.Println(longestConsequtiveSequence([]int{6, 5, 4, 3, 2, 1}))
	fmt.Println(longestConsequtiveSequence([]int{2, 2, 2, 2, 2, 3}))
}