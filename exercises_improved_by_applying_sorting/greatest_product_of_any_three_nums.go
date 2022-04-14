package main

import (
	"fmt"
	"sort"
)
// assumptio  is that numbers in list are not unique

func greatestSumOfThree(list []int) (sum int) {
	sort.Ints(list)

	count := 1
	prev := list[len(list) - 1]
	sum = prev

	for index := len(list) - 2; index >= 0; index-- {
		if count == 3 {
			break
		}
		if list[index] != prev {
			prev = list[index]
			sum += prev
			count++
		}
	}
	return
}

func main() {
	fmt.Println(greatestSumOfThree([]int{123, 6, 1, 3, 2, 1, 2, 3, 9})) //138
	fmt.Println(greatestSumOfThree([]int{1, 2, 3, 4, 5, 6})) // 15
	fmt.Println(greatestSumOfThree([]int{0, 0, 0, 1, 0, 0, 1})) // 2
	fmt.Println(greatestSumOfThree([]int{10, 10, 100, 1, 1000, 99, 1})) // 1199
}