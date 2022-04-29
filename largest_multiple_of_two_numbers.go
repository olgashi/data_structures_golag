package main

import (
	"fmt"
	"math"
)

func largest_multiple(nums []int) float64 {
	smallNeg1 := math.Inf(0)
	smallNeg2 := math.Inf(0)

	largPos1 := math.Inf(-1)
	largPos2 := math.Inf(-1)

	for _, num := range nums {
		if float64(num) <= smallNeg1 {
			smallNeg2 = smallNeg1
			smallNeg1 = float64(num)
		} else if float64(num) > smallNeg1 && float64(num) < smallNeg2 {
			smallNeg2 = float64(num)
		}

		if float64(num) >= largPos1 {
			largPos2 = largPos1
			largPos1 = float64(num)
		} else if float64(num) < largPos1 && float64(num) > largPos2 {
      largPos2 = float64(num)
		}
	}

	possibleLargest1 := smallNeg1 * smallNeg2
	possibleLargest2 := largPos1 * largPos2

	if possibleLargest1 > possibleLargest2 {
		return possibleLargest1
	}

	return possibleLargest2
}

func main() {
	fmt.Println(largest_multiple([]int{5, -4, -3, 0, 3, 4}))
	fmt.Println(largest_multiple([]int{-9, -2, -1, 2, 3, 7}))
	fmt.Println(largest_multiple([]int{-7, -4, -3, 0, 4, 6}))
	fmt.Println(largest_multiple([]int{-6, -5, -1, 2, 3, 9}))
	fmt.Println(largest_multiple([]int{-9, -4, -3, 0, 6, 7}))
}